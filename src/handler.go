package crm

import (
	"golang.org/x/net/context"
	"proto"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strconv"
)

type CRMService struct {
	uid int64
	kv  *consul.KV
}

func (s *CRMService) Init(url string) {
	fmt.Println("[CRM] 初始化系统...")
	config := consul.DefaultConfig()
	config.Address = url
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	s.kv = c.KV()
	p, _, err := s.kv.Get("rpg/latestID", nil)
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
	out.ID = string(s.uid)
	out.Token = string(s.uid)

	s.uid += 1
	newUID := string(s.uid)

	d := &consul.KVPair{Key: "rpg/latestID", Value: []byte(newUID)}
	s.kv.Put(d, nil)

	return nil
}
