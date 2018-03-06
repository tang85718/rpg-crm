package crm

import (
	"golang.org/x/net/context"
	"proto"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strconv"
)

const (
	PLAYER_START = 10000
)

type CRMService struct {
	uid int64
}

type PlayerTotal struct {
	uid int64 `bson:"uid"`
}

func (s *CRMService) Init(url string) {
	fmt.Println("[CRM] 初始化系统...")
	config := consul.DefaultConfig()
	config.Address = url
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	kv := c.KV()
	p, _, err := kv.Get("rpg/latestID", nil)
	if err != nil {
		panic(err)
	}

	s.uid, err = strconv.ParseInt(string(p.Value), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[CRM] 最大ID：%d\n", s.uid)
}

func (s *CRMService) Signup(c context.Context, in *crm_api.SignupReq, out *crm_api.SignupResponse) error {
	if s.uid <= PLAYER_START {
		s.uid = PLAYER_START
	}

	out.ID = s.uid

	return nil
}
