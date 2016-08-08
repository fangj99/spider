package spider

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/big"
	"math/rand"
	"time"
)

//ID define
type ID []byte

func (id ID) String() string {
	return hex.EncodeToString(id)
}

//Int get int
func (id ID) Int() *big.Int {
	return big.NewInt(0).SetBytes(id)
}

//Neighbor get neighbor
func (id ID) Neighbor(tableID ID) ID {
	return append(id[:6], tableID[6:]...)
}

//Neightor get neighbor
func Neightor(id, tableID string) string {
	return id[:6] + tableID[6:]
}

//GenerateID get id
func GenerateID() ID {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := sha1.New()
	io.WriteString(hash, time.Now().String())
	io.WriteString(hash, string(random.Int()))
	return hash.Sum(nil)
}

//GenerateIDList for uniform
func GenerateIDList(count int64) (ids []ID) {
	if count <= 0 {
		return
	}
	max := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	num := big.NewInt(0).SetBytes(max)
	step := big.NewInt(0).Div(num, big.NewInt(count+2))
	for i := 1; i < int(count+1); i++ {
		item := big.NewInt(0).Mul(step, big.NewInt(int64(i)))
		item.Add(item, big.NewInt(int64(rand.Intn(99))))
		ids = append(ids, item.Bytes())
	}
	return
}
