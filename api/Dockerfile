FROM registry.access.redhat.com/ubi8/go-toolset:1.18.9-4

USER root
WORKDIR /tests
COPY . /tests

RUN go get -d ./...

RUN go build -o /build /tests/cmd/server/main.go

EXPOSE 8000

CMD [ "/build" ]
