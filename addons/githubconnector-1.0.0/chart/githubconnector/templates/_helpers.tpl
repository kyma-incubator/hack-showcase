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
{{- printf "github-connector-svc" -}}
{{- end -}}
{{- $name := .Release.Name | trimAll "hb-github-connector-" | trunc 47 | lower -}}
