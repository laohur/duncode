extern "C" {
    int unicodesToDunbytes(char32_t* ,int , unsigned char* );
    int DunbytesToUnicodes(unsigned char* ,int , char32_t* );
//int unicodesToDunbytes(char32_t* unicodes,int unicodeCount, unsigned char* bytes);
//int DunbytesToUnicodes(unsigned char* bytes,int byteCount, char32_t* unicodes);
//vector<unsigned char> unicodesToBytes(vector<char32_t> unicodes)
//vector<char32_t> bytesToUnicodes(vector<unsigned char> bytes)
}
