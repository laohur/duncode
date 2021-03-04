#include <iostream>
#include <fstream>
#include <string>
#include <fstream>
#include <iostream>
#include <locale>
#include <codecvt>
#include <iosfwd>
#include <iomanip>
#include <unordered_map>
#include <cuchar>
#include <ctime>
#include "Duncoder.cpp"

using namespace std;

#define _CRT_SECURE_NO_WARNINGS


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