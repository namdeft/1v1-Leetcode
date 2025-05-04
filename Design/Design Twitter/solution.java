class Twitter {
    private Map<Integer, User> mp;

    public Twitter() {
        mp = new HashMap<Integer, User>();
    }

    public void postTweet(int userId, int tweetId) {
        if (!mp.containsKey(userId)) {
            User user = new User(userId);
            mp.put(userId, user);
        }
        mp.get(userId).postTweet(tweetId);
    }

    public List<Integer> getNewsFeed(int userId) {
        List<Integer> results = new LinkedList<Integer>();
        if (!mp.containsKey(userId)) {
            return results;
        }

        User user = mp.get(userId);
        Set<Integer> followingList = user.getFollowingList();
        Queue<Tweet> queue = new PriorityQueue<Tweet>(followingList.size(), (a, b) -> (b.getTime() - a.getTime()));
        for (int id : followingList) {
            Tweet tweet = mp.get(id).getTweetHead();
            if (tweet != null) {
                queue.add(tweet);
            }
        }

        int count = 0;
        while (count < 10 && !queue.isEmpty()) {
            Tweet tweet = queue.poll();
            results.add(tweet.getId());
            if (tweet.getNext() != null) {
                queue.add(tweet.getNext());
            }
            count++;
        }

        return results;
    }

    public void follow(int followerId, int followeeId) {
        if (!mp.containsKey(followerId)) {
            User user = new User(followerId);
            mp.put(followerId, user);
        }
        if (!mp.containsKey(followeeId)) {
            User user = new User(followeeId);
            mp.put(followeeId, user);
        }
        mp.get(followerId).follow(followeeId);
    }

    public void unfollow(int followerId, int followeeId) {
        if (!mp.containsKey(followerId)) {
            return;
        }
        mp.get(followerId).unfollow(followeeId);
    }
}

class Tweet {
    private static int timeStamp = 0;

    private int id;
    private int time;
    private Tweet next;

    Tweet(int id) {
        this.id = id;
        this.time = timeStamp++;
    }

    public int getId() {
        return this.id;
    }

    public int getTime() {
        return this.time;
    }

    public Tweet getNext() {
        return this.next;
    }

    public void setNext(Tweet tweet) {
        this.next = tweet;
    }
}

class User {
    private int id;
    private Set<Integer> followingList;
    private Tweet tweetHead;

    User(int id) {
        this.id = id;
        this.followingList = new HashSet<Integer>();
        this.tweetHead = null;
        follow(id);
    }

    public int getId() {
        return this.id;
    }

    public Set<Integer> getFollowingList() {
        return this.followingList;
    }

    public Tweet getTweetHead() {
        return this.tweetHead;
    }

    public void postTweet(int tweetId) {
        Tweet tweet = new Tweet(tweetId);
        tweet.setNext(tweetHead);
        this.tweetHead = tweet;
    }

    public void follow(int userId) {
        followingList.add(userId);
    }

    public void unfollow(int userId) {
        followingList.remove(userId);
    }
}

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */