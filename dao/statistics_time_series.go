package dao

// function to get all totals info
func GetDealsAttemptedInRange(from string, to string) (interface{}, error) {
	var dealsAttempatedInRange int64
	row := DB.Raw("select sum(cnt) as total_rows from (select count(*) as cnt from content_logs where created_at >= ? and created_at <= ? group by content_id) as t", from, to).Row()
	err := row.Scan(&dealsAttempatedInRange)
	if err != nil {
		return nil, err
	}
	return dealsAttempatedInRange, nil
}
