package util

import (
	"math/rand"
	"strings"
)

const (
	alphabet = "AaBbcDefghklmnoPqRSTOv"
)

func RandomStr(n int) string {
	var builder strings.Builder
	builder.Grow(n)
	k := len(alphabet)
	for i := 0; i < n; i++ {
		builder.WriteByte(alphabet[rand.Intn(k)])
	}
	return builder.String()
}
