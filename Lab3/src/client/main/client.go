package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)

/*
Considering ID space of 128. Index based.
*/
var Circle [128]string

type Pair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type Server struct {
	HostName string
	Key      int
}

var Servers []Server
var serverMap map[int]Server
var Data []Pair
var serverKeys []int

func main() {

	serverMap = make(map[int]Server)

	Servers = append(Servers, Server{"http://localhost:3000", 0})
	Servers = append(Servers, Server{"http://localhost:3001", 0})
	Servers = append(Servers, Server{"http://localhost:3002", 0})

	for _, server := range Servers {
		addServer(&server)
	}
	Data = append(Data, Pair{1, "a"})
	Data = append(Data, Pair{2, "b"})
	Data = append(Data, Pair{3, "c"})
	Data = append(Data, Pair{4, "d"})
	Data = append(Data, Pair{5, "e"})
	Data = append(Data, Pair{6, "f"})
	Data = append(Data, Pair{7, "g"})
	Data = append(Data, Pair{8, "h"})
	Data = append(Data, Pair{9, "i"})
	Data = append(Data, Pair{10, "j"})

	for _, data := range Data {
		addElement(data)
	}

	fmt.Println("\n==================")
	getElement(3)
}

func addServer(server *Server) {
	hash := hash(server.HostName)
	server.Key = hash
	Circle[hash] = server.HostName
	serverKeys = append(serverKeys, hash)
	serverMap[hash] = *server
	fmt.Println(serverKeys)
}

func addElement(pair Pair) {
	key := pair.Key
	hash := hash(strconv.Itoa(key))
	i := hash
	for !contains(i, serverKeys) {
		i = i + 1
		if i == 128 {
			i = 0
		}
	}
	fmt.Println("Found server:", Circle[i])
	putKeyOnServer(i, pair)
	Circle[hash] = strconv.Itoa(pair.Key)
}

func getElement(key int) {
	hash := hash(strconv.Itoa(key))
	i := hash
	for !contains(i, serverKeys) {
		i = i + 1
		if i == 128 {
			i = 0
		}
	}
	fmt.Println("Found server:", Circle[i])
	getKeyOnServer(i, key)
}

func putKeyOnServer(serverKey int, pair Pair) {
	for _, server := range serverMap {
		if server.Key == serverKey {
			url := server.HostName + "/keys/" + strconv.Itoa(pair.Key) + "/" + pair.Value
			fmt.Println("PUT URL: %s", url)
			req, _ := http.NewRequest("POST", url, nil)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			break
		}
	}
}
func getKeyOnServer(serverKey int, key int) {
	for _, server := range serverMap {
		if server.Key == serverKey {
			url := server.HostName + "/keys/" + strconv.Itoa(key)
			fmt.Println("GET URL: %s", url)
			req, _ := http.NewRequest("GET", url, nil)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("GET Response:", string(body))
			defer resp.Body.Close()
			break
		}
	}
}
func contains(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func hash(key string) int {
	bi := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(key))
	hexstr := hex.EncodeToString(h.Sum(nil))
	bi.SetString(hexstr, 16)
	hash := bi.Int64() % 128
	if hash < 0 {
		hash = -hash
	}
	fmt.Println("hash", hash)
	return int(hash)
}
