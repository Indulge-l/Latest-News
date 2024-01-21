package latest_news

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"latest-news/internal/model/entity"
)

type BaiduNews struct {
}

func NewBaiduNews() *BaiduNews {
	return &BaiduNews{}
}

func (b *BaiduNews) GetLatestNewsList(ctx context.Context, count int) (list []entity.LatestNews, err error) {
	client := g.Client()
	reqData := map[string]string{
		"form":         "news_webapp",
		"pd":           "webapp",
		"os":           "android",
		"ver":          "6",
		"category_id":  "102",
		"action":       "1",
		"display_time": "0",
		"wf":           "0",
	}
	response, err := client.PostForm(ctx, "https://news.baidu.com/sn/api/feed_channellist", reqData)
	if err != nil {
		g.Log().Errorf(ctx, "[BaiduNews GetLatestNewsList] client.PostForm error: %v", err)
		return nil, err
	}

	baiduNewsRes := &entity.BaiduNewRes{}
	err = json.Unmarshal(response.ReadAll(), &baiduNewsRes)
	if err != nil {
		g.Log().Errorf(ctx, "[BaiduNews GetLatestNewsList] json.Unmarshal error: %v", err)
		return nil, err
	}

	return ToGetBaiduNewsListRes(baiduNewsRes.Data), nil
}

func ToGetBaiduNewsListRes(data entity.BaiduNewsData) []entity.LatestNews {
	newList := make([]entity.LatestNews, 0)
	for _, item := range data.News {
		newList = append(newList, entity.LatestNews{
			Url:     item.Url,
			Source:  item.Site,
			Content: item.LongAbs,
			Title:   item.Title,
			Time:    item.Pulltime,
			Img:     item.Imageurls[0].Url,
		})
	}
	return newList
}
