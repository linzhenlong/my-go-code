package services

import (
	"context"
	"errors"
	"strconv"
)

type SearchService struct {
}

func (s *SearchService) GetArticles(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	resp := &SearchResponse{}
	articles := make([]*ArticleInfo, 0, 100)
	start := request.PageNumber * request.ResultPerPage
	end := (request.PageNumber + 1) * request.ResultPerPage
	if start > 100 {
		resp.ErrorMsg = "error"
		resp.ErrorCode = 1
		return resp, errors.New("越界了")
	}
	for i := 1; i <= 100; i++ {
		article := &ArticleInfo{
			ArticleId: int32(i),
			Title:     "this is title " + strconv.Itoa(i),
		}
		articles = append(articles, article)
	}
	respData := &ResponseData{
		Rows: articles[start:end],
	}
	resp.Data = respData
	return resp, nil
}
