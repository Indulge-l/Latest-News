package latest_news

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"latest-news/internal/model/entity"
	"latest-news/internal/service/latest_news"

	v1 "latest-news/api/latest_news/v1"
)

func (l LatestNewsControllerV1) GetLatestNewsList(ctx context.Context, req *v1.GetLatestNewsListReq) (res *v1.GetLatestNewsListRes, err error) {
	res = &v1.GetLatestNewsListRes{Code: 1}
	var newList []entity.LatestNews
	// 获取腾讯新闻
	list, err := latest_news.NewTencentNews().GetLatestNewsList(ctx, req.Count)
	if err != nil {
		g.Log().Errorf(ctx, "[LatestNewsControllerV1 GetLatestNewsList] latest_news.NewTencentNews().GetLatestNewsList error: %v", err)
		res.Code = -1
		g.RequestFromCtx(ctx).Response.WriteJson(res)
		return
	}
	newList = append(newList, list...)

	// 获取百度新闻
	baiduNews, err := latest_news.NewBaiduNews().GetLatestNewsList(ctx, req.Count)
	if err != nil {
		g.Log().Errorf(ctx, "[LatestNewsControllerV1 GetLatestNewsList] latest_news.NewBaiduNews().GetLatestNewsList error: %v", err)
		res.Code = -1
		g.RequestFromCtx(ctx).Response.WriteJson(res)
		return
	}
	newList = append(newList, baiduNews...)
	for _, item := range list {
		res.LatestNewsList = append(res.LatestNewsList, &v1.LatestNewsListItem{
			Url:     item.Url,
			Source:  item.Source,
			Content: item.Content,
			Title:   item.Title,
			Time:    item.Time,
			Img:     item.Img,
		})
	}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return
}
