package blockchain

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

type blockListStr struct {
	BlockList []blockNode `json:"blockList"`
}

var blockList *list.List = list.New()

//增加创世块
func addCreationPiece() blockNode {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d %s %s %s %d", 0, "2819", "0.0.0.0", "root", 0)))
	cipherStr := h.Sum(nil)
	_md5 := hex.EncodeToString(cipherStr)
	_perr := peerData{
		"0.0.0.0",
		"root",
	}
	_blockNode := blockNode{
		0,
		_perr,
		0,
		_md5,
		_md5,
	}
	blockList.PushBack(_blockNode)
	return _blockNode
}

func addBlockNode(req *http.Request, perr peerData) blockNode {
	v, _ := blockList.Back().Value.(blockNode)

	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d %s %s %s %d", v.Index+1, v.PreviousHash, perr.IpAddress, perr.Name, time.Now().Unix())))
	cipherStr := h.Sum(nil)
	_md5 := hex.EncodeToString(cipherStr)

	_blockNode := blockNode{
		v.Index + 1,
		perr,
		time.Now().Unix(),
		_md5,
		v.PreviousHash,
	}
	blockList.PushBack(_blockNode)
	return _blockNode
}
