package method

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/log"
	"os"
)

// fsResource 本地文件资源地址
type fsResource struct {
	*resource
}

// GetResource 通过文件地址获取文件的过程
func (resource *fsResource) GetResource() []word {
	localfile, err := os.Open(resource.resource.path)
	if err != nil {
		log.Err("open  file from file system failed.Error:\n", err)
		return nil
	}
	defer localfile.Close()
	reader := bufio.NewReader(localfile)
	return resource.readFile(reader)

}