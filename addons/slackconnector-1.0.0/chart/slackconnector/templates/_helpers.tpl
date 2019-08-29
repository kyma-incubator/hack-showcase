{{- define "slack-connector-chart.release.name" -}}
{{- printf "slack-%s" randAlphaNum | trunc 31 | trimSuffix "-" -}}
{{- end -}}

{{- define "slack-connector-chart.release.service" -}}
{{- default .Release.Service | trunc 31 | trimSuffix "-" -}}
{{- end -}}

{{- define "slack-connector-chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "slack-connector-chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}
