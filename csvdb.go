package csvdb

import (
	"encoding/csv"
	"io"
	"os"
)

var DB = make(map[string][]string)
var fname string

func Init(fn string) {
	//	var tmps []string
	//	DB = nil
	fname = fn
	f, err := os.Open(fn)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(fn)
			if err != nil {
				panic("Ошибка создания файла данных !\n")
			}
		} else {
			panic("Ошибка при открытии файла данных !\n")
		}
	}
	csvr := csv.NewReader(f)
	csvr.Comma = '|'
	for {
		tmps, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			//			fmt.Println(tmps[0], tmps[1:])
			DB[tmps[0]] = tmps[1:]
		}
	}

	f.Close()
}

func save() {
	f, err := os.Create(fname)
	if err != nil {
		panic("Ошибка создания файла данных !\n")
	}
	csvw := csv.NewWriter(f)
	csvw.Comma = '|'
	for k, v := range DB {
		tmps := append([]string{k}, v...)
		csvw.Write(tmps)
	}
	csvw.Flush()
	f.Close()

}

func Add(key string, value []string) {
	DB[key] = value
	save()
}

func Del(key string) {
	if key == "" {
		for k := range DB {
			delete(DB, k)
		}
	} else {
		delete(DB, key)
	}
	save()
}
