package hash

import (
	"crypto/sha256"
	// "fmt"
	"io"
	"os"
	"strconv"
)

func Hash(blockID int) (string, error) {
	var ret string
	var err error

	// TODO: Get hash value of taken blockID
	f, err := os.Open("history.block." + strconv.Itoa(blockID))
	if err != nil {
		return ret, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return ret, err
	}
	ret = string(h.Sum(nil))
	return ret, err
}
