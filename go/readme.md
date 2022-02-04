
Duncode v2 by golang


zone	byte1	byte2	byte3	tail	symbols	languages	bytes pre char
ascii				0xxxxxxx	x	ascii	1
双节			1xxxxxxx	0xxxxxxx	x	HanZI…	2
7位字		1nnnnnnn	1xxxxxxx	0yyyyyyy	x,y	Greek…	1.5
8位字		1111nnxx	1xxxxxxy 	0yyyyyyy	x,y	common	1.5
单字	1xxxxxxx	2xxxxxxx	3xxxxxxx	4xxxxxxx	x	rare	4
