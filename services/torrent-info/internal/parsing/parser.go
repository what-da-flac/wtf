package parsing

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/what-da-flac/wtf/services/torrent-info/internal/domain"
)

func ParseTorrent(metadata string) (*domain.Torrent, error) {
	lines := strings.Split(metadata, "\n")
	var torrent domain.Torrent
	var files []domain.File
	var trackers []domain.Tracker

	// Regex patterns
	namePattern := regexp.MustCompile(`Name: (.+)`)
	pieceCountPattern := regexp.MustCompile(`Piece Count: (\d+)`)
	pieceSizePattern := regexp.MustCompile(`Piece Size: (.+)`)
	totalSizePattern := regexp.MustCompile(`Total Size: (.+)`)
	privacyPattern := regexp.MustCompile(`Privacy: (.+)`)
	trackerPattern := regexp.MustCompile(`udp://[^\s]+`)
	filePattern := regexp.MustCompile(`(.+)\((.+)\)`)

	// Loop through each line and match patterns
	for i, line := range lines {
		line = strings.TrimSpace(line)

		// Match general info
		if nameMatch := namePattern.FindStringSubmatch(line); nameMatch != nil {
			torrent.Name = nameMatch[1]
		}
		if pieceCountMatch := pieceCountPattern.FindStringSubmatch(line); pieceCountMatch != nil {
			pieceCount, _ := strconv.Atoi(pieceCountMatch[1])
			torrent.PieceCount = pieceCount
		}
		if pieceSizeMatch := pieceSizePattern.FindStringSubmatch(line); pieceSizeMatch != nil {
			torrent.PieceSize = pieceSizeMatch[1]
		}
		if totalSizeMatch := totalSizePattern.FindStringSubmatch(line); totalSizeMatch != nil {
			torrent.TotalSize = totalSizeMatch[1]
		}
		if privacyMatch := privacyPattern.FindStringSubmatch(line); privacyMatch != nil {
			torrent.Privacy = privacyMatch[1]
		}

		// Match trackers
		if trackerMatch := trackerPattern.FindStringSubmatch(line); trackerMatch != nil {
			trackers = append(trackers, domain.Tracker{
				Tier: i + 1, // Assuming the tier is based on order
				URL:  trackerMatch[0],
			})
		}

		// Match files (after "FILES" section)
		if strings.HasPrefix(line, "FILES") {
			for j := i + 1; j < len(lines); j++ {
				fileLine := strings.TrimSpace(lines[j])
				if fileMatch := filePattern.FindStringSubmatch(fileLine); fileMatch != nil {
					files = append(files, domain.File{
						FileName: strings.TrimSpace(fileMatch[1]),
						FileSize: strings.Trim(fileMatch[2], "()"),
					})
				}
			}
			break
		}
	}

	torrent.Trackers = trackers
	torrent.Files = files

	return &torrent, nil
}
