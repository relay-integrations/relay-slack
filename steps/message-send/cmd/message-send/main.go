package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/puppetlabs/relay-sdk-go/pkg/log"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
	"github.com/slack-go/slack"
)

type ConnectionSpec struct {
	APIToken string `spec:"apiToken"`
}

type Spec struct {
	// New-form API Token.
	Connection *ConnectionSpec

	Channel string

	Message string
	Blocks  string

	Username string
}

func main() {
	u, err := taskutil.MetadataSpecURL()
	if err != nil {
		log.FatalE(err)
	}
	specURL := flag.String("spec-url", u, "url to fetch the spec from")

	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}

	var spec Spec
	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		log.FatalE(err)
	}

	if spec.Connection == nil {
		log.Fatal("specify the Slack connection to use")
	} else if spec.Connection.APIToken == "" {
		log.Fatal("the specified connection must be a Slack connection")
	} else if spec.Message == "" || spec.Blocks == "" {
		log.Fatal("specify the message, or the api block, to send to Slack")
	} else if spec.Channel == "" {
		log.Fatal("specify the channel to send the message to")
	}

	if spec.Username == "" {
		spec.Username = "Relay by Puppet"
	}

	options := make([]slack.MsgOption, 0, 1)
	if spec.Blocks != "" {
		blocks := slack.Blocks{}
		err = json.Unmarshal([]byte(spec.Blocks), &blocks)
		if err != nil {
			log.FatalE(fmt.Errorf("Cannot unmarshal %s: %w", spec.Blocks, err))
		}
		options = append(options, slack.MsgOptionBlocks(blocks.BlockSet...))
	}
	if spec.Message != "" {
		options = append(options, slack.MsgOptionText(spec.Message, false))
	}
	options = append(options, slack.MsgOptionUsername(spec.Username))

	api := slack.New(spec.Connection.APIToken)
	_, _, err = api.PostMessage(spec.Channel, options...)
	if err != nil {
		log.FatalE(err)
	}
}
