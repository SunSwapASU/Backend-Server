# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=project Version=0.0.1
# EXPOSE 3000


FROM golang:alpine AS builder
WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY ./ ./
# generate the Prisma Client Go client
RUN go run github.com/steebchen/prisma-client-go generate
# or, if you use go generate to run the generator, use the following line instead
# RUN go generate ./...

# build a fully standalone binary with zero dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

# use the scratch image for the smallest possible image size
FROM alpine

# copy over SSL certificates, so that we can make HTTPS requests
COPY --from=builder /workspace/app /app

EXPOSE 3000
ENTRYPOINT ["/app"]
