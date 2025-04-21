package server

import (
	"fmt"
	"math/rand"
	"time"
)

var IPool = map[string]ipool{
	"r": &RandomIpPool{},
}

type ipool interface {
	Init(prefix string)
	GetIp() string
}

type RandomIpPool struct {
	prefix string
}

func (p *RandomIpPool) Init(prefix string) {
	p.prefix = prefix
}
func (p *RandomIpPool) GetIp() string {
	rand.Seed(time.Now().UnixNano())
	// 固定部分
	fixedPart := p.prefix + ":"
	// 随机部分
	var randomPart string
	for i := 0; i < 4; i++ {
		randomPart += fmt.Sprintf("%04x:", rand.Intn(65536))
	}
	// 去除最后的冒号
	randomPart = randomPart[:len(randomPart)-1]
	return "[" + fixedPart + randomPart + "]:0"
}
