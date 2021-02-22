package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"path/filepath"
	"strconv"
)

var (
	helpFlag = flag.Bool("h", false, "Show this help")
	inFile = flag.String("i", "alarmList.txt", "File to read alarm entries from")
	stdinFlag = flag.Bool("si", false, "Read alarm entries from stdin")
	stdoutFlag = flag.Bool("so", false, "Write alarm filters to stdout")
	outFile = flag.String("o", "filterList.txt", "File to write alarm filters to")
	appendFlag = flag.Bool("a", false, "Overwrite entries if file exists")
	err error
)

const usage = "`filterAlarms` [options]"

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}

	var filterList []string
	if *stdinFlag {
		reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
		for {
			line := readFromStdin(reader)
			if line != "" {
				filter := parseLine(line)
				if filter != "" {
					filterList = append(filterList, filter)
				}
			} else {
				if len(filterList) < 1 {
					fmt.Println("Did not find any alarms in input"); return
				} else {
					fmt.Printf("Read %d alarms from stdin\n", len(filterList))
					break
				}
			}
		}
	} else {
		fp, err := filepath.Abs(*inFile)
		if err != nil {
			fmt.Println(err); return
		}
		f, err := os.Open(fp)
		if err != nil {
			fmt.Println(err); return
		}
		defer f.Close()
		filterList = readFromFile(f)
		if len(filterList) < 1 {
			fmt.Printf("Did not find any alarms in %s\n", f.Name()); return
		} else {
			fmt.Printf("Generated %d alarm filters from %s\n", len(filterList), f.Name())
		}
	}
	if *stdoutFlag {
		fmt.Println("\ndiagnostics")
		for _, af := range filterList {
			fmt.Printf("%s\n", af)
		}
		fmt.Println("exit")
	} else {
		var of *os.File
		ofp, err := filepath.Abs(*outFile)
		if err != nil {
			fmt.Println(err); return
		}
		if *appendFlag {
			_, err = os.Stat(ofp)
			if err == nil {
				err = os.Remove(ofp)
				if err != nil {
					fmt.Println("Cannot remove existing file at %s, error: %v\nDumping contents to stdout\n\n", ofp, err)
					fmt.Println("\ndiagnostics")
					for _, af := range filterList {
						fmt.Printf("%s\n", af)
					}
					fmt.Println("exit")
					return
				}
			}
		}
		of, err = os.OpenFile(ofp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err); return
		}
		defer of.Close()
		
		fmt.Fprintln(of, "diagnostics")
		for _, af := range filterList {
			fmt.Fprintf(of, "%s\n", af)
		}
		fmt.Fprintln(of, "exit")
		fmt.Printf("Wrote %d alarm filters to %s\n", len(filterList), ofp)
	}
	
	
}

func parseLine(line string) string {
	//fmt.Println(line)
	str := strings.Fields(line)
	if len(str) < 4 {
		return ""
	}
	// shortcut method, anticipating copy-paste from "show alarm"
	// first entry will be the alarm code, at least 3 digits long
	// not to be confused with summary of total alarms number
	if len(str[0]) < 4 {
		return ""
	}
	_, err := strconv.Atoi(str[0])
	if err != nil {
		return ""
	}
	// last entry will be the affected object
	return fmt.Sprintf("set alarm-filter code %s object %s", str[0], str[len(str)-1])
}

func readFromFile(f *os.File) []string {
	s := bufio.NewScanner(f)
	var st []string
	for s.Scan() {
		filter := parseLine(s.Text())
		if filter != "" {
			st = append(st, filter)
		}
	}
	return st
}

func readFromStdin(r *bufio.Reader) string {
	a, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	} else if err != nil {
		panic(err)
	}
	
	return strings.TrimRight(string(a), "\r\n")
}
