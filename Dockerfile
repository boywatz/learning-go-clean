# สร้าง builder stage
FROM golang:alpine AS builder

# ติดตั้ง git
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/app
COPY . .

# สร้าง binary
RUN go get -d -v
RUN go build -o /go/bin/myapp

# สร้าง scratch stage
FROM scratch

# กำหนด port
EXPOSE 80

# คัดลอก binary จาก builder stage
COPY --from=builder /go/bin/myapp /go/bin/myapp

# รัน binary
ENTRYPOINT ["/go/bin/myapp"]
