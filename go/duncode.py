
#         began,end,size,en,ch,zone_name,zone_id=l
"""

    ZoneIdMap = {
        "ascii": 0,  # BlockId 0
        "双节": 1,  # lanid 1
        "8位字": 2,
        "7位字": 3,
        "独字": 4
    }

"""
ShuangJie = open("ShuangJieSet.txt").read().splitlines()
ShuangJie2Idx = {x: i for i, x in enumerate(ShuangJie)}

Blocks = open("Blocks.txt").read().splitlines()
blocks = []
for line in Blocks:
    l = line.split('\t')
    began, end, size, en, ch, zone_name, zone_id = l
    row = [int(began, 16), int(end, 16), int(zone_id)]
    blocks.append(row)


def get_zone(char):
    if ord(chr) <= 128:
        BlockId = 0
        ZoneId = 0
        Index = ord(x)
        return ZoneId, BlockId, Index
    elif char in ShuangJie2Idx:
        BlockId = 1
        ZoneId = 1
        Index = ShuangJie2Idx[x]
        return ZoneId, BlockId, Index
    else:
        for BlockId, row in enumerate(blocks):
            Began, End, ZoneId = row
            if Began <= ord(char) <= End:
                Index = ord(char)-Began
                return ZoneId, BlockId, Index
    return -1
