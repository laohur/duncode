

import glob
import os
import sys
from logzero import logger


def show_size():
    doc = []
    tgt1_dir = "C:/data/duncoder1"
    tgt2_dir = "C:/data/duncoder2"
    files = glob.glob(
        rf"C:/data/wiki-1m/*.txt")
    for src in files:
        name = os.path.basename(src)
        lang = name[:2]
        tgt1 = f"{tgt1_dir}/{lang}.txt"
        tgt2 = f"{tgt2_dir}/{lang}.txt"
        n_chars = len(open(src).read())
        size0 = os.path.getsize(src)
        size1 = os.path.getsize(tgt1)
        size2 = os.path.getsize(tgt2)

        if n_chars > max(size1, size2):
            logger.error(
                f"{src} n_chars{n_chars} {size0} --> {size1} {size2} ")
            sys.exit()
        logger.info(f"{src} n_chars{n_chars} {size0} --> {size1} {size2} ")

        row = (name, n_chars, size0, size1, size2)
        doc.append(row)
    with open("wiki-show.txt", "w") as f:
        for row in doc:
            l = '\t'.join([str(x) for x in row])
            f.write(l+'\n')
    return doc


if __name__ == '__main__':
    show_size()
