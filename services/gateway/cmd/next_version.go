package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/what-da-flac/wtf/services/gateway/internal/pipes"
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
		if len(args) != 1 {
			return fmt.Errorf("usage: next-version <type>")
		}
		versionType := args[0]
		reader, err := pipes.ReadStdin()
		if err != nil {
			if !errors.Is(err, pipes.ErrNoPipe) {
				return err
			}
			return fmt.Errorf("no pipe in stdin, cannot execute")
		}
		writer := os.Stdout
		return nextVersion(reader, writer, versionType)
	},
}

func init() {
	rootCmd.AddCommand(nextVersionCmd)
}

func nextVersion(r io.Reader, w io.Writer, versionType string) error {
	tags, err := tagMatches(r, versionType)
	if err != nil {
		return err
	}
	lastTag, err := calcNextTag(versionType, tags)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, lastTag)
	return err
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
	tags := IString(strings.Split(string(data), "\n")).Filter(func(s string) bool {
		if strings.TrimSpace(s) == "" {
			return false
		}
		return strings.HasPrefix(s, versionType)
	})
	sort.Slice(tags, func(i, j int) bool {
		prev := Tag(tags[i]).Number()
		curr := Tag(tags[j]).Number()
		return prev < curr
	})
	return tags.Reverse(), nil
}

type Tag string

func (x Tag) Number() int {
	values := strings.Split(string(x), sep)
	if len(values) == 0 {
		return 0
	}
	value, err := strconv.Atoi(values[len(values)-1])
	if err != nil {
		return 0
	}
	return value
}

type IString []string

func (x IString) Filter(filter func(s string) bool) IString {
	var res IString
	for _, v := range x {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (x IString) Reverse() IString {
	var res IString
	for i := len(x) - 1; i >= 0; i-- {
		res = append(res, x[i])
	}
	return res
}
