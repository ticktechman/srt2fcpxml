package Resources

import "os"

type Effect struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Uid  string `xml:"uid,attr"`
	Src  string `xml:"src,attr"`
}

func NewEffect() *Effect {
	return &Effect{
		Text: "",
		ID:   "r2",
		// Name: "Basic Title",
		// Uid:  ".../Titles.localized/Bumper:Opener.localized/Basic Title.localized/Basic Title.moti",
		Name: "new-titles",
		Uid:  "~/Titles.localized/background-caption/background-caption.moti",
		// Src:  "file:///Users/ticktech/Movies/Motion%20Templates.localized/Titles.localized/background-caption/background-caption.moti",
		Src: "file://" + os.Getenv("HOME") + "/Movies/Motion%20Templates.localized/Titles.localized/background-caption/background-caption.moti",
	}
}
