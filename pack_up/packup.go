package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Handles my files while writing applications.
	// Moves the temporary files from 'docs' to a the archive 'apps'
	// TODO: Seperate into functions?

	// Parse input
	repoPath, _ := filepath.Abs(os.Args[1]) // TODO: Exists?
	companyName := os.Args[2]

	// Get directory paths
	appsPath := filepath.Join(repoPath, "apps") // TODO: Exists?
	docsPath := filepath.Join(repoPath, "docs") // TODO: Exists?

	// Create new directory
	yearNum, weekNum := time.Now().ISOWeek()
	week := fmt.Sprintf("%02d", weekNum)
	year := strconv.Itoa(yearNum)
	fmt.Println(week)
	name := strings.Join([]string{year, week, companyName}, "-")
	newDir := filepath.Join(appsPath, name)
	fmt.Println(newDir)
	os.Mkdir(newDir, 0777) // TODO: Handle ERROR

	// Copy cv and coverletter
	app := filepath.Join(docsPath, "coverletter/output/coverletter.pdf")
	cv := filepath.Join(docsPath, "cv/output/cv.pdf")
	os.Rename(app, filepath.Join(newDir, "ansøgning.pdf")) // TODO: Handle ERROR
	os.Rename(cv, filepath.Join(newDir, "cv.pdf"))         // TODO: Handle ERROR
}
