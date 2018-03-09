package crm

import (
	"golang.org/x/net/context"
	"proto/crm"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
	"github.com/tangxuyao/mongo"
)

const (
	CRM_DB    = "crm"
	PlayerCOL = "player"
	ActorCOL  = "actors"
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
	id := bson.NewObjectId()

	out.ID = strconv.FormatInt(s.pIndex, 10)
	out.Token = id.Hex()

	now := time.Now()
	player := &mongo.Player{
		ID:         id,
		DisplayID:  out.ID,
		Token:      id.Hex(),
		CreateTime: now,
		UpdateTime: now,
	}

	mc := s.mgo.DB(CRM_DB).C(PlayerCOL)
	err := mc.Insert(player)
	if err != nil {
		return err
	}

	s.pIndex += 1
	newUID := strconv.FormatInt(s.pIndex, 10)

	d := &consul.KVPair{Key: "rpg/latestID", Value: []byte(newUID)}
	s.kv.Put(d, nil)

	//todo: create elastic diary
	return nil
}

func (s *CRMService) BindPhone(c context.Context, in *crm_api.BindPhoneReq, out *crm_api.BindPhoneResponse) error {
	return nil
}

func (s *CRMService) MakeActor(c context.Context, in *crm_api.MakeActorReq, out *crm_api.MakeActorRsp) error {
	playerCol := s.mgo.DB(CRM_DB).C(PlayerCOL)
	player := mongo.Player{}
	err := playerCol.Find(bson.M{"token": in.Token}).One(&player)
	if err != nil {
		return err
	}

	actorCOL := s.mgo.DB(CRM_DB).C(ActorCOL)
	count, err := actorCOL.Find(bson.M{"player_token": player.Token}).Count()

	if count > 0 {
		return errors.New("不允许创建超过1个角色")
	}

	actor := mongo.Charactor{PlayerToken: in.Token, Name: in.Name, HP: 5, Energy: 0, EnergyType: 0}
	actorCOL.Insert(&actor)

	fmt.Printf("创建新角色%s, 属于玩家%s(%s)\n", in.Name, player.DisplayID, player.Token)
	return nil
}
