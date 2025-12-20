package interfaces

import "time"

type BeatmapsResp struct {
	Beatmaps []Beatmap `json:"beatmaps"`
}

type Beatmap struct {
	BeatmapsetID         int        `json:"beatmapset_id"`
	DifficultyRating     float64    `json:"difficulty_rating"`
	ID                   int        `json:"id"`
	Mode                 string     `json:"mode"`
	Status               string     `json:"status"`
	TotalLength          int        `json:"total_length"`
	UserID               int        `json:"user_id"`
	Version              string     `json:"version"`
	Accuracy             int        `json:"accuracy"`
	Ar                   int        `json:"ar"`
	Bpm                  float64    `json:"bpm"`
	Convert              bool       `json:"convert"`
	CountCircles         int        `json:"count_circles"`
	CountSliders         int        `json:"count_sliders"`
	CountSpinners        int        `json:"count_spinners"`
	Cs                   int        `json:"cs"`
	DeletedAt            time.Time  `json:"deleted_at"`
	Drain                int        `json:"drain"`
	HitLength            int        `json:"hit_length"`
	IsScoreable          bool       `json:"is_scoreable"`
	LastUpdated          time.Time  `json:"last_updated"`
	ModeInt              int        `json:"mode_int"`
	Passcount            int        `json:"passcount"`
	Playcount            int        `json:"playcount"`
	Ranked               int        `json:"ranked"`
	URL                  string     `json:"url"`
	Checksum             string     `json:"checksum"`
	Beatmapset           Beatmapset `json:"beatmapset"`
	CurrentUserPlaycount int        `json:"current_user_playcount"`
	Failtimes            Failtimes  `json:"failtimes"`
	MaxCombo             int        `json:"max_combo"`
	Owners               []Owners   `json:"owners"`
}

type Beatmapset struct {
	AnimeCover         bool               `json:"anime_cover"`
	Artist             string             `json:"artist"`
	ArtistUnicode      string             `json:"artist_unicode"`
	Covers             Covers             `json:"covers"`
	Creator            string             `json:"creator"`
	FavouriteCount     int                `json:"favourite_count"`
	GenreID            int                `json:"genre_id"`
	Hype               Hype               `json:"hype"`
	ID                 int                `json:"id"`
	LanguageID         int                `json:"language_id"`
	Nsfw               bool               `json:"nsfw"`
	Offset             int                `json:"offset"`
	PlayCount          int                `json:"play_count"`
	PreviewURL         string             `json:"preview_url"`
	Source             string             `json:"source"`
	Spotlight          bool               `json:"spotlight"`
	Status             string             `json:"status"`
	Title              string             `json:"title"`
	TitleUnicode       string             `json:"title_unicode"`
	TrackID            int                `json:"track_id"`
	UserID             int                `json:"user_id"`
	Video              bool               `json:"video"`
	Bpm                float64            `json:"bpm"`
	CanBeHyped         bool               `json:"can_be_hyped"`
	DeletedAt          time.Time          `json:"deleted_at"`
	DiscussionEnabled  bool               `json:"discussion_enabled"`
	DiscussionLocked   bool               `json:"discussion_locked"`
	IsScoreable        bool               `json:"is_scoreable"`
	LastUpdated        time.Time          `json:"last_updated"`
	LegacyThreadURL    string             `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary `json:"nominations_summary"`
	Ranked             int                `json:"ranked"`
	RankedDate         time.Time          `json:"ranked_date"`
	Rating             float64            `json:"rating"`
	Storyboard         bool               `json:"storyboard"`
	SubmittedDate      time.Time          `json:"submitted_date"`
	Tags               string             `json:"tags"`
	Availability       Availability       `json:"availability"`
	Ratings            []int              `json:"ratings"`
}
