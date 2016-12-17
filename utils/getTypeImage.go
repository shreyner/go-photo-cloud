package utils

func GetTypeImage(str string) string {

	switch str {

	case "image/bmp":
		return "bmp"
	case "image/gif":
		return "gif"
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	default:
		return ""
	}
}