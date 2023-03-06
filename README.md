# InstagramSpyBot

<h2> What is it? </h2>

Instagram SpyBot is a bot that automatically downloads and send to telegram every new instagram stories of users placed in the watchlist👀.

<h2> Configuration </h2>

<ol>
  <li>open configuration.env</li>
</ol>

```bash
TOKEN_BOT=<your telegram bot token>
CHAT_ID=<your user id> (you can find him on @chatIDrobot)
SESSION_ID=<sessionid cookie> (1)
APP_ID=<x-ig-app-id header> (2)
USER_AGENT=<user-agent header> (3)
```

<h2>🍪 How to get Instagram Cookie and Header?</h2>



<ul>
  <li>Go to https://www.instagram.com/yourUser</li>
  <li>Open your Browser Console (on Chrome just pressing F12)</li>
  <li>refresh page and select any request from instagram.com</li>
</ul>
