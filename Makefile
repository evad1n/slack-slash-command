all:
	GOOS=linux go build main.go
	zip function.zip main
	rm main