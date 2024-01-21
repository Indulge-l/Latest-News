package latest_news

import (
	"context"
	v1 "latest-news/api/latest_news/v1"
)

type ILatestNewsV1 interface {
	GetLatestNewsList(ctx context.Context, req *v1.GetLatestNewsListReq) (*v1.GetLatestNewsListRes, error)
}
