// Source : https://leetcode.com/problems/design-twitter
// Author : BradleyZhang
// Date   : 2026-10-03

/*****************************************************************************************************
 *
 * Design a simplified version of Twitter where users can post tweets, follow/unfollow another user,
 * and is able to see the 10 most recent tweets in the user's news feed.
 *
 * Implement the Twitter class:
 *
 * 	Twitter() Initializes your twitter object.
 * 	void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user
 * userId. Each call to this function will be made with a unique tweetId.
 * 	List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's
 * news feed. Each item in the news feed must be posted by users who the user followed or by the user
 * themself. Tweets must be ordered from most recent to least recent.
 * 	void follow(int followerId, int followeeId) The user with ID followerId started following
 * the user with ID followeeId.
 * 	void unfollow(int followerId, int followeeId) The user with ID followerId started
 * unfollowing the user with ID followeeId.
 *
 * Example 1:
 *
 * Input
 * ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow",
 * "getNewsFeed"]
 * [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
 * Output
 * [null, null, [5], null, null, [6, 5], null, [5]]
 *
 * Explanation
 * Twitter twitter = new Twitter();
 * twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
 * twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return
 * [5]
 * twitter.follow(1, 2);    // User 1 follows user 2.
 * twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
 * twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5].
 * Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
 * twitter.unfollow(1, 2);  // User 1 unfollows user 2.
 * twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since
 * user 1 is no longer following user 2.
 *
 * Constraints:
 *
 * 	1 <= userId, followerId, followeeId <= 500
 * 	0 <= tweetId <= 10^4
 * 	All the tweets have unique IDs.
 * 	At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
 * 	A user cannot follow himself.
 ******************************************************************************************************/
package designtwitter

import (
	"container/heap"
)

type Twitter struct {
	id2u map[int]*User
	time int
}
type User struct {
	Id        int
	following map[int]struct{}
	posts     []Post
}
type Post struct {
	Id   int
	time int
}

func Constructor() Twitter {
	return Twitter{
		id2u: make(map[int]*User),
	}
}
func (this *Twitter) userExist(userId int) {
	if _, ok := this.id2u[userId]; !ok {
		u := User{
			Id:        userId,
			following: map[int]struct{}{},
			posts:     []Post{},
		}
		this.id2u[userId] = &u
	}
}
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userExist(userId)

	now := this.time
	this.time++
	post := Post{
		time: now,
		Id:   tweetId,
	}
	posts := this.id2u[userId].posts
	if len(posts) > 9 {
		posts = posts[1:]
	}
	posts = append(posts, post)
	this.id2u[userId].posts = posts
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.userExist(userId)
	h := MinHeap(append([]Post{}, this.id2u[userId].posts...)) // copy and init heap
	heap.Init(&h)
	for uid := range this.id2u[userId].following {
		for _, p := range this.id2u[uid].posts {
			if h.Len() < 10 {
				heap.Push(&h, p)
			} else if p.time > h[0].time {
				heap.Pop(&h)
				heap.Push(&h, p)
			}
		}
	}
	var ans []int
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(&h).(Post).Id)
	}
	// reverse
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	this.userExist(followeeId)
	this.userExist(followerId)

	this.id2u[followerId].following[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.userExist(followeeId)
	this.userExist(followerId)
	delete(this.id2u[followerId].following, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type MinHeap []Post

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[j].time > h[i].time
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Post))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
