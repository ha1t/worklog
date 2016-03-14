
IMAGE_DIR := ./screencapture
FFMPEG := ffmpeg
DATE := $(shell date +'%Y%m%d')

all:
	renumber $(IMAGE_DIR) $(DATE)
	$(FFMPEG) -r 10 -i "$(IMAGE_DIR)/$(DATE)_%05d.png" -pix_fmt yuv420p -vcodec h264 -s 1120x700 -y "worklog_$(DATE).mp4"
clean:
	@while [ -z "$$CONTINUE" ]; do \
		read -r -p "Type anything but Y or y to exit. [y/N] " CONTINUE; \
	done ; \
	if [ ! $$CONTINUE == "y" ]; then \
	if [ ! $$CONTINUE == "Y" ]; then \
		echo "Exiting." ; exit 1 ; \
	fi \
	fi
	find $(IMAGE_DIR) -name $(DATE)_*.png | xargs rm
rename:
	rename.sh $(DATE) $(IMAGE_DIR)
