

FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN rm -rf /usr/src/app/.git

# for Devlopment with hot reload
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
#
# for Production
# RUN go build -buildvcs=false -v -o /usr/local/bin/app ./api
#
# CMD ["app"]
