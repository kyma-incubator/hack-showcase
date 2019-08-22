# GitHub Connector

## Overview

The GitHub Connector is an additional tool for handling GitHub events in Kyma. It registers GitHub API and events in Kyma Application Registry. Then, it converts events incoming from connected GitHub webhooks into format acceptable by Kyma Event Bus and forwards them. For now it handles those events:

- new pull request,
- new issue,
- new review on pull request.

## Table of Contents

- [Installation](installation.md)
- [Configuration](configuration.md)
- [Example usage](examples/demoscenario.md)
