package main

import (
	"flag"
	"net/http"

	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/github"
	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/hook"
	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/registration"

	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/events"
	"github.com/kyma-incubator/github-slack-connectors/github-connector/internal/handlers"

	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

//Config containing all program configs
type Config struct {
	GithubConnectorName string `envconfig:"GITHUB_CONNECTOR_NAME"`
	GithubToken         string `envconfig:"GITHUB_TOKEN"`
	GithubSecret        string `envconfig:"GITHUB_SECRET"`
	KymaAddress         string `envconfig:"KYMA_ADDRESS"`
	Port                string `envconfig:"PORT"`
}

func main() {
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		log.Fatal("Env error: ", err.Error())
	}
	log.Infof("Github Connector Name: %s", conf.GithubConnectorName)
	log.Infof("Github Token: %s", conf.GithubToken)
	log.Infof("Kyma Address: %s", conf.KymaAddress)
	log.Infof("Port: %s", conf.Port)

	log.Info("Registration started.")
	flag.Parse()
	repos := flag.Args()

	builder := registration.NewPayloadBuilder(registration.NewFileReader(), conf.GithubConnectorName, conf.GithubToken)
	id, err := registration.NewApplicationRegistryClient(builder, 5, 10).RegisterService()

	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	webHook := hook.NewHook(conf.KymaAddress)
	secret := conf.GithubSecret
	for i, address := range repos {
		_, err := webHook.Create(conf.GithubToken, address, secret)
		if err != nil {
			log.Printf("Can not create %d webhook on %s adress. Prease create it manualy. Error body: %s", i, address, err.Error())
		} else {
			log.Printf("Webhook %d created: %s", i, address)
		}
	}
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
	log.Info("Webhook created!")

	kyma := events.NewSender(&http.Client{}, events.NewValidator(), "http://event-publish-service.kyma-system:8080/v1/events")
	webhook := handlers.NewWebHookHandler(
		github.NewReceivingEventsWrapper(secret),
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+conf.Port, nil))

	log.Info("Happy GitHub-Connecting!")

}
