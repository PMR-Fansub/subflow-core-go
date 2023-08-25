package dto

type TeamInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	QQGroup string `json:"QQGroup"`
	Logo    string `json:"logo"`
	Desc    string `json:"desc"`
}
