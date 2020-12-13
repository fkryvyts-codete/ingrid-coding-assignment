FROM golang:1.15 as builder

ENV APP_USER ingrid
ENV APP_HOME /go/src/ingrid

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY ./ .

RUN go mod download
RUN go mod verify
RUN go build -o ingrid-coding-assignment

EXPOSE 8080
USER $APP_USER
CMD ["./ingrid-coding-assignment", "serve"]
