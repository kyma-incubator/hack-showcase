{{- define "azure-comments-analytics.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "azure-comments-analytics.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "azure-comments-analytics.repository" -}}
{{- .Values.githubURL -}}
{{- end -}}

{{- define "azure-comments-analytics.workspace" -}}
{{- .Values.workspaceName -}}
{{- end -}}