package utils

import (
	"strings"
)

var supportedExtensions = [...]string{
	// Images
	".jpg",
	".png",
	".webp",
	".heic",

	// Raw images
	".crw",
	".cr2",
	".cr3",
	".nef",
	".nrw",
	".orf",
	".raw",
	".rw2",
	".arw",
	".srf",
	".sr2",
	".raf",
	".dng",

	// Videos
	".mpg",
	".mod",
	".mmv",
	".tod",
	".wmv",
	".asf",
	".avi",
	".divx",
	".mov",
	".m4v",
	".3gp",
	".3g2",
	".mp4",
	".m2t",
	".m2ts",
	".mts",
	".mkv",
}

func isExtensionSupported(toCheck string) bool {
	for _, extension := range supportedExtensions {
		if strings.ToLower(toCheck) == strings.ToLower(extension) {
			return true
		}
	}
	return false
}
