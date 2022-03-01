
Duncode v1 by golang

| zone id | zone  | byte1     | byte2     | byte3    | tail     | symbols | languages             | bytes pre char |
|---------|-------|-----------|-----------|----------|----------|---------|-----------------------|----------------|
|       0 | ascii |           |           |          | 0xxxxxxx | x       | ascii                 |              1 |
|       1 | 双节  |           |           | 1xxxxxxx | 0xxxxxxx | x       | 0x0080~0x07ff, HanZi… |              2 |
|       4 | 孤字  |           | 1xxxxxxx  | 1xxxxxxx | 0xxxxxxx | x       | rare                  |              3 |
|       2 | 8位字 | 111nnxxx  | 1xxxxxyy  | 1yyyyyyz | 0zzzzzzz | x,y,z   | Greek…                |           1.33 |
|       3 | 7位字 | 1nnnnnnn  | 1xxxxxxx  | 1yyyyyyy | 0zzzzzzz | x,y,z   | Devanagari…           |           1.33 |
