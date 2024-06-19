package problems

import (
	"math/rand"
	"time"
)

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var songs = []string{"A", "B", "C", "D", "E"}
// var k = 2

// type MusicPlayer struct {
// 	songDict map[string]int
// 	songArr  []string
// 	timer    map[int][]string
// 	t        int
// }

// func NewMusicPlayer() *MusicPlayer {
// 	mp := &MusicPlayer{
// 		songDict: make(map[string]int),
// 		songArr:  make([]string, len(songs)),
// 		timer:    map[int][]string{},
// 	}
// 	copy(mp.songArr, songs)
// 	for i, song := range mp.songArr {
// 		mp.songDict[song] = i
// 	}
// 	return mp
// }

// func (mp *MusicPlayer) GetRandomSong() string {
// 	fmt.Println(mp.songArr)
// 	mp.t++
// 	rand.Seed(time.Now().UnixNano())
// 	songIndex := rand.Intn(len(mp.songArr) - 1)
// 	song := mp.songArr[songIndex]
// 	mp.deleteSong(song)
// 	mp.timer[mp.t+k] = append(mp.timer[mp.t+k], song)
// 	return song
// }

// func (mp *MusicPlayer) deleteSong(song string) {
// 	i := mp.songDict[song]
// 	j := len(mp.songArr) - 1

// 	mp.songArr[i], mp.songArr[j] = mp.songArr[j], mp.songArr[i]
// 	mp.songArr = mp.songArr[:len(mp.songArr)-1]

// 	delete(mp.songDict, song)

// 	mp.songDict[mp.songArr[i]] = i
// }

// func (mp *MusicPlayer) InsertAgainToPlay() {
// 	if len(mp.timer) == 0 {
// 		return
// 	}

//		if len(mp.timer[mp.t]) > 0 {
//			songs := mp.timer[mp.t]
//			mp.timer[mp.t] = []string{}
//			mp.songArr = append(mp.songArr, songs...)
//			for _, song := range songs {
//				mp.songDict[song] = len(mp.songArr) - 1
//			}
//		}
//	}
var songs = []struct {
	name     string
	duration int
}{
	{"A", 1},
	{"B", 3},
	{"C", 2},
	{"D", 4},
	{"E", 2},
}

type MusicPlayer struct {
	songDict map[string]int
	songArr  []string
	timer    []struct {
		song     string
		duration int
		time     int
	}
	t int
}

func NewMusicPlayer() *MusicPlayer {
	mp := &MusicPlayer{
		songDict: make(map[string]int),
		songArr:  make([]string, len(songs)),
	}
	for i, song := range songs {
		mp.songArr[i] = song.name
		mp.songDict[song.name] = i
	}
	return mp
}

func (mp *MusicPlayer) GetRandomSong() string {
	mp.t++
	rand.Seed(time.Now().UnixNano())
	var eligibleSongs []string
	for _, song := range mp.songArr {
		if mp.songDuration(song) <= 0 {
			eligibleSongs = append(eligibleSongs, song)
		}
	}
	if len(eligibleSongs) == 0 {
		return ""
	}
	var selectedSong string
	minDuration := int(^uint(0) >> 1)
	for _, song := range eligibleSongs {
		duration := mp.songDuration(song)
		if duration < minDuration {
			selectedSong = song
			minDuration = duration
		}
	}
	mp.deleteSong(selectedSong)
	mp.timer = append(mp.timer, struct {
		song     string
		duration int
		time     int
	}{selectedSong, minDuration, mp.t + minDuration})
	return selectedSong
}

func (mp *MusicPlayer) deleteSong(song string) {
	i := mp.songDict[song]
	j := len(mp.songArr) - 1

	mp.songArr[i], mp.songArr[j] = mp.songArr[j], mp.songArr[i]
	mp.songArr = mp.songArr[:len(mp.songArr)-1]

	delete(mp.songDict, song)

	mp.songDict[mp.songArr[i]] = i
}

func (mp *MusicPlayer) InsertAgainToPlay() {
	if len(mp.timer) == 0 {
		return
	}

	if mp.timer[0].time <= mp.t {
		song := mp.timer[0].song
		mp.timer = mp.timer[1:]
		mp.songArr = append(mp.songArr, song)
		mp.songDict[song] = len(mp.songArr) - 1
	}
}

func (mp *MusicPlayer) songDuration(song string) int {
	for _, s := range songs {
		if s.name == song {
			return s.duration
		}
	}
	return -1
}
