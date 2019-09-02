import glob
import os
import sys

def get_target_list():
    files = glob.glob('../../screencapture/*.png')

    target_list = []
    for file in files:
        filename = os.path.basename(file)
        parts = filename.split('_')
        target_list.append(parts[0])
        #print file

    print list(set(target_list))

def renumber(target_dir, target_date):
    #target_date = '20140217'
    #target_dir = '../../screencapture' + '/'
    target_dir = target_dir + '/'

    target_list = []
    #print('%s%s_00000.png' % (target_dir, target_date))
    for file in glob.glob('%s%s_*.png' % (target_dir, target_date)):
        filename = os.path.basename(file)
        target_list.append(filename)
        #print file

    target_list.sort()
    number = 1
    for filename in target_list:
        old = filename
        new = '%s_%05d.png' % (target_date, number)
        print('%s -> %s' % (old, new))
        if old != new:
            os.rename(target_dir + old, target_dir + new)
        number = number + 1

#get_target_list()
renumber(sys.argv[1], sys.argv[2])
