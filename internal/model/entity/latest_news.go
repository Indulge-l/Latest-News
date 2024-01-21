package entity

type LatestNews struct {
	Url     string `json:"url"`
	Source  string `json:"source"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Img     string `json:"img"`
}

// TencentNewsRes 腾讯新闻消息数据
type (
	TencentNewsRes struct {
		TraceId string            `json:"trace_id"`
		Data    []TencentNewsData `json:"data"`
	}

	TencentNewsData struct {
		Title          string                    `json:"title"`
		PublishTime    string                    `json:"publish_time"`
		Desc           string                    `json:"desc"`
		LinkInfo       TencentNewsLinkInfo       `json:"link_info"`
		MediaInfo      TencentNewsMediaInfo      `json:"media_info"`
		InterationInfo TencentNewsInterationInfo `json:"interation_info"`
		PicInfo        TencentNewsPicInfo        `json:"pic_info"`
	}

	TencentNewsLinkInfo struct {
		ShareUrl string `json:"share_url"`
		Url      string `json:"url"`
		ShortUrl string `json:"short_url"`
		OrgUrl   string `json:"org_url"`
	}

	TencentNewsMediaInfo struct {
		ChlId         string `json:"chl_id"`
		ChlName       string `json:"chl_name"`
		Icon          string `json:"icon"`
		Sicon         string `json:"sicon"`
		Mrk           string `json:"mrk"`
		LastArtUpdate string `json:"last_art_update"`
		Uin           string `json:"uin"`
		EncodedSuid   string `json:"encoded_suid"`
		VipType       string `json:"vip_type"`
		VipTypeNew    string `json:"vip_type_new"`
		VipDesc       string `json:"vip_desc"`
		VipIcon       string `json:"vip_icon"`
	}

	TencentNewsInterationInfo struct {
		CommentId string `json:"comment_id"`
		CommetNum int    `json:"commet_num"`
		ReadNum   int    `json:"read_num"`
	}

	TencentNewsPicInfo struct {
		BigImg       []string `json:"big_img"`
		SmallImg     []string `json:"small_img"`
		ThreeImg     []string `json:"three_img"`
		LsImgExpType string   `json:"ls_img_exp_type"`
		ShareImg     string   `json:"share_img"`
	}

	TencentReq struct {
		Qimei36   string         `json:"qimei36"`
		BaseReq   TencentBaseReq `json:"base_req"`
		FlushNum  int            `json:"flush_num"`
		ChannelId string         `json:"channel_id"`
		DeviceId  string         `json:"device_id"`
		ItemCount int            `json:"item_count"`
		Forward   string         `json:"forward"`
	}

	TencentBaseReq struct {
		From string `json:"from"`
	}
)

type (
	BaiduNewRes struct {
		Errno     int           `json:"errno"`
		Errmsg    string        `json:"errmsg"`
		Timestamp int           `json:"timestamp"`
		Data      BaiduNewsData `json:"data"`
	}

	BaiduNewsData struct {
		Top  BaiduNewsTop  `json:"top"`
		News BaiduNewsNews `json:"news"`
	}

	BaiduNewsTop []struct {
		Nid              string        `json:"nid"`
		Sourcets         string        `json:"sourcets"`
		Ts               string        `json:"ts"`
		Title            string        `json:"title"`
		Url              string        `json:"url"`
		Imageurls        []interface{} `json:"imageurls"`
		Site             string        `json:"site"`
		Type             interface{}   `json:"type"`
		Abs              string        `json:"abs"`
		DisplayType      int           `json:"display_type"`
		DisplayUrl       string        `json:"display_url"`
		Topic            []interface{} `json:"topic"`
		LongAbs          string        `json:"long_abs"`
		HasRelated       []interface{} `json:"has_related"`
		Content          []interface{} `json:"content"`
		Layout           string        `json:"layout"`
		NType            string        `json:"n_type"`
		Chinajoy         int           `json:"chinajoy"`
		SentimentDisplay int           `json:"sentiment_display"`
	}

	BaiduNewsNews []struct {
		Nid         string          `json:"nid"`
		Sourcets    string          `json:"sourcets"`
		Ts          string          `json:"ts"`
		Title       string          `json:"title"`
		Url         string          `json:"url"`
		Imageurls   []BaiduImageUrl `json:"imageurls"`
		Site        string          `json:"site"`
		Type        string          `json:"type"`
		Abs         string          `json:"abs"`
		DisplayType int             `json:"display_type"`
		DisplayUrl  string          `json:"display_url"`
		Topic       []interface{}   `json:"topic"`
		LongAbs     string          `json:"long_abs"`
		HasRelated  []interface{}   `json:"has_related"`
		Content     []interface{}   `json:"content"`
		//ContentType struct {
		//	Text int `json:"text"`
		//} `json:"content_type"`
		Layout   string `json:"layout"`
		Pulltime string `json:"pulltime"`
		NType    string `json:"n_type"`
		//Ext      struct {
		//	ContentType int `json:"content_type"`
		//	PublicTime  int `json:"public_time"`
		//} `json:"ext"`
		Chinajoy         int `json:"chinajoy"`
		SentimentDisplay int `json:"sentiment_display"`
		//Tags             struct {
		//	Own []struct {
		//		Type int    `json:"type"`
		//		Name string `json:"name"`
		//	} `json:"own"`
		//	Tag []interface{} `json:"tag"`
		//} `json:"tags"`
		//Ctag struct {
		//	Name    string `json:"name"`
		//	RimShow int    `json:"rim_show"`
		//} `json:"ctag"`
		//Emojis struct {
		//} `json:"emojis"`
		//CommentCount string `json:"comment_count"`
		//FeedbackTag  []struct {
		//	Id   int    `json:"id"`
		//	Type int    `json:"type"`
		//	Name string `json:"name"`
		//} `json:"feedback_tag"`
	}

	BaiduImageUrl struct {
		Url     string `json:"url"`
		Height  int    `json:"height"`
		Width   int    `json:"width"`
		UrlWebp string `json:"url_webp"`
	}
)
