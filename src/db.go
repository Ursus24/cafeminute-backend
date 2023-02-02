package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var crypto = []byte{}

var (
	newFile *os.File
	err     error
)

func readKey(key string, dir string) string {
	if hasKey(key, dir) {
		file, err := os.Open(dir + "/" + key)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		ret := (string(data))
		result, err := DecryptGCM(crypto, []byte(ret))
		if err != nil {
			fmt.Println(err)
		}
		return string(result)
	}
	return ""
}

func addKey(key string, value string, dir string) {
	ciphertext, err := EncryptGCM(crypto, []byte(value))
	if err != nil {
		fmt.Println(err)
	}

	if !hasKey(key, dir) {
		os.Mkdir(dir, os.ModePerm)
		newFile, err = os.Create(dir + "/" + key)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newFile)
		newFile.Close()
		err2 := ioutil.WriteFile(dir+"/"+key, []byte(ciphertext), 0666)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

}
func removeKey(key string, dir string) {
	if hasKey(key, dir) {
		err := os.Remove(dir + "/" + key)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func hasKey(key string, dir string) bool {

	_, err := os.Stat(dir + "/" + key)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func addKeyUnsafe(key string, value string, dir string) {
	if err != nil {
		fmt.Println(err)
	}

	if !hasKey(key, dir) {
		os.Mkdir(dir, os.ModePerm)
		newFile, err = os.Create(dir + "/" + key)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newFile)
		newFile.Close()
		err2 := ioutil.WriteFile(dir+"/"+key, []byte(value), 0666)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

}
func readKeyUnsafe(key string, dir string) string {
	if hasKey(key, dir) {
		file, err := os.Open(dir + "/" + key)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		ret := (string(data))
		if err != nil {
			fmt.Println(err)
		}
		return ret
	}
	return ""
}
