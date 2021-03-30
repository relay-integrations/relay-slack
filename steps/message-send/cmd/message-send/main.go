package main

import (
	"flag"

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

	Channel  string
	Message  string
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
	} else if spec.Message == "" {
		log.Fatal("specify the message to send to Slack")
	} else if spec.Channel == "" {
		log.Fatal("specify the channel to send the message to")
	}

	if spec.Username == "" {
		spec.Username = "Relay by Puppet"
	}

	api := slack.New(spec.Connection.APIToken)
	_, _, err = api.PostMessage(spec.Channel, slack.MsgOptionText(spec.Message, false), slack.MsgOptionUsername(spec.Username))
	if err != nil {
		log.FatalE(err)
	}
}
