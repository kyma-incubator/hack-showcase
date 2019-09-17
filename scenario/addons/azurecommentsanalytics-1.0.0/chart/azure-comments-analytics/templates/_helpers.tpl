{{- define "azure-comments-analytics.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "azure-comments-analytics.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "azure-comments-analytics.repository" -}}
{{- .Values.githubURL | trimAll "repos/" | trimAll "orgs/" | replace "/" "-" | trunc 47 | trimSuffix "-" | printf "github-%s" -}}
{{- end -}}

{{- define "azure-comments-analytics.workspace" -}}
{{- regexReplaceAll "[^a-zA-Z\\d\\w:]" (printf "slack-connector-%s" .Values.workspaceName | lower) "-" | replace "--" "-" | trunc 31 | trimSuffix "-" -}}
{{- end -}}