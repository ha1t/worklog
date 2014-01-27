# これはなに？

[Macのスクリーンショットで1日の作業を記録して動画で振り返り - 西尾泰和のはてなダイアリー](http://d.hatena.ne.jp/nishiohirokazu/20120731/1343745529)

# Install

```
$ brew install ffmpeg
```

# How to use

```
$ pyrhon worklog.py
```

```
$ make DATE=20131220 IMAGE_DIR=../../screencapture/
$ make clean DATE=20131220 IMAGE_DIR=../../screencapture/
```

# TODO
- マシンがフリーズして再起動したときに自動的にスクリプトが再実行されてほしい(crontab?)
- 二重起動防止したい
