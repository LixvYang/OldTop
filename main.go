package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
)

var (
    // Specify the keystore file in the -config parameter
    config = flag.String("config", "", "keystore file path")
)

func main() {
    // Use flag package to parse the parameters
    flag.Parse()

    // Open the keystore file
    f, err := os.Open(*config)
    if err != nil {
        log.Panicln(err)
    }

    // Read the keystore file as json into mixin.Keystore, which is a go struct
    var store mixin.Keystore
    if err := json.NewDecoder(f).Decode(&store); err != nil {
        log.Panicln(err)
    }

    // Create a Mixin Client from the keystore, which is the instance to invoke Mixin APIs
    client, err := mixin.NewFromKeystore(&store)
    if err != nil {
        log.Panicln(err)
    }

	target := GetZhihuHot()
	// go log.Println(target)
    var title []string
    for _, v := range target {
        title = append(title, v.Title)
    }
    titleT := strings.Join(title, "")
    
    go log.Println(titleT)
    // Prepare the message loop that handle every incoming messages,
    // and reply it with the same content.
    // We use a callback function to handle them.
    h := func(ctx context.Context, msg *mixin.MessageView, userID string) error {
        // if there is no valid user id in the message, drop it
        if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
            return nil
        }

        // The incoming message's message ID, which is an UUID.
        id, _ := uuid.FromString(msg.MessageID)

		log.Println(msg.Data)
		log.Println(msg.DataBase64)
        // Create a request
        reply := &mixin.MessageRequest{
            // Reuse the conversation between the sender and the bot.
            // There is an unique UUID for each conversation.
            ConversationID: msg.ConversationID,
            // The user ID of the recipient.
            // Our bot will reply messages, so here is the sender's ID of each incoming message.
            RecipientID: msg.UserID,
            // Create a new message id to reply, it should be an UUID never used by any other message.
            // Create it with a "reply" and the incoming message ID.
            MessageID: uuid.NewV5(id, "reply").String(),
            // Our bot just reply the same category and the sam content of the incoming message
            // So, we copy the category and data
            Category: msg.Category,
            Data:     msg.Data,
        }
        // Send the response
        return client.SendMessage(ctx, reply)
    }

    ctx := context.Background()

    // Start the message loop.
    for {
        // Pass the callback function into the `BlazeListenFunc`
        if err := client.LoopBlaze(ctx, mixin.BlazeListenFunc(h)); err != nil {
            log.Printf("LoopBlaze: %v", err)
        }

        // Sleep for a while
        time.Sleep(time.Second)
    }
}



type ZhiHuHot struct {
	Data []Data `json:"data"`
	Paging Paging `json:"paging"`
	FreshText string `json:"fresh_text"`
	DisplayNum int `json:"display_num"`
}
type Author struct {
	Type string `json:"type"`
	UserType string `json:"user_type"`
	ID string `json:"id"`
	URLToken string `json:"url_token"`
	URL string `json:"url"`
	Name string `json:"name"`
	Headline string `json:"headline"`
	AvatarURL string `json:"avatar_url"`
}
type Target struct {
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	Type string `json:"type"`
	Created int `json:"created"`
	AnswerCount int `json:"answer_count"`
	FollowerCount int `json:"follower_count"`
	Author Author `json:"author"`
	BoundTopicIds []int `json:"bound_topic_ids"`
	CommentCount int `json:"comment_count"`
	IsFollowing bool `json:"is_following"`
	Excerpt string `json:"excerpt"`
}
type Children struct {
	Type string `json:"type"`
	Thumbnail string `json:"thumbnail"`
}
type Data struct {
	Type string `json:"type"`
	StyleType string `json:"style_type"`
	ID string `json:"id"`
	CardID string `json:"card_id"`
	Target Target `json:"target"`
	AttachedInfo string `json:"attached_info"`
	DetailText string `json:"detail_text"`
	Trend int `json:"trend"`
	Debut bool `json:"debut"`
	Children []Children `json:"children"`
}
type Paging struct {
	IsEnd bool `json:"is_end"`
	Next string `json:"next"`
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

	var target []Target
	for _, v := range zhihuHot.Data {
		target = append(target, v.Target)
	}
	return target
}