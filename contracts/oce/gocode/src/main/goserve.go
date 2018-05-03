package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"html/template"
	"math/big"
	"encoding/json"
	"github.com/gorilla/mux"
	"OperContract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type BinDataReq struct {
	Name string
	Measurement string
	Base int
}

type PageBinDataReq struct {
	Host string
	Name string
	Measurement string
	Base int
}

type BinDataResp struct {
	Name string
	Measurement string
	Base int
	Result bool
}

type HostDataReq struct {
	Name string
	Data []BinDataReq
}

type HostDataResp struct {
	Name string
	Data []BinDataResp
	Compliance int
}

type HostDataReqResp struct {
	Status bool
}

type Host struct {
	Name string
	BinMeasureMap map[string]string
	BinStatusMap map[string]bool
}

var hostMap map[string]*Host
var Title string

var operContHdl *OperContract.OperCont

var id int64 = 1697

type PageVariables struct {
	PageTitle string
	PageBinaryDetails []HostDataResp
	Compliance int
}

func main() {

	Title = "NXOS INTEGRITY STATION"

	hostMap = make(map[string]*Host)

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

	router.HandleFunc("/display", Index).Methods("GET")
//	router.HandleFunc("/display", DisplayDataPageHandler).Methods("POST")

	router.HandleFunc("/verifyPage", AddData).Methods("GET")
	router.HandleFunc("/verifyPage", VerifyRequestPage).Methods("POST")

	router.HandleFunc("/verify", VerifyRequest).Methods("POST")
	router.HandleFunc("/record", RecordRequest).Methods("POST")
	fmt.Println("Server deployed")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetAllData() (res []HostDataResp, compliance int) {
	num_total,num_compliant := 0,0

	var allData []HostDataResp
	for _,host := range hostMap {
		var hostData []BinDataResp
		host_num_total,host_num_compliant := 0,0

		num_total += len(host.BinMeasureMap)
		host_num_total = len(host.BinMeasureMap)

		for key,val := range host.BinMeasureMap {
			hostData = append(hostData, BinDataResp{key,val,0,host.BinStatusMap[key]})
			if host.BinStatusMap[key] {
				num_compliant++
				host_num_compliant++
			}
		}
		allData = append(allData,HostDataResp{host.Name,hostData,host_num_compliant*100/host_num_total})
	}

	if num_total == 0 {
		return nil,0
	}

	return allData,num_compliant*100/num_total
}


func parsePage(htmlFile string,MyPageVariables PageVariables, w http.ResponseWriter) {
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	pageBinaryDetails,compliance := GetAllData()

	MyPageVariables := PageVariables{Title, pageBinaryDetails,compliance}

	parsePage("Index.html",MyPageVariables,w)

	fmt.Println("Get request Recieved")
}

func AddData(w http.ResponseWriter, r *http.Request) {

	pageBinaryDetails,compliance := GetAllData()

	MyPageVariables := PageVariables{Title, pageBinaryDetails,compliance}

	parsePage("AddData.html",MyPageVariables,w)

	fmt.Println("Get request Recieved")
}

func VerifyRequest(w http.ResponseWriter, r *http.Request) {

	var hostReq HostDataReq
	var measure big.Int

	fmt.Println("POST request Recieved")
	err := json.NewDecoder(r.Body).Decode(&hostReq)
	if err != nil {
		fmt.Println("VerifyRequest: json decoding failed",err)
		return
	}

	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}

	fmt.Println("decoded request",hostReq)

	hostName := hostReq.Name

	host,ok := hostMap[hostName] 
	if !ok {
		host = &Host{
			Name:hostName,
			BinMeasureMap: make(map[string]string),
			BinStatusMap: make(map[string]bool)}
		hostMap[hostName] = host
	}

	for _,binReq := range hostReq.Data {
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

		host.BinMeasureMap[binReq.Name] = measure.String()
		host.BinStatusMap[binReq.Name] = res
	}

	fmt.Println(hostMap)

	json.NewEncoder(w).Encode(HostDataReqResp{true})
}


func RecordRequest(w http.ResponseWriter, r *http.Request) {

	var hostReq HostDataReq
	var measure big.Int

	key := "{\"address\":\"bbfd141b3daaffe1bbfe34e999c85fc7ef264efb\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"44b9ba7f40c9b417ea29ed511457f1515ab948aa94663d58e71a2b17c4da3323\",\"cipherparams\":{\"iv\":\"475936860cc1f556c68aad8d27668fbb\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"cde65553b78fd8f06f90539a5003f085a9b66a7c98a1aee70104579aa23030cd\"},\"mac\":\"3e965e2216728391ee7ab97b4180852c0d20bad33d9a55eaab89434c0fd6d349\"},\"id\":\"87499b4e-cf69-41a8-99e0-602064539d52\",\"version\":3}"

	auth, err := bind.NewTransactor(strings.NewReader(key), "ins2")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	fmt.Println("POST request Recieved")
	err = json.NewDecoder(r.Body).Decode(&hostReq)
	if err != nil {
		fmt.Println("VerifyRequest: json decoding failed",err)
		return
	}

	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}

	fmt.Println("decoded request",hostReq)

	for _,binReq := range hostReq.Data {
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

		// Wrap the Token contract instance into a session
		session := &OperContract.OperContSession{
			Contract: operContHdl,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: big.NewInt(18000000000*1.5),
				GasLimit: 100000,
				Nonce : big.NewInt(id),
			},
		}

		id++

		tx,err := session.AddBinary(binReq.Name,&measure)
		if err != nil {
			fmt.Println("error while recording",err)
			json.NewEncoder(w).Encode(HostDataReqResp{false})
			return
		}
		fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())

	}

	fmt.Println(hostMap)

	json.NewEncoder(w).Encode(HostDataReqResp{true})
}

func VerifyRequestPage (w http.ResponseWriter, r *http.Request) {

	var binReq PageBinDataReq
	var measure big.Int

	fmt.Println("POST request Recieved")
	r.ParseForm()

	binReq = PageBinDataReq{r.Form.Get("hostName"),r.Form.Get("binName"),r.Form.Get("binMeasure"),10}

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

	res := true

	hostName := binReq.Host

	host,ok := hostMap[hostName] 
	if !ok {
		host = &Host{
			Name:hostName,
			BinMeasureMap: make(map[string]string),
			BinStatusMap: make(map[string]bool)}
		hostMap[hostName] = host
	}

	host.BinMeasureMap[binReq.Name] = measure.String()
	host.BinStatusMap[r.Form.Get("binName")] = res

	pageBinaryDetails, compliance := GetAllData()

	MyPageVariables := PageVariables {Title,pageBinaryDetails,compliance}

	parsePage("Index.html",MyPageVariables,w)
}

func DisplayDataPageHandler(w http.ResponseWriter, r *http.Request) {
	var pageBinaryDetails []HostDataResp
	var compliance int

	r.ParseForm()

	pageBinaryDetails, compliance = GetAllData()


	MyPageVariables := PageVariables {Title,pageBinaryDetails,compliance}

	parsePage("Index.html",MyPageVariables,w)

}
