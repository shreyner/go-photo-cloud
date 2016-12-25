package utils

func GetTypeImage(str string) (string, bool) {

	switch str {
	case "image/bmp":
		return "bmp", false
	case "image/gif":
		return "gif", false
	case "image/jpeg":
		return "jpg", false
	case "image/png":
		return "png", false
	default:
		return "", true
	}
}