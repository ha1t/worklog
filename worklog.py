from time import sleep
import datetime
import subprocess
import os
outdir = 'screencapture'

now = datetime.datetime.today()
counter = 1

def capture():
    global counter
    global now
    filename = '%s%02d%02d_%08d.png' % (now.year, now.month, now.day, counter)
    outpath = os.path.join(outdir, filename)
    subprocess.call(['screencapture', '-x', '-C', outpath])
    counter += 1

while True:
    capture()
    sleep(10)
