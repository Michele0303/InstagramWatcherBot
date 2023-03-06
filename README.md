# InstagramSpyBot

<h2> What is it? </h2>

Instagram SpyBot is a bot that automatically downloads and send to telegram every new instagram stories of users placed in the watchlist👀.

<h2> 🛠 Configuration </h2>

<ol>
  <li>open configuration.env</li>
</ol>

```bash
TOKEN_BOT=<your telegram bot token>
CHAT_ID=<your user id> (you can find him on @chatIDrobot)
SESSION_ID=<sessionid cookie>
APP_ID=<x-ig-app-id header>
USER_AGENT=<user-agent header>
```

<h2>🍪 How to get Instagram Cookie and Header?</h2>

<ul>
  <li>Go to https://www.instagram.com/yourUser</li>
  <li>Open your Browser Console (on Chrome just pressing F12)</li>
  <li>refresh page and select any request from instagram.com</li>
</ul>

<h4>SESSION_ID </h4>
<img src="/assets/cookie.png" width="1000px">

<h4>APP_ID & USER_AGENT </h4>
<img src="/assets/headers.png" width="1000px">



