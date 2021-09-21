all:
	GOOS=linux go build -o main main.go
	zip echo.zip main
	rm main