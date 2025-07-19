package types

type TitleType string

const (
	MovieType    TitleType = "movie"
	TVSeriesType TitleType = "tv"
)

type Title struct {
	ID          uint64    `json:"title_id"`
	Title       string    `json:"title"`
	Type        TitleType `json:"type"`
	Description *string   `json:"description"`
	IsVerified  bool      `json:"is_verified"`
	SubmittedBy *uint64   `json:"-"`
	RunningTime *int      `json:"running_time"`
	Seasons     *int      `json:"seasons"`
	Episodes    *int      `json:"episodes"`
}
