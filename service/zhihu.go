package service

import (
	"encoding/json"
	"log"
	"net/http"
)

type OldTopService interface {
	GetOldTop() []Target
}

type ZhiHuHot struct {
	Data       []Data `json:"data"`
	Paging     Paging `json:"paging"`
	FreshText  string `json:"fresh_text"`
	DisplayNum int    `json:"display_num"`
}
type Target struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Excerpt       string `json:"excerpt"`
}
type Children struct {
	Type      string `json:"type"`
	Thumbnail string `json:"thumbnail"`
}
type Data struct {
	Type         string     `json:"type"`
	StyleType    string     `json:"style_type"`
	ID           string     `json:"id"`
	CardID       string     `json:"card_id"`
	Target       Target     `json:"target"`
	AttachedInfo string     `json:"attached_info"`
	DetailText   string     `json:"detail_text"`
	Trend        int        `json:"trend"`
	Debut        bool       `json:"debut"`
	Children     []Children `json:"children"`
}
type Paging struct {
	IsEnd    bool   `json:"is_end"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

const (
	ZHIHU_HOT_URL = "https://api.zhihu.com/topstory/hot-list?limit=10&reverse_order=0"
)

func GetZhihuHot() []Target {
	resp, err := http.Get(ZHIHU_HOT_URL)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	var zhihuHot ZhiHuHot
	if err := json.NewDecoder(resp.Body).Decode(&zhihuHot); err != nil {
		log.Println(err)
	}

	var targetList []Target
	for _, v := range zhihuHot.Data {
		targetList = append(targetList, v.Target)
	}

	return targetList
}
