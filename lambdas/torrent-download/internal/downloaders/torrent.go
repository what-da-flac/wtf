package downloaders

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/what-da-flac/wtf/go-common/ifaces"
)

type TorrentDownloader struct {
	logger  ifaces.Logger
	timeout time.Duration
}

func NewTorrentDownloader(logger ifaces.Logger, timeout time.Duration) *TorrentDownloader {
	return &TorrentDownloader{
		logger:  logger,
		timeout: timeout,
	}
}

// Start starts transmission daemon in background.
func (x *TorrentDownloader) Start() error {
	x.logger.Info("starting torrent downloader")
	cmd := exec.Command("transmission-daemon")
	if err := cmd.Start(); err != nil {
		return err
	}
	return x.waitForStart(x.timeout)
}

func (x *TorrentDownloader) waitForStart(wait time.Duration) error {
	interval := time.Second * 2
	timeout := time.After(wait)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	checkFn := func() bool {
		cmd := exec.Command("transmission-remote", "--list")
		data, err := cmd.Output()
		if err != nil {
			return false
		}
		str := string(data)
		if strings.Contains(str, "ID") {
			x.logger.Infof("daemon responding: %s", str)
			return true
		}
		return false
	}

	for {
		select {
		case <-timeout:
			err := fmt.Errorf("torrent download timed out")
			x.logger.Errorf("torrent download timed out: %v", err)
			return err
		case <-ticker.C:
			x.logger.Info("waiting for torrent download")
			if checkFn() {
				x.logger.Info("torrent download completed")
				return nil
			}
		}
	}
}

// AddTorrent adds a torrent file to download stack.
func (x *TorrentDownloader) AddTorrent(targetDir, torrentFileName string) error {
	x.logger.Infof("adding torrent: %s", torrentFileName)
	cmd := exec.Command(
		"transmission-remote",
		"--download-dir", targetDir,
		"--add", torrentFileName,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		x.logger.Errorf("failed to add torrent: %s, error: %v", torrentFileName, err)
		x.logger.Debugf("command output: %s", output)
		return err
	}
	return nil
}

// checkCompleted reads from current torrents being processed,
// and returns true if download has completed.
func (x *TorrentDownloader) checkCompleted() bool {
	cmd := exec.Command("transmission-remote", "--list")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if x.checkLine(line) {
			return true
		}
	}
	return false
}

func (x *TorrentDownloader) checkLine(line string) bool {
	expected := map[string]struct{}{
		"100%": {},
		"Done": {},
	}
	values := strings.Fields(line)
	for _, v := range values {
		delete(expected, v)
	}
	return len(expected) == 0
}

// WaitForDownload checks for torrent file to be downloaded
// using a retry mechanism.
func (x *TorrentDownloader) WaitForDownload(wait, interval time.Duration) bool {
	timeout := time.After(wait)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			x.logger.Info("torrent download timed out")
			return false
		case <-ticker.C:
			x.logger.Info("waiting for torrent download")
			if x.checkCompleted() {
				x.logger.Info("torrent download completed")
				return true
			}
		}
	}
}

func (x *TorrentDownloader) ClearAll() error {
	cmd := exec.Command("transmission-remote", "-t", "all", "--remove-and-delete")
	output, err := cmd.CombinedOutput()
	if err != nil {
		x.logger.Errorf("failed to clear torrents: %v", err)
		x.logger.Debugf("command output: %s", output)
		return err
	}
	return nil
}