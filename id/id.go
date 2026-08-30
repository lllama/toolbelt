package id

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/rzajac/zflake"
	"github.com/zeebo/xxh3"
)

var (
	flake     *zflake.Gen
	flakeOnce sync.Once
)

const filenameChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func NextID() int64 {
	flakeOnce.Do(func() {
		id, err := machineid.ID()
		if err != nil {
			id = time.Now().Format(time.RFC3339Nano)
		}
		h := xxh3.HashString(id) % (1 << zflake.BitLenGID)
		flake = zflake.NewGen(
			zflake.Epoch(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			zflake.GID(uint16(h)),
		)
	})
	return flake.NextFID()
}

func NextEncodedID() string {
	return EncodeID(NextID())
}

func EncodeID(id int64) string {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(id))
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(buf)
}

func EncodedIDToInt64(s string) (int64, error) {
	trimmed := strings.TrimRight(s, "=")
	buf, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(trimmed)
	if err != nil {
		return 0, err
	}
	if len(buf) != 8 {
		return 0, fmt.Errorf("encoded id must decode to 8 bytes, got %d", len(buf))
	}
	return int64(binary.LittleEndian.Uint64(buf)), nil
}

func AliasHash(alias string) int64 {
	return int64(xxh3.HashString(alias) & 0x7fffffffffffffff)
}

func AliasHashf(format string, args ...interface{}) int64 {
	return AliasHash(fmt.Sprintf(format, args...))
}

func AliasHashEncoded(alias string) string {
	h := AliasHash(alias)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(h))

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(buf)
}

func AliasHashEncodedf(format string, args ...interface{}) string {
	return AliasHashEncoded(fmt.Sprintf(format, args...))
}

func SafeFilename(str string) string {
	b := strings.Builder{}

	for _, char := range str {
		if strings.ContainsRune(filenameChars, char) {
			b.WriteRune(char)
		}
	}
	return b.String()
}

func SafeSubject(str string) string {
	b := strings.Builder{}

	for _, char := range str {
		if strings.ContainsRune(filenameChars, char) {
			b.WriteRune(char)
		}
		if char == ' ' {
			b.WriteRune('-')
		}
	}
	return b.String()
}
