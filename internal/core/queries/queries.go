package queries

const GetAllQuery = `
	SELECT *
	FROM bq_video_games_dataset.bq_video_games_table
	WHERE year_of_release IS NOT NULL
	ORDER BY year_of_release
	LIMIT 100`

const GetAdventureQuery = `
	SELECT *
	FROM bq_video_games_dataset.bq_video_games_table
	WHERE genre = "Adventure"
		AND year_of_release IS NOT NULL
	ORDER BY year_of_release
	LIMIT 100`
