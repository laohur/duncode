package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func EncodeFile(src, tgt string, debug bool) {
	log.Printf("EncodeFile src:%s --> tgt:%s started \n", src, tgt)
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.Create(tgt)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	var n_line = 0
	var n_src = 0
	var n_tgt = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line = scanner.Text()
		var bytes = Encode(line)
		if debug {
			var decoded = Decode(bytes)
			if line != decoded {
				var d = strings.Compare(line, decoded)
				log.Fatalf("n_line:%d %s != %s %d", n_line, line, decoded, d)
				os.Exit(n_line)
			}
		}
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
		n_line += 1
		n_src += len(line)
		n_tgt += len(bytes)
		if n_line%1000000 == 0 {
			log.Printf("line %d src %d tgt %d", n_line, n_src, n_tgt)
		}
	}
	w.Flush()
	log.Printf("line %d src %d tgt %d done \n", n_line, n_src, n_tgt)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// var s = "Aÿぃ好乇αβζЖѰѾѾक़ऄळ४ॐꌊ걹"
	// var r="বিৰোধাভাস, কূট, কূটাভাস বা পেৰাডক্স (ইংৰাজী: Paradox) হৈছে দেখাত স্ববিৰোধী যেন লগা কিছুমান উক্তি বা মত, যিবোৰ চালি-জাৰি চালে শুদ্ধ যেনও লাগে বা এনে এক উক্তি যি নিজৰে বিৰোধ কৰে অথবা এনে এক উক্তি যি সাধাৰণ ধাৰণাৰ বিৰোধ কৰে। ই হৈছে সূত্ৰৰ অন্তৰ্বিৰোধ, মিমাংসাহীন দ্বন্দ্বভাব বা বিৰোধভাস, যিটোক সাধাৰণতে অসম্ভৱ দেখা যায় কিন্তু আচলতে সঁচা। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুনিবলৈ অযুক্তিকৰ কিন্তু বাস্তৱত সত্য। উদাহৰণস্বৰূপে 'মন্তি হল পেৰাডক্স'। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুনিবলৈয়ো অযুক্তিকৰ আৰু বাস্তৱতো অসত্য। উদাহৰণস্বৰূপে ১=২ ৰ গাণিতিক প্ৰমাণ য'ত শূন্যৰে শূন্যক হৰণ কৰা হয়। অৰ্থাৎ এনে উক্তিসমূহ কিছুমান ভুল তথ্যৰ ওপৰত নিৰ্ভৰ কৰে যাৰ বাস্তৱত কোনো অস্তিত্ব নাই। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুদ্ধও নহয়, ভুলো নহয় বা শুদ্ধও হ'ব নোৱাৰে আৰু ভুলো হ'ব নোৱাৰে। উদাহৰণস্বৰূপে ককাদেউতাৰ কূটাভাস (The Grandfather Paradox) ৰ কথা ক'ব পাৰি। (The Berber Paradox) গাঁৱৰ নাপিত এজনে নিয়ম কৰি লৈছে যে তেওঁ কেৱল সেইবিলাক মানুহৰহে ডাঢ়ি খুৰাব যিবিলাকে নিজে নিজৰ ডাঢ়ি নুখুৰায়। প্ৰশ্নটো হ’ল নাপিতজনে নিজৰ ডাঢ়ি খুৰাবনে নাই? যদি নাপিতজনে নিজৰ ডাৰি খুৰায়, তেওঁ কৰি লোৱা নিয়মমতে নিজৰ ডাঢ়ি খুৰাব নালাগিছিল (কাৰণ তেওঁ নিজেই নাপিত)। আকৌ তেওঁ যদি নিজৰ ডাঢ়ি নুখুৰায়, তেওঁৰ নিয়মমতে তেওঁ নিজৰ ডাঢ়ি খুৰাব লাগিছিল। এই পৰিস্থিতি নাপিতজনে কৰিব কি? “মই এটা কথাই জানো যে, মই একো নাজানো।” কূট বা পেৰাডক্সৰ তালিকা ভ্ৰান্তি ভ্ৰম মিছলীয়া কূট কুটাভাস (Paradox), অণু-কলন গোহাঞি, ড॰ হিৰেন. Anglo-Assamese Dictionary. বনলতা. পৃষ্ঠা. ৩০; ৬০২.  বৰুৱা, হেমচন্দ্ৰ. হেমকোষ. হেমকোষ প্ৰকাশন. পৃষ্ঠা. ৭৯৫.  বি. আৰ শৰ্মা, প্ৰণৱ কুমাৰ পেগু. Ajanta Comprehensive Dictionary. অজন্তা প্ৰকাশন. পৃষ্ঠা. ৬৫৯.  ৱিল্লাৰ্দ ভান্ অৰ্মান কুইন "what is russells paradox". scientificamerican.comhttps://scientificamerican.com. https://www.scientificamerican.com/article/what-is-russells-paradox/। আহৰণ কৰা হৈছে: May 22, 2020."
	var r = "বিৰোধাভাস, কূট, কূটাভাস বা পেৰাডক্স (ইংৰাজী: Paradox) হৈছে দেখাত স্ববিৰোধী যেন লগা কিছুমান উক্তি বা মত, যিবোৰ চালি-জাৰি চালে শুদ্ধ যেনও লাগে বা এনে এক উক্তি যি নিজৰে বিৰোধ কৰে অথবা এনে এক উক্তি যি সাধাৰণ ধাৰণাৰ বিৰোধ কৰে। ই হৈছে সূত্ৰৰ অন্তৰ্বিৰোধ, মিমাংসাহীন দ্বন্দ্বভাব বা বিৰোধভাস, যিটোক সাধাৰণতে অসম্ভৱ দেখা যায় কিন্তু আচলতে সঁচা। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুনিবলৈ অযুক্তিকৰ কিন্তু বাস্তৱত সত্য। উদাহৰণস্বৰূপে 'মন্তি হল পেৰাডক্স'। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুনিবলৈয়ো অযুক্তিকৰ আৰু বাস্তৱতো অসত্য। উদাহৰণস্বৰূপে ১=২ ৰ গাণিতিক প্ৰমাণ য'ত শূন্যৰে শূন্যক হৰণ কৰা হয়। অৰ্থাৎ এনে উক্তিসমূহ কিছুমান ভুল তথ্যৰ ওপৰত নিৰ্ভৰ কৰে যাৰ বাস্তৱত কোনো অস্তিত্ব নাই। এনে কূটাভাসৰ অন্তৰ্গত উক্তিসমূহ শুদ্ধও নহয়, ভুলো নহয় বা শুদ্ধও হ'ব নোৱাৰে আৰু ভুলো হ'ব নোৱাৰে। উদাহৰণস্বৰূপে ককাদেউতাৰ কূটাভাস (The Grandfather Paradox) ৰ কথা ক'ব পাৰি। (The Berber Paradox) গাঁৱৰ নাপিত এজনে নিয়ম কৰি লৈছে যে তেওঁ কেৱল সেইবিলাক মানুহৰহে ডাঢ়ি খুৰাব যিবিলাকে নিজে নিজৰ ডাঢ়ি নুখুৰায়। প্ৰশ্নটো হ’ল নাপিতজনে নিজৰ ডাঢ়ি খুৰাবনে নাই? যদি নাপিতজনে নিজৰ ডাৰি খুৰায়, তেওঁ কৰি লোৱা নিয়মমতে নিজৰ ডাঢ়ি খুৰাব নালাগিছিল (কাৰণ তেওঁ নিজেই নাপিত)। আকৌ তেওঁ যদি নিজৰ ডাঢ়ি নুখুৰায়, তেওঁৰ নিয়মমতে তেওঁ নিজৰ ডাঢ়ি খুৰাব লাগিছিল। এই পৰিস্থিতি নাপিতজনে কৰিব কি? “মই এটা কথাই জানো যে, মই একো নাজানো।” কূট বা পেৰাডক্সৰ তালিকা ভ্ৰান্তি ভ্ৰম মিছলীয়া কূট কুটাভাস (Paradox), অণু-কলন গোহাঞি, ড॰ হিৰেন. Anglo-Assamese Dictionary. বনলতা. পৃষ্ঠা. ৩০; ৬০২.  বৰুৱা, হেমচন্দ্ৰ. হেমকোষ. হেমকোষ প্ৰকাশন. পৃষ্ঠা. ৭৯৫.  বি. আৰ শৰ্মা, প্ৰণৱ কুমাৰ পেগু. Ajanta Comprehensive Dictionary. অজন্তা প্ৰকাশন. পৃষ্ঠা. ৬৫৯.  ৱিল্লাৰ্দ ভান্ অৰ্মান কুইন "
	var s = string([]rune(r)[:])
	// testDuncode(s)
	// testDuncodeCompress(s)
	testLine(s)

	// var src = "C:/data/wiki-1m/as.txt"
	// var tgt = src + ".duncode1"
	// EncodeFile(src, tgt, true)
	fmt.Print("done")

}

func main0() {
	// var s = "Aÿぃ好乇αβζЖѰѾѾक़ऄळ४ॐꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	// var bytes = Encode(s)
	// var t = Decode(bytes)
	// fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
	args := os.Args
	var src = args[1]
	var tgt = args[2]
	var debug = false
	if len(args) > 3 {
		debug = true
	}
	// var src = "C:/data/sentences.csv"
	// var tgt = src + ".duncode1"
	// log.Printf("done")
	EncodeFile(src, tgt, debug)
	// log.Printf("done")

}
