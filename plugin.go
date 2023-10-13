package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/CatchZeng/feishu/pkg/feishu"
)

type (
	// Repo repo base info
	Repo struct {
		ShortName string //  short name
		GroupName string //  group name
		FullName  string //  repository full name
		OwnerName string //  repo owner
		RemoteURL string //  repo remote url
	}

	// Build info
	Build struct {
		Status     string //  providers the current build status
		Link       string //  providers the current build link
		Event      string //  trigger event
		StartAt    uint64 //  build start at ( unix timestamp )
		FinishedAt uint64 //  build finish at ( unix timestamp )
	}

	// Commit info
	Commit struct {
		Branch  string //  providers the branch for the current commit
		Link    string //  providers the http link to the current commit in the remote source code management system(e.g.GitHub)
		Message string //  providers the commit message for the current build
		Sha     string //  providers the commit sha for the current build
		Ref     string //  commit ref
		Author  CommitAuthor
	}

	// Stage drone stage env
	Stage struct {
		StartedAt  uint64
		FinishedAt uint64
	}

	// CommitAuthor commit author info
	CommitAuthor struct {
		Avatar   string //  providers the author avatar for the current commit
		Email    string //  providers the author email for the current commit
		Name     string //  providers the author name for the current commit
		Username string //  the author username for the current commit
	}

	// Drone drone info
	Drone struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Stage  Stage
	}

	// Config plugin private config
	Config struct {
		Debug       bool
		AccessToken string
		Secret      string
		IsAtALL     bool
		Mobiles     string
		Username    string
		MsgType     string
		TipsTitle   string
	}

	// Plugin plugin all config
	Plugin struct {
		Drone  Drone
		Config Config
	}

	// Status status
	Status struct {
		Success string
		Failure string
	}

	// Consuming custom consuming env
	Consuming struct {
		StartedEnv  string
		FinishedEnv string
	}
)

// Exec execute WebHook
func (p *Plugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error
	if "" == p.Config.AccessToken {
		msg := "missing feishu access token"
		return errors.New(msg)
	}

	if p.Config.TipsTitle == "" {
		p.Config.TipsTitle = "you have a new message"
	}

	client := feishu.NewClient(p.Config.AccessToken, p.Config.Secret)

	msg := feishu.NewInteractiveMessage()

	pluginEnv := GetPluginEnv()

	card := (Card{}).Build(
		p.Drone.Repo.ShortName,
		p.Drone.Commit.Branch,
		p.Drone.Commit.Author.Username,
		p.Drone.Commit.Author.Email,
		p.Drone.Build.Status,
		p.Drone.Commit.Message,
		p.Drone.Commit.Link,
		p.Drone.Build.Link,
		pluginEnv.PluginCardTitle,
		pluginEnv.PluginSuccessImgKey,
		pluginEnv.PluginFailureImgKey,
		pluginEnv.PluginPoweredByImgKey,
		pluginEnv.PluginPoweredByImgAlt,
	)

	buf, err := json.Marshal(card)

	if err != nil {
		log.Println("json marshal error:", err)
		return err
	}

	msg.SetCard(string(buf))

	log.Println("send message:", string(buf))

	_, _, err = client.Send(msg)

	if err == nil {
		log.Println("send message success!")
	}

	return err
}

// getStatus
func (p *Plugin) getStatus() string {

	return p.Drone.Build.Status
}

// get emoticon
func (p *Plugin) getEmoticon() string {
	emoticons := make(map[string]string)
	emoticons["success"] = ":)"
	emoticons["failure"] = ":("

	emoticon, ok := emoticons[p.Drone.Build.Status]
	if ok {
		return emoticon
	}

	return ":("
}
