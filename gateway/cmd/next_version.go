package cmd

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/what-da-flac/wtf/go-common/istring"
	"github.com/what-da-flac/wtf/go-common/pipes"
)

const (
	defaultTagValue = "0.0.0"
	sep             = "."
)

var nextVersionCmd = &cobra.Command{
	Use:   "next-version",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		// example
		// git tag | go run . next-version cdk
		return nextVersion(args)
	},
}

func init() {
	rootCmd.AddCommand(nextVersionCmd)
}

func nextVersion(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: next-version <type>")
	}
	versionType := args[0]
	r, err := pipes.ReadStdin()
	if err != nil {
		if !errors.Is(err, pipes.ErrNoPipe) {
			return err
		}
		return fmt.Errorf("no pipe in stdin, cannot execute")
	}
	tags, err := tagMatches(r, versionType)
	if err != nil {
		return err
	}
	lastTag, err := calcNextTag(versionType, tags)
	if err != nil {
		return err
	}
	fmt.Println(lastTag)
	return nil
}

func calcNextTag(versionType string, tags []string) (string, error) {
	if len(tags) == 0 {
		tags = []string{
			versionType + sep + defaultTagValue,
		}
	}
	lastTag := tags[0]
	values := strings.Split(lastTag, sep)
	last, err := strconv.Atoi(values[len(values)-1])
	if err != nil {
		return "", err
	}
	last++
	values[len(values)-1] = strconv.Itoa(last)
	lastTag = strings.Join(values, sep)
	return lastTag, nil
}

func tagMatches(r io.Reader, versionType string) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	tags := istring.IString(strings.Split(string(data), "\n")).Filter(func(s string) bool {
		if strings.TrimSpace(s) == "" {
			return false
		}
		return strings.HasPrefix(s, versionType)
	})
	sort.Strings(tags)
	return tags.Reverse(), nil
}
