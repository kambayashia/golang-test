package twitter_ads

type Response struct{
	DataType string `json:"data_type"`
	TotalCount int `json:"total_count"`
	NextCursor string `json:"next_cursor"`
}
