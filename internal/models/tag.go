package models

import "github.com/go-playground/validator/v10"

type Tag struct {
	TagName string `bson:"tag_name" json:"tag_name"`
	TagType string `bson:"tag_type" json:"tag_type" validate:"required,tag_oneof"`
}

func (t *Tag) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("tag_oneof", validateTagType)
	return validate.Struct(t)
}

var tagTypes = [6]string{"tag_PL", "tag_framework", "tag_package", "tag_technology", "tag_tool", "tag_other"}

func validateTagType(fl validator.FieldLevel) bool {
	tagType := fl.Field().String()
	for i := range tagTypes {
		if tagType == tagTypes[i] {
			return true
		}
	}
	return false
}
