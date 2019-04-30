package models

func AllComments() ([]string, error) {
	return db.LRange("comments", 0, 10).Result()
}

func Create(comment string) error {
	return db.LPush("comments", comment).Err()
}
