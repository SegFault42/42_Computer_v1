NAME = computer_v1

SRC = main.go

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRC)

clean:

fclean:
	/bin/rm $(NAME)

re: fclean all

dep:
	go get "github.com/kr/pretty"
