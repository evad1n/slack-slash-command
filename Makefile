NAME=echo

all:
	GOOS=linux go build main.go
	zip $(NAME).zip $(NAME)
	rm $(NAME)