#!/bin/bash
original_images=screencapture
ldate=$(date +%Y%m%d)
#ffmpeg -y -f image2 -r 10 -i "$original_images/%08d.png" -s 720x450 "screencapture$ldate.mp4"
ffmpeg -y -f image2 -r 10 -i "${original_images}/%08d.png" -s 720x450 "worklog_${ldate}.mp4"
