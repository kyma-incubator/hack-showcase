{{- define "github-connector-chart.release.name" -}}
{{- default .Release.Name | trunc 40 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-connector-chart.release.service" -}}
{{- default .Release.Service | trunc 40 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-connector-chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-connector-chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-connector-chart.repository" -}}
{{- .Values.githubEndpoint | first | trimAll "repos/" | trimAll "orgs/" | replace "/" "-" | trunc 47 | trimSuffix "-" | trimPrefix "-" | lower | printf "github-%s" -}}
{{- end -}}
