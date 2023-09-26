package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTagValidation(t *testing.T) {
	tag := Tag{
		TagName: "golang",
		TagType: "tag_PL",
	}
	err := tag.Validate()
	require.NoError(t, err)

	tag.TagType = "tag_unknown"
	err = tag.Validate()
	require.Error(t, err)
}
