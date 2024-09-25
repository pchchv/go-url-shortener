package domain

const (
	PlatformUnknown Platform = iota
	PlatformInstagram
	PlatformTwitter
	PlatformYouTube
)

type Platform int

func (p Platform) String() string {
	switch p {
	case PlatformInstagram:
		return "Instagram"
	case PlatformTwitter:
		return "Twitter"
	case PlatformYouTube:
		return "YouTube"
	default:
		return "Unknown"
	}
}
