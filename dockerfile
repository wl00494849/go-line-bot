FROM golang
WORKDIR /go/src/go-line-bot
ADD . /go/src/go-line-bot

RUN cd /go/src/go-line-bot \
    && go get github.com/joho/godotenv \
    && go get github.com/gin-gonic/gin \
    && go build 

EXPOSE 6666
ENTRYPOINT /go/src/go-line-bot/go-line-bot