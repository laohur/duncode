import os
from logzero import logger


duncoder1 = "C:/doc/duncode/code/duncode/go1/go1.exe"
duncoder2 = "C:/doc/duncode/code/duncode/go2/go2.exe"


def encoderFile(src, tgt, duncoder):
    cmd = f"{duncoder} {src} {tgt}"
    os.system(cmd)


if __name__ == "__main__":
    src = "C:/data/sentences.csv"
    size0 = os.path.getsize(src)
    n_chars = len(open(src).read())
    tgt = src + ".duncode1"
    # encoderFile(src, tgt, duncoder1)
    size1 = os.path.getsize(tgt)
    tgt = src + ".duncode2"
    # encoderFile(src, tgt, duncoder2)
    size2 = os.path.getsize(tgt)
    logger.info(f"{src} 0:{size0} --> 1:{size1} 2:{size2}")

"""
[I 220206 12:44:50 test_global:24] C:/data/sentences.csv 0:543514581 --> 1:504062313 2:511564327
"""
