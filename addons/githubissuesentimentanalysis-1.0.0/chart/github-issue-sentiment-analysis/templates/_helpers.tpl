{{- define "github-issue-sentiment-analysis.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-issue-sentiment-analysis.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "github-issue-sentiment-analysis.repository" -}}
{{- .Values.githubURL -}}
{{- end -}}

{{- define "github-issue-sentiment-analysis.workspace" -}}
{{- .Values.workspaceName -}}
{{- end -}}
