package filehelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/sys/windows/registry"

	"github.com/compico/osutools/encoding/database"
	"github.com/compico/osutools/osu"
)

var ErrUnknownPath = errors.New("unknown game path")

func (osufolder *OsuFolder) getAllPaths() error {
	if osufolder.GamePath == "" {
		return ErrUnknownPath
	}
	osufolder.initSongsPath()
	osufolder.initSkinsPath()
	return nil
}

func (osufolder *OsuFolder) SetGamePath(gamepath string) {
	osufolder.GamePath = gamepath
}

func (osufolder *OsuFolder) initSongsPath() {
	osufolder.SongsPath = filepath.Join(osufolder.GamePath, "Songs")
}

func (osufolder *OsuFolder) initSkinsPath() {
	osufolder.SkinsPath = filepath.Join(osufolder.GamePath, "Skins")
}

//For windows only
func (osufolder *OsuFolder) InitGamePathByReg() error {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `osu\DefaultIcon`, registry.QUERY_VALUE)
	if err != nil {
		return fmt.Errorf("cannot open registry key: %w", err)
	}
	defer k.Close()

	path, _, err := k.GetStringValue("")
	if err != nil {
		return fmt.Errorf("cannot read registry key value: %w", err)
	}
	path = path[1:]
	path = filepath.Dir(path)

	osufolder.GamePath = path

	err = osufolder.getAllPaths()
	if err != nil {
		return err
	}
	return nil
}

func (osufolder *OsuFolder) ReadOsudbFile() error {
	osufolder.DataBase = new(osu.OsuDB)
	osufolder.DataBase.Beatmaps = make([]osu.Beatmap, 0)
	err := database.Unmarshal(osufolder.GamePath+"/osu!.db", osufolder.DataBase)
	if err != nil {
		return fmt.Errorf("cannot decode osu database file: %w", err)
	}
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
