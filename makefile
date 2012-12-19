

IMAGE_DIR := ./screencapture
FFMPEG := ffmpeg

# TARGETS := $(wildcard ./screencapture/*.png)

all:
	$(FFMPEG) -y -f image2 -r 10 -i "$(IMAGE_DIR)/%08d.png" -s 720x450 "worklog_$(shell date +'%Y%m%d').mp4"
clean:
	rm $(IMAGE_DIR)/*.png
renumber:
	php renumber.php
