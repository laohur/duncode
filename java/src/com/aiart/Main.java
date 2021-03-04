package com.aiart;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class Main {
    public static void main0(String[] args) {
//        testDir("D:/wiki/20190801/txt");
//        testLines();
        String dir="D:/wiki/20190801/";
        String[] files={"arwiki","arwikisource","betawikiversity","enwikisource","frwikisource","jawiki","jawikisource","ruwiki","ruwikisource","wikisource","zhwiki","zhwikisource"};
//        for(String file:files) {
//            testFile("u",dir+"txt/"+file+".txt",dir+"dun/"+file+".dun");
////            testFile("d",dir+"dun/"+file+".dun",dir+"utf8/"+file+".utf8");
//        }
//         testPair("D:/wiki/20190801/txt/jawikisource.txt");
        testPair("data/betawikiversity.txt");
//        testFile("d","D:/wiki/20190801/txt/jawiki.txt.dun","D:/wiki/20190801/txt/jawiki.txt.dun.utf8");
//        testClear();
    }

    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();   //获取开始时间
        if(args.length<2){
            System.out.println("usage :\n  1[u/d]:source file type, utf8 or duncode \n 2[inpath]  3[outpath]\n ");
        }
        String type=args[0];
        String inpath=args[1];
        String outpath=args.length==2?null:args[2];
        testFile(type,inpath,outpath);
//        testFile("u","D:/doc/wiki/wikitext_java/zhwiki.txt",null);
//        testFile("u","D:/doc/wiki/wikitext_java/zhwikisource.txt",null);
//        testLine();
//        testDir("D:/doc/wikitext");
//        testLine();
        long endTime = System.currentTimeMillis(); //获取结束时间
        System.out.println("程序运行时间： " + (endTime - startTime)/1000.0 + "s");
    }

    static void testClear() {
        String s = "1075694eiatnsorluhdcmgk.оpанv,итеіbрвاс40867yfкwлjلд/1му:0пTم-و=я2نàзđäر3ưьõгه9)(يتبáчدع8Sбف<ی45Lх7>6CüцйPNxIїוqיếAHMạжôệسêהюVDâEóộz?حKBơقốਕớجềأờєਾфة،ਿלấكתậшרĐـủảב'ợíOịWشאầì_öصמשăਸFਰểВСUщزRựGਨПکਮੀởىứضਂخנử–إổ;ụãУọਆםიაІéùữКדú\u200CקýМطذОНਵכעắੇਤòъừQਦJਟפეੁồਜਹਲثРੈАחצXДЛ’ũگსТסਗặੱ“Б”YੋءਅîГЗĩằਪრגè*پੰ«ਣੂآן»ỉŭטმغ{+}َო][ბФХỏდ!—਼ئ%ظლĝıễਇыזਭĉნვИỗਡუşਬთÜẫЕẩ다Цچçੌਚਖਫ•ُ ẳЯგЧਧךZਥẻųė老師ẽّəÔკШẹტშţ੍ًٔ|ؤِ&āხš„Õỳ@ЄЮਛ്।ْწžਈ$าĵЙ하ỷਉਊცąαზ№ÂਓÁłïץ…ūนŝਢპף#ਏ¬，Äīτქਐč사εรęνι؟ỡژåûỹფЖ어οա이ിỵ\u200Eਯกσ李ēക้`ยง는ǎਔ‘Ĉყ자ґ니ρјẤ에į지ാżੜэ중의ηնനǐ่〉ğั×İπφวอีเ。£ม·→ുิძ～ŋωത가ő국ե제기Щё나ღςր강ÉλทỞര대서（）ίവ\uF05Aบțκμล로리Þ۲ห陳Ðñਝ고를ěո☞ٍയ~χЬด한γ조ÎōυЇ不라?ëപคส人王\u200DĜիไ娘ÇÍńśșψ요ਘ간黃시오ല姑帝۰۱ÀÓ여ٌ을া그ใ儀문수ːθ蔡아明음њള之天子부ỨƯόЭ있ข以^∨्ปჩ해´Ýǒ∧습ւְ은것แ단동\\³́έ용인校체െ道ćմڪाംട배Şǔή위िษู皇�δ敏؛क্武과主도大마만學而文전റි들°Ö教ּ우۵்മസต명?ẵð√կսք۳으일র་長國보ĥάָ정词ภะ也仁任내학미번‚լ۴۹입িව院히婉德Ò선க게ങേพ可小ßℏ慧ў장梁ব구上ึุ公래夫年\u00AD¿Ûώ٫朝ਠധ一三ำ치、的基家병者\u2028€성ļű소յִ۶衍मय개ദ글余信濤ჭ터其吳화知好穎무æ세Ґֶֹ日有ेকโ濟太바维；怡¥생成−알գվַ是नে氏趙්จ็漢〈ჯჰー늘합陽르高宗방²상ǃ실տ적पয진거ഷ금ณื于佩の劉두표和神女学발鳳분心\u200BÈŽɪ英안와ە林주言ണถ中ས無醫金韓祖婷字寶山常麗黎Ỷ徵±÷식ξђљ쓰դծհ양ֵ영曾작行रीন語ഗഭ七ซ何鄭할면물ẢỔ聖¡øČɛ신ύћ씨政ֲٽ۸۽저ुল각공越මස下ธ个ོ타然톤則副당더后生때형国름말宇?州비齊美¹řŠů芝심ց없원वह감차水ചයු侯키イ邱털년先前달十南玲포君등甜嘉雅在람头秀모孝본Ế徐志思†職Ú胡Ơ所ΑΓΨΩ蓮էխ蕭언方연星智움ۆ۷유杜栩तस楊এশ欣மి초교ญฐ五達て날ィ통勤錦瑩璐員關현雪順며安定少반ṭ！張義必ÊĂ손手ɑ≠范ү앞զը약晉운月机재종ंखদপম民처계ක์你车通いすな傑낮ジ태郭ン都儿关데편理된唐따电雄러科머官經ṇ育þŜ我舒戴ʊЪѝՎ՛╞ևׁ역올١景워۾本杨점लशष楚要টত়權譚று같ీ건천貞경法ഒ洛ജഴോනප为亮代གད体潘使남為탄プ鄧노卿原叫球필행四真地랑页夏如馬禮많姚매秦竇魏?紹불麥¢설Ś속憲自與興∗ɾ스̄莊Ββ董բ敬数պ蘭ٿ暖울ڻ육来읽果접表एठणदংউ짜正டபకవ貝결貴곧അഎശഹീ추දරศ么二軍仲伍མིེྡ農这進ま邑をキシテトバ郝ロ너네照共里重父片닫답勿卓玄只各合名听周디用商떻호白盈相뜻非面矣音᠋려료루外多奇禄妙程马笔第孫宮寧对封將展복鴻ṣ度得耀怨耳恭聰悅\u0096§©산ĮŌ셨慶ź慾⇒∙戰良ʂʰ芳芷ʼ싸̌ΔΧζ葉않թռ新ֱ▲藏旵蘅昇蘇예易春時晟虹ڊ욱ہێ曦服잡柱柳절梅裘줄आडृै西見ধস১樂让லவி책汉్象겠관ഇഞ洲ൂൊ浩海起권ටල귀ැ淑ී丁ชฒ且世ฤ温길ๆ까九书침亂了身从令们仮滿会輝潔འ潤澄ྣ係过近瀞\uF02D過こし遠とに火はる邢ウェク邵点탁郡・元兒煒녕六凱列初到動勞파鍾珊珍同되琬琳瑛田由피핑함開嗣또활百監盧雯雷眉霞靜瞿᠀렇레ᠠ력题碧奉奕른馮娟婆穗몇목竹符孙筠宁守完富민籍紀봅綺緯빠府经Ố廖羅龍快怕性쁘聯您살想意感ķœ慢Ŭ↔ǚ숙房打쉽ʃ报草ΜΣΥΦ放Մձղ整չջ야փ斤열时昭虎٪웬曹會期木잘衣उगजधफो極ই즉ডণষ직집৭詞謀欢欧議死段ை母ం갑갔气ట谢దనమాేొ貳賀ആഈ洗活හ跟清급与东游乐친习깨事什車介他커伯ང位轩ུ辕\uF03Dで크側備난那邦\uF0A7낸タ탕ペメラ焉光克煥八높农燕出누分利爱加능动특化千危去돌发叔瑗呢钟드甌错长門善喺门阮환陵陶登發횟후回因目휴眼희락략량韋럼研堯報렵?록飞声류票祿妃妮림막秋始姐種究媒媧못立章孔客容받尺버벳변?終組工己已ṛ빨平黄Ộ弘弟当忠怀怎\u200F″恒恥息聽\u0080肃\u0093º삼Å색惠能愛석ĭ愿셀ņ셈慎⅛憂Ə송應Ǎǜ臣臨懷或战순舜舞숫户才ɔ쉬色承把护슨苦십拾指쌍써쎄ΆΉΕΘΚΝΞΠ接ќ蒙操ԱԳ收攸改Տճ얕억얻얼엄族׳״엽였옷왔晚왜虞٢٣晴ٻ虽ڌڏړھ更望朵응杰東익잃임잠柴样根桀桓존袁좋\u089E被준च褚थগজহী见질览觉짝째訟許誠說쪽調論諸謂譯讓த认யளழ殷许ா毅每比毛话诞该课갈갖గజడ착ప찮య참స江汽척겨견청河油治波ഉ촌ഥബ贺派资\u0D45赖赤ൾ궁超ඉ축趕출충ළා跑深跳淵근万踐专丘측긴층乎칭카칸乾躂源亲亶今께件།伏休輪꽃但低코པ佗ཙ作软轻辛便俊迎运还候倚速逵遂く끓さ做ただっば灵べ끼りれ炎ガ킬\uF0B8냄ドベ택\uF0DE냥ル儂儒煚入兪兰兵농놓写토决量刚別别투爸刻版剛牛特느늦助튼님独티勾包팔北猶匹午华協卜厅玉펜参평及돕史現吃向吧됩吾呀告둘풀품풍뒤哎哗哥哪프甘唯딱啊畫異했喜떠향病험间嘅阳혹随際雒盛뛰雲直零固看圣뜨眾着靈흐睡란垂鞅埧石럽頁ᠢᠭ顏顯项ᢀᢁᢇ风食夜夠祚奪奭릅妈妹离秉망맞駱騎먹먼穀멈空穿骂媚站嫩鬼策孟季简算宣뮤箱節範寬寸밖尚簡尼簾居?类鱼벌범粵별紧북左巧巾希鸟縣ḥ織干并鹿續应麟庸结给Ử开张缪뽑彝形齿很從御群뿐習翰翻忽考‐‒耕急怪 总恐聞聪⁴肆肉悉\u008F肚\u009F悦ª®肯¶삶悺½情새ËÌ惑Ù惟惡Āď脑ĒĕĚ脚愚Ğ™ģĤĪ脱脸ľŎŢťŪ慮腿膏←솔膽懂Ǐ臏ǑǓǕǖǗǘǙǛ臥ǹ臻懼∅∆∈舉舌숨∩航般戶船扇≈扈ɕɡ扫折花슷승抽ˈ拉拍拜若拨싶苹拿̀挂按挑⌒̝̠̕茶̹挺挾捌荏药掉Ό排ΙΟΤ推莽菁提揚菜握菩華援Ј搜搞营搬萬ѓџ蒜播蒯Ө쓴쓸┴씻Հ故Ն蕉앉Ս救敗՝애敢ժշ薛□薪斬於施־旅업엇었엌藍既藥旧早昆옛昞昨옮昼완♈晋왕處晝٠٧٨٩普외ٺ왼ڍ暑욕ڙڳ웃曆曇蛇曉蛋曌월書替最蜀朋✓末朱杂条杲极잉잊잖枝柒蟜査젊树젓젼格血桂桌桑術衡补족졸좀袋\u088C袜梨梯梵棄棒죠裤裴죽अ椅इईऔ褘ञट襄쥐業概আ視榜증ঢ覧ফভ榮觀ূ视構ো짐様解০짧৪৫৬৯訁樣訥訳橋쩌機話誓쫒諗諮次識此譬護讀殆இ變讠讨ர记讲殿识ே诉译毕语说请读谁调అ찌찍ఒతఫభరష창ు채豊ౌ汗豫汴걷걸決검겁겉沙沛철没격첫겸겹賁곤注泮곱泳ഏഖ괜贝총洞ഠഡ贰괴贵流浅ൈ赋赛赢走굴굽涂涇趟තධ足涿ෘෙ규淮극淺丈ฎ踏丕业ผ渠丢两並測丰渴丹丽举渾김久깊乌乒乓乘깝乞칠买乱亏云些亞交亦溫躺꺼亿껍껏컬仰컴漂輒伕众伝伞켤꼭似꼿住ཏནབ콘潙佛རལ콩轰佽较辆來侍输辦侮辰꾸边进远保连濡迦꿨修送适々〇倉《個》「」倍倒瀕借造倡連週\uF03E倾あ끄え끈運が遍けご끗せ停끝ちつ健灯も炒ん낙낡傢ォ邪邮邯グ傳スッ烈냉僊フブポム烤ュ烦ョ烧部レ热鄂널넘鄰鄱鄲儼녀녁兄酇테免념텔煤內煮酱兴具典녹놀冀内熊册再冬冰冷净释野釒几処눈刑눕刘爬ሺሻ爾牙뉜鉞牟鉤牧늙力办功务㊣努트势틀勇勣닥勳담錄猪献獄升半卖던덩덮却厉玖厨玩厲玻叁又友双폐受变叟珠鏡口珣古句台右号돼吁琅吉吕吗될琰鐵吹呂瑞瑰瑶璃钅针咖璜咯咸铁铅哉득哟든듣哦퓨哭듯链唉唔锢딪唱申딸画땅問畏啓留镜啡啦항喂喝떤떨單閻疼떼허헤闭问闹혀협똑嘗嘴阴阵附陆홍噔확癡除器险황陪陰隆隋회皓뚜嚟嚣嚳难囂盐盖雖盜团難雨困囲圆圍霍휘뜰霸睇흉坏坐靑青块睛睢鞋랗힘瞽城韦韩短矮렀堅᠌᠍領련ᠨᠩᠬᠰ場破ᠴ렷령례硃?塗론硬롭顶须顼顿ᢂ领颗墙颛颜墟風墨碰飛飠磨士壬壹飽磾夋复餐夕館礼社祁套奘륙饣饭饱奶祸她饿禅馆妍香릴私种秒姒맛租姬移駿威騭娱娲稷멀驊멋積멍穷骑窗骨媪窬體ﬁ笑鬚嬴笵묻筆等答魚孛孤孩筷它宋宙宜宝实鮮害므寄寒篤篮寺导対믿밀将밝밤尤밥밭就백尾层屆屋籌屠岁岐법벗粗粤粮岳벽糊볍糖鳥볶峻볼봄崇素索累?統붉붓綠鶩網巢緱布师鸡帥带帮帽ṃṅ빈빌幫빵幹广纘Ạ红Ầ麦级Ẩ纪Ậ庭康纸麻线Ề组织终绍黑Ớ廢Ủ黨弄뼈弋ἔ5:01:00？ἡἸ鼻ὄ齐网齒罔录置影往羊律後龙龜龟忍忘忙ῦ念";
        char[] chars=s.toCharArray();
        long t0 = System.currentTimeMillis();   //获取开始时间
        StringBuilder sb=new StringBuilder();
        for(int i=0;i<1000000;i++){
            sb=new StringBuilder();
            for(char c:chars)
                sb.append(c);
            sb.toString();
        }

        long t1 = System.currentTimeMillis(); //获取结束时间
        System.out.println("clear程序运行时间： " + (t1 - t0)/1000.0 + "s");
        List<Character> list=new ArrayList<>();
        for(int i=0;i<1000000;i++){
            list=new ArrayList<>();
            for(char c:chars)
                list.add(c);

        }
        long t2 = System.currentTimeMillis(); //获取结束时间
        System.out.println("new程序运行时间： " + (t2 - t1)/1000.0 + "s");

    }

    static void testLines() {
        Duncoder duncoder = new Duncoder();
        String s = "1075694eiatnsorluhdcmgk.оpанv,итеіbрвاс40867yfкwлjلд/1му:0пTم-و=я2نàзđäر3ưьõгه9)(يتبáчدع8Sбف<ی45Lх7>6CüцйPNxIїוqיếAHMạжôệسêהюVDâEóộz?حKBơقốਕớجềأờєਾфة،ਿלấكתậшרĐـủảב'ợíOịWشאầì_öصמשăਸFਰểВСUщزRựGਨПکਮੀởىứضਂخנử–إổ;ụãУọਆםიაІéùữКדú\u200CקýМطذОНਵכעắੇਤòъừQਦJਟפეੁồਜਹਲثРੈАחצXДЛ’ũگსТסਗặੱ“Б”YੋءਅîГЗĩằਪრגè*پੰ«ਣੂآן»ỉŭטმغ{+}َო][ბФХỏდ!—਼ئ%ظლĝıễਇыזਭĉნვИỗਡუşਬთÜẫЕẩ다Цچçੌਚਖਫ•ُ ẳЯგЧਧךZਥẻųė老師ẽّəÔკШẹტშţ੍ًٔ|ؤِ&āხš„Õỳ@ЄЮਛ്।ْწžਈ$าĵЙ하ỷਉਊცąαზ№ÂਓÁłïץ…ūนŝਢპף#ਏ¬，Äīτქਐč사εรęνι؟ỡژåûỹფЖ어οա이ിỵ\u200Eਯกσ李ēക้`ยง는ǎਔ‘Ĉყ자ґ니ρјẤ에į지ാżੜэ중의ηնനǐ่〉ğั×İπφวอีเ。£ม·→ുิძ～ŋωത가ő국ե제기Щё나ღςր강ÉλทỞര대서（）ίവ\uF05Aบțκμล로리Þ۲ห陳Ðñਝ고를ěո☞ٍയ~χЬด한γ조ÎōυЇ不라?ëപคส人王\u200DĜիไ娘ÇÍńśșψ요ਘ간黃시오ല姑帝۰۱ÀÓ여ٌ을া그ใ儀문수ːθ蔡아明음њള之天子부ỨƯόЭ있ข以^∨्ปჩ해´Ýǒ∧습ւְ은것แ단동\\³́έ용인校체െ道ćմڪाംട배Şǔή위िษู皇�δ敏؛क্武과主도大마만學而文전റි들°Ö教ּ우۵்മസต명?ẵð√կսք۳으일র་長國보ĥάָ정词ภะ也仁任내학미번‚լ۴۹입িව院히婉德Ò선க게ങേพ可小ßℏ慧ў장梁ব구上ึุ公래夫年\u00AD¿Ûώ٫朝ਠധ一三ำ치、的基家병者\u2028€성ļű소յִ۶衍मय개ദ글余信濤ჭ터其吳화知好穎무æ세Ґֶֹ日有ेকโ濟太바维；怡¥생成−알գվַ是नে氏趙්จ็漢〈ჯჰー늘합陽르高宗방²상ǃ실տ적पয진거ഷ금ณื于佩の劉두표和神女学발鳳분心\u200BÈŽɪ英안와ە林주言ണถ中ས無醫金韓祖婷字寶山常麗黎Ỷ徵±÷식ξђљ쓰դծհ양ֵ영曾작行रीন語ഗഭ七ซ何鄭할면물ẢỔ聖¡øČɛ신ύћ씨政ֲٽ۸۽저ुল각공越මස下ธ个ོ타然톤則副당더后生때형国름말宇?州비齊美¹řŠů芝심ց없원वह감차水ചයු侯키イ邱털년先前달十南玲포君등甜嘉雅在람头秀모孝본Ế徐志思†職Ú胡Ơ所ΑΓΨΩ蓮էխ蕭언方연星智움ۆ۷유杜栩तस楊এশ欣மి초교ญฐ五達て날ィ통勤錦瑩璐員關현雪順며安定少반ṭ！張義必ÊĂ손手ɑ≠范ү앞զը약晉운月机재종ंखদপম民처계ක์你车通いすな傑낮ジ태郭ン都儿关데편理된唐따电雄러科머官經ṇ育þŜ我舒戴ʊЪѝՎ՛╞ևׁ역올١景워۾本杨점लशष楚要টত়權譚று같ీ건천貞경法ഒ洛ജഴോනප为亮代གད体潘使남為탄プ鄧노卿原叫球필행四真地랑页夏如馬禮많姚매秦竇魏?紹불麥¢설Ś속憲自與興∗ɾ스̄莊Ββ董բ敬数պ蘭ٿ暖울ڻ육来읽果접表एठणदংউ짜正டபకవ貝결貴곧അഎശഹീ추දරศ么二軍仲伍མིེྡ農这進ま邑をキシテトバ郝ロ너네照共里重父片닫답勿卓玄只各合名听周디用商떻호白盈相뜻非面矣音᠋려료루外多奇禄妙程马笔第孫宮寧对封將展복鴻ṣ度得耀怨耳恭聰悅\u0096§©산ĮŌ셨慶ź慾⇒∙戰良ʂʰ芳芷ʼ싸̌ΔΧζ葉않թռ新ֱ▲藏旵蘅昇蘇예易春時晟虹ڊ욱ہێ曦服잡柱柳절梅裘줄आडृै西見ধস১樂让லவி책汉్象겠관ഇഞ洲ൂൊ浩海起권ටල귀ැ淑ී丁ชฒ且世ฤ温길ๆ까九书침亂了身从令们仮滿会輝潔འ潤澄ྣ係过近瀞\uF02D過こし遠とに火はる邢ウェク邵点탁郡・元兒煒녕六凱列初到動勞파鍾珊珍同되琬琳瑛田由피핑함開嗣또활百監盧雯雷眉霞靜瞿᠀렇레ᠠ력题碧奉奕른馮娟婆穗몇목竹符孙筠宁守完富민籍紀봅綺緯빠府经Ố廖羅龍快怕性쁘聯您살想意感ķœ慢Ŭ↔ǚ숙房打쉽ʃ报草ΜΣΥΦ放Մձղ整չջ야փ斤열时昭虎٪웬曹會期木잘衣उगजधफो極ই즉ডণষ직집৭詞謀欢欧議死段ை母ం갑갔气ట谢దనమాేొ貳賀ആഈ洗活හ跟清급与东游乐친习깨事什車介他커伯ང位轩ུ辕\uF03Dで크側備난那邦\uF0A7낸タ탕ペメラ焉光克煥八높农燕出누分利爱加능动특化千危去돌发叔瑗呢钟드甌错长門善喺门阮환陵陶登發횟후回因目휴眼희락략량韋럼研堯報렵?록飞声류票祿妃妮림막秋始姐種究媒媧못立章孔客容받尺버벳변?終組工己已ṛ빨平黄Ộ弘弟当忠怀怎\u200F″恒恥息聽\u0080肃\u0093º삼Å색惠能愛석ĭ愿셀ņ셈慎⅛憂Ə송應Ǎǜ臣臨懷或战순舜舞숫户才ɔ쉬色承把护슨苦십拾指쌍써쎄ΆΉΕΘΚΝΞΠ接ќ蒙操ԱԳ收攸改Տճ얕억얻얼엄族׳״엽였옷왔晚왜虞٢٣晴ٻ虽ڌڏړھ更望朵응杰東익잃임잠柴样根桀桓존袁좋\u089E被준च褚थগজহী见질览觉짝째訟許誠說쪽調論諸謂譯讓த认யளழ殷许ா毅每比毛话诞该课갈갖గజడ착ప찮య참స江汽척겨견청河油治波ഉ촌ഥബ贺派资\u0D45赖赤ൾ궁超ඉ축趕출충ළා跑深跳淵근万踐专丘측긴층乎칭카칸乾躂源亲亶今께件།伏休輪꽃但低코པ佗ཙ作软轻辛便俊迎运还候倚速逵遂く끓さ做ただっば灵べ끼りれ炎ガ킬\uF0B8냄ドベ택\uF0DE냥ル儂儒煚入兪兰兵농놓写토决量刚別别투爸刻版剛牛特느늦助튼님独티勾包팔北猶匹午华協卜厅玉펜参평及돕史現吃向吧됩吾呀告둘풀품풍뒤哎哗哥哪프甘唯딱啊畫異했喜떠향病험间嘅阳혹随際雒盛뛰雲直零固看圣뜨眾着靈흐睡란垂鞅埧石럽頁ᠢᠭ顏顯项ᢀᢁᢇ风食夜夠祚奪奭릅妈妹离秉망맞駱騎먹먼穀멈空穿骂媚站嫩鬼策孟季简算宣뮤箱節範寬寸밖尚簡尼簾居?类鱼벌범粵별紧북左巧巾希鸟縣ḥ織干并鹿續应麟庸结给Ử开张缪뽑彝形齿很從御群뿐習翰翻忽考‐‒耕急怪 总恐聞聪⁴肆肉悉\u008F肚\u009F悦ª®肯¶삶悺½情새ËÌ惑Ù惟惡Āď脑ĒĕĚ脚愚Ğ™ģĤĪ脱脸ľŎŢťŪ慮腿膏←솔膽懂Ǐ臏ǑǓǕǖǗǘǙǛ臥ǹ臻懼∅∆∈舉舌숨∩航般戶船扇≈扈ɕɡ扫折花슷승抽ˈ拉拍拜若拨싶苹拿̀挂按挑⌒̝̠̕茶̹挺挾捌荏药掉Ό排ΙΟΤ推莽菁提揚菜握菩華援Ј搜搞营搬萬ѓџ蒜播蒯Ө쓴쓸┴씻Հ故Ն蕉앉Ս救敗՝애敢ժշ薛□薪斬於施־旅업엇었엌藍既藥旧早昆옛昞昨옮昼완♈晋왕處晝٠٧٨٩普외ٺ왼ڍ暑욕ڙڳ웃曆曇蛇曉蛋曌월書替最蜀朋✓末朱杂条杲极잉잊잖枝柒蟜査젊树젓젼格血桂桌桑術衡补족졸좀袋\u088C袜梨梯梵棄棒죠裤裴죽अ椅इईऔ褘ञट襄쥐業概আ視榜증ঢ覧ফভ榮觀ূ视構ো짐様解০짧৪৫৬৯訁樣訥訳橋쩌機話誓쫒諗諮次識此譬護讀殆இ變讠讨ர记讲殿识ே诉译毕语说请读谁调అ찌찍ఒతఫభరష창ు채豊ౌ汗豫汴걷걸決검겁겉沙沛철没격첫겸겹賁곤注泮곱泳ഏഖ괜贝총洞ഠഡ贰괴贵流浅ൈ赋赛赢走굴굽涂涇趟තධ足涿ෘෙ규淮극淺丈ฎ踏丕业ผ渠丢两並測丰渴丹丽举渾김久깊乌乒乓乘깝乞칠买乱亏云些亞交亦溫躺꺼亿껍껏컬仰컴漂輒伕众伝伞켤꼭似꼿住ཏནབ콘潙佛རལ콩轰佽较辆來侍输辦侮辰꾸边进远保连濡迦꿨修送适々〇倉《個》「」倍倒瀕借造倡連週\uF03E倾あ끄え끈運が遍けご끗せ停끝ちつ健灯も炒ん낙낡傢ォ邪邮邯グ傳スッ烈냉僊フブポム烤ュ烦ョ烧部レ热鄂널넘鄰鄱鄲儼녀녁兄酇테免념텔煤內煮酱兴具典녹놀冀内熊册再冬冰冷净释野釒几処눈刑눕刘爬ሺሻ爾牙뉜鉞牟鉤牧늙力办功务㊣努트势틀勇勣닥勳담錄猪献獄升半卖던덩덮却厉玖厨玩厲玻叁又友双폐受变叟珠鏡口珣古句台右号돼吁琅吉吕吗될琰鐵吹呂瑞瑰瑶璃钅针咖璜咯咸铁铅哉득哟든듣哦퓨哭듯链唉唔锢딪唱申딸画땅問畏啓留镜啡啦항喂喝떤떨單閻疼떼허헤闭问闹혀협똑嘗嘴阴阵附陆홍噔확癡除器险황陪陰隆隋회皓뚜嚟嚣嚳难囂盐盖雖盜团難雨困囲圆圍霍휘뜰霸睇흉坏坐靑青块睛睢鞋랗힘瞽城韦韩短矮렀堅᠌᠍領련ᠨᠩᠬᠰ場破ᠴ렷령례硃?塗론硬롭顶须顼顿ᢂ领颗墙颛颜墟風墨碰飛飠磨士壬壹飽磾夋复餐夕館礼社祁套奘륙饣饭饱奶祸她饿禅馆妍香릴私种秒姒맛租姬移駿威騭娱娲稷멀驊멋積멍穷骑窗骨媪窬體ﬁ笑鬚嬴笵묻筆等答魚孛孤孩筷它宋宙宜宝实鮮害므寄寒篤篮寺导対믿밀将밝밤尤밥밭就백尾层屆屋籌屠岁岐법벗粗粤粮岳벽糊볍糖鳥볶峻볼봄崇素索累?統붉붓綠鶩網巢緱布师鸡帥带帮帽ṃṅ빈빌幫빵幹广纘Ạ红Ầ麦级Ẩ纪Ậ庭康纸麻线Ề组织终绍黑Ớ廢Ủ黨弄뼈弋ἔ5:01:00？ἡἸ鼻ὄ齐网齒罔录置影往羊律後龙龜龟忍忘忙ῦ念";
        List<Byte> bytes = duncoder.stringToBytes(s);
        char[] chars=s.toCharArray();
        long t0 = System.currentTimeMillis();   //获取开始时间

        for(int i=0;i<10000;i++)
            duncoder.charsToBytes(chars);
        long t1 = System.currentTimeMillis(); //获取结束时间
        System.out.println("程序运行时间： " + (t1 - t0)/1000.0 + "s");

        for(int i=0;i<10000;i++)
            duncoder.bytesToString(bytes, null, null, null, null);
        long t2 = System.currentTimeMillis(); //获取结束时间
        System.out.println("程序运行时间： " + (t2 - t1)/1000.0 + "s");
//结论，只有ArrayList.clear() 时间较低
    }

    static void testDir(String dir) {
        File files = new File(dir);
        Duncoder duncoder = new Duncoder();
        //u2d
        for (File f : files.listFiles()) {
            if (!f.isDirectory()) {
                String path = f.getPath();
                if (path.endsWith(".txt")) {
                    long startTime = System.currentTimeMillis();   //获取开始时间
                    duncoder.UfileToDfile(path, path + ".dun");
                    long endTime = System.currentTimeMillis(); //获取结束时间
                    System.out.println(" ---UfileToDfile程序运行时间： " + (endTime - startTime) / 1000.0 + "s");
                }
            }
        }
        //d2u
        for (File f : files.listFiles()) {
            if (!f.isDirectory()) {
                String path = f.getPath();
                if (path.endsWith(".txt")) {
                    long startTime = System.currentTimeMillis();   //获取开始时间
                    duncoder.DfileToUfile(path + ".dun", path + ".dun.utf8");
                    long endTime = System.currentTimeMillis(); //获取结束时间
                    System.out.println(" ---DfileToUfile程序运行时间： " + (endTime - startTime) / 1000.0 + "s");
                }
            }

        }
    }

    public static void testPair(String path) {
        testFile("u",path,path+".dun");
        testFile("d",path+".dun",path+".dun.utf8");
    }
    static void testFile(String type, String inpath,String outpath) {
        Duncoder decoder = new Duncoder();
        long startTime = System.currentTimeMillis();   //获取开始时间
        if(type.equals("u")){
            if(outpath==null)
                outpath=inpath+".dun";
            decoder.UfileToDfile(inpath, outpath);}
        else {
            if(outpath==null)
                outpath=inpath+".utf8";
            decoder.DfileToUfile(inpath, outpath);
        }
        long endTime = System.currentTimeMillis(); //获取结束时间
        System.out.println( type+" ---testFile程序运行时间： " + (endTime - startTime) / 1000.0 + "s");
    }

    static void testLine() {
        Duncoder decoder = new Duncoder();
        String r = "᠀᠋ ᠀᠌ ᠀᠍ ᠀\u200D";
        r = "翻訳す";
        String s = r.substring(1);
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode duncode = new Duncode(tool);
        List<Byte> bytes = decoder.charsToBytes(chars);
        String t = decoder.bytesToString(bytes, null, null, null, null);
        System.out.println(s);
        System.out.println(t);
    }

    void testLineBytes() {
        // char[]->Duncode[]
        String t = "A§Ŭÿぃ好乇ЖΘक़ꌊ걹ओऔकखगघङАБВГДЕЖشصضطظعغػ0";
        String s = t.substring(0);
        System.out.println(s);
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            if (lastDun.compress(nowDun)) //char[] test
                continue;//只需要原来Duncode附加即可  有更新,加完了再处理
            else {
                lastDun = nowDun;
                dunList.add(lastDun);
            }
        }

        ArrayList<Byte> byteList = new ArrayList<>();
        for (Duncode d : dunList) {
            for (Byte b : d.getBytes())
                byteList.add(b);
        }
        ArrayList<Byte> bucket = new ArrayList<>();
        ArrayList<Character> charList = new ArrayList<>();
        ArrayList<Duncode> dunList2 = new ArrayList<>();

        for (Byte b : byteList) {
            bucket.add(b);
            if ((b & 0x80) == 0) { //末尾字节
                Duncode nowDun = new Duncode(tool);
                nowDun.readDunchars(bucket);
                dunList2.add(nowDun);
//                System.out.println(bucket);
                nowDun.printDun();
                charList.addAll(nowDun.getChars(null));
                bucket.clear();
            }
        }

        System.out.println();
        for (int i = 0; i < dunList.size(); i++) {
            System.out.println("  " + i + "   ");
            Duncode d1 = dunList.get(i);
            Duncode d2 = dunList2.get(i);
            d1.printDun();
            d2.printDun();
        }
        System.out.println();


        for (int i = 0; i < chars.length; i++)
            System.out.println(chars[i] + "\t" + charList.get(i));
    }

    void testOneDunBytes() {
        //教训  对象更改引用要注意
        // char[]->Duncode[]
        String t = "0A§Ŭÿぃ好乇ЖΘक़ꌊ걹ओऔकखगघङАБВГДЕЖشصضطظعغػ0";
        String s = t.substring(0);
        System.out.println(s);
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
        ArrayList<Byte> byteList = new ArrayList<>();
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            if (lastDun.compress(nowDun)) //char[] test
                continue;//只需要原来Duncode附加即可
            byte[] bytes = nowDun.getBytes();
            for (byte b : bytes)
                byteList.add(b);
            Duncode reDun = new Duncode(tool);

            reDun.readDunchars(bytes); //会破坏
            System.out.println(c + "\t" + reDun.getChars(null).get(0));

            dunList.add(nowDun);
            lastDun = nowDun;
        }
        ArrayList<Byte> bucket = new ArrayList<>();
        ArrayList<Character> charList = new ArrayList<>();
        for (Byte b : byteList) {
            bucket.add(b);
            if ((b & 0x80) == 0) { //末尾字节
                Duncode nowDun = new Duncode(tool);
                nowDun.readDunchars(bucket);
//                nowDun.printDun();
                charList.addAll(nowDun.getChars(null));
                bucket.clear();
            }
        }
    }

    void testZoro() {
        // char[]->Duncode[]
        String t = "A§Ŭÿぃ好乇ЖΘक़ꌊ걹ओऔकखगघङАБВГДЕЖشصضطظعغػ";
        String s = t.substring(0);
        System.out.println(s);
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
//        ArrayList<Byte> byteList = new ArrayList<>();
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            if (lastDun.compress(nowDun)) //char[] test
                continue;//只需要原来Duncode附加即可
//            byte[] bytes = nowDun.getBytes();
//            for (byte b : bytes)
//                byteList.add(b);
//            tmpDun=new Duncode(tool);
//            tmpDun.readDunchars(bytes);
//            System.out.println(chars[i] + "\t " + nowDun.zoneId + "\t" + nowDun.langId + "\t" + nowDun.symbleIds + "\t" + nowDun.ge);

            dunList.add(nowDun);
            lastDun = nowDun;
        }

        //Duncode[]->byte[]

        //byte[]->Duncode[]

        //Duncode[]->char[]
//        ArrayList<Duncode> dunList1 = new ArrayList<>();
//        ArrayList<Byte> bucket = new ArrayList<>();
//        for (Byte b : byteList) {
//            bucket.add(b);
//            if ((b & 0x80) == 0) { //末尾字节
//                Duncode nowDun = new Duncode(tool);
//                nowDun.readDunchars(bucket);
//                dunList1.add(nowDun);
//                bucket.clear();
//            }
//        }
        ArrayList<Character> charList = new ArrayList<>();
        for (Duncode d : dunList)
            charList.addAll(d.getChars(null));
//            d.printDun();
        for (int i = 0; i < chars.length; i++)
            System.out.println(chars[i] + "\t" + charList.get(i));
    }

    void testBytes() {
        // char[]->Duncode[]
//        String s = "A§ÿぃ好乇ЖΘक़ꌊ걹";
//        String s="ओऔकखगघङАБВГДЕЖشصضطظعغػ";
        String s = "A§Ŭÿぃ好乇ЖΘक़ꌊ걹ओऔकखगघङАБВГДЕЖشصضطظعغػ";
//        String s="ДЕЖشصضطظع";
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
        Duncode tmpDun = new Duncode(tool);
//        Duncode nowDun = new Duncode(tool);
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            System.out.println(chars[i] + "\t " + nowDun.zoneId + "\t" + nowDun.langId + "\t" + nowDun.symbleIds + "\t" + nowDun.getChars(null).get(0));
//            if(lastDun.compress(nowDun)) //char[] test
//                continue;//只需要原来Duncode附加即可
            byte[] bytes = nowDun.getBytes();
            tmpDun = new Duncode(tool);
            tmpDun.readDunchars(bytes);
            System.out.println(chars[i] + "\t " + tmpDun.zoneId + "\t" + tmpDun.langId + "\t" + tmpDun.symbleIds + "\t" + tmpDun.getChars(null).get(0));

            //            dunList.add(nowDun);
//            lastDun=nowDun;
//            nowDun=tmpDun;
        }

        //Duncode[]->byte[]

        //byte[]->Duncode[]

        //Duncode[]->char[]
        ArrayList<Character> charList = new ArrayList<>();
        for (Duncode d : dunList)
            charList.addAll(d.getChars(null));
        for (int i = 0; i < chars.length; i++)
            System.out.println(chars[i] + "\t" + charList.get(i));
//            System.out.println(chars[i] + "\t " + tmpDun.zoneId + "\t" + tmpDun.langId + "\t" + tmpDun.symbleIds + "\t" + charList.get(i));
    }

    void testCompress() {
        // char[]->Duncode[]
//        String s = "Aÿぃ好乇ЖΘक़ꌊ걹";
//        String s="ओऔकखगघङАБВГДЕЖشصضطظعغػ";
        String s = "Aÿぃ好乇ЖΘक़ꌊ걹ओऔकखगघङАБВГДЕЖشصضطظعغػ";
//        String s="ДЕЖشصضطظع";
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
        Duncode tmpDun = new Duncode(tool);
//        Duncode nowDun = new Duncode(tool);
        for (int i = 0; i < chars.length; i++) {
//            if(i!=6)
//                continue;
            char c = chars[i];
            Duncode nowDun = new Duncode(tool);
            nowDun.readUnichar(c);
            if (lastDun.compress(nowDun)) //char[] test
                continue;//只需要原来Duncode附加即可
            dunList.add(nowDun);
            lastDun = nowDun;
//            nowDun=tmpDun;
        }

        //Duncode[]->byte[]

        //byte[]->Duncode[]

        //Duncode[]->char[]
        ArrayList<Character> charList = new ArrayList<>();
        for (Duncode d : dunList)
            charList.addAll(d.getChars(null));
        for (int i = 0; i < chars.length; i++)
            System.out.println(chars[i] + "\t " + tmpDun.zoneId + "\t" + tmpDun.langId + "\t" + tmpDun.symbleIds + "\t" + charList.get(i));

    }

    void testSingle() {
        String s = "Aÿぃ好乇ЖΘक़ꌊ걹";
        char[] chars = s.toCharArray();
        Tool tool = new Tool();
        Duncode lastDun = new Duncode(tool);

        ArrayList<Duncode> dunList = new ArrayList<>();
        Duncode tmpDun = new Duncode(tool);
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            tmpDun.readUnichar(c);
            if (!lastDun.compress(tmpDun)) //char[] test
                dunList.add(tmpDun);
            System.out.println(c + "\t " + tmpDun.zoneId + "\t" + tmpDun.langId + "\t" + tmpDun.symbleIds + "\t" + tmpDun.getChars(null).get(0));
        }

    }
}
