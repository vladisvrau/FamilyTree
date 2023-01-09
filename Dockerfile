FROM golang:1.19.4 as builder
COPY . /FamilyTree
WORKDIR /FamilyTree

RUN go build -gcflags=all=-l -o api main.go

FROM gcr.io/distroless/base
COPY --from=builder /FamilyTree/api /bin/api
EXPOSE 3000
CMD ["/bin/api"]