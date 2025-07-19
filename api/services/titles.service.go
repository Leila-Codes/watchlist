package services

import (
	"fmt"

	"lunasphere.co.uk/watchlist/api/types"
)

func ListTitles(requestingUserId uint64) ([]types.Title, error) {
	rows, err := db.Query(`SELECT title_id, title, type, description, running_time, 
	seasons, episodes, is_verified, submitted_by
	FROM titles
	WHERE is_verified = true OR submitted_by = $1
	LIMIT 10`, requestingUserId)

	if err != nil {
		return nil, err
	}

	titles := make([]types.Title, 0)

	for i := 0; rows.Next(); i++ {
		var (
			id          uint64
			titleName   string
			titleType   types.TitleType
			description *string
			runningTime *int
			seasons     *int
			episodes    *int
			verified    bool
			submittedBy *uint64
		)

		err := rows.Scan(&id, &titleName, &titleType, &description, &runningTime, &seasons, &episodes, &verified, &submittedBy)
		if err != nil {
			return nil, err
		}

		titles = append(titles, types.Title{
			ID:          id,
			Title:       titleName,
			Type:        types.TitleType(titleType),
			Description: description,
			RunningTime: runningTime,
			Seasons:     seasons,
			Episodes:    episodes,
			IsVerified:  verified,
			SubmittedBy: submittedBy,
		})
	}

	if len(titles) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return titles, err
}

func SubmitTitle(title types.Title, submittingUserId uint64) error {
	_, err := db.Exec(`INSERT INTO titles ("title", "type", "description", "running_time", "seasons", "episodes", "submitted_by")
	VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		title.Title,
		title.Type,
		title.Description,
		title.RunningTime,
		title.Seasons,
		title.Episodes,
		submittingUserId,
	)

	return err
}

func EditTitle(title types.Title, submittingUserId uint64) error {
	res, err := db.Exec(
		`UPDATE titles SET title=$1, description=$2, running_time=$3, seasons=$4, episodes=$5 WHERE title_id=$6 AND submitted_by=$7`,
		title.Title,
		title.Description,
		title.RunningTime,
		title.Seasons,
		title.Episodes,
		title.ID,
		submittingUserId)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("not found")
	}

	return nil
}
