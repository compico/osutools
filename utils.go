package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	ErrNotOsuFile = errors.New("input file extention is not '.osu'")
	ErrNoFiles    = errors.New("no files")
)

func test(s []string) {
	for _, l := range s {
		fmt.Println(l)
	}
}

func osufiletest(l []string) error {
	for i := 0; i < len(l); i++ {
		if filepath.Ext(l[i]) == ".osu" {
			continue
		}

		return fmt.Errorf("%w: %s", ErrNotOsuFile, l[i])
	}

	return nil
}

func notexisterror(f []string, format string) error {
	if len(f) == 0 {
		return ErrNoFiles
	}

	fmt.Println("[ERROR] In folder(s):")

	for _, l := range f {
		fmt.Println("►", l)
	}

	fmt.Printf("...no have %s file. But you can ignore it.\n", format)

	return nil
}

func parsefile(f string) (r []string) {
	r = strings.Split(f, "\n")
	return r
}

//This function take path directory and return directory contents
func lsDir(songsPath string) ([]string, error) {
	songFolders, err := ioutil.ReadDir(songsPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s directory: %w", songsPath, err)
	}
	var strList []string
	for i := 0; i < len(songFolders); i++ {
		strList = append(strList, filepath.Join(songsPath, songFolders[i].Name()))
	}
	return strList, nil
}
