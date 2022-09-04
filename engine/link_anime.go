package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	fields "github.com/baranek44/findAniemByScene/cmd/fields"
	"github.com/baranek44/findAniemByScene/cmd/helper"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/muesli/termenv"
)

const (
	linkSearchURL = "https://api.trace.moe/search?anilistInfo&url="
)

func SearchingAnimeByLink(img string) {
	_, err := url.ParseRequestURI(img)
	if err != nil {
		log.Panic("Error URL")
	}

	termenv.HideCursor()
	defer termenv.ShowCursor()

	s := spinner.New(spinner.CharSets[10], 100*time.Millisecond)
	s.Prefix = "Looking for your anime!!!: "
	s.FinalMSG = "Complete!\n"
	s.Start()
	time.Sleep(4 * time.Second)
	s.UpdateCharSet(spinner.CharSets[1])
	s.Prefix = "Printing output: "
	s.Stop()

	go Interrupt(s)

	s.Start()

	requestBody, err := json.Marshal(map[string]string{})
	helper.PrintError(err)

	response, err := http.Post(linkSearchURL+img, "application/json", bytes.NewBuffer(requestBody))
	helper.PrintError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	helper.PrintError(err)

	var animeRespone fields.Response
	json.Unmarshal(body, &animeRespone)

	s.Stop()

	fmt.Println("Title: ", animeRespone.Result[0].Anilist.Title.Native)
	helper.Timer()
	fmt.Printf("Similarity: \n")
	helper.SimilarityScene(strconv.FormatFloat(float64(animeRespone.Result[0].Similariti), 'f', 6, 64))
	_, err = fmt.Fprintln(color.Output, "Episode Number: "+color.MagentaString(strconv.Itoa(animeRespone.Result[0].Episode)))
	helper.PrintError(err)
	helper.Timer()
	fmt.Println("Scene from: ", animeRespone.Result[0].From)
	fmt.Println("Scene to: ", animeRespone.Result[0].To)
}
