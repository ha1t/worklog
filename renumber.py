import glob
import os

def get_target_list():
    files = glob.glob('../../screencapture/*.png')

    target_list = []
    for file in files:
        filename = os.path.basename(file)
        parts = filename.split('_')
        target_list.append(parts[0])
        #print file

    print list(set(target_list))

def renumber():
    target_date = 20140217
    target_dir = '../../screencapture' + '/'

    target_list = []
    for file in glob.glob('%s%d_*.png' % (target_dir, target_date)):
        filename = os.path.basename(file)
        target_list.append(filename)
        #print file

    target_list.sort()
    number = 1
    for filename in target_list:
        old = filename
        new = '%d_%05d.png' % (target_date, number)
        print('%s -> %s' % (old, new))
        if old != new:
            os.rename(target_dir + old, target_dir + new)
        number = number + 1

#get_target_list()
renumber()
