FROM golang:alpine

RUN apk add curl

COPY ./src /app

WORKDIR  /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["air"]