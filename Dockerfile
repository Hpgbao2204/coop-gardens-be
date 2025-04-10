# Sử dụng Golang với Alpine để build nhẹ hơn
FROM golang:alpine AS builder

# Đặt thư mục làm việc
WORKDIR /app

# Copy toàn bộ source code vào container
COPY . .

# Lấy dependencies và build ứng dụng
RUN go mod tidy && go build -o main main.go

# Đảm bảo file có quyền thực thi
RUN chmod +x main

# Chạy ứng dụng
CMD ["/app/main"]
