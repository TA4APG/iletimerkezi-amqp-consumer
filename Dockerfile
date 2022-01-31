############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS build
WORKDIR /go/build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/program .

############################
# STEP 2 build a small image
############################
FROM scratch
COPY --from=build /go/bin/program /program
COPY *.yml .
COPY certs/ certs/
CMD ["/program"]