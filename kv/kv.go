package kv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"../history"
)

type DatabaseInterface interface {
	GetState(key string) (string, error)
	PutState(key, value string) error
}

type Database struct {
	state     KV              //TODO: state DB
	tempBlock history.History //TODO: history of transactions
}

type KV map[string]string

func (db Database) GetState(key string) (string, error) {
	var ret string
	var err error

	// TODO: Get the value of an input key and return error if it has
	ret, ok := db.state[key]
	if !ok {
		err = errors.New("existing input key")
	}

	return ret, err
}

func (db *Database) PutState(key, value string) error {
	var err error
	// TODO: Put a key-value pair into state DB and return error if it has
	_, ok := db.state[key]
	if ok {
		// question: error generate?
	} else {
		db.state[key] = value
	}

	// TODO: Append history into temp block
	err = db.tempBlock.Append(key + "\t" + value)

	return err
}

func Init(obj *Database) {
	obj.state = make(KV)

	// TODO: Load state DB
	statePath := "state.db"
	data, err := ioutil.ReadFile(statePath)
	if err != nil {
		// fmt.Println("kviniterr1")
		// fmt.Print(err)
		return
	}
	datas := strings.Split(string(data), "\n")
	for _, s := range datas {
		if s == "" {
			// fmt.Println("s is empty")
			continue
		}
		ss := strings.Split(s, "\t")
		obj.state[ss[0]] = ss[1]
	}

	// TODO: Initialize history of transactions
	err = obj.tempBlock.Init()
	if err != nil {
		fmt.Println("kviniterr")
		fmt.Print(err)
	}
}

func Finalize(obj *Database) {
	// TODO: Store current state DB
	statePath := "state.db"
	s := ""
	// err := ioutil.WriteFile(statePath, []byte(s), 0644)
	for key, value := range obj.state {
		s = s + key + "\t" + value + "\n"
	}
	err := ioutil.WriteFile(statePath, []byte(s), 0644)
	if err != nil {
		fmt.Println("kvfinalizeerr1")
		fmt.Print(err)
	}
	// TODO: Store tempBlock
	err = obj.tempBlock.Write()
	if err != nil {
		fmt.Println("kvfinalizeerr2")
		fmt.Print(err)
	}
}
