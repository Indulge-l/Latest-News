package latest_news

import (
	api_latest_news "latest-news/api/latest_news"
)

type LatestNewsControllerV1 struct{}

func NewLatestNewsControllerV1() api_latest_news.ILatestNewsV1 {
	return &LatestNewsControllerV1{}
}
