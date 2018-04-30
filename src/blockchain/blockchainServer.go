package blockchain

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func Run() {
	http.HandleFunc("/blocks", getBlocks)
	http.HandleFunc("/mineBlock", http.HandlerFunc(mineBlock))
	http.HandleFunc("/addBlock", http.HandlerFunc(addBlock))
	addCreationPiece()
	http.ListenAndServe("127.0.0.1:28199", nil)
	log.Println("start blocks server")
	select {}
}

func getBlocks(w http.ResponseWriter, req *http.Request) {
	if blockList.Len() > 0 {
		var rbody []map[string]interface{}
		for p := blockList.Front(); p != nil; p = p.Next() {
			jsons, errs := json.Marshal(p.Value)
			if errs == nil {
				t := make(map[string]interface{})
				t["data"] = string(jsons)
				rbody = append(rbody, t)
			}
		}
		b, errs := json.Marshal(rbody)

		if errs == nil {
			send(w, string(b))
		} else {
			send(w, "{}")
		}
	} else {
		send(w, "{}")
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
