package main

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/jmoiron/jsonq"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const (
	PHOTO = 1
	VIDEO = 2
)

var baseUrl = "https://www.instagram.com"

var client *req.Client = req.C()

func runBot(watchlist map[string]string) {

	fmt.Println("\n[*] Bot Started!\n")

	// get ids from usernames
	for username := range watchlist {
		id := getIdFromUser(username)
		if id == "" {
			fmt.Printf("\n[-] Unable to find the ID for user: " + username)
			continue
		}
		watchlist[username] = id
	}

	for {
		// start monitoring
		scanUsers(watchlist)
		time.Sleep(60 * time.Second)
	}

}

func scanUsers(watchlist map[string]string) {
	for username := range watchlist {
		time.Sleep(30 * time.Second)

		mediaUrl, mediaType := getLastReel(username, watchlist[username])
		if mediaUrl == "" {
			continue
		}

		if mediaType == PHOTO {
			sendPhoto(mediaUrl, username)
		} else {
			sendVideo(mediaUrl, username)
		}
		fmt.Printf("[+] [%s] New instagam story sent!\n", username)
	}
}

func getIdFromUser(user string) string {
	// Realizar la solicitud HTTP para obtener los datos del usuario
	resp, err := client.R().
		Get("/web/search/topsearch/?context=blended&query=" + user + "&rank_token=0.3953592318270893&count=1")
	if err != nil {
		log.Fatalf("Error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Leer y decodificar la respuesta JSON
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
	}

	// Verificar si la respuesta contiene la estructura esperada
	users, ok := data["users"].([]interface{})
	if !ok || len(users) == 0 {
		log.Fatalf("No se encontró el campo 'users' en la respuesta JSON")
	}

	userObj, ok := users[0].(map[string]interface{})
	if !ok {
		log.Fatalf("No se pudo convertir 'users[0]' a un objeto map[string]interface{}")
	}

	userNode, ok := userObj["user"].(map[string]interface{})
	if !ok {
		log.Fatalf("No se pudo convertir 'user' a un objeto map[string]interface{}")
	}

	id, ok := userNode["pk"].(string)
	if !ok {
		log.Fatalf("No se encontró la clave 'pk' o no es una cadena")
	}

	return id
}

func parseJson(jsonStr string, username string) (mediaUrl string, mediaType int) {

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	// get last stories
	lastReel, err := jq.Int("reels_media", "0", "latest_reel_media")
	if err != nil {
		return "", 0
	}

	// check if I have already downloaded it
	if !IsStoryNew(username, int64(lastReel)) {
		return "", 0
	}

	items, _ := jq.Array("reels_media", "0", "items")
	itemsCount := len(items)
	lastItem := fmt.Sprintf("%v", items[itemsCount-1])

	// check if it's a video or a photo
	if strings.Contains(lastItem, "media_type:1") {
		mediaType = PHOTO
		mediaUrl = strings.Split(lastItem, "image_versions2")[1]
	} else {
		mediaType = VIDEO
		mediaUrl = strings.Split(lastItem, "video_versions")[1]
	}

	mediaUrl = strings.Split(mediaUrl, "url:")[1]
	mediaUrl = strings.Split(mediaUrl, " width")[0]
	mediaUrl = strings.ReplaceAll(mediaUrl, "\\u0026", "&")

	return mediaUrl, mediaType
}

func getLastReel(username string, id string) (string, int) {
	resp, err := client.R().Get("/api/v1/feed/reels_media/?reel_ids=" + id)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return parseJson(string(content), username)
}

func setupHttpClient(sessionID string, appID string, userAgent string) {
	// make cookieJar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	// set cookie in cookieJar
	urlObj, _ := url.Parse(baseUrl)
	jar.SetCookies(urlObj, []*http.Cookie{
		{
			Name:   "sessionid",
			Value:  sessionID,
			Path:   "/",
			Domain: "www.instagram.com",
		},
		{
			Name:   "ds_user_id",
			Value:  strings.Split(sessionID, "%")[0],
			Path:   "/",
			Domain: "www.instagram.com",
		},
	})

	// make client http
	client = req.C().
		SetCommonHeaderNonCanonical("x-ig-app-id", appID).
		SetUserAgent(userAgent).
		SetCookieJar(jar).
		SetBaseURL(baseUrl)
}
