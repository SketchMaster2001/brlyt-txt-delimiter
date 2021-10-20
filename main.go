package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

/*
To whom ever is reading this
----------------------------
WiiLink translators aren't the brightest and won't use Wii Layout Editor to max out the char limit
on the string inside the txt1 header. This tool maxes out the char limit for every txt1 header it can find.
----------------------------
- Sketch
 */

const textLimit = 78

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: brlyttool <input>")
		os.Exit(1)
	}

	input := os.Args[1]

	inputData, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Now we find all occurrences of txt1.
	txt1 := findAllOccurrences(inputData)

	for _, offset := range txt1 {
		// Jump to the offset of the txt1 header, then add by 78 to get the beginning of the text limit offset.
		// First set to 0x27. Jump one more offset and set to 0x0F. Since the text limit is an u16,
		// the combined values will be 0x270F, or 9999.
		inputData[offset+textLimit] = 0x27
		inputData[offset+textLimit+1] = 0x0F
	}

	err = ioutil.WriteFile(input, inputData, 0666)
	if err != nil {
		return
	}
}


// findAllOccurrences finds the offsets of the txt1 header.
func findAllOccurrences(data []byte) []int {
	var results []int
	searchData := data
	term := []byte("txt1")
	for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d+x+1 {
		results = append(results, x+d)
		searchData = searchData[x+1:]
	}

	return results
}