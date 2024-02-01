.PHONY: all
all:clean archive push

clean:
	del .\kz_bot.zip

archive:
	7z a kz_bot.zip kz_bot.exe

push:
	git push