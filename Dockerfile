FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avitoTechInternship ./cmd/main/app.go

CMD ["./avitoTechInternship"]
