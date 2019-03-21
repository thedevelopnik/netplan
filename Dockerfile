FROM golang:1.12.1-alpine3.9

RUN apk add --update nodejs npm curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN npm i -g @vue/cli

WORKDIR /go/src/github.com/thedevelopnik/netplan

COPY . .

RUN npm i

RUN dep ensure

RUN npm run build

RUN go build .

RUN chmod +x netplan

RUN ls

CMD ./netplan
