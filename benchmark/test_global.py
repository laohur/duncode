import os
from logzero import logger


duncoder1 = "C:/doc/duncode/code/duncode/go1/go1.exe"
duncoder2 = "C:/doc/duncode/code/duncode/go2/go2.exe"


def encoderFile(src, tgt, duncoder):
    cmd = f"{duncoder} {src} {tgt}"
    os.system(cmd)
    size = os.path.getsize(tgt)
    return size


if __name__ == "__main__":
    src = "C:/data/sentences.csv"
    size0 = os.path.getsize(src)
    tgt = src + ".duncode1"
    size1 = encoderFile(src, tgt, duncoder1)
    tgt = src + ".duncode2"
    size2 = encoderFile(src, tgt, duncoder2)
    logger.info(f"{src} {size0} -->{size1} {size1}")

"""

"""