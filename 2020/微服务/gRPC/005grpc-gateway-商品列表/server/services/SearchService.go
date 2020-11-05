package services

import (
	"context"
	"strconv"
)

type SearchService struct {

}

func(s *SearchService)GetArticles(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	resp := &SearchResponse{
		ErrorMsg: "success",
		ErrorCode: 0,
	}
	respData := make([]*ArticleInfo,0)
	for i:=0;i<30;i++ {
		article := &ArticleInfo{
			ArticleId: int32(i),
			Title: "文章标题-->"+strconv.Itoa(i)+" this is very good",
		}
		respData = append(respData, article)
	}

	resp.Data = &ResponseData{
		Rows: respData,
	}
	return resp,nil
}