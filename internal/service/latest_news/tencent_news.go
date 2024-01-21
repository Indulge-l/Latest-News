package latest_news

import (
	"context"
	"encoding/json"

	"latest-news/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type TencentNews struct {
}

func NewTencentNews() *TencentNews {
	return &TencentNews{}
}

func (s *TencentNews) GetLatestNewsList(ctx context.Context, count int) ([]entity.LatestNews, error) {
	client := g.Client()
	reqData := &entity.TencentReq{
		Qimei36:   "1_2234942705",
		BaseReq:   entity.TencentBaseReq{From: "pc"},
		FlushNum:  1,
		ChannelId: "news_news_top",
		DeviceId:  "1_2234942705",
		ItemCount: 1,
		Forward:   "2",
	}
	reqDataBytes, err := json.Marshal(reqData)
	if err != nil {
		g.Log().Errorf(ctx, "[TencentNews GetLatestNewsList] json.Marshal error: %v", err)
		return nil, err
	}

	response, err := client.Post(ctx, "https://i.news.qq.com/web_feed/getHotModuleList", reqDataBytes)
	tencentNewRes := &entity.TencentNewsRes{}

	err = json.Unmarshal(response.ReadAll(), &tencentNewRes)
	if err != nil {
		g.Log().Errorf(ctx, "[TencentNews GetLatestNewsList] json.Unmarshal error: %v", err)
		return nil, err
	}
	return ToGetTencentNewsListRes(tencentNewRes.Data), nil
}

func ToGetTencentNewsListRes(list []entity.TencentNewsData) []entity.LatestNews {
	newList := make([]entity.LatestNews, 0)
	for _, item := range list {
		newList = append(newList, entity.LatestNews{
			Url:     item.LinkInfo.Url,
			Source:  item.MediaInfo.ChlName,
			Content: item.Desc,
			Title:   item.Title,
			Time:    item.PublishTime,
			Img:     item.PicInfo.SmallImg[0],
		})
	}
	return newList
}
