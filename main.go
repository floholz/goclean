package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var cleanupPaths = os.Getenv("GO_CLEAN_PATHS")
var schedule = os.Getenv("GO_CLEAN_SCHEDULE")
var maxAge = os.Getenv("GO_CLEAN_MAX_AGE")
var maxAgeD time.Duration

func main() {
	cleanupPaths = strings.Trim(cleanupPaths, "\"")
	schedule = strings.Trim(schedule, "\"")
	maxAge = strings.Trim(maxAge, "\"")

	if cleanupPaths == "" {
		log.Fatal("❌  No paths to clean up are set. Set the Paths in the 'GO_CLEAN_PATHS' environment variable. Multiple paths should be separated by ';'.")
	} else {
		fmt.Printf("📁  Paths to celan up: '%s'\n", cleanupPaths)
	}

	if schedule == "" {
		schedule = "0 0 * * *" // daily at 00:00
		fmt.Printf("🕗  No schedule set. Defaulting to '%s'\n", schedule)
	} else {
		fmt.Printf("🕗  Schedule set: '%s'\n", schedule)
	}

	if maxAge == "" {
		maxAge = "7d" // 7 days
		fmt.Printf("⏳  No max file age set. Defaulting to '%s'\n", maxAge)
	} else {
		fmt.Printf("⏳  Max file age set to: '%s'\n", maxAge)
	}
	maxAgeD = parseMaxAge()

	c := cron.New(
		cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)),
	)

	_, _ = c.AddFunc(schedule, func() { fmt.Printf("[%s] 🗑️🧹  Cleaning up ...\n", time.Now().Format(time.RFC3339)) })
	_, err := c.AddFunc(schedule, cleanup)
	if err != nil {
		log.Fatal("❌  Error while cleaning up!\n", err)
	}

	c.Run()
}

func cleanup() {
	threshold := time.Now().Add(-maxAgeD)
	for _, path := range strings.Split(cleanupPaths, ";") {

		entries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("❌  Error cleaning up '%s'", path)
		}
		for _, entry := range entries {
			rmR(filepath.Clean(path), entry, threshold)
		}
	}
}

func rmR(path string, entry os.DirEntry, ageThr time.Time) bool {
	path = filepath.Join(path, entry.Name())
	if entry.IsDir() {
		subEntries, _ := os.ReadDir(path)
		empty := true
		for _, sub := range subEntries {
			empty = empty && rmR(path, sub, ageThr)
		}
		if empty {
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("⚠️  Failed to delete '%s'", path)
				return false
			} else {
				fmt.Printf("📁  Deleted dir  '%s'", path)
			}
		}
		return empty
	}
	fi, _ := entry.Info()
	age := fi.ModTime().Sub(ageThr)
	if age <= 0 {
		err := os.Remove(path)
		if err != nil {
			fmt.Printf("⚠️  Failed to delete '%s'", path)
		} else {
			fmt.Printf("📄  Deleted file '%s'", path)
			return true
		}
	}
	return false
}

func parseMaxAge() time.Duration {
	r, err := regexp.Compile("^(?:(\\d+)Y)*(?:(\\d+)M)*(?:(\\d+)d)*(?:(\\d+)h)*(?:(\\d+)m)*(?:(\\d+)s)*$")
	if err != nil {
		log.Fatal("❌  Error parsing set max age in 'GO_CLEAN_MAX_AGE'.\n", err)
	}
	matches := r.FindStringSubmatch(maxAge)

	year, err := parseStrToInt(matches[1])
	month, err := parseStrToInt(matches[2])
	day, err := parseStrToInt(matches[3])
	hour, err := parseStrToInt(matches[4])
	minute, err := parseStrToInt(matches[5])
	second, err := parseStrToInt(matches[6])

	if err != nil {
		log.Fatal("❌  Error parsing set max age in 'GO_CLEAN_MAX_AGE'.\n", err)
	}

	return time.Duration(year)*time.Hour*8760 +
		time.Duration(month)*time.Hour*731 + //730.485h = avg Month
		time.Duration(day)*time.Hour*24 +
		time.Duration(hour)*time.Hour +
		time.Duration(minute)*time.Minute +
		time.Duration(second)*time.Second
}

func parseStrToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	num, err := strconv.Atoi(str)
	return num, err
}
