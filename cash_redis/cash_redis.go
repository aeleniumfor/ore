package cash_redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// func main() {
// 	c, err := redis.Dial("tcp", ":6379")
// 	if err != nil {
// 		// handle error
// 	}
// 	defer c.Close()

// 	r, _ := c.Do("SET", "hosts", "web.svc.user.localhost")
// 	fmt.Println(r)
// 	s, _ := redis.String(c.Do("GET", "hosts"))
// 	fmt.Println(s)
// }

type Cach struct {
	protocol string
	host     string
	client   redis.Conn
}

func Init(host, protocol string) *Cach {
	fmt.Println("Init")
	return &Cach{host: host, protocol: protocol}
}

func (c *Cach) Connect() {
	c.client, _ = redis.Dial(c.protocol, c.host)
}

func (c Cach) GetCash(host string) (res string, err error) {
	//	"web.test.user.localhost"
	s, err := redis.String(c.client.Do("GET", host))
	return s, err
}

func (c *Cach) SetCach(request, revers_host string) {
	c.client.Do("SET", request, revers_host)
}
