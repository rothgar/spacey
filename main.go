package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Space struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	IsTicketed       bool      `json:"is_ticketed"`
	Lang             string    `json:"lang"`
	HostIds          []string  `json:"host_ids"`
	ParticipantCount int       `json:"participant_count"`
	SpeakerIds       []string  `json:"speaker_ids"`
	StartedAt        time.Time `json:"started_at"`
	State            string    `json:"state"`
	UpdatedAt        time.Time `json:"updated_at"`
	Title            string    `json:"title"`
	InvitedUserIds   []string  `json:"invited_user_ids"`
	CreatorID        string    `json:"creator_id"`
}

type Spaces struct {
	Spaces []Space `json:"data"`
}

type args struct {
	Queries             []string `arg:"positional,required" help:"Search words in space titles" placeholder:"QUERY"`
	MinimumParticipants int      `arg:"--minp" help:"Minimum number of participants to match" default:"5"`
	MinimumSpeakers     int      `arg:"--mins" help:"Minimum number of speakers to match" default:"1"`
	Output              string   `arg:"-o,--output" help:"Output for notification [table, text]" default:"table"`
}

func (args) Description() string {
	return "Spacey will find Twitter Spaces based on CLI criteria."
}

func main() {

	var args args
	arg.MustParse(&args)

	if len(args.Queries) == 0 {
		fmt.Println("Please add a term to search.")
		return
	}

	t := table.NewWriter()
	if args.Output == "table" {
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.AppendHeader(table.Row{"QUERY", "SPACE", "lang", "Participants", "SPEAKERS", "URL"})
	}

	for _, v := range args.Queries {
		query := strings.ReplaceAll(v, " ", "%20")
		url := "https://api.twitter.com/2/spaces/search?query=" + query + "&state=live&expansions=speaker_ids&space.fields=host_ids%2Cparticipant_count%2Ccreated_at%2Ccreator_id%2Cid%2Clang%2Cinvited_user_ids%2Cspeaker_ids%2Cstarted_at%2Cstate%2Ctitle%2Cupdated_at%2Cscheduled_start%2Cis_ticketed&user.fields=created_at%2Cdescription%2Centities%2Cid%2Clocation%2Cname%2Cpinned_tweet_id%2Cprofile_image_url%2Cprotected%2Cpublic_metrics%2Curl%2Cusername%2Cverified%2Cwithheld"
		method := "GET"
		bearer := "Bearer " + os.Getenv("TWITTER_BEARER_TOKEN")
		req, _ := http.NewRequest(method, url, nil)
		req.Header.Add("Authorization", bearer)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var result Spaces
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		// fmt.Println(result)
		// fmt.Printf("%+v\n", result.Spaces)
		// fmt.Println(args[i], len(result.Spaces))
		if len(result.Spaces) > 0 {
			for _, space := range result.Spaces {
				if space.ParticipantCount >= args.MinimumParticipants && len(space.SpeakerIds) >= args.MinimumSpeakers {
					if args.Output == "text" {
						fmt.Printf("%s with %d participants at https://twitter.com/i/spaces/%s\n", space.Title, space.ParticipantCount, space.ID)
					} else if args.Output == "table" {
						t.AppendRows([]table.Row{
							{v, space.Title, space.Lang, space.ParticipantCount, len(space.SpeakerIds), "https://twitter.com/i/spaces/" + space.ID},
						})
					}
				}
			}
		}
	}

	if args.Output == "table" {
		t.SortBy([]table.SortBy{
			{Name: "Participants", Mode: table.DscNumeric},
		})
		t.AppendSeparator()
		t.Render()
	}
}
