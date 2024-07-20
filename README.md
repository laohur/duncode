# duncode    
an efficient universal character encoder, implement of paper ["Duncode Characters Shorter"](./Duncode Characters Shorter.pdf).   

## corpus
https://github.com/laohur/wiki2txt

## v0
cpp

## v1
java

## v2
* go1: default Duncoder
* go2: is another style of Duncode with different zone distribution, worse than go1


## performance of encoders 
### v1
| corpus          | local encoder   | local (MB) | UTF-8 (MB) | Duncode (MB) | UTF-8/ local | Duncode/ local |
| --------------- | --------------- | ---------- | ---------- | ------------ | ------------ | -------------- |
| enwikisource    | ISO 8859-1      | 1,558      | 1,564      | 1,562        | 1.00         | 1.00           |
| frwikisource    | ISO 8859-1      | 154        | 159        | 158          | 1.03         | 1.03           |
| betawikiversity | ISO 8859-1      | 7          | 9          | 8            | 1.27         | 1.17           |
| zhwiki          | GB 18030        | 917        | 1,294      | 955          | 1.41         | 1.04           |
| zhwikisource    | GB 18030        | 2,088      | 3,095      | 2,173        | 1.48         | 1.04           |
| jawiki          | JIS X 0208–1990 | 1,930      | 2,787      | 1,952        | 1.44         | 1.01           |
| jawikisource    | JIS X 0208–1990 | 52         | 76         | 53           | 1.47         | 1.02           |
| arwiki          | ISO 8859-6      | 762        | 1,297      | 1,086        | 1.70         | 1.43           |
| arwikisource    | ISO 8859-6      | 496        | 886        | 738          | 1.79         | 1.49           |
| ruwiki          | ISO 8859-5      | 3,156      | 5,566      | 4,546        | 1.76         | 1.44           |
| ruwikisource    | ISO 8859-5      | 624        | 1,075      | 890          | 1.72         | 1.43           |

### v2
|    language   | wiki_file_1m |   n_chars  | n_bytes      (utf8) | n_bytes      (duncode) | n_bytes/n_chars      (utf8) | n_bytes/n_chars      (duncoder) | utf8/duncode      (size) |
|:-------------:|:------------:|:----------:|:-------------------:|:----------------------:|:---------------------------:|:-------------------------------:|:------------------------:|
|    English    |    en.txt    | 1,054,002  |          1,066,474  |             1,061,949  |            1.01             |              1.01               |          100.43%         |
|     French    |    fr.txt    | 1,054,065  |          1,096,721  |             1,094,085  |            1.04             |              1.04               |          100.24%         |
|     Arabic    |    ar.txt    | 1,103,308  |          1,855,164  |             1,462,890  |            1.68             |              1.33               |          126.82%         |
|    Russian    |    ru.txt    | 1,049,337  |          1,821,554  |             1,398,275  |            1.74             |              1.33               |          130.27%         |
|    Chinese    |    zh.txt    | 1,052,649  |          2,420,113  |             1,740,409  |            2.30             |              1.65               |          139.05%         |
|    Japanese   |    ja.txt    | 1,051,113  |          2,689,017  |             1,872,561  |            2.56             |              1.78               |          143.60%         |
|     Korean    |    ko.txt    | 1,048,759  |          2,103,649  |             2,087,029  |            2.01             |              1.99               |          100.80%         |
|   Abkhazian   |    ab.txt    | 1,049,130  |          1,846,144  |             1,411,219  |            1.76             |              1.35               |          130.82%         |
|    Burmese    |    my.txt    | 1,052,890  |          2,820,647  |             1,456,799  |            2.68             |              1.38               |          193.62%         |
| Central Khmer |    km.txt    | 1,081,258  |          2,890,359  |             1,516,227  |            2.67             |              1.40               |          190.63%         |
|    Tibetan    |    bo.txt    | 1,053,038  |          3,108,029  |             2,080,294  |            2.95             |              1.98               |          149.40%         |
|     Yoruba    |    yo.txt    | 1,050,996  |          1,230,927  |             1,193,098  |            1.17             |              1.14               |          103.17%         |


## Citation
```bib
@article{Xue2023DuncodeCS,
  title={Duncode Characters Shorter},
  author={Changshan Xue},
  journal={ArXiv},
  year={2023},
  volume={abs/2307.05414},
  url={https://api.semanticscholar.org/CorpusID:259766322}
}

```


## License
[Anti-996 License](https://github.com/996icu/996.ICU/blob/master/LICENSE)