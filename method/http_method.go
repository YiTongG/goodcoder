package method

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/log"
	"net/http"
)

// httpResource http类型资源地址
type httpResource struct {
	*resource
}

const (
	// RetryNum http 超时最大重试次数
	RetryNum = 3
)

// Call 超时重试
func (resource *httpResource) Call() (resp *http.Response, err error) {
	for i := 0; i < RetryNum; i++ {
		resp, err = http.Get(resource.path)
		if err != nil {
			continue
		}
		return resp, nil
	}
	return nil, err
}

// GetResource 通过http地址获取资源
func (resource *httpResource) GetResource() []word {
	resp, err := resource.Call()
	if err != nil {
		log.Err("fail to get http resource.Error:\n", err)
		return nil
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	return resource.readFile(reader)
}