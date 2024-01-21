package latest_news

import (
	"context"
	"latest-news/internal/model/entity"
)

type LatestNewsV1 interface {
	GetLatestNewsList(ctx context.Context, count int) ([]entity.LatestNews, error)
}
