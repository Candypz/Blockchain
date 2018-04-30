package blockchain

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type clientData struct {
	Name string `json:"name"`
}

type peerData struct {
	IpAddress string `json:"ipAddress"`
	Name      string `json:"name"`
}

type blockNode struct {
	Index        int      `json:"index"`
	Data         peerData `json:"data"`
	Timestamp    int64    `json:"timestamp"`
	Hash         string   `json:"hash"`
	PreviousHash string   `json:"previousHash"`
}

func send(w http.ResponseWriter, info string) {
	w.Write([]byte(info))
	log.Println(info)
}

func checkRequestValue(w http.ResponseWriter, req *http.Request) (bool, peerData) {
	keys, ok := req.URL.Query()["key"]
	_data := peerData{}
	if !ok || len(keys) < 1 {
		send(w, "Request value error")
		return false, _data
	}
	log.Println(keys[0])
	_clientData := clientData{}
	err := json.Unmarshal([]byte(keys[0]), &_clientData)
	if err != nil {
		send(w, "Request value error")
		return false, _data
	}

	_data.IpAddress = req.RemoteAddr
	_data.Name = _clientData.Name
	return true, _data
}

func addBlockNode(req *http.Request, perr peerData) blockNode {
	_peerData := perr
	_blockNode := blockNode{
		1,
		_peerData,
		time.Now().Unix(),
		"ss",
		"22",
	}
	return _blockNode
}

func Run() {
	http.HandleFunc("/blocks", getBlocks)
	http.HandleFunc("/mineBlock", http.HandlerFunc(mineBlock))
	http.HandleFunc("/addBlock", http.HandlerFunc(addBlock))
	http.ListenAndServe("127.0.0.1:28199", nil)
	log.Println("start blocks server")
	select {}
}

func getBlocks(w http.ResponseWriter, req *http.Request) {
	status, _peerData := checkRequestValue(w, req)
	if !status {
		return
	}

	_node := addBlockNode(req, _peerData)
	jsons, errs := json.Marshal(_node)
	if errs == nil {
		send(w, string(jsons))
	} else {
		send(w, "getBlocks error")
	}
}

func addBlock(w http.ResponseWriter, req *http.Request) {
	status, _peerData := checkRequestValue(w, req)
	if !status {
		return
	}

	_node := addBlockNode(req, _peerData)
	jsons, errs := json.Marshal(_node)
	if errs == nil {
		send(w, string(jsons))
	} else {
		send(w, "addBlock error")
	}
}

func mineBlock(w http.ResponseWriter, req *http.Request) {
	status, _peerData := checkRequestValue(w, req)
	if !status {
		return
	}

	_node := addBlockNode(req, _peerData)
	jsons, errs := json.Marshal(_node)
	if errs == nil {
		send(w, string(jsons))
	} else {
		send(w, "mineBlock error")
	}
}
