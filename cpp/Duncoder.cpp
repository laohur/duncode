#ifndef DUNCODE_LIBRARY_H
#define DUNCODE_LIBRARY_H
#include <stdio.h>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <iterator>
#include <dirent.h>
#include <set>
#include <unistd.h>
#include <string.h>
#include "Duncode.h"
#include "Duncode.cpp"

using namespace std;

class Duncoder{
public:
    Tool* tool;
    Duncoder(Tool* tool){
        this->tool=tool;
    }
    ~Duncoder(){
        printf("~Duncoder ");
        delete tool;
    }
    vector<unsigned char> unicodesToBytes(vector<char32_t> unicodes){
//        Duncode[] duncodes=charsToDuncodes(chars);
//        Byte[] bytes=duncodesToBytes(duncodes);
//        return bytes;
        return duncodesToBytes(unicodesToDuncodes(unicodes));
    }
    vector<Duncode*> unicodesToDuncodes(vector<char32_t> unicodes){
        vector<Duncode*> duncodes;
        Duncode* lastDun=new Duncode(tool);
        lastDun->readUnicode(unicodes[0]);
        duncodes.push_back(lastDun);

        for(int i=1;i<unicodes.size();i++){
            char32_t c=unicodes[i];
            Duncode* nowDun = new Duncode(tool);
            nowDun->readUnicode(c);
            if (lastDun->compress(nowDun)) {
                delete nowDun;
                continue;
            }
            else{
                duncodes.push_back(nowDun);
                lastDun = nowDun;
            }
        }
        vector<char32_t>().swap(unicodes);
        return  duncodes;
    }
    vector<unsigned char> duncodesToBytes(vector<Duncode*> duncodes){
        vector<unsigned char> bytes;
        vector<unsigned char> bs;
        for(Duncode* d:duncodes){
             bs=d->getBytes();
            bytes.insert(bytes.end(),bs.begin(),bs.end());
            bs.clear();
        }
        vector<unsigned char>().swap(bs);
        for(vector<Duncode*>::iterator it=duncodes.begin();it!=duncodes.end();it++){
            if(NULL!=*it){
                delete *it;
                *it=NULL;
            }
        }
//        auto f=duncodes[0];
//        int g=f->unichar;
//        int h=sizeof(*duncodes[0]);
//        vector<Duncode*>().swap(duncodes);

        return  bytes;
    }
    vector<char32_t > bytesToUnicodes(vector<unsigned char> bytes){
        vector<char32_t > unicodes;
        vector<unsigned char> bucket;
        Duncode* nowDun = new Duncode(tool);
        for (unsigned char b : bytes) {
            bucket.push_back(b);
            if (b < 0x80) { //末尾字节
                nowDun->readDunbytes(bucket);
                vector<char32_t > us=nowDun->getUnicodes();
                unicodes.insert(unicodes.end(),us.begin(),us.end());
                bucket={};

            }
        }
        delete nowDun;
        nowDun=NULL;
        vector<unsigned char>().swap(bucket);
        vector<unsigned char>().swap(bytes);
        return unicodes;
    }

};
int utf8FileDun(string inpath,string outpath) {
    clock_t t0 = clock();
    printf("\n\n  ---load %s  \n", inpath.c_str());
    ifstream fin(inpath,ios::binary);
    if (fin.fail()) {
        printf(" read utf8file %s fail \n", inpath.c_str());
        perror(" error! %s  \n");
    } else {
        printf("read utf8file %s ok \n", inpath.c_str());
    }
    ofstream dout(inpath+".dun",ios::binary);
//    dout.rdbuf()->pubsetbuf(&buff2.front(),buff2.size());
    Tool *tool = new Tool();
    Duncoder *duncoder = new Duncoder(tool);
    string line;
    vector<char32_t> unicodes;
    int nline=0;
    while(fin && fin.peek()!=EOF) {
        getline(fin, line);
        line+="\n";
        nline++;

        tool->utf8lineToUnicoides(line, &unicodes);
//        line2 = tool->UnicodesToUtf8line(unicodes);

        vector<unsigned char> bytes = duncoder->unicodesToBytes(unicodes);
//        vector<char32_t> unicodes2 = duncoder->bytesToUnicodes(bytes);
//        line3 = tool->UnicodesToUtf8line(unicodes2);
//        printf("%s \n%s \n%s \n", line.c_str(), line2.c_str(), line3.c_str());
        dout.write(reinterpret_cast<const char *>(bytes.data()), bytes.size());
//        dout<<bytes;

//        if(nline%10000==0)
//            printf("%d %d %d  \n",nline,line.length(),bytes.size());
//        for(vecort<unsigned char>)

//        if(nline%100==0)
//            printf(" utf8FileDun   %d %d \n",sizeof(unicodes),sizeof(bytes));
        vector<char32_t >() .swap(unicodes);
        vector<unsigned char>() .swap(bytes);
        string().swap(line);

//        unicodes={};
//        bytes={};
//        line="";

    }
    delete duncoder;
    fin.close();
    dout.close();
    printf(" write utf8 to %s",outpath.c_str());
    clock_t t1 = clock();
    printf(" utf8FileDun耗时总计 %f \n", (t1 - t0)/1000.0);
    return 0;
}
int dunFileUtf8(string inpath,string outpath) {
    clock_t t0 = clock();
    printf("\n\n  ---load %s  \n", inpath.c_str());
    ifstream fin(inpath,ios::binary);
    ofstream u8out(outpath,ios::binary);
    if (fin.fail()) {
        printf(" read dunfile %s fail \n", inpath.c_str());
        perror(" error!  \n");
    } else {
        printf("read dunfile %s ok \n", inpath.c_str());
    }
    Tool *tool = new Tool();
    Duncoder *duncoder = new Duncoder(tool);
    string line,line2;
    int nline=0;
    while(fin && fin.peek()!=EOF) {
        getline(fin, line);
        line+="\n";
        nline++;
//        tool->utf8lineToUnicoides(line, &unicodes);
//        line2 = tool->UnicodesToUtf8line(unicodes);
//        vector<unsigned char> bytes = duncoder->unicodesToBytes(unicodes);
        std::vector<unsigned char> bytes(line.begin(), line.end());
        vector<char32_t> unicodes = duncoder->bytesToUnicodes(bytes);
        line2=tool->UnicodesToUtf8line(unicodes);
//        printf("%s \n%s \n%s \n", line.c_str(), line2.c_str(), line3.c_str());
//        if(nline%100000==0)
//            printf(" dunFileUtf8 %d  %d %d %d\n",nline,unicodes.size(),bytes.size(),line2.length());
//        u8out.write(line2.c_str(), line2.size());
        u8out<<line2;
        std::vector<unsigned char>().swap(bytes);
        vector<char32_t>().swap(unicodes);
//        bytes={};
//        unicodes={};
        line2="";
        line="";
//        bytes={};
//        unicodes={};
//        line2={};
    }
    fin.close();
    u8out.close();
    printf(" write utf8 to %s",outpath.c_str());
    clock_t t1 = clock();
    printf(" dunFileUtf8耗时总计 %f \n", (t1 - t0)/1000.0);
    return 0;
}


int testDir(string dirpath="D:/doc/wikitext_cpp"){
    std::vector<std::string> utf8Files;
    std::vector<std::string> dunFiles;
    DIR *dir;
    struct dirent *ent;
    if ((dir = opendir (dirpath.c_str())) != NULL) {
        /* print all the files and directories within directory */
        while ((ent = readdir (dir)) != NULL) {
            string inpath=dirpath+"/"+ent->d_name;
            if(inpath.length()<=4)
                continue;
            if(inpath[inpath.length()-1]=='.'){
                printf("-- bad inpath %s %c \n",inpath.c_str(),inpath[inpath.length()-1]);
                continue;
            }
            string sufix=inpath.substr(inpath.size()-4);
            if(strcmp(sufix.c_str(),".txt")==0){
                utf8Files.push_back(inpath);
            }else  if(strcmp(sufix.c_str(),".dun")==0) {
                dunFiles.push_back(inpath);
            }
            }
        closedir (dir);
    } else {
        /* could not open directory */
        perror ("");
        return EXIT_FAILURE;
    }

    printf ("---try utf8FileDuns \n ");
    for(string inpath:utf8Files){
        printf ("---try utf8FileDun %s\n ", inpath.c_str());

        try {
//            utf8FileDun(inpath,inpath+".dun");
        }catch (string msg){
            cerr<<msg;
            continue;
        }
    }
    sleep(1);

    printf ("---try dunFileUtf8s \n ");
    for(string inpath:dunFiles){
        printf ("---try dunFileUtf8 %s\n ", inpath.c_str());

        try {
            dunFileUtf8(inpath,inpath+".utf8");
        }catch (string msg){
            cerr<<msg;
            continue;
        }
    }
    return 0;
}

int testpair(){
    clock_t t0 = clock();
    string dirpath="D:/doc/wikitext_cpp";
    string inpath=dirpath+"/betawikiversity.txt";
//    inpath="../data/test1.txt";
    utf8FileDun(inpath,inpath+".dun");
    sleep(1);
    dunFileUtf8(inpath+".dun",inpath+".dun.utf8");
    clock_t t1 = clock();
    printf(" 耗时总计 %f ", (t1 - t0)/1000.0);
    return 0;
}
int testDuncoder(){
    clock_t t0 = clock();
    string dirpath="D:/doc/wikitext_cpp";
//Tool* tool=new Tool();
//int t=sizeof(*tool);
//Duncode* d=new Duncode(tool);
//int a=sizeof(*d);
//delete d;
//d=NULL;   //*D=
//int b=sizeof(*d);
//int c=d->unichar;

    sleep(3);
//    testpair();
    testDir();
    clock_t t1 = clock();
    printf(" 耗时总计  %.2f s", (t1 - t0)/1000.00);
    return 0;
}
int main0(int argc, char** argv){
//    string inpath0="../data/test1.txt";
//    utf8FileDun(inpath0, inpath0+".dun");
//    dunFileUtf8(inpath0+".dun",inpath0+".utf8");

    // u/d inpath outpath
    if(argc<2 || argc>4 ){
        printf("usage :\n  1[u/d]:source file type, utf8 or duncode \n 2[inpath]  3[outpath]");
        return -1;}
    char type=*argv[1];

    string inpath=argv[2];
    string outpath;
    if(argc==4) {
         outpath = argv[3];
    }else if(type=='d' || type=='D'){
        outpath=inpath+".utf8";
    }else{
        outpath=inpath+".dun";
    }
    if(type=='d' || type=='D' ){
        dunFileUtf8(inpath,outpath);
    }    else{
        utf8FileDun(inpath,outpath);
    }

    return 0;
}

static Tool* toolInstance = new Tool();
static Duncoder* duncoderInstance = new Duncoder(toolInstance);

extern "C" {
    int unicodesToDunbytes(char32_t* unicodes,int unicodeCount, unsigned char* bytes){
        vector<char32_t> us;
        for(int i=0;i<unicodeCount;i++)
            us.push_back(unicodes[i]);
        vector<unsigned char> re= duncoderInstance->unicodesToBytes(us);
        bytes= re.data();
        vector<char32_t> empty;
        us.swap(empty);
        return 0;
    }
    int DunbytesToUnicodes(unsigned char* bytes,int byteCount, char32_t* unicodes){
        vector<unsigned char> bs;
        for(int i=0;i<byteCount;i++)
            bs.push_back(unicodes[i]);
        vector<char32_t > re= duncoderInstance->bytesToUnicodes(bs);
        unicodes= re.data();
        vector<unsigned char> empty;
        bs.swap(empty);
        return 0;
    }
//    vector<unsigned char> unicodesToBytes(vector<char32_t> unicodes)
//    vector<char32_t> bytesToUnicodes(vector<unsigned char> bytes)
    }
#endif //DUNCODE_LIBRARY_H

int main(int argc, char** argv){
//    string inpath0="../data/test1.txt";
//    utf8FileDun(inpath0, inpath0+".dun");
//    dunFileUtf8(inpath0+".dun",inpath0+".utf8");

    // u/d inpath outpath
    if(argc<2 || argc>4 ){
        printf("usage :\n  1[u/d]:source file type, utf8 or duncode \n 2[inpath]  3[outpath]");
        return -1;}
    char type=*argv[1];

    string inpath=argv[2];
    string outpath;
    if(argc==4) {
        outpath = argv[3];
    }else if(type=='d' || type=='D'){
        outpath=inpath+".utf8";
    }else{
        outpath=inpath+".dun";
    }
    if(type=='d' || type=='D' ){
        dunFileUtf8(inpath,outpath);
    }    else{
        utf8FileDun(inpath,outpath);
    }

    return 0;
}