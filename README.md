# 🤖 InstagramWatcherBot 

InstagramWatcherBot is a bot that automatically downloads and send to telegram every new instagram stories of users placed in the watchlist👀.

<img src="https://github.com/Michele0303/InstagramWatcherBot/blob/main/assets/home.png" width="850px">

<h2> 🛠 Getting Started </h2>

<ol>
  <li>Download and install <a href="https://go.dev/doc/install">Go 1.20</a></li>
  <li> Install the program by running the command in a shell:</li>
</ol>

<!-- go install github.com/Michele0303/InstagramWatcherBot@latest -->
```bash
# Clone this repository
git clone https://github.com/Michele0303/InstagramWatcherBot
# Go into the repository
cd InstagramWatcherBot
# Build project
go build .
# run on windows
InstagramWatcherBot.exe
# run on linux
./InstagramWatcherBot
```
<ol start="3">
  <li>Open configuration.env and follow this <a href="https://github.com/Michele0303/InstagramWatcherBot#-how-to-get-instagram-cookie-and-header">guide</a></li>
</ol>

```bash
TOKEN_BOT=<your telegram bot token>
CHAT_ID=<your telegram user id> (you can find it on @chatIDrobot)
SESSION_ID=<sessionid cookie>
APP_ID=<x-ig-app-id header>
USER_AGENT=<user-agent header>
WATCH_LIST=<user1;user2;user3>
```

<h2>🍪 How to get Instagram Cookie and Header?</h2>

<ul>
  <li>Go to https://www.instagram.com/yourUser</li>
  <li>Open your Browser Console (on Chrome just pressing F12)</li>
  <li>refresh page and select any request from instagram.com</li>
</ul>

<img src="https://github.com/Michele0303/InstagramWatcherBot/blob/main/assets/cookie.png" width="1000px">

<img src="https://github.com/Michele0303/InstagramWatcherBot/blob/main/assets/headers.png" width="1000px">



