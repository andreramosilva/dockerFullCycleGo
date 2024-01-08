#ramosivlandre/fullcycle

#full cycle rocks

#Go Ola mundo

#imagem menos de 2MB

FROM golang:alpine as builder

WORKDIR /go/workdir/app

COPY . .

RUN CGO_ENABLED=0 go build -o /app main.go

FROM scratch 

COPY --from=builder /app /app

CMD ["/app"]