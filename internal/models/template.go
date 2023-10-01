package models

type Data struct {
	Project  *Project
	Projects []Project
	Admin    bool
	Error    string
}
