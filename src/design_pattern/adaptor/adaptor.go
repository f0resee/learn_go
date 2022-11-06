package adaptor

import (
	"fmt"
	"strings"
)

type MediaPlayer interface {
	Play(string, string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(string)
	PlayMp4(string)
}

type VlcPlayer struct {
}

func (v VlcPlayer) PlayVlc(fileName string) {
	fmt.Println("Playing vlc file.Name: ", fileName)
}

func (v VlcPlayer) PlayMp4(fileName string) {

}

type Mp4Player struct {
}

func (m Mp4Player) PlayVlc(fileName string) {

}

func (m Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing vlc file.Name: ", fileName)
}

func NewMediaAdaptor(audioType string) *MediaAdaptor {
	adaptor := MediaAdaptor{}
	if strings.EqualFold(audioType, "vlc") {
		adaptor.advancedMediaPlayer = VlcPlayer{}
	} else if strings.EqualFold(audioType, "mp4") {
		adaptor.advancedMediaPlayer = Mp4Player{}
	} else {
		return nil
	}
	return &adaptor
}

type MediaAdaptor struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m MediaAdaptor) Play(audioType string, fileName string) {
	if strings.EqualFold(audioType, "vlc") {
		m.advancedMediaPlayer.PlayVlc(fileName)
	} else if strings.EqualFold(audioType, "mp4") {
		m.advancedMediaPlayer.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdaptor *MediaAdaptor
}

func (a AudioPlayer) Play(audioType string, fileName string) {
	if strings.EqualFold(audioType, "mp3") {
		fmt.Println("Playing mp3 file.Name: ", fileName)
	} else if strings.EqualFold(audioType, "vlc") || strings.EqualFold(audioType, "mp4") {
		a.mediaAdaptor = NewMediaAdaptor(audioType)
		a.mediaAdaptor.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media.", audioType, " format not supported")
	}
}
