


#Stage 1: Builder Stage
FROM golang:alpine AS builder
WORKDIR /build

COPY ../go.mod .
RUN go mod download # giong voi npm install
#Biên dịch chương trình Go từ thư mục ./cmd/server, và tạo ra file thực thi có tên crm.shopdev.com đặt trong /build.
RUN go build -o crm.shopdev.com ./cmd/server
#Stage 2: Production Stage
#Sử dụng image scratch (trống hoàn toàn), không có hệ điều hành hoặc thư viện mặc định nào, giúp giảm dung lượng tối đa cho image.
#Thích hợp cho ứng dụng Go vì Go tạo ra file thực thi tĩnh.
FROM scratch

COPY ../config /config
#Sao chép chứng chỉ SSL từ builder sang container. Điều này rất cần thiết nếu ứng dụng cần thực hiện các kết nối HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#Sao chép file thực thi crm.shopdev.com từ builder vào thư mục gốc của container.
COPY --from=builder /build/crm.shopdev.com /
#Thiết lập lệnh khởi chạy container. Khi container bắt đầu, nó sẽ thực thi file crm.shopdev.com với tham số là config/local.yaml.
ENTRYPOINT ["/crm.shopdev.com", "config/local.yaml"]


