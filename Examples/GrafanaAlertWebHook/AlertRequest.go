package main

type AlertRequest struct {
	PanelID     int    `json:"panelId" gorm:"column:panelId"`
	DashboardID int    `json:"dashboardId" gorm:"column:dashboardId"`
	ImageUrl    string `json:"imageUrl" gorm:"column:imageUrl"`
	RuleName    string `json:"ruleName" gorm:"column:ruleName"`
	State       string `json:"state" gorm:"column:state"`
	Message     string `json:"message" gorm:"column:message"`
	RuleID      int    `json:"ruleId" gorm:"column:ruleId"`
	Title       string `json:"title" gorm:"column:title"`
	RuleUrl     string `json:"ruleUrl" gorm:"column:ruleUrl"`
	OrgID       int    `json:"orgId" gorm:"column:orgId"`

	EvalMatches []struct {
		Metric string   `json:"metric" gorm:"column:metric"`
		Value  int      `json:"value" gorm:"column:value"`
		Tags   struct{} `json:"tags" gorm:"column:tags"`
	} `json:"evalMatches" gorm:"column:evalMatches"`

	Tags struct {
		TagName string `json:"tag name" gorm:"column:tag name"`
	} `json:"tags" gorm:"column:tags"`
}
