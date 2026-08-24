// auto-license. by u/Lazy-Physics3198
// this is a tool used by chat_now v2 (a passion project, private) in order to properly acknowledge that it
// is under the AGPL3.

// however, i value this part as nothing.
// therefore, it is under the Apache license.

package main

import (
	"os"
	"time"
	"strings"
	"slices"
	"bytes"
	"io"
	"io/fs"
	"encoding/json"
	"log"
	"fmt"
	path "path/filepath"
)

// on running, it will read "license.txt" (assuming its not pre-commented.)
// along with "auto_license-config.json"/

// do not expect this code to be clean. i am writing it as a black box.
type Config map[string]string
var License string
var Conf Config
var Memoize_Formats map[string][]byte = make(map[string][]byte)
var Counter int64

func ReadFileClean(file_name string) []byte {
	File, err := os.Open(file_name)
	defer File.Close()
	if err != nil {
		log.Fatal(err)
	}
	Data, err2 := io.ReadAll(File)
	if err2 != nil {
		log.Fatal(err2)
	}
	return Data
}

func WriteFileClean(file_name string, writeFrom []byte) {
	var Buffer *bytes.Buffer = bytes.NewBuffer(writeFrom)
	File, err := os.OpenFile(file_name, os.O_RDWR | os.O_TRUNC, 0666)
	defer File.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(File, Buffer)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartTime := time.Now()
	log.Println("Reading Files...")
	License = string(ReadFileClean("license.txt"))
	if err := json.Unmarshal(ReadFileClean("auto_license-config.json"), &Conf); err != nil {
		log.Fatal(err)
	}
	log.Println("Walking...")
	path.WalkDir(".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error while walking: %w\n", err)
			return nil
		}
		Extension := path.Ext(d.Name())
		if CommentFormat, Exists := Conf[Extension]; Exists {
			log.Printf("Now Writing %s...\n", p)
			Counter++
			OrigData := ReadFileClean(p)
			if Prelude, Exists2 := Memoize_Formats[Extension]; Exists2 {
				if bytes.HasPrefix(OrigData, Prelude) {
					log.Printf("%s is already licensed. Skipping...\n", p)
					Elapsed := time.Since(StartTime).Seconds()
					var Speed float64
					if Elapsed > 0 {
						Speed = float64(Counter) / Elapsed
					}
					log.Printf("Finished with file %s. %d files processed. (%.1f files/sec) \r", p, Counter, Speed)
					return nil
				}
				var WriteData []byte = slices.Concat(Prelude, OrigData)
				WriteFileClean(p, WriteData)
			} else {
				SplitData := strings.Split(License, "\n")
				var Commented_License_Aux []string = make([]string, len(SplitData))
				for i, v := range SplitData {
					Commented_License_Aux[i] = fmt.Sprintf("%s %s", CommentFormat, v)
				}
				var Commented_License []byte
				Commented_License = []byte(strings.Join(Commented_License_Aux, "\n"))
				Memoize_Formats[Extension] = Commented_License
				
				if bytes.HasPrefix(OrigData, Commented_License) {
					log.Printf("%s is already licensed. Skipping... \n", p)
					Elapsed := time.Since(StartTime).Seconds()
					var Speed float64
					if Elapsed > 0 {
						Speed = float64(Counter) / Elapsed
					}
					log.Printf("Finished with file %s. %d files processed. (%.1f files/sec) \r", p, Counter, Speed)
					return nil
				}
				var WriteData []byte = slices.Concat(Commented_License, OrigData)
				WriteFileClean(p, WriteData)
			}
			
		
		Elapsed := time.Since(StartTime).Seconds()
		var Speed float64
		if Elapsed > 0 {
			Speed = float64(Counter) / Elapsed
		}
		log.Printf("Finished with file %s. %d files processed. (%.1f files/sec) \n", p, Counter, Speed)
		
		} else {
			log.Printf("Extension %s not defined by config.\n", Extension)
			return nil
		}
		return nil
	})
	Elapsed := time.Since(StartTime).Seconds()
	var Speed float64
	if Elapsed > 0 {
		Speed = float64(Counter) / Elapsed
	}
	log.Printf("Finished in %.1f seconds, with a speed of %.1f files/sec. \r", Elapsed, Speed)
}
