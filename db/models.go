package db

type Tag struct {
	Name string `bson:"tag_name"`
	Type string
}

type Project struct {
	Name        string `bson:"name"`
	Description string `bson:"description,omitempty"`
	Url         string `bson:"url,omitempty"`
	IsFinished  bool   `bson:"finished,omitempty"`
}
