package main

import (
	"fmt"
	"go-nhl/nhlApi"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main(){
	// help timing request time
	now := time.Now()

	rosterFile,err := os.OpenFile("rosters.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err!=nil{
		log.Fatal("error opening rosters.txt: %v",err)
	}
	defer rosterFile.Close()

	rfWritter := io.MultiWriter(os.Stdout,rosterFile)
	log.SetOutput(rfWritter)

	teams, err := nhlApi.GetAlltTeams()
	if err!=nil{
		log.Fatal("error while retrieving teams: %v",err)
	}

	var wg sync.WaitGroup

	wg.Add(len(teams))
	fmt.Println(len(teams))
	results := make(chan []nhlApi.Roster)

	for _,team:=range teams{
		go func(team nhlApi.Team) {
			fmt.Println("fetching roster")
			roster,err := nhlApi.GetRoster(team.ID)
			if err!=nil{
				log.Fatalf("error getting roster: %v",err)
			}

			results <- roster
			wg.Done()
		}(team)
	}

	go func(){
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v",time.Now().Sub(now).String())
}

func display(res chan []nhlApi.Roster) {
	fmt.Println("printing roster")
	for r := range res{
		for _, roster := range r{
			log.Println("------------")
			log.Printf("Name: %s", roster.Person.FullName)
			log.Printf("ID: %d", roster.Person.ID)
			log.Println("------------")
		}
	}
}
