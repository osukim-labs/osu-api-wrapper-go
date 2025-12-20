package interfaces

type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type RequiredMeta struct {
	MainRuleset    int `json:"main_ruleset"`
	NonMainRuleset int `json:"non_main_ruleset"`
}

type NominationsSummary struct {
	Current              int          `json:"current"`
	EligibleMainRulesets []string     `json:"eligible_main_rulesets"`
	RequiredMeta         RequiredMeta `json:"required_meta"`
}

type Availability struct {
	DownloadDisabled bool `json:"download_disabled"`
	MoreInformation  any  `json:"more_information"`
}

type Failtimes struct {
	Fail []int `json:"fail"`
	Exit []int `json:"exit"`
}

type Owners struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Hype struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}
