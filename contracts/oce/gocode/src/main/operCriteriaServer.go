package main

import (
	"fmt"
	"log"
	"net/http"
	"math/big"
	"encoding/json"
	"OperContract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

type BinData struct {
	Name string
	Measurement string
	Base int
}

type BinDataResp struct {
	Result bool
}

var operContHdl *OperContract.OperCont

func main() {
	conn, err := ethclient.Dial("/home/sgr/blockchain/ins2/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	operContHdl, err = OperContract.NewOperCont(common.HexToAddress("0xeca033804febe945f1827bba2c5345cd3e1ff807"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/verify", VerifyRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func VerifyRequest(w http.ResponseWriter, r *http.Request) {

	var binReq BinData
	var measure big.Int

	err := json.NewDecoder(r.Body).Decode(&binReq)
	if err != nil {
		fmt.Println("VerifyRequest: json decoding failed")
		return
	}

	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}

	fmt.Println("decoded request",binReq)

	base := binReq.Base
	if base == 0 {
		base = 16
	}
	_,rv := measure.SetString(binReq.Measurement,binReq.Base)

	if !rv {
		fmt.Println("conversion failed, req",binReq.Measurement,"base",binReq.Base)
		return
	}

	fmt.Println("measure:",&measure)

	res,err := operContHdl.VerifyBinary(nil,binReq.Name,&measure)

	if err != nil {
		fmt.Println("error",err)
	}

	fmt.Println("result: ",res)

	binResp := BinDataResp{res}

	json.NewEncoder(w).Encode(binResp)
}
