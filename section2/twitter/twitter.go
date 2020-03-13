package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY_TWITTER")
	consumerSecret = os.Getenv("CONSUMER_SECRET_TWITTER")
	accessToken    = os.Getenv("ACCESS_KEY_TWITTER")
	accessSecret   = os.Getenv("ACCESS_SECRET_TWITTER")
)

type cleanTweet struct {
	ID       string
	Text     string
	Likes    int
	Retweets int
	language string
	URL      string
}

func getCleanTweet(tweet anaconda.Tweet) cleanTweet {
	t := cleanTweet{
		tweet.IdStr,
		tweet.Text,
		tweet.FavoriteCount,
		tweet.RetweetCount,
		tweet.Lang,
		"www.twitter.com/i/web/status" + tweet.IdStr,
	}

	return t
}

func prettyPrintTweet(tweet anaconda.Tweet) {
	t := getCleanTweet(tweet)

	tweetJSON, _ := json.MarshalIndent(t, "", "   ")
	fmt.Println(tweetJSON)
}

func saveTweetsJSON(tweetsJSON []cleanTweet, fileName string) error {
	formatedTweets, _ := json.MarshalIndent(tweetsJSON, "", "   ")
	err := ioutil.WriteFile(fileName, formatedTweets, 0644)

	if err != nil {
		return err
	}

	return nil
}

func loadTweetsFile(fileName string) ([]cleanTweet, error) {
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var cleanTweets []cleanTweet

	err = json.Unmarshal(fileData, &cleanTweets)
	if err != nil {
		return nil, err
	}

	return cleanTweets, nil
}

func main() {
	tweets, err := loadTweetsFile("tweets.json")
	if err == nil {
		fmt.Println("Loading tweets first...")
		for _, t := range tweets {
			formatedTweet, _ := json.MarshalIndent(t, "", "\t")
			fmt.Println(formatedTweet)
		}
	}

	api := anaconda.NewTwitterApiWithCredentials(
		accessToken, accessSecret, consumerKey, consumerSecret)

	fmt.Println("Started the API...")

	searchResult, _ := api.GetSearch("deep learning",
		url.Values{"result_type": []string{"popular"}})

	fmt.Printf("Retrieved %v tweets\n", len(searchResult.Statuses))

	var tweetsForFile []cleanTweet
	for _, tweet := range searchResult.Statuses {
		if tweet.FavoriteCount > 5000 && tweet.RetweetCount > 2000 {
			tweetsForFile = append(tweetsForFile, getCleanTweet(tweet))
		} else {
			fmt.Println("Skipping tweet")
			prettyPrintTweet(tweet)
		}
	}

	err = saveTweetsJSON(tweetsForFile, "tweets.json")
	if err != nil {
		fmt.Println("Error saving tweet")
	} else {
		fmt.Println("Tweets saved")
	}

}
