package main

import (
	"C"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Api struct {
	Kind     string          `json:"kind"`
	Etag     string          `json:"etag"`
	PageInfo json.RawMessage `json:"pageInfo"`
	Items    []Items         `json:"items"`
}

type Items struct {
	Kind       string          `json:"kind"`
	Etag       string          `json:"etag"`
	Id         string          `json:"id"`
	Statistics json.RawMessage `json:"statistics"`
}

type Status struct {
	ViewCount             string `json:"viewCount"`
	CommentCount          string `json:"commentCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

var (
	channelInfo        []byte                                                                                                                         /*Extract json file*/
	userApi            Api                                                                                                                            /*Entire data from json file*/
	userStatus         Status                                                                                                                         /*As metadata which will get useful data*/
	err                error                                                                                                                          /*Get error message*/
	yourYouTubeApiSite string = "https://www.googleapis.com/youtube/v3/channels?part=statistics&id=(Your channel ID)&key=(Your YouTube Data API key)" /*Site of YouTube Data API*/
)

func requestJson() {
	//Download data
	getJson, err := http.Get(yourYouTubeApiSite)
	if err != nil {
		panic("Something wrong from your site, please ensure you are connected Internet. If connected, please double check that you have correct ID and a valid key from Google API")
	}
	defer getJson.Body.Close()
	//Create json file
	saveJson, _ := os.Create("channels.json")
	defer saveJson.Close()
	//Writing latest data
	_, _ = io.Copy(saveJson, getJson.Body)
	//Import json file from downloaded Google API result
	channelInfo, _ = ioutil.ReadFile("./channels.json")
	//Storing in memory
	userApi = Api{}
	//Ripping all content on json file
	_ = json.Unmarshal([]byte(channelInfo), &userApi)
	//Locate data which contain useful data
	userItems := userApi.Items[0]
	//Ripping actual user's data
	_ = json.Unmarshal([]byte(userItems.Statistics), &userStatus)
}

//export getSubscriber
func getSubscriber() int {
	requestJson()
	sub, _ := strconv.Atoi(userStatus.SubscriberCount)
	return sub
}

//export getViews
func getViews() int {
	requestJson()
	views, _ := strconv.Atoi(userStatus.ViewCount)
	return views
}

//export getAvgViewsVideos
func getAvgViewsVideos() int {
	requestJson()
	views, _ := strconv.Atoi(userStatus.ViewCount)
	videos, _ := strconv.Atoi(userStatus.VideoCount)
	return views / videos
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
