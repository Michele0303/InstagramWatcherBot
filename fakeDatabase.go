package InstagramWatcherBot

var users = make(map[string]int64)

func IsStoryNew(username string, lastReel int64) bool {
	if users[username] == 0 || users[username] < lastReel {
		users[username] = lastReel
		return true
	}

	return false
}
