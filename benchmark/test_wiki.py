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
    if os.path.exists(tgt):
        logger.info(f"{tgt} exists ")
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
        name = os.path.basename(src)
        lang = name[:2]
        tgt = f"{tgt1_dir}/{lang}.txt"
        encoderFile(src, tgt, duncoder1)
        tgt = f"{tgt2_dir}/{lang}.txt"
        encoderFile(src, tgt, duncoder2)
    return doc


if __name__ == '__main__':
    encode_wikis()
