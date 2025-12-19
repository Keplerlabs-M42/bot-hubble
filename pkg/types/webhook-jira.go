package types

type JiraUser struct {
	DisplayName string `json:"displayName"`
}

type JiraIssue struct {
	Key    string `json:"key"`
	Fields struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
		Summary string `json:"summary"`
	} `json:"fields"`
}

type WebhookPayload struct {
	IssueEventTypeName string    `json:"issue_event_type_name"`
	User               JiraUser  `json:"user"`
	Issue              JiraIssue `json:"issue"`
}
