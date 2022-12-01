package shared

import "io/ioutil"

func ReadFileContents(path string) string {
	if contents, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(contents)
	}

}
