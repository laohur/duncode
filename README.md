# duncode    
an efficient universal character encoder, implement of paper "Duncode Charasters Shorter".   

## v0
cpp

## v1
java

## v2
python


## performance of encoders (v1)
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

## License
[Anti-996 License](https://github.com/996icu/996.ICU/blob/master/LICENSE)