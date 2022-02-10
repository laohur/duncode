#!/usr/bin/env python
# -*- coding: utf-8 -*-
import lzma
import sys
import traceback
import glob
import subprocess
import json
import os
import requests
import re
import random
import time
import shutil
from logzero import logger

import gzip

duncoder1 = "C:/doc/duncode/code/duncode/go1/go1.exe"
duncoder2 = "C:/doc/duncode/code/duncode/go2/go2.exe"


def encoderFile(src, tgt, duncoder):
    # if os.path.exists(tgt):
    # logger.info(f"{tgt} exists ")
    cmd = f"{duncoder} encode {src} {tgt} debug"
    with os.popen(cmd) as p:
        for l in p:
            if 'err' in l:
                logger.error(f"{cmd} ")
                sys.exit()
    # os.system(cmd)
    size = os.path.getsize(tgt)
    return size


def encode_wikis():
    doc = []
    tgt1_dir = "C:/data/duncoder1"
    tgt2_dir = "C:/data/duncoder2"
    files = glob.glob(
        rf"C:/data/wiki-1m/*.txt")
    files = list(files)
    srcs = ["C:/data/sentences.csv"]
    srcs += files
    for src in srcs:
        size0 = os.path.getsize(src)
        n_chars = len(open(src).read())
        # n_chars = sum(len(x) for x in open(src).read())

        name = os.path.basename(src)
        # lang = name[:2]
        # if lang != 'bi':
        # continue
        tgt = f"{tgt1_dir}/{name}"
        encoderFile(src, tgt, duncoder1)
        size1 = os.path.getsize(tgt)

        tgt = f"{tgt2_dir}/{name}"
        encoderFile(src, tgt, duncoder2)
        size2 = os.path.getsize(tgt)

        if n_chars > max(size1, size2):
            logger.error(
                f"{src} n_chars{n_chars} {size0} --> {size1} {size2} ")
            sys.exit()
        logger.info(f"{src} n_chars{n_chars} {size0} --> {size1} {size2} ")

    return doc


if __name__ == '__main__':
    encode_wikis()
"""
n_char 479362758  src 543514581 tgt 513994997  tgt 517119260 done

"""
