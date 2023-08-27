FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./

# download dependencies and modules
RUN go mod download

COPY *.go ./

# build the go app
RUN go build -o /multi_arch_sample

# network port at runtime
EXPOSE 8000

CMD [ "/multi_arch_sample" ]