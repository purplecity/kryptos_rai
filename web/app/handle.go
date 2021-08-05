package app

import (
	"encoding/json"
	"io/ioutil"
	"kryptos_rai/use/logger"
	"net/http"
)

func HandlePoolInfo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqStruct := new(OperatePoolInfoReq)
	readBytes, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal([]byte(readBytes), reqStruct)
	logger.Info("receive %+v  HandlePoolInfo request,param is %+v", req.RemoteAddr, *reqStruct)
	resp := handlePoolInfo(reqStruct)
	ret(w, http.StatusOK, resp)
}

func AddressAnalysis(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqStruct := new(AddressAnalysisReq)
	readBytes, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal([]byte(readBytes), reqStruct)
	logger.Info("receive %+v  AddressAnalysis request,param is %+v", req.RemoteAddr, *reqStruct)
	resp := addressAnalysis(reqStruct)
	ret(w, http.StatusOK, resp)
}

func HandleLogoInfo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqStruct := new(LogoReq)
	readBytes, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal([]byte(readBytes), reqStruct)
	logger.Info("receive %+v  HandleLogoInfo request,param is %+v", req.RemoteAddr, *reqStruct)
	resp := handleLogoInfo(reqStruct)
	ret(w, http.StatusOK, resp)
}

func PoolComboBox(w http.ResponseWriter, req *http.Request) {
	resp := poolComboBox()
	ret(w, http.StatusOK, resp)
}

func LogoComboBox(w http.ResponseWriter, req *http.Request) {
	resp := logoComboBox()
	ret(w, http.StatusOK, resp)
}

func HandleSubcribe(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqStruct := new(SubscribeInfoReq)
	readBytes, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal([]byte(readBytes), reqStruct)
	logger.Info("receive %+v  HandleSubcribe request,param is %+v", req.RemoteAddr, *reqStruct)
	resp := handleSubcribe(reqStruct)
	ret(w, http.StatusOK, resp)
}
