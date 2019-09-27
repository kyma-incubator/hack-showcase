## Roadmap

Our scenario consists of two main components: GitHub Connector and Slack Connector. Use them with Kyma's lambda function to facilitate community management. 

---
### Github Connector
* [x] The Connector as a Kyma Add-On 
* [X] Converting all Github webhooks' payloads to the AsyncAPI specification standard
* [X] Improving security
* [x] Setting up the GitHub webhooks from Kyma

### Kyma's Lambda Usage
* [x] Connecting to Azure Text Analytics to analyze the sentiment of issues on the connected GitHub repository
* [x] Communication with the Slack Connector
* [x] Communication with the GitHub API

### Slack Connector
* [x] The Connector as a Kyma Add-On 
* [x] Sending notifications to corresponding channels on Slack
* [x] Handling all Slack Events 
