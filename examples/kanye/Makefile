NAME=kanye

all:
	GOOS=linux go build -o $(NAME) main.go
	zip $(NAME).zip $(NAME)
	rm $(NAME)