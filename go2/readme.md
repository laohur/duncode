
Duncode v2 by golang

| zone id | zone  | byte1    | byte2    | byte3     | tail     | symbols | languages | bytes pre char |
|---------|-------|----------|----------|-----------|----------|---------|-----------|----------------|
|       0 | ascii |          |          |           | 0xxxxxxx | x       | ascii     |              1 |
|       1 | 双节  |          |          | 1xxxxxxx  | 0xxxxxxx | x       | HanZI…    |              2 |
|       2 | 8位字 |          | 1111nnxx | 1xxxxxxy  | 0yyyyyyy | x,y     | common    |            1.5 |
|       3 | 7位字 |          | 1nnnnnnn | 1xxxxxxx  | 0yyyyyyy | x,y     | Greek…    |            1.5 |
|       4 | 孤字  | 10xxxxxx | 2xxxxxxx | 3xxxxxxx  | 4xxxxxxx | x       | rare      |              4 |
|       5 | 私用  | 11xxxxxx | 2xxxxxxx | 3xxxxxxx  | 4xxxxxxx | x       | custom    |              4 |
