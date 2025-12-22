package interfaces

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	MainRuleset    *int `json:"main_ruleset"`
	NonMainRuleset *int `json:"non_main_ruleset"`
}

type NominationsSummary struct {
	Current              *int          `json:"current"`
	EligibleMainRulesets *[]string     `json:"eligible_main_rulesets"`
	RequiredMeta         *RequiredMeta `json:"required_meta"`
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
	Current  *int `json:"current"`
	Required *int `json:"required"`
}

type FlexibleTime struct {
	time.Time
}

func (ft *FlexibleTime) UnmarshalJSON(b []byte) error {
	// Try to unmarshal as a number (Unix milliseconds)
	var ms int64
	if err := json.Unmarshal(b, &ms); err == nil {
		ft.Time = time.Unix(0, ms*int64(time.Millisecond))
		return nil
	}

	// Try to unmarshal as a string (ISO 8601)
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		if s == "" {
			ft.Time = time.Time{}
			return nil
		}
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return err
		}
		ft.Time = t
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s as FlexibleTime", string(b))
}
