package admin

var tagTypes = [6]string{"tag_PL", "tag_framework", "tag_package", "tag_technology", "tag_tool", "tag_other"}

func IsValidTagType(tagtype string) bool {
	for i := range tagTypes {
		if tagtype == tagTypes[i] {
			return true
		}
	}
	return false
}
