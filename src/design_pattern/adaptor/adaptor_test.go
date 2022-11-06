package adaptor

import "testing"

func TestAdaptor(t *testing.T) {
	audioPlayer := AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
