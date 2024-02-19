FROM golang

WORKDIR /go-backend

COPY . .
RUN go mod tidy

EXPOSE 1234

CMD go run .

# docker rmi (id image)
# docker image ls