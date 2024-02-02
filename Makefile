.PHONY: all
all:clean build archive push

clean:
	del .\kz_bot.zip
build:
	go build
archive:
	7z a kz_bot.zip kz_bot.exe
push:
	git push