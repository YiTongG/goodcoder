package input

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"net/http"
)

type httpAddress struct {
	*BaseAddress
}
const (
	// RetryNum http 超时最大重试次数
	RetryNum = 3
)
// Call 超时重试
func (address *httpAddress) Call() (resp *http.Response, err error) {
	for i := 0; i < RetryNum; i++ {
		resp, err = http.Get(address.Address.Path)
		if err != nil {
			continue
		}
		return resp, nil
	}
	return nil, err
}

// GetResource 通过http地址获取资源
func (address *httpAddress) GetResource() []Word {
	resp, err := address.Call()
	if err != nil {
		log.Err("fail to get http resource.Error:%v\n",err)
		return nil
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	return address.ReadFile(reader)
}

