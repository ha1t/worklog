# これはなに？

[Macのスクリーンショットで1日の作業を記録して動画で振り返り - 西尾泰和のはてなダイアリー](http://d.hatena.ne.jp/nishiohirokazu/20120731/1343745529)

# Install

```
$ brew install ffmpeg
$ go build renumber.go
$ go build worklog.go
```

# How to use

```
$ make DATE=20131220 IMAGE_DIR=../../screencapture/
$ make clean DATE=20131220 IMAGE_DIR=../../screencapture/
```

# setup

Macで使う場合、以下のような内容の worklog.plist を作り、ユーザーのLibrary/LaunchAgents/ におく。

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>worklog</string>
	<key>OnDemand</key>
	<false/>
	<key>ProgramArguments</key>
	<array>
		<string>/Users/XXXXXXXXXX/src/worklog/worklog</string>
	</array>
</dict>
</plist>
```

おいたら

launchctl load ~/Library/LaunchAgents/worklog.plist

を実行するだけ

参考 : https://qiita.com/rsahara/items/7d37a4cb6c73329d4683

# TODO
- マシンがフリーズして再起動したときに自動的にスクリプトが再実行されてほしい(crontab?)
- 二重起動防止したい
