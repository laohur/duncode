package com.aiart;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

public class Duncoder {
    Tool tool;
//    int bufferSize=1024*8;
    Duncoder(){
        if(tool==null)
            tool=new Tool();
    }

    List<Byte> stringToBytes(String s){
        return duncodeToBytes(charsToDuncodes(s.toCharArray()));
    }

    List<Byte> charsToBytes(char[] chars){
        return duncodeToBytes(charsToDuncodes(chars));
    }
    List<Duncode> charsToDuncodes(char[] chars){
        Duncode lastDun = new Duncode(tool);
        List<Duncode> dunList = new ArrayList<>();
        for (char c:chars) {
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            if (lastDun.compress(nowDun)) //char[] test
                continue;//只需要原来Duncode附加即可  有更新,加完了再处理
            else{
                dunList.add(nowDun);
                lastDun = nowDun;
            }
        }
        return  dunList;
    }

    List<Byte> duncodeToBytes(List<Duncode> duncodes){
        List<Byte> byteList = new ArrayList<>();
        for(Duncode d:duncodes){
            for(Byte b:d.getBytes())
                byteList.add(b);
        }
        return  byteList;
    }

    String bufferToString(byte[] bytes,List<Byte> bucket, StringBuilder sb , Duncode nowDun,List<Character> chars){
        if (bucket==null)
            bucket= new ArrayList<>();
//      left      bucket.clear();
        if(sb==null)
            sb=new StringBuilder();
        else
            sb.setLength(0);;
        if(nowDun==null)
            nowDun = new Duncode(tool);
        for (Byte b : bytes) {
            bucket.add(b);
            if ((b & 0x80) == 0) { //末尾字节
                nowDun.readDunchars(bucket);
                for(char c:nowDun.getChars(chars))
                    sb.append(c);
                bucket.clear();
            }
        }
        return sb.toString();
    }

    String bytesToString(List<Byte> bytes,List<Byte> bucket, StringBuilder sb , Duncode nowDun,List<Character> chars){
        if (bucket==null)
            bucket= new ArrayList<>();
        else
            bucket.clear();
        if(sb==null)
            sb=new StringBuilder();
        else
            sb.setLength(0);;
        if(nowDun==null)
            nowDun = new Duncode(tool);
        for (Byte b : bytes) {
            bucket.add(b);
            if ((b & 0x80) == 0) { //末尾字节
                nowDun.readDunchars(bucket);
                for(char c:nowDun.getChars(chars))
                    sb.append(c);
                bucket.clear();
            }
        }
//        bytes.clear();
        return sb.toString();
    }

    char[] bytesToChars(Byte[] bytes,List<Character> re){
        List<Byte> bucket = new ArrayList<>();
        StringBuilder sb=new StringBuilder();

        for (Byte b : bytes) {
            bucket.add(b);
            if ((b & 0x80) == 0) { //末尾字节
                Duncode nowDun = new Duncode(tool);
                nowDun.readDunchars(bucket);
                for(char c:nowDun.getChars(re))
                    sb.append(c);
                bucket.clear();
            }
        }
        return sb.toString().toCharArray();
    }

    List<List<Byte>> readU(String path)   {
        List<List<Byte>>byteMatrix=new ArrayList<>();
        try {
            BufferedReader br = new BufferedReader(new FileReader(path));
            String line;
            char[] chars;
            while ((line = br.readLine()) != null) {
                chars = line.toCharArray();
                byteMatrix.add(charsToBytes(chars));
            }
            br.close();
        }catch (Exception e){
            e.printStackTrace();
        }
        return byteMatrix;
    }

    char[][] readD(String path)   {
        List<char[]> charMatrix = new ArrayList<>();
        try {
            DataInputStream dis = new DataInputStream(new BufferedInputStream( new FileInputStream(path)));
            char[] chars;
            List<Byte> bytes = new ArrayList<>();
            while (dis.available() > 0) {
                byte b = dis.readByte();
                if (b != '\n') {
                    bytes.add(b);
                } else {
                    chars = bytesToChars(bytes.toArray(new Byte[bytes.size()]),null);
                    charMatrix.add(chars);
                    bytes.clear();
                }
            }
            dis.close();
        }catch (Exception e){
            e.printStackTrace();
        }
        return charMatrix.toArray(new char[charMatrix.size()][]);
    }

    void UfileToDfile(String inpath,String outpath)  {
        System.out.println("unicode文件开始读入"+inpath);
        try {
            File in=new File(inpath);
            long size=in.length();
//            long processed=0;
            BufferedReader br = new BufferedReader(new FileReader(in));
            String line;
            char[] chars;
            List<Byte> bytes;
            byte[] buffer;
            DataOutputStream dos = new DataOutputStream(new BufferedOutputStream(new FileOutputStream(outpath)));  //后置,万一上步出错不会生成空文件

            while (( line = br.readLine())!=null) {
                chars = line.toCharArray();
                bytes=charsToBytes(chars);
                buffer=new byte[bytes.size()+1];
                for(int i=0;i<bytes.size();i++)
                    buffer[i]=bytes.get(i);
                buffer[buffer.length-1]='\n';
                dos.write(buffer);
//                for (byte b : bytes)
//                    dos.write(b);
////                processed+=(bytes.size()+1);
////                if(processed%(bufferSize)==0){
////                    System.out.println(processed*100.0/size+"%已经写入 "+1.0*processed/(bufferSize)+"m "+(new Date()).toString());
////                }//                if(line!=null)//末尾空行会忽略,弃用
//                dos.write('\n'); //末尾无换行也会加上
            }
            br.close();
            dos.close();
            System.out.println("dunbyte文件已写入" + outpath);
        }catch (Exception e){
            e.printStackTrace();
        }
    }
    void DfileToUfile(String inpath,String outpath) {
        System.out.println("dunbyte文件开始读入"+inpath);
        try {
            File in=new File(inpath);
            long size=in.length();
            long processed=0;
            DataInputStream dis = new DataInputStream(new BufferedInputStream( new FileInputStream(in)));

            List<Byte> bytes = new ArrayList<>();
            List<Character> chars=new ArrayList<>();
            List<Byte> bucket=new ArrayList<>();
            StringBuilder sb=new StringBuilder();
            Duncode nowDun=new Duncode(tool);
            BufferedWriter bw = new BufferedWriter(new FileWriter(outpath));
            String line;
            byte[] buffer;
            int bufferSize=1024*8;
            int len;
            int count=dis.available();
            while (count > 0) {
//                len=Math.min(bufferSize,count);
                len=bufferSize<count?bufferSize:count;
                buffer = new byte[len];
                dis.read(buffer);
                line=bufferToString(buffer,bucket,sb,nowDun,chars);
                bw.write(line);
//                for(byte b:buffer){
//                    bytes.add(b);
//                    if( b=='\n' ){  //// 至少是是尾部字节,可能是换行,也可能是多字节尾
//                        bw.write(bytesToString(bytes,bucket,sb,nowDun,chars));
//    //                    processed+=bytes.size();
//    //                    if(processed%(bufferSize)==0)
//    //                        System.out.println(processed*100.0/size+"%已经写入 "+1.0*processed/(bufferSize)+"m "+(new Date()).toString());
//                        bytes.clear();    //文件可能不以'\n'结尾
//                    }
//                }
                count=dis.available();
            }
            dis.close();
            bw.close();
            System.out.println("unicode文件已写入" + outpath);
        }catch (Exception e){
            e.printStackTrace();
        }
    }


}
