
IMAGE_DIR := ./screencapture
FFMPEG := ffmpeg
DATE := $(shell date +'%Y%m%d')
# TARGETS := $(wildcard ./screencapture/*.png)

all:
	$(FFMPEG) -r 10 -i "$(IMAGE_DIR)/$(DATE)_%05d.png" -pix_fmt yuv420p -vcodec h264 -s 1120x700 -y "worklog_$(DATE).mp4"
clean:
	rm $(IMAGE_DIR)/*.png
