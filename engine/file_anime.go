package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	fields "github.com/baranek44/findAniemByScene/cmd/fields"
	"github.com/baranek44/findAniemByScene/cmd/helper"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/muesli/termenv"
)

const (
	file = "https://api.trace.moe/search?anilistInfo"
)

func SearchingAnimeByFile(img string) {
	if _, err := os.Stat(img); os.IsNotExist(err) {
		if err != nil {
			log.Printf("Invalid file path")
		}
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

	imageFile, err := os.Open(img)
	helper.PrintError(err)

	payload := &bytes.Buffer{}
	copy := multipart.NewWriter(payload)
	part, _ := copy.CreateFormFile("image", filepath.Base(img))

	_, err = io.Copy(part, imageFile)
	helper.PrintError(err)

	err = copy.Close()
	helper.PrintError(err)

	resp, err := http.Post(file, copy.FormDataContentType(), payload)
	helper.PrintError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helper.PrintError(err)

	var animeRespone fields.Response
	json.Unmarshal(body, &animeRespone)

	s.Stop()

	fmt.Println("Title: ", animeRespone.Result[0].Anilist.Title.Native)
	helper.Timer()
	fmt.Printf("Similarity: \n")
	helper.SimilarityScene(strconv.FormatFloat(float64(animeRespone.Result[0].Similariti), 'f', 6, 64))
	helper.Timer()
	_, err = fmt.Fprintln(color.Output, "Episode Number: "+color.MagentaString(strconv.Itoa(animeRespone.Result[0].Episode)))
	helper.PrintError(err)
	helper.Timer()
	fmt.Println("Scene from: ", animeRespone.Result[0].From)
	fmt.Println("Scene to: ", animeRespone.Result[0].To)
}
