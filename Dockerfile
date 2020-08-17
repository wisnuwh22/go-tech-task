FROM amd64/golang:1.14
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build main.go ingredient.go recipe.go
CMD ["/app/main"]