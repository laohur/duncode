package com.aiart;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Duncode {
    int unichar = -1;  //symbleIds[0]
    int langId = -1;
    int zoneId = -1;
    List<Integer> symbleIds;
    Tool tool;

    Duncode(Tool tool) {
        this.tool = tool;
    }

    //读入unichar初始化
    void readUnichar(char u) {
        langId = tool.langId(u);
        zoneId = tool.zoneId(langId, u);
        switch (zoneId) {              // 1 4 5 需要 语言偏移量+字母偏移量+首位取1 //Ỿ 超过128 还要分配
            case 0:
                unichar=u;
                return;
            case 1:
                unichar=u - tool.heads[langId];
                return;
            case 2:
                unichar=(tool.han2idx.get(u));
                return;
            case 3:
                unichar=u;
                return;
            case 4: //四节三字
                symbleIds=  new ArrayList<>();
                symbleIds.add(u - tool.heads[langId]);
                return;
            case 5://四节双字 10+16
                symbleIds=  new ArrayList<>();
                symbleIds.add(u - tool.heads[langId]); //字母偏移
                return;
            default:
                unichar=u;
        }
    }

    //读入dunbyte[]初始化
    void readDunchars(List<Byte> list) {
        byte[] array = new byte[list.size()];
        for (int i = 0; i < array.length; i++)
            array[i] = list.get(i);
        readDunchars(array);
    }

    void readDunchars(byte[] bytes) {
        switch (bytes.length) {  //zone 1 4 5 要计算偏移量
            case 1:
                unichar=bytes[0];
                langId = 0;
                zoneId = 0;
                return ;
            case 2:
                bytes[0] &= 0x7f; //首位置1
                if ((bytes[0] >>> 5) == 0) {  //zone1   1 00 ~
                    zoneId = 1;
                    for (int i = 0; i < 7; i++) {  //可以考虑散列表
                        if (bytes[0] >= tool.dheads[i] && bytes[0] < tool.dtails[i]) {  //匹配成功
                            langId = tool.ilangs[i];
                            int symbleBlock = bytes[0] - tool.dheads[i];
                            symbleBlock <<= 7;//此为实值偏移量
                            symbleBlock += bytes[1];  //256 会占用同一位
                            unichar=symbleBlock;
                            return ;
                        }
                    }
                } else { //zone2   1 01+ ~
                    zoneId = 2;
                    unichar = asemble(bytes) - 4096;
                    langId = tool.langId(tool.idx2han[unichar]);
                    return ;
                }
                break;
            case 3: //zone3 其余三字节
                zoneId = 3;
                unichar = asemble(bytes);
                langId = tool.langId(unichar);
                return ;
            case 4:
                bytes[0] &= 0x7f;  //解压缩            //三字压缩形式
                if (bytes[0] < 0x40) { //1 0~
                    zoneId = 4; // 3+8+8+8
                    if(symbleIds==null)
                        symbleIds=  new ArrayList<>();
                    else
                        symbleIds.clear();;
                    int symbleBlock = asemble(bytes);
                    langId = symbleBlock >> 24;
                    langId = tool.ilangs[langId + 7];
                    int h = (symbleBlock >> 16) & 0xff;
                    if (h != 0)  //除非全0表示第一个，负责不会0首字
                        symbleIds.add(h);
                    h = (symbleBlock >> 8) & 0xff;
                    if (!symbleIds.isEmpty() || h != 0)
                        symbleIds.add(h);
                    h = symbleBlock & 0xff;
                    symbleIds.add(h);
                    return;
                } else if (bytes[0] < 0x60) { //两字压缩
                    zoneId = 5; //5+5+8+8
                    if(symbleIds==null)
                        symbleIds=  new ArrayList<>();
                    else
                        symbleIds.clear();
                    int symbleBlock = asemble(bytes);
                    langId = symbleBlock >> 16;
                    langId &= 0x03ff;  //10
                    int symble = (symbleBlock >> 8) & 0xff;
                    if (symble > 0)
                        symbleIds.add(symble);
                    symble = symbleBlock & 0xff;
                    symbleIds.add(symble);
                    return;
                } else {
                    zoneId = 6;
                    //转义解析
                }
                return;
        }
    }

    boolean compress(Duncode later) { //zone 4 5
        if (langId == later.langId) {
            if (((zoneId == 4) && (symbleIds.size() <= 2)) || ((zoneId == 5) && (symbleIds.size() <= 1))) {
                if (symbleIds.size() == 1 && symbleIds.get(0) == 0)
                    return false; //第一个字母独占 全为零表示  //应该改为很多不是256 ,应该改为特殊值表示
                symbleIds.add(later.symbleIds.get(0));
                return true;
            }
        }
        return false;
    }

    //当前Duncode输出char[]
    List<Character> getChars( List<Character> chars) {
        if(chars==null)
            chars = new ArrayList<>();
        else
            chars.clear();
        if (zoneId != 4 && zoneId != 5 ) { //非压缩
            chars.add(genChar(zoneId,langId,unichar));  //更省
            return chars;
        }
        for (int s : symbleIds) {
            chars.add(genChar(zoneId, langId, s));
        }
        return chars;
    }

    //生成单个unicchar
    char genChar(int zoneId, int langId, int symbleId) { //单个  1 4 5 需要
        switch (zoneId) {
            case 0:
                return (char) symbleId;
            case 1:
                return (char) (tool.heads[langId] + symbleId);
            case 2: //idx2han
                return tool.idx2han[symbleId];
            case 3:
                return (char) symbleId;
            case 4:
                return (char) (tool.heads[langId] + symbleId);
            case 5:
                return (char) (tool.heads[langId] + symbleId);
            default: //6 转义区
                return (char) symbleId;
        }
    }

    //取出DunByte[]形式
    byte[] getBytes() { //1 4 5 要压缩

        if(zoneId==4 || zoneId==5){
            unichar=0;
            for (int s : symbleIds) {
                unichar <<= 8;
                unichar |= s;
            }
        }
        return genBytes(zoneId, langId, unichar);
    }

    //生成dunbyte[]数组
    byte[] genBytes(int zoneId, int langId, int symbleId) {
        byte[] re;
        switch (zoneId) {
            case 0:
                return new byte[]{(byte) symbleId};
            case 1:
                for (int i = 0; i < 7; i++) {
                    if (tool.ilangs[i] == langId) {
                        int h = tool.dheads[i];
                        h <<= 7;
                        h += symbleId;
                        re=allocate(h,2);
                        return re;
                    }
                }
                return null;
            case 2://常用汉字
                char f = getChars(null).get(0);
                int h = tool.han2idx.get(f);
                int j = h + 0x1000; // 0x2000>>1
                re = allocate(j, 2);
                return re;
            case 3:
                re = allocate(symbleId, 3);
                return re;
            case 4:
                for (int i = 7; i <= 9; i++) {
                    if (tool.ilangs[i] == langId) {
                        h = (i - 7) << 24;
                        h += symbleId;
                        re = allocate(h, 4);
                        return re;
                    }
                }
                return null;
            case 5:
                h = (langId << 16);
                h += symbleId;
                re = allocate(h, 4);
                re[0] |= 0x40; //前缀 110
                return re;
            case 6:
                re = allocate(symbleId, 4);
                re[0] |= 0x60; //
                return re;
            default:
                return null;
        }
    }

    //将值分配给数个字节，不包含桩
    byte[] allocate(int n, int l) {
        int h = n & 0x7f; //末7位
//        int j = 16834 & 0x7f;
        byte a = (byte) h;
        if (l == 1)
            return new byte[]{a};

        h = (n >> 7) & 0x7f; // 至第三字节后七位
        h |= 0x80;
        byte b = (byte) h;
        if (l == 2)
            return new byte[]{b, a};

        h = (n >> 14) & 0x7f; // 至第二字节后七位
        h |= 0x80;
        byte c = (byte) h;
        if (l == 3)
            return new byte[]{c, b, a};

        h = (n >> 21) & 0x7f; // 至第一字节
        h |= 0x80;
        byte d = (byte) h;
        if (l == 4)
            return new byte[]{d, c, b, a};
        return null;
    }

    static int asemble(byte[] bytes) {  //只管集合，不管桩位 根据dunbyte[]取值
        switch (bytes.length) {
            case 1:
                return bytes[0];
            case 2:
                return (bytes[0] & 0x7f) << 7 | bytes[1];
            case 3:
                return (bytes[0] & 0x7f) << 14 | (bytes[1] & 0x7f) << 7 | bytes[2];
            case 4:
                return (bytes[0] & 0x7f) << 21 | (bytes[1] & 0x7f) << 14 | (bytes[2] & 0x7f) << 7 | bytes[3];
        }
        return -1;
    }

    void printDun() {
        String symble = "\t symbles:";
        String chars = "\t chars:";
        for (int i = 0; i < symbleIds.size(); i++) {
//            System.out.println(zoneId + "\t" + langId + "\t " + i+symbleIds.get(i) + "\t" + getChars().get(i)+"\t"+symbleIds.size()+ Arrays.toString(getBytes()));
            symble += symbleIds.get(i) + " ";
            chars = getChars(null).get(i) + " ";
        }
        String s = zoneId + "\t" + langId + "\t " + symble + "\t" + chars + "\t" + symbleIds.size() + "\t" + Arrays.toString(getBytes());
        System.out.println(s);
    }
}
