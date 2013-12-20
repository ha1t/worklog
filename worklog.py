from time import sleep
import datetime
import subprocess
import os
outdir = '../../screencapture'

now = datetime.datetime.today()
counter = 1

def capture():
    global counter
    global now

    if now.day != datetime.datetime.today().day:
        now = datetime.datetime.today()
        counter = 1

    filename = '%s%02d%02d_%05d.png' % (now.year, now.month, now.day, counter)
    outpath = os.path.join(outdir, filename)
    subprocess.call(['screencapture', '-x', '-C', outpath])
    counter += 1

while True:
    capture()
    sleep(10)
