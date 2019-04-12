// Copyright (c) - Damien Fontaine <damien.fontaine@lineolia.net>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package record

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/DamienFontaine/lunarc/datasource/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/gridfs"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

//Record audio
type Record struct {
	Metadata Metadata           `json:"metadata" bson:"metadata"`
	Filename string             `json:"filename" bson:"filename"`
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

//Metadata audio text
type Metadata struct {
	Text string `json:"text" bson:"text"`
	Set  string `json:"set" bson:"set"`
}

//Manager manage records
type Manager interface {
	Add(r Record, reader io.Reader) (Record, error)
	FindAll() ([]Record, error)
	FindPerPage(int32, int32) ([]Record, error)
	Delete(string) error
	Update(id string, set string) (Record, error)
	Upload(id string) (io.Reader, error)
	Count() (int64, error)
	Get(id string) (Record, error)
	Export() (io.Reader, error)
}

//Service works with models.Thing
type Service struct {
	MongoService mongo.Service
}

//NewService creates a new Service
func NewService(ms mongo.Service) Service {
	s := Service{MongoService: ms}
	return s
}

//Count records in datasource
func (s *Service) Count() (count int64, err error) {
	c := s.MongoService.Mongo.Database.Collection("records.files", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err = c.Count(ctx, bsonx.Doc{})
	if err != nil {
		return 0, err
	}
	return
}

//Add a new Record in DataSource
func (s *Service) Add(r Record, reader io.Reader) (Record, error) {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return r, err
	}

	doc := bsonx.Doc{}
	doc = doc.Append("text", bsonx.String(r.Metadata.Text))
	doc = doc.Append("set", bsonx.String(r.Metadata.Set))

	optsUpload := options.GridFSUpload()
	optsUpload.SetMetadata(doc)
	t := time.Now()
	r.Filename = fmt.Sprintf("%v.wav", t.Unix())
	r.ID, err = bucket.UploadFromStream(r.Filename, reader, optsUpload)
	if err != nil {
		return r, err
	}
	return r, nil
}

//Delete a Record in DataSource
func (s *Service) Delete(id string) error {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return err
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	err = bucket.Delete(objectID)
	if err != nil {
		return err
	}
	return nil
}

//FindAll Record in DataSource
func (s *Service) FindAll() (r []Record, err error) {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return r, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := bucket.Find(ctx, nil)
	if err != nil {
		return r, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var record Record
		err := res.Decode(&record)
		if err != nil {
			log.Printf("Error: %v", err)
			return r, err
		}
		r = append(r, record)
	}
	return r, nil
}

//FindPerPage Record in DataSource
func (s *Service) FindPerPage(page int32, limit int32) (r []Record, err error) {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return r, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	optsFindOptions := options.GridFSFind()
	optsFindOptions.SetSkip(page * limit)
	optsFindOptions.SetLimit(limit)
	res, err := bucket.Find(ctx, optsFindOptions)
	if err != nil {
		return r, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var record Record
		err := res.Decode(&record)
		if err != nil {
			log.Printf("Error: %v", err)
			return r, err
		}
		r = append(r, record)
	}
	return r, nil
}

//Upload a record
func (s *Service) Upload(id string) (r io.Reader, err error) {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return r, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	r, err = bucket.OpenDownloadStream(objectID)
	return
}

// Get a Record
func (s *Service) Get(id string) (r Record, err error) {
	optsBucket := options.GridFSBucket()
	optsBucket.SetName("records")
	bucket, err := gridfs.NewBucket(s.MongoService.Mongo.Database, optsBucket)
	if err != nil {
		return r, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)

	res, err := bucket.Find(bsonx.Doc{bsonx.Elem{Key: "_id", Value: bsonx.ObjectID(objectID)}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for res.Next(ctx) {
		err := res.Decode(&r)
		if err != nil {
			log.Printf("Error: %v", err)
			return r, err
		}
	}
	return r, nil
}

//Update Record in DataSource
func (s *Service) Update(id string, set string) (r Record, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)

	res, err := s.MongoService.Mongo.Database.Collection("records.files").UpdateOne(ctx, bsonx.Doc{bsonx.Elem{Key: "_id", Value: bsonx.ObjectID(objectID)}}, bsonx.Doc{bsonx.Elem{Key: "$set", Value: bsonx.Document(bsonx.Doc{bsonx.Elem{Key: "metadata.set", Value: bsonx.String(set)}})}})
	if err != nil {
		return r, err
	}

	if res.MatchedCount == 0 {
		return r, errors.New("file with this ID not found")
	}

	r, err = s.Get(id)
	if err != nil {
		return r, err
	}
	return r, nil
}

// Export all records
func (s *Service) Export() (r io.Reader, err error) {

	records, err := s.FindAll()
	if err != nil {
		return
	}

	// Create structure
	tmpDirPath, err := ioutil.TempDir("", "export-")
	if err != nil {
		return
	}

	testDirPath := filepath.Join(tmpDirPath, "test")
	err = os.Mkdir(testDirPath, 0700)
	if err != nil {
		return
	}

	test, err := os.Create(filepath.Join(testDirPath, "test.csv"))
	defer test.Close()
	if err != nil {
		return
	}
	test.WriteString("wav_filename,wav_filesize,transcript\n")

	devDirPath := filepath.Join(tmpDirPath, "dev")
	err = os.Mkdir(devDirPath, 0700)
	if err != nil {
		return
	}

	dev, err := os.Create(filepath.Join(devDirPath, "dev.csv"))
	defer dev.Close()
	if err != nil {
		return
	}
	dev.WriteString("wav_filename,wav_filesize,transcript\n")

	trainDirPath := filepath.Join(tmpDirPath, "train")
	err = os.Mkdir(trainDirPath, 0700)
	if err != nil {
		return
	}

	train, err := os.Create(filepath.Join(trainDirPath, "train.csv"))
	defer train.Close()
	if err != nil {
		return
	}
	train.WriteString("wav_filename,wav_filesize,transcript\n")

	// Create vocabulary
	voc, err := os.Create(filepath.Join(tmpDirPath, "vocabulary.txt"))
	defer voc.Close()
	if err != nil {
		return
	}

	for _, record := range records {
		// Add in vocabulary
		voc.WriteString(record.Metadata.Text + "\n")

		// Create wav
		reader, err := s.Upload(record.ID.Hex())
		if err != nil {
			return nil, err
		}

		var path string
		if strings.Compare(record.Metadata.Set, "train") == 0 {
			path = trainDirPath
		} else if strings.Compare(record.Metadata.Set, "dev") == 0 {
			path = devDirPath
		} else {
			path = testDirPath
		}

		file, err := os.Create(filepath.Join(path, record.Filename))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		written, err := io.Copy(file, reader)
		if err != nil {
			return nil, err
		}

		if strings.Compare(record.Metadata.Set, "train") == 0 {
			train.WriteString(fmt.Sprintf("/data/train/%s,%d,%s\n", record.Filename, written, record.Metadata.Text))
		} else if strings.Compare(record.Metadata.Set, "dev") == 0 {
			dev.WriteString(fmt.Sprintf("/data/dev/%s,%d,%s\n", record.Filename, written, record.Metadata.Text))
		} else {
			test.WriteString(fmt.Sprintf("/data/test/%s,%d,%s\n", record.Filename, written, record.Metadata.Text))
		}
	}

	//Create ZIP
	zipFile, err := ioutil.TempFile("", "export-*.zip")
	if err != nil {
		return
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	baseDir := filepath.Base(tmpDirPath)

	filepath.Walk(tmpDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, tmpDirPath))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)

		return err
	})

	zipFile, err = os.Open(filepath.Join(filepath.Dir(tmpDirPath), filepath.Base(zipFile.Name())))

	return zipFile, err
}
