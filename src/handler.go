package crm

import (
	"golang.org/x/net/context"
	"proto"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strconv"
	"gopkg.in/mgo.v2"
	"github.com/pborman/uuid"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type CRMService struct {
	pIndex int64
	kv     *consul.KV
	mgo    *mgo.Session
}

func (s *CRMService) Init(consulURL string, mgoUrl string) {
	fmt.Println("[CRM] 初始化系统...")
	config := consul.DefaultConfig()
	config.Address = consulURL
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	s.kv = c.KV()
	p, _, err := s.kv.Get("rpg/latestID", nil)
	if err != nil {
		panic(err)
	}

	s.pIndex, err = strconv.ParseInt(string(p.Value), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[CRM] 最大ID：%d\n", s.pIndex)

	s.mgo, err = mgo.Dial(mgoUrl)
	if err != nil {
		panic(err)
	}

	s.mgo.SetMode(mgo.Monotonic, true)

	fmt.Println("[CRM] 连接 mongo 成功")
}

func (s *CRMService) Signup(c context.Context, in *crm_api.SignupReq, out *crm_api.SignupResponse) error {
	out.ID = strconv.FormatInt(s.pIndex, 10)
	out.Token = uuid.NewUUID().String()

	s.pIndex += 1
	newUID := strconv.FormatInt(s.pIndex, 10)

	d := &consul.KVPair{Key: "rpg/latestID", Value: []byte(newUID)}
	s.kv.Put(d, nil)

	now := time.Now()
	player := &Player{
		ID:         bson.NewObjectId(),
		DisplayID:  out.ID,
		Token:      out.Token,
		CreateTime: now,
		UpdateTime: now,
	}

	mc := s.mgo.DB("crm").C("player")
	err := mc.Insert(player)
	if err != nil {
		return err
	}

	//mc.Insert(bson.M{
	//	"id":          out.ID,
	//	"token":       out.Token,
	//	"create_time": now,
	//	"update_time": now,
	//})

	//todo: create elastic diary

	return nil
}
