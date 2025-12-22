package interfaces

type BeatmapSet struct {
	AnimeCover         bool                 `json:"anime_cover"`
	Artist             string               `json:"artist"`
	ArtistUnicode      string               `json:"artist_unicode"`
	Covers             Covers               `json:"covers"`
	Creator            string               `json:"creator"`
	FavouriteCount     int                  `json:"favourite_count"`
	GenreID            int                  `json:"genre_id"`
	Hype               Hype                 `json:"hype"`
	ID                 int                  `json:"id"`
	LanguageID         int                  `json:"language_id"`
	Nsfw               bool                 `json:"nsfw"`
	Offset             int                  `json:"offset"`
	PlayCount          int                  `json:"play_count"`
	PreviewURL         string               `json:"preview_url"`
	Source             string               `json:"source"`
	Spotlight          bool                 `json:"spotlight"`
	Status             string               `json:"status"`
	Title              string               `json:"title"`
	TitleUnicode       string               `json:"title_unicode"`
	TrackID            int                  `json:"track_id"`
	UserID             int                  `json:"user_id"`
	Video              bool                 `json:"video"`
	Bpm                float64              `json:"bpm"`
	CanBeHyped         bool                 `json:"can_be_hyped"`
	DeletedAt          FlexibleTime         `json:"deleted_at"`
	DiscussionEnabled  bool                 `json:"discussion_enabled"`
	DiscussionLocked   bool                 `json:"discussion_locked"`
	IsScoreable        bool                 `json:"is_scoreable"`
	LastUpdated        FlexibleTime         `json:"last_updated"`
	LegacyThreadURL    string               `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary   `json:"nominations_summary"`
	Ranked             int                  `json:"ranked"`
	RankedDate         FlexibleTime         `json:"ranked_date"`
	Rating             float64              `json:"rating"`
	Storyboard         bool                 `json:"storyboard"`
	SubmittedDate      FlexibleTime         `json:"submitted_date"`
	Tags               string               `json:"tags"`
	Availability       Availability         `json:"availability"`
	Beatmaps           []Beatmaps           `json:"beatmaps"`
	Converts           []Converts           `json:"converts"`
	CurrentNominations []CurrentNominations `json:"current_nominations"`
	Description        Description          `json:"description"`
	Genre              Genre                `json:"genre"`
	Language           Language             `json:"language"`
	PackTags           []string             `json:"pack_tags"`
	Ratings            []int                `json:"ratings"`
	RecentFavourites   []RecentFavourites   `json:"recent_favourites"`
	RelatedUsers       []RelatedUsers       `json:"related_users"`
	RelatedTags        []RelatedTags        `json:"related_tags"`
	User               User                 `json:"user"`
	VersionCount       int                  `json:"version_count"`
}

type TopTagIds struct {
	TagID int `json:"tag_id"`
	Count int `json:"count"`
}

type Beatmaps struct {
	BeatmapsetID         int          `json:"beatmapset_id"`
	DifficultyRating     float64      `json:"difficulty_rating"`
	ID                   int          `json:"id"`
	Mode                 string       `json:"mode"`
	Status               string       `json:"status"`
	TotalLength          int          `json:"total_length"`
	UserID               int          `json:"user_id"`
	Version              string       `json:"version"`
	Accuracy             int          `json:"accuracy"`
	Ar                   int          `json:"ar"`
	Bpm                  float64      `json:"bpm"`
	Convert              bool         `json:"convert"`
	CountCircles         int          `json:"count_circles"`
	CountSliders         int          `json:"count_sliders"`
	CountSpinners        int          `json:"count_spinners"`
	Cs                   int          `json:"cs"`
	DeletedAt            FlexibleTime `json:"deleted_at"`
	Drain                int          `json:"drain"`
	HitLength            int          `json:"hit_length"`
	IsScoreable          bool         `json:"is_scoreable"`
	LastUpdated          FlexibleTime `json:"last_updated"`
	ModeInt              int          `json:"mode_int"`
	Passcount            int          `json:"passcount"`
	Playcount            int          `json:"playcount"`
	Ranked               int          `json:"ranked"`
	URL                  string       `json:"url"`
	Checksum             string       `json:"checksum"`
	CurrentUserPlaycount int          `json:"current_user_playcount"`
	CurrentUserTagIds    []any        `json:"current_user_tag_ids"`
	Failtimes            Failtimes    `json:"failtimes"`
	MaxCombo             int          `json:"max_combo"`
	Owners               []Owners     `json:"owners"`
	TopTagIds            []TopTagIds  `json:"top_tag_ids"`
}

type Converts struct {
	BeatmapsetID     int          `json:"beatmapset_id"`
	DifficultyRating float64      `json:"difficulty_rating"`
	ID               int          `json:"id"`
	Mode             string       `json:"mode"`
	Status           string       `json:"status"`
	TotalLength      int          `json:"total_length"`
	UserID           int          `json:"user_id"`
	Version          string       `json:"version"`
	Accuracy         int          `json:"accuracy"`
	Ar               int          `json:"ar"`
	Bpm              float64      `json:"bpm"`
	Convert          bool         `json:"convert"`
	CountCircles     int          `json:"count_circles"`
	CountSliders     int          `json:"count_sliders"`
	CountSpinners    int          `json:"count_spinners"`
	Cs               int          `json:"cs"`
	DeletedAt        FlexibleTime `json:"deleted_at"`
	Drain            int          `json:"drain"`
	HitLength        int          `json:"hit_length"`
	IsScoreable      bool         `json:"is_scoreable"`
	LastUpdated      FlexibleTime `json:"last_updated"`
	ModeInt          int          `json:"mode_int"`
	Passcount        int          `json:"passcount"`
	Playcount        int          `json:"playcount"`
	Ranked           int          `json:"ranked"`
	URL              string       `json:"url"`
	Checksum         string       `json:"checksum"`
	Failtimes        Failtimes    `json:"failtimes"`
	Owners           []Owners     `json:"owners"`
	TopTagIds        []TopTagIds  `json:"top_tag_ids"`
}

type CurrentNominations struct {
	BeatmapsetID int      `json:"beatmapset_id"`
	Rulesets     []string `json:"rulesets"`
	Reset        bool     `json:"reset"`
	UserID       int      `json:"user_id"`
}

type Description struct {
	Description string `json:"description"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RecentFavourites struct {
	AvatarURL     string       `json:"avatar_url"`
	CountryCode   string       `json:"country_code"`
	DefaultGroup  string       `json:"default_group"`
	ID            int          `json:"id"`
	IsActive      bool         `json:"is_active"`
	IsBot         bool         `json:"is_bot"`
	IsDeleted     bool         `json:"is_deleted"`
	IsOnline      bool         `json:"is_online"`
	IsSupporter   bool         `json:"is_supporter"`
	LastVisit     FlexibleTime `json:"last_visit"`
	PmFriendsOnly bool         `json:"pm_friends_only"`
	ProfileColour string       `json:"profile_colour"`
	Username      string       `json:"username"`
}

type RelatedUsers struct {
	AvatarURL     string       `json:"avatar_url"`
	CountryCode   string       `json:"country_code"`
	DefaultGroup  string       `json:"default_group"`
	ID            int          `json:"id"`
	IsActive      bool         `json:"is_active"`
	IsBot         bool         `json:"is_bot"`
	IsDeleted     bool         `json:"is_deleted"`
	IsOnline      bool         `json:"is_online"`
	IsSupporter   bool         `json:"is_supporter"`
	LastVisit     FlexibleTime `json:"last_visit"`
	PmFriendsOnly bool         `json:"pm_friends_only"`
	ProfileColour string       `json:"profile_colour"`
	Username      string       `json:"username"`
}

type RelatedTags struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	RulesetID   int          `json:"ruleset_id"`
	Description string       `json:"description"`
	CreatedAt   FlexibleTime `json:"created_at"`
	UpdatedAt   FlexibleTime `json:"updated_at"`
}

type User struct {
	AvatarURL     string       `json:"avatar_url"`
	CountryCode   string       `json:"country_code"`
	DefaultGroup  string       `json:"default_group"`
	ID            int          `json:"id"`
	IsActive      bool         `json:"is_active"`
	IsBot         bool         `json:"is_bot"`
	IsDeleted     bool         `json:"is_deleted"`
	IsOnline      bool         `json:"is_online"`
	IsSupporter   bool         `json:"is_supporter"`
	LastVisit     FlexibleTime `json:"last_visit"`
	PmFriendsOnly bool         `json:"pm_friends_only"`
	ProfileColour string       `json:"profile_colour"`
	Username      string       `json:"username"`
}
