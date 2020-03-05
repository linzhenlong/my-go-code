package parser
import(
	"github.com/linzhenlong/my-go-code/2020/study/crawler/engine"
	"regexp"

)

const cityListRegexp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// CityList 解析城市列表.
func CityList(contents []byte) engine.ParserResult {
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/shanghai" data-v-5fa74e39>上海</a>
	re := regexp.MustCompile(cityListRegexp)
	match := re.FindAllStringSubmatch(string(contents), -1)
	result := engine.ParserResult{}
	for _, m := range match {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			URL: m[1],
			ParserFunc: engine.NilParser,
		})
	}
	return result
}