echo "protoc payment"
protoc --proto_path=/home/anth/go/src/uit_payment/proto/payment --go_out=plugins=grpc:/home/anth/go/src/uit_payment/services/payment/ /home/anth/go/src/uit_payment/proto/payment/payment.proto

echo "protoc dom"
protoc --proto_path=/home/anth/go/src/uit_payment/proto/dom --go_out=plugins=grpc:/home/anth/go/src/uit_payment/services/dom/ /home/anth/go/src/uit_payment/proto/dom/dom.proto
