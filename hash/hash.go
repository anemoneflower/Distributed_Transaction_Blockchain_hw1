package hash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
	"os"
)

func Hash(blockID int) (string, error){
	var ret string
	var err error

	// TODO: Get hash value of taken blockID
	//	f, err := os.Open("history.block."+strconv.Itoa(blockID))
	//	defer f.Close()
	return ret, err
}