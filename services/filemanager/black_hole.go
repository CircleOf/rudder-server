package filemanager

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type TimeBounds struct {
	LowerBound int
	UpperBound int
}

//BlackHoleFileManager simply absorbs all the data that is
//sent its way to capture and returns a consistent outcome
//of events.
type BlackHoleFileManager struct {
	// This file manager also captures the time-bounds for upload and download operations.
	// TODO: At this moment, we have same timebounds for upload and download
	// TODO: which can be separated in future.
	Config *TimeBounds
}

func GetBlackHoleConfig(config map[string]interface{}) *TimeBounds {

	lowerBound, _ := strconv.Atoi(config["objectStorageLowerBound"].(string))
	upperBound, _ := strconv.Atoi(config["objectStorageUpperBound"].(string))

	pkgLogger.Infof("LoadTest: creating black hole config with bounds: lower: %d and upper: %d",
		lowerBound,
		upperBound)

	return &TimeBounds{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	}
}

func (bh *BlackHoleFileManager) Upload(ctx context.Context, f *os.File, prefixes ...string) (UploadOutput, error) {
	pkgLogger.Infof("LoadTest(ObjectStorage)(BH): Received a call to upload to storage: %#v", prefixes)

	location, objectName := "", ""

	if sliceContains(prefixes, "rudder-warehouse-staging-logs") {
		location, objectName = "staging-logs", "stagingfile.json.gz"
	}

	if sliceContains(prefixes, "rudder-warehouse-load-objects") {

		if sliceContains(prefixes, "demo_track") {
			location, objectName = "load-objects", "demo-track-load-objects.csv.gz"
		}

		if sliceContains(prefixes, "tracks") {
			location, objectName = "load-objects", "tracks-load-objects.csv.gz"
		}
	}

	// We were unable to identify the upload request.
	// error out at this time.
	if location == "" {
		panic("Unable to identify the location of upload")
	}

	bh.delay() // Introduce some delay in upload
	return UploadOutput{
		Location:   location,
		ObjectName: objectName,
	}, nil
}

func sliceContains(entries []string, elem string) bool {

	for _, entry := range entries {
		if entry == elem {
			return true
		}
	}

	return false
}

func (bh *BlackHoleFileManager) Download(ctx context.Context, output *os.File, key string) error {
	pkgLogger.Infof("LoadTest(ObjectStorage)(BH): Received a call to download file, prefixes: %v", key)

	fname := ""
	switch {

	case strings.Contains(key, "staging"):
		fname = "./services/filemanager/load-test/staging.json.gz"

	case strings.Contains(key, "demo-track-load-objects"):
		fname = "./services/filemanager/load-test/demo-track-load-objects.csv.gz"

	case strings.Contains(key, "tracks-load-objects"):
		fname = "./services/filemanager/load-test/tracks-load-objects.csv.gz"

	default:
		panic("Unrecognized key: " + key)
	}

	bh.delay() // Introduce some delay in download.
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Open(fmt.Sprintf("%s/%s", dir, fname))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(output, f)
	return err
}

func (bh *BlackHoleFileManager) GetObjectNameFromLocation(string) (string, error) {
	return "", nil
}

func (bh *BlackHoleFileManager) GetDownloadKeyFromFileLocation(string) string {
	return "dummy-name"
}

func (bh *BlackHoleFileManager) DeleteObjects(context.Context, []string) error {
	return nil
}

func (bh *BlackHoleFileManager) ListFilesWithPrefix(context.Context, string, int64) (fileObjects []*FileObject, err error) {
	return nil, nil
}

func (bh *BlackHoleFileManager) GetConfiguredPrefix() string {
	return ""
}

func (bh *BlackHoleFileManager) SetTimeout(*time.Duration) {
}

// delay simply delays the execution of function in which it is called
// based on the added lower and upper bounds.
func (bh *BlackHoleFileManager) delay() {
	interval := bh.Config.LowerBound + rand.Intn(bh.Config.UpperBound-bh.Config.LowerBound)
	time.Sleep(time.Millisecond * time.Duration(interval))
}