import json
import subprocess
import sys
import os
import sys
import hashlib
import logging


def init_logger(level):
    logger_fmt = '%(message)s'
    logging.basicConfig(format=logger_fmt, stream=sys.stdout, level=level)


def hashFiles(files, folder):
    hashMap = {}
    for file in files:
        if os.path.isdir(file):
            parsedFiles = [os.path.join(dirpath,f) for (dirpath, dirnames, filenames) in os.walk(file) for f in filenames]
            for parsedFile in parsedFiles:
                hashMap[parsedFile] = hashFile(os.path.join(folder, parsedFile))
        else:
            hashMap[file] = hashFile(os.path.join(folder, file))
    return hashMap


def hashFile(file):
    BUF_SIZE = 65536

    md5 = hashlib.md5()

    with open(file, 'rb') as f:
        while True:
            data = f.read(BUF_SIZE)
            if not data:
                break
            md5.update(data)

    return md5.hexdigest()


def readConfig():
    with open(sys.argv[1], 'r') as f:
        return json.load(f)


def setup(testCase):
    logging.debug('Test config: ' + json.dumps(testCase))
    os.mkdir(testCase['outputDeratFolder'])


def build():
    result = subprocess.run('go build .', shell=True, cwd='cmd/rat')
    if result.returncode:
        logging.error("Error while building.")
        exit(1)


def rat(testCase):
    outputRatFile = testCase['outputRatFile'] 
    inputRatFiles = ''
    for file in testCase['inputRatFiles']:
        inputRatFiles += file + ' '
    
    command = f'cmd/rat/main {outputRatFile} {inputRatFiles}'
    logging.debug('Running: ' + command)
    result = subprocess.run(command, shell=True)
    if result.returncode:
        logging.error("Error while rating.")
        exit(1)


def derat(testCase):
    outputDeratFolder = testCase['outputDeratFolder']
    inputFile = testCase['outputRatFile']
    
    command = f'cmd/rat/main -x {inputFile} -C {outputDeratFolder}'
    logging.debug('Running: ' + command)
    result = subprocess.run(command, shell=True)
    if result.returncode:
        logging.error("Error while rating.")
        exit(1)


def validate(ratMap, deratMap):
    if ratMap != deratMap:
        logging.error("Hash maps are different.")
        exit(1)


def clean(testCase):
    outputRatFile = testCase['outputRatFile']
    outputDeratFolder = testCase['outputDeratFolder']

    command = f'rm -rf {outputRatFile} {outputDeratFolder}'
    logging.debug('Running: ' + command)
    result = subprocess.run(command, shell=True)
    if result.returncode:
        logging.error("Error while cleaning.")
        exit(1)


config = readConfig()
init_logger(config['logLevel'])

build()
for case in config['tests']:
    logging.info('===> Test: ' + case['description'])

    setup(case)
    ratMap = hashFiles(case['inputRatFiles'], '')
    logging.debug('Rat map ' + str(ratMap))
    rat(case)
    derat(case)
    deratMap = hashFiles(case['inputRatFiles'], case['outputDeratFolder'])
    logging.debug('Derat map: ' + str(deratMap))
    validate(ratMap, deratMap)
    clean(case)

    logging.info("<=== Test succeded\n")

