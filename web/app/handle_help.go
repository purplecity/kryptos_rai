package app

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	internal_data "kryptos_rai/block_parse/simple_internal_tx/bsc/logic"
	logic "kryptos_rai/block_parse/simple_tx/bsc/logic"
	"kryptos_rai/use/const_def"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	analysis "kryptos_rai/use/pb/address_analysis"
	"kryptos_rai/web/config"
	"kryptos_rai/web/rpc"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ret(w http.ResponseWriter, code int, resp []byte) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(resp)
}

//删改查所用
func handleLogoReq(req *LogoReq) bson.M {
	t := reflect.TypeOf(*req)
	v := reflect.ValueOf(*req)
	s := bson.M{}
	for k := 0; k < t.NumField(); k++ {
		switch t.Field(k).Type.Kind() {
		case reflect.String:
			if v.Field(k).String() != "" {
				key := string(t.Field(k).Tag.Get("json"))
				if key == "token_address" || key == "name" {
					s[key] = strings.ToLower(v.Field(k).String())
				} else {
					s[key] = v.Field(k).String()
				}

			}
		default:
			//
		}
	}

	return s
}

func checkLogoReq(req *LogoReq) error {

	if req.TokenAddress != "" {
		if !strings.HasPrefix(req.TokenAddress, "0x") {
			logger.Error("checkLogoReq::TokenAddress error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	reg := regexp.MustCompile("^[A-Za-z]+$")
	if !reg.MatchString(req.Name) {
		logger.Error("checkLogoReq:name only a-zA-Z")
		return errors.New(ErrorMap[InvalidArgument])
	}

	chainInfo := config.GetChainCheck()
	poolType := config.GetLogoType()
	if req.Chain != "" {
		_, ok := chainInfo[req.Chain]
		if !ok {
			logger.Error("CheckLogoReq::chain error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	if req.LogoType != "" {
		_, ok := poolType[req.LogoType]
		if !ok {
			logger.Error("CheckLogoReq::logotype error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	switch req.Method {
	case const_def.ADD:
		//因为没有池子id那样的单个字段决定为一条一条记录
		if req.Name == "" || req.URL == "" || req.LogoType == "" || req.Chain == "" {
			logger.Error("CheckLogoReq::ADD Method Must contain name,url,logotype,chain")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.EDIT:
		if req.ID == "" {
			logger.Error("CheckLogoReq::EDIT Method Must contain  id")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.QUERY:
		if req.Page < 1 {
			logger.Error("CheckLogoReq::QUERY Method page Must gte 1")
			return errors.New(ErrorMap[InvalidArgument])
		}
	case const_def.DELETE:
		if req.ID == "" {
			logger.Error("CheckLogoReq::DELETE Method Must contain  id")
			return errors.New(ErrorMap[InvalidArgument])
		}

	default:
		logger.Error("CheckLogoReq::invalid method")
		return errors.New(ErrorMap[InvalidArgument])
	}

	return nil
}

func handleLogoInfo(req *LogoReq) []byte {
	response := Response{}
	(&response).ResponseInit()
	err := checkLogoReq(req)
	if err != nil {
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
		d, _ := json.Marshal(response)
		return d
	}

	switch req.Method {
	case const_def.ADD:
		//唯一性
		m := handleLogoReq(req)
		delete(m, "method")
		delete(m, "page")
		delete(m, "_id")
		u := bson.M{}

		u["chain"] = m["chain"]
		u["type"] = m["type"]
		u["name"] = m["name"]
		u["token_address"] = strings.ToLower(req.TokenAddress)

		count, err := mongo.LogoMongoClient.Find(context.Background(), u).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo::Find logo failed:%+v,u :%+v", err, u)
		} else if count > 0 {
			response.Code = LogoExist
			response.Msg = ErrorMap[LogoExist]
			logger.Error("handleLogoInfo::LogoExist already Exist:%+v,u :%+v", err, u)
		} else {
			addInfo := Logo{
				Name:         strings.ToLower(req.Name),
				URL:          req.URL,
				LogoType:     req.LogoType,
				TokenAddress: strings.ToLower(req.TokenAddress),
				Chain:        req.Chain,
				UpdateTime:   time.Now().Unix(),
			}
			_, err := mongo.LogoMongoClient.InsertOne(context.Background(), addInfo)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handleLogoInfo::InsertOne failed:%+v,data is :%+v", err, addInfo)
			}

			key := req.Chain + "_" + req.LogoType + "_" + strings.ToLower(req.Name) + "_" + strings.ToLower(req.TokenAddress)
			logoMap.Set(key, addInfo)
		}

	case const_def.EDIT:

		//[12]byte
		tmpBytes, terr := hex.DecodeString(req.ID)
		if terr != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo::hex.DecodeString :%+v", err)
		}
		idBytes := [12]byte{}
		copy(idBytes[:], tmpBytes)
		count, err := mongo.LogoMongoClient.Find(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)}).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo:: Find failed:%+v,ID :%+v", err, req.ID)
		} else if count < 1 {
			response.Code = LogoNotExist
			response.Msg = ErrorMap[LogoNotExist]
			logger.Error("handleLogoInfo::LogoNotExist not Exist:%+v,ID :%+v", err, req.ID)
		} else {
			m := handleLogoReq(req)
			m["update_time"] = time.Now().Unix()
			delete(m, "method")
			delete(m, "page")
			delete(m, "_id")
			u := bson.M{}
			u["$set"] = m

			err := mongo.LogoMongoClient.UpdateOne(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)}, u)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handleLogoInfo::UpdateOne failed:%+v,data is :%+v", err, m)
			}

			//没写事物了
			logoInfo := Logo{}
			mongo.LogoMongoClient.Find(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)}).One(&logoInfo)
			key := logoInfo.Chain + "_" + logoInfo.LogoType + "_" + strings.ToLower(logoInfo.Name) + "_" + strings.ToLower(logoInfo.TokenAddress)
			logoMap.Replace(key, logoInfo)
		}

	case const_def.QUERY:

		m := handleLogoReq(req)
		delete(m, "method")
		delete(m, "page")
		delete(m, "_id")

		data := []Logo{}
		count, err := mongo.LogoMongoClient.Find(context.Background(), m).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo::Get Count failed:%+v,data is :%+v", err, m)
		} else {
			err := mongo.LogoMongoClient.Find(context.Background(), m).Sort("-update_time").Skip((req.Page - 1) * 10).Limit(10).All(&data)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handleLogoInfo::Find failed:%+v,data is :%+v", err, m)
			}
		}
		retData := map[string]interface{}{}
		retData["Total"] = count
		if len(data) > 0 {
			retData["LogoList"] = data
		}
		response.Data = retData

	case const_def.DELETE:

		//[12]byte
		tmpBytes, terr := hex.DecodeString(req.ID)
		if terr != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo::hex.DecodeString :%+v", err)
		}
		idBytes := [12]byte{}
		copy(idBytes[:], tmpBytes)
		count, err := mongo.LogoMongoClient.Find(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)}).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleLogoInfo:: Find failed:%+v,ID :%+v", err, req.ID)
		} else if count < 1 {
			response.Code = LogoNotExist
			response.Msg = ErrorMap[LogoNotExist]
			logger.Error("handleLogoInfo::LogoNotExist not Exist:%+v,ID :%+v", err, req.ID)
		} else {

			logoInfo := Logo{}
			mongo.LogoMongoClient.Find(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)}).One(&logoInfo)
			key := logoInfo.Chain + "_" + logoInfo.LogoType + "_" + strings.ToLower(logoInfo.Name) + "_" + strings.ToLower(logoInfo.TokenAddress)

			err := mongo.LogoMongoClient.Remove(context.Background(), bson.M{"_id": primitive.ObjectID(idBytes)})
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handleLogoInfo::Remove failed:%+v,ID is :%+v", err, req.ID)
			}

			logoMap.Delete(key)
		}
	default:
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
	}
	d, _ := json.Marshal(response)
	return d
}

//删改查所用
func handleOperatePoolInfoReq(req *OperatePoolInfoReq) bson.M {
	t := reflect.TypeOf(*req)
	v := reflect.ValueOf(*req)
	s := bson.M{}
	for k := 0; k < t.NumField(); k++ {
		switch t.Field(k).Type.Kind() {
		case reflect.String:
			if v.Field(k).String() != "" {
				key := string(t.Field(k).Tag.Get("json"))
				if key == "pool_address" {
					s[key] = strings.ToLower(v.Field(k).String())
				} else {
					s[key] = v.Field(k).String()
				}

			}
		case reflect.Slice:
			if v.Field(k).Len() != 0 {
				key := string(t.Field(k).Tag.Get("json"))
				s[key] = v.Field(k).Interface()
			}
		default:
			//
		}
	}

	return s
}

func checkOperatePoolInfoReq(req *OperatePoolInfoReq) error {
	reg := regexp.MustCompile("^[1-9][0-9]*$")
	req.InputTokenNameList = []string{}
	if len(req.InputTokenList) > 0 {
		ts := []string{}

		for _, e := range req.InputTokenList {
			s := strings.Split(e, "=")
			if len(s) != 3 {
				logger.Error("CheckOperatePoolInfoReq:input token format:length error")
				return errors.New(ErrorMap[InvalidArgument])
			}
			if !strings.HasPrefix(s[1], "0x") {
				logger.Error("CheckOperatePoolInfoReq:input token format:address error")
				return errors.New(ErrorMap[InvalidArgument])
			}

			if !reg.MatchString(s[2]) {
				logger.Error("CheckOperatePoolInfoReq:input token format:precision error")
				return errors.New(ErrorMap[InvalidArgument])
			}

			num, _ := strconv.Atoi(s[2])
			if num < 2 {
				logger.Error("CheckOperatePoolInfoReq:input token format:precision error")
				return errors.New(ErrorMap[InvalidArgument])
			}
			ts = append(ts, strings.ToLower(s[0]))
		}
		req.InputTokenNameList = ts
	}

	if len(req.RewardTokenList) > 0 {

		for _, e := range req.RewardTokenList {
			s := strings.Split(e, "=")
			if len(s) != 3 {
				logger.Error("CheckOperatePoolInfoReq:reward token format:length error")
				return errors.New(ErrorMap[InvalidArgument])
			}
			if !strings.HasPrefix(s[1], "0x") {
				logger.Error("CheckOperatePoolInfoReq:reward token format:address error")
				return errors.New(ErrorMap[InvalidArgument])
			}

			if !reg.MatchString(s[2]) {
				logger.Error("CheckOperatePoolInfoReq:reward token format:precision error")
				return errors.New(ErrorMap[InvalidArgument])
			}

			num, _ := strconv.Atoi(s[2])
			if num < 2 {
				logger.Error("CheckOperatePoolInfoReq:reward token format:precision error")
				return errors.New(ErrorMap[InvalidArgument])
			}
		}
	}

	if req.CertificateToken != "" {
		s := strings.Split(req.CertificateToken, "=")
		if len(s) != 3 {
			logger.Error("CheckOperatePoolInfoReq:CertificateToken  format:length error")
			return errors.New(ErrorMap[InvalidArgument])
		}
		if !strings.HasPrefix(s[1], "0x") {
			logger.Error("CheckOperatePoolInfoReq:CertificateToken  format:address error")
			return errors.New(ErrorMap[InvalidArgument])
		}

		if !reg.MatchString(s[2]) {
			logger.Error("CheckOperatePoolInfoReq:CertificateToken  format:precision error")
			return errors.New(ErrorMap[InvalidArgument])
		}

		num, _ := strconv.Atoi(s[2])
		if num < 2 {
			logger.Error("CheckOperatePoolInfoReq:CertificateToken  format:precision error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	if len(req.FundFlowList) > 0 {
		for _, x := range req.FundFlowList {
			s := strings.Split(x, "=")
			if len(s) != 2 {
				logger.Error("CheckOperatePoolInfoReq:FundFlowList  format:length error")
				return errors.New(ErrorMap[InvalidArgument])
			}
			if !strings.HasPrefix(s[1], "0x") {
				logger.Error("CheckOperatePoolInfoReq:FundFlowList  format:address error")
				return errors.New(ErrorMap[InvalidArgument])
			}
		}
	}

	if req.PoolAddress != "" {
		if !strings.HasPrefix(req.PoolAddress, "0x") {
			logger.Error("CheckOperatePoolInfoReq:Pooladdress error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	platformInfo := config.GetPlatformCheck()
	chainInfo := config.GetChainCheck()
	poolPropertyInfo := config.GetPoolProperty()
	if req.Chain != "" {
		_, ok := chainInfo[req.Chain]
		if !ok {
			logger.Error("CheckOperatePoolInfoReq:chain error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	if req.Platform != "" {
		_, ok := platformInfo[req.Platform]
		if !ok {
			logger.Error("CheckOperatePoolInfoReq:platform error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	if req.PoolProperty != "" {
		_, ok := poolPropertyInfo[req.PoolProperty]
		if !ok {
			logger.Error("CheckOperatePoolInfoReq:property error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	switch req.Method {
	case const_def.ADD:
		if (req.IsWatching != const_def.TRUE && req.IsWatching != const_def.FALSE) || req.PoolAddress == "" || req.PoolProperty == "" {
			logger.Error("CheckOperatePoolInfoReq:ADD Method Must contain iswatching and pooladdress and poolproperty")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.EDIT:
		if req.PoolAddress == "" {
			logger.Error("CheckOperatePoolInfoReq:EDIT Method Must contain  pooladdress")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.QUERY:
		if req.Page < 1 {
			logger.Error("CheckOperatePoolInfoReq:QUERY Method page Must gte 1")
			return errors.New(ErrorMap[InvalidArgument])
		}
	case const_def.DELETE:
		if req.PoolAddress == "" {
			logger.Error("CheckOperatePoolInfoReq:DELETE Method Must contain  pooladdress")
			return errors.New(ErrorMap[InvalidArgument])
		}
	default:
		logger.Error("CheckOperatePoolInfoReq:invalid method")
		return errors.New(ErrorMap[InvalidArgument])
	}

	return nil
}

//删改查所用
//pool_address是直接合约交互的地址 比如pancake pair合约 凭证代币合约pool_address是一样的
func handlePoolInfo(req *OperatePoolInfoReq) []byte {
	response := Response{}
	(&response).ResponseInit()
	err := checkOperatePoolInfoReq(req)
	if err != nil {
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
		d, _ := json.Marshal(response)
		return d
	}

	switch req.Method {
	case const_def.ADD:
		count, err := mongo.MongoClient.Find(context.Background(), bson.M{"pool_address": req.PoolAddress}).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handlePoolInfo::Find pooladdress failed:%+v,pooladdress :%+v", err, req.PoolAddress)
		} else if count > 0 {
			response.Code = PoolExist
			response.Msg = ErrorMap[PoolExist]
			logger.Error("handlePoolInfo::PoolExist already Exist:%+v,pooladdress :%+v", err, req.PoolAddress)
		} else {
			m := handleOperatePoolInfoReq(req)
			m["update_time"] = time.Now().Unix()
			delete(m, "method")
			delete(m, "page")
			delete(m, "query_token_name")

			poolInfo := PoolInfo{
				Platform:           req.Platform,
				Chain:              req.Chain,
				PoolAddress:        strings.ToLower(req.PoolAddress),
				PoolProperty:       req.PoolProperty,
				IsWatching:         req.IsWatching,
				WatchDescribe:      req.WatchDescribe,
				InputTokenList:     req.InputTokenList,
				CertificateToken:   req.CertificateToken,
				FundFlowList:       req.FundFlowList,
				InputTokenNameList: req.InputTokenNameList, //handled
				RewardTokenList:    req.RewardTokenList,
				UpdateTime:         time.Now().Unix(),
			}

			_, err := mongo.MongoClient.InsertOne(context.Background(), poolInfo)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handlePoolInfo::InsertOne failed:%+v,data is :%+v", err, m)
			}

			key := strings.ToLower(req.PoolAddress)

			poolMap.Set(key, poolInfo)
		}

	case const_def.EDIT:

		count, err := mongo.MongoClient.Find(context.Background(), bson.M{"pool_address": req.PoolAddress}).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handlePoolInfo::Find pooladdress failed:%+v,pooladdress :%+v", err, req.PoolAddress)
		} else if count < 1 {
			response.Code = PoolNotExist
			response.Msg = ErrorMap[PoolNotExist]
			logger.Error("handlePoolInfo::PoolExist not Exist:%+v,pooladdress :%+v", err, req.PoolAddress)
		} else {
			m := handleOperatePoolInfoReq(req)
			m["update_time"] = time.Now().Unix()
			delete(m, "method")
			delete(m, "page")
			delete(m, "query_token_name")
			u := bson.M{}
			u["$set"] = m
			err := mongo.MongoClient.UpdateOne(context.Background(), bson.M{"pool_address": req.PoolAddress}, u)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handlePoolInfo::UpdateOne failed:%+v,data is :%+v", err, m)
			}

			poolInfo := PoolInfo{}
			mongo.MongoClient.Find(context.Background(), bson.M{"pool_address": req.PoolAddress}).One(&poolInfo)
			poolMap.Replace(strings.ToLower(req.PoolAddress), poolInfo)
		}

	case const_def.QUERY:

		m := handleOperatePoolInfoReq(req)
		delete(m, "method")
		delete(m, "page")
		delete(m, "input_token_name_list")
		delete(m, "query_token_name")
		if req.QueryTokenName != "" {
			m["input_token_name_list"] = bson.M{"$in": []string{strings.ToLower(req.QueryTokenName)}}
		}

		data := []PoolInfo{}
		count, err := mongo.MongoClient.Find(context.Background(), m).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handlePoolInfo::Get Count failed:%+v,data is :%+v", err, m)
		} else {
			err := mongo.MongoClient.Find(context.Background(), m).Sort("-update_time").Skip((req.Page - 1) * 10).Limit(10).All(&data)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handlePoolInfo::Find failed:%+v,data is :%+v", err, m)
			}
		}
		retData := map[string]interface{}{}
		retData["Total"] = count
		if len(data) > 0 {
			retData["PoolList"] = data
		}
		response.Data = retData
	case const_def.DELETE:
		count, err := mongo.MongoClient.Find(context.Background(), bson.M{"pool_address": strings.ToLower(req.PoolAddress)}).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handlePoolInfo::Get Count failed:%+v,data is :%+v", err, req.PoolAddress)
		} else if count < 1 {
			response.Code = PoolNotExist
			response.Msg = ErrorMap[PoolNotExist]
			logger.Error("handlePoolInfo::PoolExist not Exist:%+v,pooladdress :%+v", err, req.PoolAddress)
		} else {

			err := mongo.MongoClient.Remove(context.Background(), bson.M{"pool_address": strings.ToLower(req.PoolAddress)})
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handlePoolInfo::remove failed:%+v,data is :%+v", err, req.PoolAddress)
			}

			poolMap.Delete(strings.ToLower(req.PoolAddress))
		}

	default:
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
	}
	d, _ := json.Marshal(response)
	return d
}

func addressAnalysis(req *AddressAnalysisReq) []byte {
	response := Response{}
	(&response).ResponseInit()

	poolPropertyInfo := config.GetPoolProperty()
	chainCheckInfo := config.GetChainCheck()

	//有交互才有订阅操作

	targetList := []PoolInfo{} //我们支持的  直接交互过的池子列表
	//以pancakeswap为例 pair合约在这里面 farm即masterchelf合约地址也在这里

	masterChelfMap := map[string]map[string]map[string]interface{}{}
	//记录交互过的masterchelf类型合约 具体的池子地址
	//比如在pancakeswap 0x73feaa1ee314f8c655e354234017be2193c9e24e mastterchelf类型中 哪些交互过的池子 即冲了哪些币种
	//masterchelf address->stake token address -> stake token info

	if req.Chain != "" {
		_, ok := chainCheckInfo[req.Chain]
		if !ok {
			response.Code = InvalidArgument
			response.Msg = ErrorMap[InvalidArgument]
			logger.Error("addressAnalysis::invalid argument:%+v", req.Chain)
			d, _ := json.Marshal(response)
			return d
		}
	}

	//TODO 不管是simple_tx还是simple_internal_tx 事实上如果数据全的话有to就意味着from 但是测试仅保留10万区块数据 可能有to但是没有from的情况
	txlist := []logic.TxData{}
	err1 := mongo.TxMongoClient.Find(context.Background(), bson.M{"from": strings.ToLower(req.Address)}).Sort("-block_number").All(&txlist)

	poolInfoList := []PoolInfo{}
	err2 := mongo.MongoClient.Find(context.Background(), bson.M{}).All(&poolInfoList)

	intx1 := []internal_data.InTxRecord{}
	err3 := mongo.InternalTxClient.Find(context.Background(), bson.M{"from": strings.ToLower(req.Address)}).Sort("-block_number").All(&intx1)

	if err1 != nil || err2 != nil || err3 != nil {
		response.Code = InnerError
		response.Msg = ErrorMap[InnerError]
		logger.Error("addressAnalysis::addr:%+v,Get data from db failed::err1:%+v,err2:%+v,err3:%+v", req.Address, err1, err2, err3)
	} else {
		//空间换时间
		m1 := map[string]bool{}
		m2 := map[string]int{}
		m := map[string]bool{}

		for _, e := range txlist {
			m1[strings.ToLower(e.To)] = true
		}
		for k, e := range poolInfoList {
			m2[strings.ToLower(e.PoolAddress)] = k
		}

		if len(m1) > len(m2) {
			for k, index := range m2 {
				if _, ok := m1[k]; ok {
					targetList = append(targetList, poolInfoList[index])
					m[k] = true
				}
			}
		} else {
			for k, _ := range m1 {
				if index, ok := m2[k]; ok {
					targetList = append(targetList, poolInfoList[index])
					m[k] = true
				}
			}
		}

		m3 := map[string]bool{}
		m4 := []string{}

		for _, e := range intx1 {
			m3[strings.ToLower(e.To)] = true
		}

		if len(m3) > len(m2) {
			for k, _ := range m2 {
				_, ok1 := m3[k]
				_, ok2 := m[k]
				if ok1 && !ok2 {
					m4 = append(m4, k)
				}
			}
		} else {
			for k, _ := range m3 {
				_, ok1 := m2[k]
				_, ok2 := m[k]
				if ok1 && !ok2 {
					m4 = append(m4, k)
				}
			}
		}

		intxPoolInfo := []PoolInfo{}
		err := mongo.MongoClient.Find(context.Background(), bson.M{"pool_address": bson.M{"$in": m4}}).Sort("-update_time").All(&intxPoolInfo)
		if err != nil {
			logger.Error("addressAnalysis::Get intxPoolInfo failed:%+v", err1)
		} else {
			targetList = append(targetList, intxPoolInfo...)
		}

		mastchelfInfo := config.GetMasterChelf()

		//masterchelf类型都是在simple_internal_tx中交互的
		for _, e := range targetList {
			if e.PoolProperty == const_def.MasterChelf {
				key := strings.ToLower(e.PoolAddress)
				platfromMasterChelfInfo, ok := mastchelfInfo[key]
				if !ok {
					logger.Error("addressAnalysis::find masterchelf addr:%+v,plat:%+v but not config it!", e.PoolAddress, e.Platform)
					continue
				}
				itemMap := map[string]map[string]interface{}{}

				for _, y := range intx1 {
					//往masterchelf合约打钱 to地址就是PoolAddress Address就是币种合约地址
					//TODO 不同masterchelf处理可能不一样
					if y.To == e.PoolAddress {
						addrMasterChelfInfo, found := platfromMasterChelfInfo[strings.ToLower(y.Address)]
						if !found {
							logger.Error("addressAnalysis::find masterchelf addr:%+v,plat:%+v but not config  special address :%+v !", e.PoolAddress, e.Platform, y.Address)
							continue
						}
						_, ok := itemMap[strings.ToLower(y.Address)]
						if !ok {
							itemMap[strings.ToLower(y.Address)] = addrMasterChelfInfo
						}

					}
				}
				masterChelfMap[e.PoolAddress] = itemMap
			}
		}
	}
	logger.Debug("addressAnalysis::mastchelfMap:%+v", masterChelfMap)
	logger.Debug("addressAnalysis::targetList:%+v,ts:%+v", targetList, time.Now().UnixNano())

	userSubMap := map[string]subData{}
	if req.ConnectAddress != "" {
		subMaps := subMap.Items()
		for k, e := range subMaps {
			v := e.(SubscribeInfo)
			if strings.HasSuffix(k, "_"+req.ConnectAddress) {
				item := subData{
					HasBalance:      false,
					MasterChelfAddr: v.MasterChelfAddr,
				}
				userSubMap[strings.ToLower(v.PoolAddress)] = item
			}
		}
	}

	//库里的地址都已经小写转换过
	platformInfo := map[string]string{}
	relativeInfo := map[string]map[string]string{}
	inputInfo := map[string][]string{}
	if len(targetList) > 0 {
		aaReq := new(analysis.HA_Request)
		aaReq.RequestInfo = []*analysis.HA_Single_Request{}
		aaReq.Chain = req.Chain

		for _, e := range targetList {

			platformInfo[e.PoolAddress] = e.Platform
			relativeMap := map[string]string{}
			for _, y := range e.FundFlowList {
				s := strings.Split(y, "=")
				relativeMap[strings.ToLower(s[1])] = s[0]
			}
			relativeInfo[e.PoolAddress] = relativeMap
			inputInfo[e.PoolAddress] = e.InputTokenNameList

			item := new(analysis.HA_Single_Request)
			if e.CertificateToken != "" {
				certificateInfo := strings.Split(e.CertificateToken, "=")
				certificatePrecision, _ := strconv.Atoi(certificateInfo[2])
				item.CertificateAddress = certificateInfo[1]
				item.CertificateTokenName = certificateInfo[0]
				item.CertificatePrecision = int32(certificatePrecision)
				item.PoolHasCertificate = true
			} else {
				item.PoolHasCertificate = false
			}

			item.PoolAddress = e.PoolAddress
			item.QueryAdddress = strings.ToLower(req.Address)
			item.PoolType = poolPropertyInfo[e.PoolProperty]
			item.Platform = e.Platform

			inputTokenList := []*analysis.InputToken{}
			for _, y := range e.InputTokenList {
				s := strings.Split(y, "=")
				it := new(analysis.InputToken)
				it.Address = s[1]
				it.Name = s[0]
				inputPrecision, _ := strconv.Atoi(s[2])
				it.Precision = int32(inputPrecision)
				inputTokenList = append(inputTokenList, it)
			}
			item.InputTokenList = inputTokenList

			rewardTokenList := []*analysis.RewardToken{}
			for _, y := range e.RewardTokenList {
				s := strings.Split(y, "=")
				it := new(analysis.RewardToken)
				it.Address = s[1]
				it.Name = s[0]
				inputPrecision, _ := strconv.Atoi(s[2])
				it.Precision = int32(inputPrecision)
				rewardTokenList = append(rewardTokenList, it)
			}
			item.RewardTokenList = rewardTokenList

			masterChelfTokenList := []*analysis.MasterChelfToken{}
			masterChelfInfo, ok := masterChelfMap[e.PoolAddress]
			if ok {
				for _, y := range masterChelfInfo {
					item := new(analysis.MasterChelfToken)
					item.Address = y["Address"].(string)
					item.Name = y["Symbol"].(string)
					item.Pid = y["PID"].(int64)
					item.Precision = y["Precision"].(int64)
					masterChelfTokenList = append(masterChelfTokenList, item)
				}
			}
			item.MasterChelfList = masterChelfTokenList

			aaReq.RequestInfo = append(aaReq.RequestInfo, item)
		}

		logger.Debug("addressAnalysis::platformInfo:%+v,relativeInfo:%+v,inputInfo:%+v", platformInfo, relativeInfo, inputInfo)
		logger.Debug("addressAnalysis::reqInfo:%+v,ts:%+v", aaReq, time.Now().UnixNano())
		rt, err := rpc.RPCAddressAnalysis(aaReq)
		logger.Debug("addressAnalysis::RPCAddressAnalysis:%+v,ts:%+v", rt, time.Now().UnixNano())
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			data := map[string]interface{}{}
			response.Data = data
		} else {

			rd := map[string]map[string]interface{}{}
			for _, e := range rt.ReplyList {

				//最外层平台字段
				_, ok := rd[platformInfo[e.PoolAddress]]
				if !ok {
					platformRetInfo := map[string]interface{}{}
					rd[platformInfo[e.PoolAddress]] = platformRetInfo
				}
				platformRetInfo := rd[platformInfo[e.PoolAddress]]

				//平台里面的平台PlatformLogo字段
				_, ok = platformRetInfo["PlatformLogo"]
				if !ok {
					logoKey := req.Chain + "_" + "项目" + "_" + strings.ToLower(platformInfo[e.PoolAddress]) + "_" + ""
					platLogoInfo := logoMap.Get(logoKey)
					if platLogoInfo != nil {
						platformRetInfo["PlatformLogo"] = platLogoInfo.(Logo).URL
					} else {
						platformRetInfo["PlatformLogo"] = ""
						logger.Error("addressAnalysis::not found  plarform logo,chain:%+v,name:%+v,poolAddress:%+v", req.Chain, platformInfo[e.PoolAddress], e.PoolAddress)
					}
				}

				//平台里面的PoolList字段
				_, ok = platformRetInfo["PoolList"]
				if !ok {
					poolList := []map[string]interface{}{}
					platformRetInfo["PoolList"] = poolList
				}

				poolList := platformRetInfo["PoolList"].([]map[string]interface{})

				//PoolList中每一个池子的信息
				poolRetItem := map[string]interface{}{}

				//池子 资金池名字
				nameList := inputInfo[e.PoolAddress]
				//TODO 关于每个池子的名字 不同项目可能不同处理 这里是pancake的处理
				if len(nameList) > 0 {
					sort.Strings(nameList)
					name := ""
					for _, y := range nameList {
						name += y + "-"
					}
					name = name[:len(name)-1]

					poolRetItem["PoolInputTokenName"] = name
				} else {
					//masterchelf
					masterChelfInfo, ok1 := masterChelfMap[e.PoolAddress]
					if ok1 {
						masterChelfAddrInfo, ok2 := masterChelfInfo[strings.ToLower(e.MasterChelfAddress)]
						if ok2 {
							poolRetItem["PoolInputTokenName"] = masterChelfAddrInfo["Symbol"]
						}
					}
				}

				//池子地址
				poolRetItem["PoolAddress"] = e.PoolAddress

				//池子是否监控
				tmpPoolInfo := poolMap.Get(strings.ToLower(e.PoolAddress))
				if tmpPoolInfo != nil {
					poolRetItem["PoolWatching"] = (tmpPoolInfo.(PoolInfo).IsWatching == const_def.TRUE)
				} else {
					poolRetItem["PoolWatching"] = false
				}
				//用户是否订阅
				subKey := strings.ToLower(e.PoolAddress) + "_" + strings.ToLower(e.MasterChelfAddress) + "_" + strings.ToLower(req.ConnectAddress)
				hasSub := subMap.Check(subKey)
				poolRetItem["UserSub"] = hasSub
				if hasSub {
					v, ok := userSubMap[subKey]
					if ok {
						vv := subData{
							HasBalance:      true,
							MasterChelfAddr: v.MasterChelfAddr,
						}
						userSubMap[subKey] = vv
					}
				}

				//余额 和 奖励信息
				out := map[string][]map[string]string{}
				outRewardList := []map[string]string{}
				outBalanceList := []map[string]string{}
				for _, y := range e.OutInfo {
					outItem := map[string]string{}
					outItem["TokenName"] = y.TokenName
					outItem["TokenAmount"] = y.TokenAmount

					tokenLogoInfo := logoMap.Get(req.Chain + "_" + "币种" + "_" + strings.ToLower(y.TokenName) + "_" + strings.ToLower(y.TokenAddress))

					if tokenLogoInfo != nil {
						outItem["TokenLogo"] = tokenLogoInfo.(Logo).URL
					} else {
						outItem["TokenLogo"] = ""
						logger.Error("addressAnalysis::not found  token logo,chain:%+v,name:%+v,tokenaddress:%+v,poolAddress:%+v", req.Chain, y.TokenName, y.TokenAddress, e.PoolAddress)
					}
					if y.TokenType == const_def.REWARD {
						outRewardList = append(outRewardList, outItem)
					} else {
						outBalanceList = append(outBalanceList, outItem)
					}

				}

				out["Balance"] = outBalanceList
				out["Reward"] = outRewardList
				poolRetItem["OutList"] = out

				//关联合约信息
				relativeList := []map[string]interface{}{}
				//k 已经小写
				for k, v := range relativeInfo[e.PoolAddress] {
					//关联合约的masterchelf地址
					relativeMasterchelfAddrList := []string{}
					if poolMap.Check(k) {
						relativePoolInfo := poolMap.Get(k).(PoolInfo)

						//TODO 不同的masterchelf lp地址可能的处理方式不同panckae 就是PoolAddress的input信息

						if relativePoolInfo.PoolProperty == const_def.MasterChelf && tmpPoolInfo != nil {
							for _, y := range tmpPoolInfo.(PoolInfo).InputTokenList {
								s := strings.Split(y, "=")
								relativeMasterchelfAddrList = append(relativeMasterchelfAddrList, strings.ToLower(s[1]))
							}
						} else {
							relativeMasterchelfAddrList = append(relativeMasterchelfAddrList, "")
						}

					}

					for _, y := range relativeMasterchelfAddrList {
						item := map[string]interface{}{}
						item["ProjectName"] = v
						item["RelativePoolAddress"] = k
						item["RelativeMasterChelfAddr"] = y

						relativeProjectLogoInfo := logoMap.Get(req.Chain + "_" + "项目" + "_" + strings.ToLower(v) + "_" + "")

						if relativeProjectLogoInfo != nil {
							item["Logo"] = relativeProjectLogoInfo.(Logo).URL
						} else {
							item["Logo"] = ""
							logger.Error("addressAnalysis::not found  plarform logo,chain:%+v,name:%+v,poolAddress:%+v", req.Chain, v, k)
						}

						tmpRelativePoolInfo := poolMap.Get(k)
						if tmpRelativePoolInfo != nil {
							item["PoolWatching"] = (tmpRelativePoolInfo.(PoolInfo).IsWatching == const_def.TRUE)
						} else {
							item["PoolWatching"] = false
						}
						//用户是否订阅
						relativeSubKey := strings.ToLower(k) + "_" + y + "_" + strings.ToLower(req.ConnectAddress)
						relativeHasSub := subMap.Check(relativeSubKey)
						item["UserSub"] = relativeHasSub
						if relativeHasSub {
							rv, ok := userSubMap[relativeSubKey]
							if ok {
								vv := subData{
									HasBalance:      true,
									MasterChelfAddr: rv.MasterChelfAddr,
								}
								userSubMap[relativeSubKey] = vv
							}
						}

						relativeList = append(relativeList, item)
					}

				}

				poolRetItem["PoolRelativeInfo"] = relativeList

				//池子的masterchelf信息
				poolRetItem["MasterChelfAddr"] = strings.ToLower(e.MasterChelfAddress)

				poolList = append(poolList, poolRetItem)
				platformRetInfo["PoolList"] = poolList
				rd[platformInfo[e.PoolAddress]] = platformRetInfo

			}

			//没有钱但是订阅过的信息
			for k, v := range userSubMap {
				if !v.HasBalance && poolMap.Check(k) {
					//项目名称
					zeroPoolInfo := poolMap.Get(k).(PoolInfo)
					_, ok := rd[zeroPoolInfo.Platform]
					if !ok {
						platformRetInfo := map[string]interface{}{}
						rd[zeroPoolInfo.Platform] = platformRetInfo
					}
					platformRetInfo := rd[zeroPoolInfo.Platform]

					//项目logo
					_, ok = platformRetInfo["PlatformLogo"]
					if !ok {
						logoKey := req.Chain + "_" + "项目" + "_" + strings.ToLower(zeroPoolInfo.Platform) + "_" + ""
						platLogoInfo := logoMap.Get(logoKey)
						if platLogoInfo != nil {
							platformRetInfo["PlatformLogo"] = platLogoInfo.(Logo).URL
						} else {
							platformRetInfo["PlatformLogo"] = ""
							logger.Error("addressAnalysis::not found  plarform logo,chain:%+v,name:%+v,poolAddress:%+v", req.Chain, zeroPoolInfo.Platform, k)
						}
					}

					_, ok = platformRetInfo["PoolList"]
					if !ok {
						poolList := []map[string]interface{}{}
						platformRetInfo["PoolList"] = poolList
					}
					poolList := platformRetInfo["PoolList"].([]map[string]interface{})
					poolRetItem := map[string]interface{}{}

					if len(zeroPoolInfo.InputTokenNameList) > 0 {
						sort.Strings(zeroPoolInfo.InputTokenNameList)
						name := ""
						for _, y := range zeroPoolInfo.InputTokenNameList {
							name += y + "-"
						}
						name = name[:len(name)-1]

						poolRetItem["PoolInputTokenName"] = name
					} else {
						//masterchelf
						masterChelfInfo, ok1 := masterChelfMap[k]
						if ok1 {
							masterChelfAddrInfo, ok2 := masterChelfInfo[strings.ToLower(v.MasterChelfAddr)]
							if ok2 {
								poolRetItem["PoolInputTokenName"] = masterChelfAddrInfo["Symbol"]
							}
						}
					}

					poolRetItem["PoolAddress"] = k
					poolRetItem["UserSub"] = true
					poolRetItem["MasterChelfAddr"] = strings.ToLower(v.MasterChelfAddr)
					//订阅相关信息
					//池子是否被监控

					tmpPoolInfo := poolMap.Get(strings.ToLower(k))
					if tmpPoolInfo != nil {
						poolRetItem["PoolWatching"] = (tmpPoolInfo.(PoolInfo).IsWatching == const_def.TRUE)
					} else {
						poolRetItem["PoolWatching"] = false
					}

					out := map[string][]map[string]string{}

					out["Balance"] = []map[string]string{}
					out["Reward"] = []map[string]string{}
					poolRetItem["OutList"] = out

					relativeList := []map[string]interface{}{}
					for _, y := range zeroPoolInfo.FundFlowList {
						s := strings.Split(y, "=")
						projectName := s[0]
						realtiveAddress := s[1]
						relativeMasterChelfAddrList := []string{}
						if poolMap.Check(realtiveAddress) {
							relativePoolInfo := poolMap.Get(realtiveAddress).(PoolInfo)

							//TODO 不同的masterchelf lp地址可能的处理方式不同panckae 就是PoolAddress的input信息

							if relativePoolInfo.PoolProperty == const_def.MasterChelf {
								for _, y := range zeroPoolInfo.InputTokenList {
									s := strings.Split(y, "=")
									relativeMasterChelfAddrList = append(relativeMasterChelfAddrList, strings.ToLower(s[1]))
								}
							} else {
								relativeMasterChelfAddrList = append(relativeMasterChelfAddrList, "")
							}

						}

						for _, y := range relativeMasterChelfAddrList {
							item := map[string]interface{}{}
							item["ProjectName"] = projectName
							item["RelativePoolAddress"] = realtiveAddress
							item["RelativeMasterChelfAddr"] = y
							relativeProjectLogoInfo := logoMap.Get(req.Chain + "_" + "项目" + "_" + strings.ToLower(projectName) + "_" + "")

							if relativeProjectLogoInfo != nil {
								item["Logo"] = relativeProjectLogoInfo.(Logo).URL
							} else {
								item["Logo"] = ""
								logger.Error("addressAnalysis::not found  plarform logo,chain:%+v,name:%+v,poolAddress:%+v", req.Chain, k, v)
							}

							tmpRelativePoolInfo := poolMap.Get(strings.ToLower(realtiveAddress))
							if tmpRelativePoolInfo != nil {
								item["PoolWatching"] = (tmpRelativePoolInfo.(PoolInfo).IsWatching == const_def.TRUE)
							} else {
								item["PoolWatching"] = false
							}
							//用户是否订阅
							relativeSubKey := strings.ToLower(realtiveAddress) + "_" + strings.ToLower(y) + "_" + strings.ToLower(req.ConnectAddress)
							relativeHasSub := subMap.Check(relativeSubKey)
							item["UserSub"] = relativeHasSub

							relativeList = append(relativeList, item)
						}
					}

					poolRetItem["PoolRelativeInfo"] = relativeList
					poolList = append(poolList, poolRetItem)
					platformRetInfo["PoolList"] = poolList
					rd[zeroPoolInfo.Platform] = platformRetInfo
				}
			}

			response.Data = rd
		}
	} else {
		data := map[string]interface{}{}
		response.Data = data
	}

	d, _ := json.Marshal(response)
	return d
}

func poolComboBox() []byte {
	platform := config.GetPlatformCheck()
	platforms := []string{}
	for k, _ := range platform {
		platforms = append(platforms, k)
	}
	chain := config.GetChainCheck()
	chains := []string{}
	for k, _ := range chain {
		chains = append(chains, k)
	}

	poolProperty := config.GetPoolProperty()
	poolPropertys := []string{}
	for k, _ := range poolProperty {
		poolPropertys = append(poolPropertys, k)
	}
	response := Response{}
	(&response).ResponseInit()
	data := map[string][]string{}
	data["Platform"] = platforms
	data["Chain"] = chains
	data["PoolProperty"] = poolPropertys
	response.Data = data
	d, _ := json.Marshal(response)
	return d
}

func logoComboBox() []byte {
	logoType := config.GetLogoType()
	logoTypes := []string{}
	for k, _ := range logoType {
		logoTypes = append(logoTypes, k)
	}
	chain := config.GetChainCheck()
	chains := []string{}
	for k, _ := range chain {
		chains = append(chains, k)
	}

	response := Response{}
	(&response).ResponseInit()
	data := map[string][]string{}
	data["LogoChain"] = chains
	data["LogoType"] = logoTypes
	response.Data = data
	d, _ := json.Marshal(response)
	return d
}

//删改查所用
func handleSubscribeReq(req *SubscribeInfoReq) bson.M {
	t := reflect.TypeOf(*req)
	v := reflect.ValueOf(*req)
	s := bson.M{}
	for k := 0; k < t.NumField(); k++ {
		switch t.Field(k).Type.Kind() {
		case reflect.String:
			if v.Field(k).String() != "" {
				key := string(t.Field(k).Tag.Get("json"))
				if key == "pool_address" || key == "user_address" || key == "master_chelf_addr" {
					s[key] = strings.ToLower(v.Field(k).String())
				} else {
					s[key] = v.Field(k).String()
				}

			}
		default:
			//
		}
	}

	return s
}

func checkSubscribeReq(req *SubscribeInfoReq) error {

	if req.PoolAddress != "" {
		if !strings.HasPrefix(req.PoolAddress, "0x") {
			logger.Error("CheckSubscribeReq::PoolAddress error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	if req.UserAddress != "" {
		if !strings.HasPrefix(req.UserAddress, "0x") {
			logger.Error("CheckSubscribeReq::UserAddress error")
			return errors.New(ErrorMap[InvalidArgument])
		}
	}

	//检查一下是否是在map中
	if req.MasterChelfAddr != "" {
		masterchelfMap := config.GetMasterChelf()
		key := strings.ToLower(req.PoolAddress)
		platfromMasterChelfInfo, ok1 := masterchelfMap[key]
		if !ok1 {
			logger.Error("CheckSubscribeReq:: not exist this masterchelf contract:%+v", req.PoolAddress)
			return errors.New(ErrorMap[InvalidArgument])
		}
		_, ok2 := platfromMasterChelfInfo[strings.ToLower(req.MasterChelfAddr)]
		if !ok2 {
			logger.Error("CheckSubscribeReq:: not exist this stake pool :%+v,of masterchelf contract :%+v", req.MasterChelfAddr, req.PoolAddress)
			return errors.New(ErrorMap[InvalidArgument])
		}

	}

	if (req.TVLDec < 0 || req.TVLDec > 100) || (req.TVLInc < 0 || req.TVLInc > 100) {
		logger.Error("CheckSubscribeReq:TVL arg error")
		return errors.New(ErrorMap[InvalidArgument])
	}

	if req.WatchFlashLoan != const_def.TRUE && req.WatchFlashLoan != const_def.FALSE && req.WatchFlashLoan != "" {
		logger.Error("CheckSubscribeReq:watchflashloan error")
		return errors.New(ErrorMap[InvalidArgument])
	}

	switch req.Method {
	case const_def.ADD:
		//因为没有池子id那样的强唯一
		if req.PoolAddress == "" || req.UserAddress == "" || req.Email == "" || req.WatchFlashLoan == "" || (req.TVLDec == 0 && req.TVLInc == 0) {
			logger.Error("CheckSubscribeReq:ADD Method Must contain pooladdress useraddress email and watch info,watchFlashLoan")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.EDIT:
		if req.PoolAddress == "" || req.UserAddress == "" {
			logger.Error("CheckSubscribeReq:EDIT Method Must contain pooladdres ad useraddress")
			return errors.New(ErrorMap[InvalidArgument])
		}
	case const_def.QUERY:
		if req.PoolAddress == "" || req.UserAddress == "" {
			logger.Error("CheckSubscribeReq:QUERY Method Must contain pooladdres ad useraddress")
			return errors.New(ErrorMap[InvalidArgument])
		}

	case const_def.DELETE:
		if req.PoolAddress == "" || req.UserAddress == "" {
			logger.Error("CheckSubscribeReq:DELETE Method Must contain pooladdres ad useraddress")
			return errors.New(ErrorMap[InvalidArgument])
		}
	default:
		logger.Error("CheckSubscribeReq:invalid method")
		return errors.New(ErrorMap[InvalidArgument])
	}

	return nil
}

func handleSubcribe(req *SubscribeInfoReq) []byte {
	response := Response{}
	(&response).ResponseInit()
	err := checkSubscribeReq(req)
	if err != nil {
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
		d, _ := json.Marshal(response)
		return d
	}

	switch req.Method {
	case const_def.ADD:
		//唯一性
		m := handleSubscribeReq(req)
		m["update_time"] = time.Now().Unix()
		delete(m, "method")
		u := bson.M{}
		u["pool_address"] = strings.ToLower(req.PoolAddress)
		u["user_address"] = strings.ToLower(req.UserAddress)
		u["master_chelf_addr"] = strings.ToLower(req.MasterChelfAddr)

		count, err := mongo.SubMongoClient.Find(context.Background(), u).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleSubcribe::Find logo failed:%+v,u :%+v", err, u)
		} else if count > 0 {
			response.Code = SubExist
			response.Msg = ErrorMap[SubExist]
			logger.Error("handleSubcribe::SubInfo already Exist:%+v,u :%+v", err, u)
		} else {
			subInfo := SubscribeInfo{
				PoolAddress:     strings.ToLower(req.PoolAddress),
				MasterChelfAddr: strings.ToLower(req.MasterChelfAddr),
				UserAddress:     strings.ToLower(req.UserAddress),
				Email:           req.Email,
				WatchFlashLoan:  req.WatchFlashLoan,
				TVLInc:          req.TVLInc,
				TVLDec:          req.TVLDec,
				UpdateTime:      time.Now().Unix(),
			}
			_, err = mongo.SubMongoClient.InsertOne(context.Background(), subInfo)
			if err != nil {
				response.Code = InnerError
				response.Msg = ErrorMap[InnerError]
				logger.Error("handleSubcribe::InsertOne failed:%+v,data is :%+v", err, m)
			}

			key := strings.ToLower(req.PoolAddress) + "_" + strings.ToLower(req.MasterChelfAddr) + "_" + strings.ToLower(req.UserAddress)

			subMap.Set(key, subInfo)
		}

	case const_def.EDIT:

		u := bson.M{}
		u["pool_address"] = strings.ToLower(req.PoolAddress)
		u["user_address"] = strings.ToLower(req.UserAddress)
		u["master_chelf_addr"] = strings.ToLower(req.MasterChelfAddr)
		count, err := mongo.SubMongoClient.Find(context.Background(), u).Count()
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleSubcribe::Find  failed:%+v,u :%+v", err, u)
		} else if count < 1 {
			response.Code = SubNotExist
			response.Msg = ErrorMap[SubNotExist]
			logger.Error("handleSubcribe::sub not Exist:%+v,u :%+v", err, u)
		} else {
			m := handleSubscribeReq(req)
			m["update_time"] = time.Now().Unix()
			delete(m, "method")

			subInfo := SubscribeInfo{}
			mongo.SubMongoClient.Find(context.Background(), u).One(&subInfo)
			subInfo.TVLDec = req.TVLDec
			subInfo.TVLInc = req.TVLInc
			subInfo.WatchFlashLoan = req.WatchFlashLoan
			if subInfo.TVLDec == 0 && subInfo.TVLInc == 0 && subInfo.WatchFlashLoan == const_def.FALSE {
				//删除
				v := bson.M{}
				v["pool_address"] = strings.ToLower(req.PoolAddress)
				v["user_address"] = strings.ToLower(req.UserAddress)
				v["master_chelf_addr"] = strings.ToLower(req.MasterChelfAddr)
				err := mongo.SubMongoClient.Remove(context.Background(), v)
				if err != nil {
					response.Code = InnerError
					response.Msg = ErrorMap[InnerError]
					logger.Error("handleSubcribe::remove failed:%+v,data is :%+v", err, v)
				}
				key := strings.ToLower(req.PoolAddress) + "_" + strings.ToLower(req.MasterChelfAddr) + "_" + strings.ToLower(req.UserAddress)
				subMap.Delete(key)
			} else {
				v := bson.M{}
				v["$set"] = m

				err := mongo.SubMongoClient.UpdateOne(context.Background(), u, v)
				if err != nil {
					response.Code = InnerError
					response.Msg = ErrorMap[InnerError]
					logger.Error("handleSubcribe::UpdateOne failed:%+v,data is :%+v", err, m)
				}

				key := strings.ToLower(req.PoolAddress) + "_" + strings.ToLower(req.MasterChelfAddr) + "_" + strings.ToLower(req.UserAddress)
				subInfo := SubscribeInfo{}
				mongo.SubMongoClient.Find(context.Background(), u).One(&subInfo)
				subMap.Replace(key, subInfo)
			}
		}
	case const_def.QUERY:

		u := bson.M{}
		u["pool_address"] = strings.ToLower(req.PoolAddress)
		u["user_address"] = strings.ToLower(req.UserAddress)
		u["master_chelf_addr"] = strings.ToLower(req.MasterChelfAddr)

		subInfo := SubscribeInfo{}
		mongo.SubMongoClient.Find(context.Background(), u).One(&subInfo)
		data := map[string]interface{}{}
		if subInfo.PoolAddress != "" {
			data["PoolAddress"] = subInfo.PoolAddress
			data["UserAddress"] = subInfo.UserAddress
			data["Email"] = subInfo.Email
			data["WatchFl"] = subInfo.WatchFlashLoan
			data["TvlInc"] = subInfo.TVLInc
			data["TvlDec"] = subInfo.TVLDec
		}
		response.Data = data

	case const_def.DELETE:

		u := bson.M{}
		u["pool_address"] = strings.ToLower(req.PoolAddress)
		u["user_address"] = strings.ToLower(req.UserAddress)
		u["master_chelf_addr"] = strings.ToLower(req.MasterChelfAddr)
		err := mongo.SubMongoClient.Remove(context.Background(), u)
		if err != nil {
			response.Code = InnerError
			response.Msg = ErrorMap[InnerError]
			logger.Error("handleSubcribe::remove failed:%+v,data is :%+v", err, u)
		}

		key := strings.ToLower(req.PoolAddress) + "_" + strings.ToLower(req.MasterChelfAddr) + "_" + strings.ToLower(req.UserAddress)
		subMap.Delete(key)

	default:
		response.Code = InvalidArgument
		response.Msg = ErrorMap[InvalidArgument]
	}
	d, _ := json.Marshal(response)
	return d
}
