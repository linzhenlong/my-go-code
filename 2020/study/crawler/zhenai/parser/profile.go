package parser

import(
	"github.com/linzhenlong/my-go-code/2020/study/crawler/engine"

)
const ageRegexp = `<div [^>]* class="des f-cl">[.+] | ([\d]+)岁 | [.+] | [.+] | ([\d]+)cm | ([\d]+-[\d]+)元</div>`

// Profile 解析用户信息.
func Profile(contents []byte) engine.ParserResult {
	return engine.ParserResult{}
}