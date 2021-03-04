#include <stdio.h>
//#include "Duncode.h"
#include "Tool.cpp"

using namespace std;
//union Symble{
//    int unichar;
//    char symbles[4];
//};
//
//struct Dun{
//    unsigned __int8  count,zoneId;
//    unsigned __int16 langId;
//    union  Symble symbleBlock;
//};

class Duncode{
public:
    int zoneId=-1;
    int langId=-1;
//    char count;
    int unichar=-1;
    vector<unsigned int> symbols;
    Tool* tool;
    Duncode(Tool* tool){
        this->tool=tool;
    }
    ~Duncode(){
//        printf(" ~Duncode ");
//        delete tool;
//        delete symbols;
//        symbols.clear();
        vector<unsigned int>().swap(symbols);
    }
    int readUnicode(char32_t u){
        langId=tool->langId(u);
        zoneId=tool->zoneId(langId,u);
        switch(zoneId){
            case 0:
                unichar=u;
                break;
            case 1:
                unichar=u-tool->heads[langId];
                break;
            case 2:
                unichar=tool->han2idx[u];
                break;
            case 3:
                unichar=u;
                break;
            case 4:
                symbols={};
                symbols.push_back(u-tool->heads[langId]);
                break;
            case 5:
                symbols={};
                symbols.push_back(u-tool->heads[langId]);
                break;
            case 6:
                unichar=u;
                break;
            default:
                unichar=u;
        }
        return 0;
    }
    void readDunbytes(vector<unsigned char> bytes) {
        switch (bytes.size()) {  //zone 1 4 5 要计算偏移量
            case 1:
                unichar=bytes[0];
                langId = 0;
                zoneId = 0;
                break;
            case 2:
                bytes[0] &= 0x7f; //首位置1
                if ((bytes[0] >> 5) == 0) {  //zone1   1 00 ~
                    zoneId = 1;
                    for (int i = 0; i < 7; i++) {
                        if (bytes[0] >= tool->dheads[i] && bytes[0] < tool->dtails[i]) {  //匹配成功
                            langId = tool->ilangs[i];
                            int symbleBlock = bytes[0] - tool->dheads[i];
                            symbleBlock <<= 7;//此为实值偏移量
                            symbleBlock += bytes[1];  //256 会占用同一位
                            unichar=symbleBlock;
                        }
                    }
                } else { //zone2   1 01+ ~
                    zoneId = 2;
                    unichar = asemble(bytes) - 4096;
                    langId = tool->langId(tool->idx2han[unichar]);
                }
                break;
            case 3: //zone3 其余汉字 三字节   //rare symbols
                zoneId = 3;
                unichar = asemble(bytes);
                langId = tool->langId(unichar);
                break;
            case 4:
                bytes[0] &= 0x7f;  //解压缩            //三字压缩形式
                if (bytes[0] < 0x40) { //1 0~
                    zoneId = 4; // 3+8+8+8
                    symbols={};;
                    int symbleBlock = asemble(bytes);
                    langId = symbleBlock >> 24;
                    langId = tool->ilangs[langId + 7];
                    int h = (symbleBlock >> 16) & 0xff;
                    if (h != 0)  //除非全0表示第一个，负责不会0首字
                        symbols.push_back(h);
                    h = (symbleBlock >> 8) & 0xff;
                    if (!symbols.empty() || h != 0)
                        symbols.push_back(h);
                    h = symbleBlock & 0xff;
                    symbols.push_back(h);
                    break;
                } else if (bytes[0] < 0x60) { //两字压缩
                    zoneId = 5; //5+5+8+8
                    symbols={};;
                    int symbleBlock = asemble(bytes);
                    langId = symbleBlock >> 16;
                    langId &= 0x03ff;  //10
                    int symble = (symbleBlock >> 8) & 0xff;
                    if (symble > 0)
                        symbols.push_back(symble);
                    symble = symbleBlock & 0xff;
                    symbols.push_back(symble);
                    break;
                } else if (bytes[0] < 0x68) {
                    zoneId = 6; //罕见字 11100xxx  -->zone byte3
                    bytes[0] &= 0x07;
                    int symbleBlock = asemble(bytes);
                    unichar =  symbleBlock;
                    langId = tool->langId(symbleBlock);
                    break;
                } else if (bytes[0] < 0x70) {
                    zoneId = 7; //新文字
                    bytes[0] &= 0x07;
                    int symbleBlock = asemble(bytes);
                    unichar = symbleBlock;
                    langId = tool->langId(symbleBlock);
                    break;
                } else if (bytes[0] < 0x78) {
                    zoneId = 8;//临时区
                    bytes[0] &= 0x07;
                    int symbleBlock = asemble(bytes);
                    unichar = symbleBlock;
                    langId = tool->langId(symbleBlock);
                    break;
                } else {
                    zoneId = 9;
                }

        }
    }

    bool compress(Duncode* later){
        if(langId==later->langId){
            if( (zoneId==4)&& symbols.size()<=2  || (zoneId==5)&&symbols.size()<=1 ){
                if(symbols.size()==1 && symbols[0]==0)
                    return false;
                symbols.push_back(later->symbols[0]);
//                delete later;
                return true;
            }
        }
        return false;
    }
//    void readDunbytes(char[] bytes)
    char32_t genUnicode(int zoneId,int langId, int symbol ){
        switch (zoneId) {
            case 0:
                return symbol;
            case 1:
                return (tool->heads[langId] + symbol);
            case 2: //idx2han
                return tool->idx2han[symbol];
            case 3:
                return symbol;
            case 4:
                return (tool->heads[langId] + symbol);
            case 5:
                return  (tool->heads[langId] + symbol);
            case 9://私用区
                return '!';
            default: //6 7 8
                return  symbol;
        }
    }
    vector<char32_t>  getUnicodes(){
        vector<char32_t> unicodes;
        if(zoneId!=4 && zoneId!=5){
            char32_t u=genUnicode(zoneId,langId,unichar);
            unicodes.push_back(u);
            return unicodes;
        }
        for(int s:symbols){
            unicodes.push_back(genUnicode(zoneId,langId,s));
        }
        return unicodes;
    }

    vector<unsigned  char> getBytes(){
        if(zoneId==4 || zoneId==5){
            unichar=0;
            for (int s : symbols) {
                unichar <<= 8;
                unichar |= s;
            }
        }
        return genBytes(zoneId, langId, unichar);
    }

    vector<unsigned  char> genBytes(int zoneId,int langId,int symbol){
        vector<unsigned  char> re;
        switch (zoneId) {
            case 0:
                return {(unsigned char)symbol};
            case 1:
                for (int i = 0; i < 7; i++) {
                    if (tool->ilangs[i] == langId) {
                        int h = tool->dheads[i];
                        h <<= 7;
                        h += symbol;
                        re=allocate(h,2);
                        return re;
                    }
                }
                return {};
            case 2:{
                char32_t f = getUnicodes()[0];
                int h = tool->han2idx[f];
                int j = h + 0x1000; // 0x2000>>1
                re = allocate(j, 2);
                return re;
            }
            case 3:
                re = allocate(symbol, 3);
                return re;
            case 4: {
                for (int i = 7; i <= 9; i++) {
                    if (tool->ilangs[i] == langId) {
                        char32_t h = (i - 7) << 24;
                        h += symbol;
                        re = allocate(h, 4);
                        return re;
                    }
                }
                return {};
            }
            case 5: {
                char32_t h = (langId << 16);
                h += symbol;
                re = allocate(h, 4);
                re[0] |= 0x40; //前缀 110
                return re;
            }
            case 6:
                re = allocate(symbol, 4);
                re[0] |= 0x60; //
                return re;
            case 7:
                re = allocate(symbol, 4);
                re[0] |= 0x68;
                return re;
            case 8:
                re = allocate(symbol, 4);
                re[0] |= 0x70;
                return re;
            case 9:
                re = allocate(symbol, 4);
                re[0] |= 0x78;
                return re;
            default:
                return {};
        }
    }

    //将值分配给数个字节，不包含桩
    vector<unsigned char> allocate(int n, int l) {
        vector<unsigned char> re;
        int h = n & 0x7f; //末7位
//        int j = 16834 & 0x7f;
        unsigned char a =  h;
        if (l == 1){
            re={a};
            return re;
        }

        h = (n >> 7) & 0x7f; // 至第三字节后七位
        h |= 0x80;
        unsigned char b =  h;
        if (l == 2) {
            re={b,a};
            return re;
        }
        h = (n >> 14) & 0x7f; // 至第二字节后七位
        h |= 0x80;
        unsigned char c = h;
        if (l == 3){
            re={c,b,a};
            return re;
        }

        h = (n >> 21) & 0x7f; // 至第一字节
        h |= 0x80;
        unsigned char d = h;
        if (l == 4){
            re={d,c,b,a};
            return re;
        }
        return {};
    }

    int asemble(vector<unsigned char> bytes) {  //只管集合，不管桩位 根据dunbyte[]取值
        switch (bytes.size()) {
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
//    toString()

};
//
//int testDuncode() {
//    string b = "哼aâu Tiễn (chữ Hán: 越王勾踐; trị vì 496 T";
//    string c = u8"哼âu Tiễn (chữ Hán: 越王勾踐; trị vì 496 T";
////        string inpath="../data/betawikiversity.txt";
//    string inpath = "../data/test1.txt";
//    printf("\n\n  ---load %s  \n", inpath.c_str());
//    ifstream fin;
//    fin.open(inpath);
//    if (fin.fail()) {
//        printf(" read %s fail \n", inpath.c_str());
//        perror(" error! %s  \n");
//    } else {
//        printf("read %s ok \n", inpath.c_str());
//    }
//    string line;
//    getline(fin, line);
//    vector<char32_t> unicodes;
//    vector<char> bytes;
//    Tool* tool=new Tool();
//    Duncode* d=new Duncode(tool);
//    tool->utf8lineToUnicoides(line, &unicodes);
//    string line2, line3;
//    Duncode* e=new Duncode(tool);
//    vector<char32_t> re;
//    for (char32_t u:unicodes) {
//        string s = tool->to_utf8(u);
//        int langid = tool->langId(u);
//        bool ishan = tool->isHan(langid);
//        int zoneid = tool->zoneId(langid, u);
////        printf(" %d \t %s \t %d \t %d \t %d \n", u, s.c_str(), langid, ishan, zoneid);
//        d->readUnicode(u);
//        char32_t  v=d->getUnicodes()[0];
//        s=tool->to_utf8(v);
//        line2+=s;
//
//        vector<unsigned char> bytes=d->getBytes();
//        e->readDunbytes(bytes);
//        v=v=d->getUnicodes()[0];
//        s=tool->to_utf8(v);
//        line3+=s;
//
////        printf(" %d \t %s \t %d \t %d \t %d \n", u, s.c_str(), langid, ishan, zoneid);
//        printf("%s \n%s \n%s \n",line.c_str(),line2.c_str(),line3.c_str());
//
//    }
//    return 0;
//}

//int main(){
//    clock_t t0 = clock();
//    clock_t t1 = clock();
//    testDuncode();
//    printf(" 耗时总计 %f ", (t1 - t0));
//    return 0;
//}