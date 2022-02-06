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


def extract_wiki(src, tgt):
    with lzma.open(src, mode="rt") as f:
        doc = f.readlines(1024*1024*1)
    with open(tgt, "w") as w:
        w.writelines(doc)

    logger.info(f"{src}  -->  {tgt} lines:{len(doc)}")
    return tgt


def extract_wikis():
    doc = []
    tail = "wiki-20220131-cirrussearch-content.txt.xz"
    tgt_dir = "C:/data/wiki-1m"
    files = glob.glob(
        rf"F:/data/wiki-20220131-cirrussearch-content-txt-xz/*{tail}")
    for src in files:
        name = os.path.basename(src)
        if name[2:] != tail:
            continue
        lang = name[:2]
        tgt = f"{tgt_dir}/{lang}.txt"
        extract_wiki(src, tgt)
    return doc


if __name__ == '__main__':
    extract_wikis()
