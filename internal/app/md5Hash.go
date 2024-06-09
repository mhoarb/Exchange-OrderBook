package app

import (
	"Exchange-OrderBook/domain/orderBook"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log/slog"
)

func Md5Hash(order orderBook.OrderBook) {
	hashes := md5.New()
	orderByte := []byte(fmt.Sprintf("%+v", order))
	hashes.Write(orderByte)
	hash := hashes.Sum(nil)
	hashHex := hex.EncodeToString(hash)
	slog.Info("md5Hash generate", "hashhex", hashHex)
}
