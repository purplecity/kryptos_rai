package app

import (
	"context"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type StringMap struct {
	Lock *sync.RWMutex
	Bm   map[string]interface{}
}

func NewStringMap() *StringMap {
	return &StringMap{
		Lock: new(sync.RWMutex),
		Bm:   make(map[string]interface{}),
	}
}

func (m *StringMap) Get(k string) interface{} {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	if val, ok := m.Bm[k]; ok {
		return val
	}
	return nil
}

func (m *StringMap) Items() map[string]interface{} {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	r := make(map[string]interface{})
	for k, v := range m.Bm {
		r[k] = v
	}
	return r
}

func (m *StringMap) Set(k string, v interface{}) bool {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	if val, ok := m.Bm[k]; !ok {
		m.Bm[k] = v
	} else if val != v {
		m.Bm[k] = v
	} else {
		return false
	}
	return true
}

func (m *StringMap) Delete(k string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	delete(m.Bm, k)
}

func (m *StringMap) Replace(k string, v interface{}) bool {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	delete(m.Bm, k)
	if val, ok := m.Bm[k]; !ok {
		m.Bm[k] = v
	} else if val != v {
		m.Bm[k] = v
	} else {
		return false
	}
	return true
}

func (m *StringMap) Check(k string) bool {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	if _, ok := m.Bm[k]; !ok {
		return false
	}
	return true
}

var poolMap *StringMap
var subMap *StringMap
var logoMap *StringMap

func LoadDB() error {
	poolMap = NewStringMap()
	subMap = NewStringMap()
	logoMap = NewStringMap()

	poolInfoList := []PoolInfo{}
	err := mongo.MongoClient.Find(context.Background(), bson.M{}).All(&poolInfoList)
	if err != nil {
		logger.Error("LoadDB::load pool info failed:%+v", err)
		return err
	}

	subList := []SubscribeInfo{}
	err = mongo.SubMongoClient.Find(context.Background(), bson.M{}).All(&subList)
	if err != nil {
		logger.Error("LoadDB::load sub info failed:%+v", err)
		return err
	}

	logoList := []Logo{}
	err = mongo.LogoMongoClient.Find(context.Background(), bson.M{}).All(&logoList)
	if err != nil {
		logger.Error("LoadDB::load logo info failed:%+v", err)
		return err
	}

	for _, e := range poolInfoList {
		poolMap.Set(strings.ToLower(e.PoolAddress), e)
	}

	for _, e := range subList {
		key := strings.ToLower(e.PoolAddress) + "_" + strings.ToLower(e.MasterChelfAddr) + "_" + strings.ToLower(e.UserAddress)
		subMap.Set(key, e)
	}

	for _, e := range logoList {
		key := e.Chain + "_" + e.LogoType + "_" + e.Name + "_" + strings.ToLower(e.TokenAddress)
		logoMap.Set(key, e)
	}

	return nil

}
