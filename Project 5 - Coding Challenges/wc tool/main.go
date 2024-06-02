package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func initFlag() {
	c := flag.Bool("c", false, "print the bytes count")
	l := flag.Bool("l", false, "print the lines count")
	w := flag.Bool("w", false, "print the words count")
	m := flag.Bool("m", false, "print the character count")
	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "" {
				break // Exit loop if an empty line is entered
			}
			files = append(files, scanner.Text())
		}
	}
	if *c {
		result := cflag(files)
		if len(files) != 1 {
			fmt.Println(result, "total")
		}
	}

	if *l {
		result := lflag(files)
		if len(files) != 1 {
			fmt.Println(result, "total")
		}
	}
	if *w {
		result := wflag(files)
		if len(files) != 1 {
			fmt.Println(result, "total")
		}
	}
	if *m {
		result := mflag(files)
		if len(files) != 1 {
			fmt.Println(result, "total")
		}
	}
	if !*m && !*w && !*l && !*c {
		lv := lflag(files)
		wv := wflag(files)
		cv := cflag(files)
		fmt.Println(lv, wv, cv, files)
	}
}

func main() {
	initFlag()
}

func cflag(files []string) int64 {
	var result int64
	for _, arg := range files {
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Cannot find fileInfo for: ", fileInfo.Name(), err)
			continue
		}
		fmt.Println(fileInfo.Size(), fileInfo.Name())
		result += fileInfo.Size()
	}
	return result
}

func lflag(files []string) int {
	var result int
	for _, arg := range files {
		l := 0
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			errStatement(err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			l++
		}
		fmt.Println(l-1, arg)
		result += l - 1
	}
	return result
}

func wflag(files []string) int {
	var result int
	for _, arg := range files {
		w := 0
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			errStatement(err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			w++
		}
		fmt.Println(w, arg)
		result += w
	}
	return result
}

func mflag(files []string) int {
	var result int
	for _, arg := range files {
		m := 0
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			errStatement(err)
			continue
		}
		fileInfo, err := file.Stat()
		characters := make([]byte, fileInfo.Size())
		_, err = file.Read(characters)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		m = utf8.RuneCountInString(string(characters))
		fmt.Println(m, arg)
		result += m
	}
	return result
}

func errStatement(err error) {
	fmt.Println(err)
}
