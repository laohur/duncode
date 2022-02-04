#!/usr/bin/env python
# -*- coding: utf-8 -*-
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


def read_gz(p):
    input = gzip.open(p)
    while True:
        line = input.readline()
        if not line:
            break
        yield line


def extract_wiki(param):
    src, tgt = param
    xz = tgt+'.xz'
    # if not os.path.exists(tgt) and os.path.exists(xz):
    # return f" {xz} exists"
    # return ''
    for p in [tgt, xz]:
        if os.path.exists(p):
            os.remove(p)
    logger.info(f"{src}  -->  {tgt} ...")
    i = 0
    with open(tgt, 'w') as f:
        input = gzip.open(src, mode='rt', errors='ignore')
        while True:
            a = ''
            i += 1
            try:
                a = input.readline()
                if not a:
                    break
                # b = input.readline()
                content = json.loads(a)
                if 'text' not in content:
                    continue
                # type = index['index']['_type']
                # id = index['index']['_id']
                # language = content['language']
                # revision = content['version']
                # if type == 'page' and content['namespace'] == 0:
                # title = content['title']
                # text = content['text']
                # drop references:
                # ^ The Penguin Dictionary
                if 'text' in content:
                    text = content['text']
                    text = re.sub(r'  \^ .*', '', text).strip()
                    f.write(text+'\n\n')
            except Exception as e:
                logger.error(e)
            if not a:
                break
    os.system(f"xz {tgt}")
    logger.info(f"{src}  -->  {tgt} lines:{i}")
    return tgt


def mparse():

    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('--lang', default="global",  type=str)
    args = parser.parse_args()
    print(args)
    lang = args.lang
    gzs = list(glob.iglob(
        f"F:/data/wiki-20220124-cirrussearch-content.json.gz/*.gz", recursive=True))

    params = []
    for src in gzs:
        name = os.path.basename(src)
        src = "F:/data/wiki-20220124-cirrussearch-content-json-gz/"+name
        if not os.path.exists(src):
            continue
        t = name.rstrip(".json.gz")
        tgt = "F:/data/wiki-20220124-cirrussearch-content-txt-xz/"+t+'.txt'
        param = (src, tgt)
        params.append(param)

    random.shuffle(params)
    import multiprocessing
    with multiprocessing.Pool(6) as pool:
        re = pool.imap_unordered(extract_wiki, params)
        for x in re:
            logger.info(x)


def parse_all():
    from download import get_names
    names = get_names()
    for name in names:
        src = "F:/data/wiki-20220124-cirrussearch-content-json-gz/"+name
        if not os.path.exists(src):
            continue
        t = name.rstrip(".json.gz")
        tgt = "F:/data/wiki-20220124-cirrussearch-content-txt-xz/"+t+'.txt'
        extract_wiki((src, tgt))
        # break


if __name__ == '__main__':
    # mparse()
    parse_all()
