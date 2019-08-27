# GitHub Connector

## Overview

The GitHub Connector is an additional tool to Kyma for handling GitHub events. It registers GitHub API and events in the Application Registry. Then, it converts events incoming from connected GitHub webhooks to the format acceptable by Kyma and forwards them to the Event Bus. The GitHub Connector handles these events:

- new pull request
- new pull request review
- new issue

## Table of Contents

- [Installation](installation.md)
- [Configuration](configuration.md)
- [Example usage](examples/demoscenario.md)
