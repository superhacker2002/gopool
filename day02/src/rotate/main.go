package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	archiveDir string
)

func init() {
	flag.StringVar(&archiveDir, "a", "", "directory to store archived files")
	flag.Parse()
}

func main() {
	logFiles := flag.Args()
	if len(logFiles) == 0 {
		log.Fatal("no log files provided")
	}

	var wg sync.WaitGroup
	for _, logFile := range logFiles {
		wg.Add(1)
		go func(logFile string) {
			defer wg.Done()
			err := rotateLogFile(logFile)
			if err != nil {
				log.Printf("error rotating log file %s: %v", logFile, err)
			}
		}(logFile)
	}

	wg.Wait()
}

func rotateLogFile(logFile string) error {
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		return fmt.Errorf("failed to get file info for %s: %w", logFile, err)
	}

	modTime := fileInfo.ModTime().Unix()
	archiveFile := fmt.Sprintf("%s_%d.tag.gz", filepath.Base(logFile), modTime)

	err = createArchive(archiveFile, logFile)
	if err != nil {
		return fmt.Errorf("failed to create archive %s: %w", archiveFile, err)
	}

	err = os.Rename(logFile, archiveFile)
	if err != nil {
		return fmt.Errorf("failed to move log file %s to archive: %w", logFile, err)
	}

	return nil
}

func createArchive(archiveFile, sourceFile string) error {
	outFile, err := os.Create(archiveFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	gw := gzip.NewWriter(outFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	stat, err := source.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    filepath.Base(sourceFile),
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
		ModTime: stat.ModTime(),
	}

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, source)
	if err != nil {
		return err
	}

	return nil
}
