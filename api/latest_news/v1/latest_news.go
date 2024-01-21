package v1

import "github.com/gogf/gf/v2/frame/g"

type GetLatestNewsListReq struct {
	g.Meta `path:"/v1/getLatestNewsList" tags:"Hello" method:"post" summary:"You first hello api"`
	Count  int `json:"count"`
}
type GetLatestNewsListRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	Code           int                   `json:"code"`
	LatestNewsList []*LatestNewsListItem `json:"latest_news_list"`
}

type LatestNewsListItem struct {
	Url     string `json:"url"`
	Source  string `json:"source"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Img     string `json:"img"`
}
