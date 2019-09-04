## Roadmap

Our scenario consists of two main components: GitHub Connector and Slack Connector, which will be used with Kyma's lambda to help community management. 

---
### Github Connector
* [x] Connector as a Kyma Add-On 
* [ ] Convert all Github webhooks' payloads to AsyncAPI specification standard
* [ ] Improve security
* [x] Setting up the GitHub webhooks from Kyma

### Kyma's Lambda Usage
* [x] Connect to Azure Text Analytics to measure the attitude of comments on connected GitHub repository
* [x] Communicate with Slack Connector

### Slack Connector
* [x] Connector as a Kyma Add-On 
* [x] Send notifications to corresponding channels on Slack
* [x] Handle all of the Slack events 
