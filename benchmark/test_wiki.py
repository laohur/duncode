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
    cmd = f"{duncoder} {src} {tgt} debug"
    os.system(cmd)
    size = os.path.getsize(tgt)
    return size


def encode_wikis():
    doc = []
    tgt1_dir = "C:/data/duncoder1"
    tgt2_dir = "C:/data/duncoder2"
    files = glob.glob(
        rf"C:/data/wiki-1m/*.txt")
    for src in files:
        size0 = os.path.getsize(src)
        n_chars = len(open(src).read())
        # n_chars = sum(len(x) for x in open(src).read())

        name = os.path.basename(src)
        lang = name[:2]
        if lang != 'bi':
            continue
        tgt = f"{tgt1_dir}/{lang}.txt"
        encoderFile(src, tgt, duncoder1)
        size1 = os.path.getsize(tgt)

        tgt = f"{tgt2_dir}/{lang}.txt"
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
