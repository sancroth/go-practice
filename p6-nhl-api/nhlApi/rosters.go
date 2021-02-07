package nhlApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Roster struct {
	Person struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}

type nhlRosterResponse struct{
	Rosters []Roster `json:"roster"`
}

func GetRoster(teamID int) ([]Roster, error){
	res,err := http.Get(fmt.Sprintf("%s/teams/%d/roster",BASE_URL,teamID))
	log.Println(fmt.Sprintf("%s/teams/%d/roster",BASE_URL,teamID))
	if err!=nil{
		return nil,err
	}
	defer res.Body.Close()
	//bodyBytes, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bodyString := string(bodyBytes)
	//log.Printf("body : %v",bodyString)
	var rosterResp nhlRosterResponse
	err = json.NewDecoder(res.Body).Decode(&rosterResp)
	if err!=nil{
		log.Fatal(err)
	}

	return rosterResp.Rosters,err
}