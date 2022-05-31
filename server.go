// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type ZhiHuHot struct {
// 	Data []Data `json:"data"`
// 	Paging Paging `json:"paging"`
// 	FreshText string `json:"fresh_text"`
// 	DisplayNum int `json:"display_num"`
// }
// type Author struct {
// 	Type string `json:"type"`
// 	UserType string `json:"user_type"`
// 	ID string `json:"id"`
// 	URLToken string `json:"url_token"`
// 	URL string `json:"url"`
// 	Name string `json:"name"`
// 	Headline string `json:"headline"`
// 	AvatarURL string `json:"avatar_url"`
// }
// type Target struct {
// 	ID int `json:"id"`
// 	Title string `json:"title"`
// 	URL string `json:"url"`
// 	Type string `json:"type"`
// 	Created int `json:"created"`
// 	AnswerCount int `json:"answer_count"`
// 	FollowerCount int `json:"follower_count"`
// 	Author Author `json:"author"`
// 	BoundTopicIds []int `json:"bound_topic_ids"`
// 	CommentCount int `json:"comment_count"`
// 	IsFollowing bool `json:"is_following"`
// 	Excerpt string `json:"excerpt"`
// }
// type Children struct {
// 	Type string `json:"type"`
// 	Thumbnail string `json:"thumbnail"`
// }
// type Data struct {
// 	Type string `json:"type"`
// 	StyleType string `json:"style_type"`
// 	ID string `json:"id"`
// 	CardID string `json:"card_id"`
// 	Target Target `json:"target"`
// 	AttachedInfo string `json:"attached_info"`
// 	DetailText string `json:"detail_text"`
// 	Trend int `json:"trend"`
// 	Debut bool `json:"debut"`
// 	Children []Children `json:"children"`
// }
// type Paging struct {
// 	IsEnd bool `json:"is_end"`
// 	Next string `json:"next"`
// 	Previous string `json:"previous"`
// }

// const (
// 	ZHIHU_HOT_URL = "https://api.zhihu.com/topstory/hot-list?limit=10&reverse_order=0"
// )

// func GetZhihuHot() []Target {
// 	resp, err := http.Get(ZHIHU_HOT_URL)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer resp.Body.Close()

// 	var zhihuHot ZhiHuHot
// 	if err := json.NewDecoder(resp.Body).Decode(&zhihuHot); err != nil {
// 		log.Println(err)
// 	}

// 	var target []Target
// 	for _, v := range zhihuHot.Data {
// 		target = append(target, v.Target)
// 	}

// 	log.Println(target)
// 	return target
// }