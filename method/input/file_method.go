package input

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"os"
)

// fsAddress 本地文件资源地址
type fsAddress struct {
	*BaseAddress
}

// 通过文件地址获取文件的过程
func (address *fsAddress) GetResource() []Word {
	localfile, err := os.Open(address.Address.Path)
	if err != nil {
		log.Err("open  file from file system failed.Error:%v\n",err)
		return nil
	}
	defer localfile.Close()
	reader := bufio.NewReader(localfile)
	if err != nil {
		log.Err("open  file from file system failed.Error:%v\n",err)
		return nil

	}

	return address.ReadFile(reader)

}
