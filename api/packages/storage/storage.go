package storage

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// Naive in memory cache
var cachedContents [][]byte
var cachedLength int

// Read all files in the given bucket and return a byte slice for each file found
func ReadFiles(bucketName string) ([][]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	filenames, err := listFiles(ctx, bucket)
	if err != nil {
		return nil, err
	}

	// @note
	// There does not seem to exist a good way of getting the last updated date for a bucket
	// through the SDK so we'll resort to using the number of read files as our cache distinguisher
	if len(filenames) == cachedLength {
		return cachedContents, nil
	}

	contents, err := readFiles(ctx, bucket, filenames)
	if err != nil {
		return nil, err
	}

	// Update the "cache"
	cachedContents = contents
	cachedLength = len(filenames)

	return contents, nil
}

// Return a list of all filenames found in the bucket
func listFiles(ctx context.Context, bucket *storage.BucketHandle) ([]string, error) {
	filenames := []string{}

	it := bucket.Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []string{}, err
		}

		filenames = append(filenames, attrs.Name)
	}

	return filenames, nil
}

// Return a byte slice of data for each given filename
func readFiles(ctx context.Context, bucket *storage.BucketHandle, filenames []string) ([][]byte, error) {
	contents := make([][]byte, len(filenames))

	for i, filename := range filenames {
		reader, err := bucket.Object(filename).NewReader(ctx)
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		reader.Close()

		contents[i] = body
	}

	return contents, nil
}
