package engine

import(
	"github.com/linzhenlong/my-go-code/2020/study/crawler/fetcher"
	"log"

)
// Run 跑.
func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:] // 移除掉第一个.
		log.Printf("Fetch url:%s", r.URL)
		body , err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("engine Run fetcher url:%s,error:%s",r.URL, err.Error())
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item %v", item)
		}
	}
}