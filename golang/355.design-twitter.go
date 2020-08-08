import "sort"

/*
 * @lc app=leetcode id=355 lang=golang
 *
 * [355] Design Twitter
 */

// @lc code=start

type Tweet struct {
	id   int
	time int
}

type Tweets []Tweet

type Twitter struct {
	timeCnt     int
	usrFollowee map[int]map[int]bool
	usrTweets   map[int][]Tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	var twitter = Twitter{0, make(map[int]map[int]bool, 0), make(map[int][]Tweet, 0)}
	return twitter
}

func ConstructorTweet(tweetId, cnt int) Tweet {
	var tweet = Tweet{tweetId, cnt}
	return tweet
}

/** Initialize your data structure here. */

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.usrTweets[userId] = append(this.usrTweets[userId], ConstructorTweet(tweetId, this.timeCnt))
	this.timeCnt++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	allTweets := []Tweet{}
	allTweets = append(allTweets, this.usrTweets[userId]...)
	if this.usrFollowee != nil {
		for i := range this.usrFollowee[userId] {
			if this.usrTweets[i] != nil && this.usrFollowee[userId][i] && i != userId {
				allTweets = append(allTweets, this.usrTweets[i]...)
			}
		}
	}

	sort.Sort(Tweets(allTweets))
	if len(allTweets) < 10 {
		return Convert(allTweets)
	}
	return Convert(allTweets[:10])
}

func Convert(tweets []Tweet) []int {
	res := make([]int, len(tweets))
	for i := range tweets {
		res[i] = tweets[i].id
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.usrFollowee[followerId] == nil {
		this.usrFollowee[followerId] = make(map[int]bool)
	}
	this.usrFollowee[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.usrFollowee[followerId] == nil {
		this.usrFollowee[followerId] = make(map[int]bool)
	}
	this.usrFollowee[followerId][followeeId] = false
}

/* tweets sort interface */
func (tweets Tweets) Len() int {
	return len(tweets)
}

func (tweets Tweets) Swap(i, j int) {
	tweets[i], tweets[j] = tweets[j], tweets[i]
}

func (tweets Tweets) Less(i, j int) bool {
	if tweets[i].time > tweets[j].time {
		return true
	}
	return false
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end
