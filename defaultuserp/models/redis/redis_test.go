package redis

import "testing"

func TestName(t *testing.T) {
	client, _ := NewRedis("16.163.34.138:16379", "nft666", 0)
	client.Set("15101131912", "6666")
	//t.Log(client.Get("17733452183"))
	//client.Del("17733452183")
	//t.Log(client.Get("17733452183"))
}
