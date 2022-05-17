package filehelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/compico/osutools/encoding/database"
	"github.com/compico/osutools/osu"
)

var ErrUnknownPath = errors.New("Unknown game path.")

func (osufolder *OsuFolder) initSongsPath() {
	osufolder.SongsPath = filepath.Join(osufolder.GamePath, "Songs")
}

func (osufolder *OsuFolder) initSkinsPath() {
	osufolder.SkinsPath = filepath.Join(osufolder.GamePath, "Skins")
}

func (osufolder *OsuFolder) ReadOsudbFile() error {
	osufolder.DataBase = new(osu.OsuDB)
	osufolder.DataBase.Beatmaps = make([]osu.Beatmap, 0)
	err := database.Unmarshal(osufolder.GamePath+"/osu!.db", osufolder.DataBase)
	if err != nil {
		return fmt.Errorf("cannot decode osu database file: %w", err)
	}
	osufolder.hashing()
	return nil
}

func (osufolder *OsuFolder) JsonToDatabase(file string) error {
	osufolder.DataBase = new(osu.OsuDB)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("cannot read JSON file: %w", err)
	}

	if err = json.Unmarshal(b, osufolder.DataBase); err != nil {
		return fmt.Errorf("cannot decode JSON input to osu database: %w", err)
	}
	return err
}

func (osufolder *OsuFolder) hashing() {
	osufolder.SongsHash = make(map[int32]map[int32]int)
	for i := 0; i < len(osufolder.DataBase.Beatmaps); i++ {
		if len(osufolder.SongsHash[osufolder.DataBase.Beatmaps[i].BeatmapID]) == 0 {
			osufolder.SongsHash[osufolder.DataBase.Beatmaps[i].BeatmapID] = make(map[int32]int)
		}
		osufolder.SongsHash[osufolder.DataBase.Beatmaps[i].BeatmapID][osufolder.DataBase.Beatmaps[i].DifficultyID] = i
	}
	osufolder.DirectorySorting()
}
