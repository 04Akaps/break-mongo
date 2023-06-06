package document

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
)

func (d *Document) UploadFile(file *os.File) {
	uploadOpts := options.GridFSUpload().SetMetadata(bson.M{"metadata_tag": "first Upload"})

	if objectID, err := d.fileBucket.UploadFromStream("file.txt", io.Reader(file), uploadOpts); err != nil {
		panic(err)
	} else {
		log.Println(objectID)
	}
}

func (d *Document) UploadFileByStraem(file *os.File) {
	uploadOpts := options.GridFSUpload().SetChunkSizeBytes(200000)

	if uploadStream, err := d.fileBucket.OpenUploadStream("streamFile.txt", uploadOpts); err != nil {
		panic(err)
	} else {
		defer uploadStream.Close() // 업로드 스트림 닫기

		if _, err := io.Copy(uploadStream, file); err != nil {
			panic(err)
		} else {
			fmt.Printf("Stream Upload: %v\n", uploadStream.FileID)
		}
	}
}

func (d *Document) RetrieveFile() {
	filter := bson.M{
		"filename": "streamFile.txt",
	}

	cursor, err := d.fileBucket.Find(filter)

	if err != nil {
		panic(err)
	}

	type gridfsFile struct {
		Name   string             `bson:"filename"`
		Length int64              `bson:"length"`
		Id     primitive.ObjectID `bson:"_id"`
	}

	var foundFiles []gridfsFile

	if err = cursor.All(context.TODO(), &foundFiles); err != nil {
		panic(err)
	}

	for _, file := range foundFiles {
		fmt.Printf("filename: %s, length: %d, Id : %v \n", file.Name, file.Length, file.Id)
	}
}

func (d *Document) DownloadFile() {
	id, _ := primitive.ObjectIDFromHex("647ec47989598ac125402b95")
	fileBuffer := bytes.NewBuffer(nil)
	if _, err := d.fileBucket.DownloadToStream(id, fileBuffer); err != nil {
		panic(err)
	}

	outputFilePath := "./outputText.txt"
	if err := os.WriteFile(outputFilePath, fileBuffer.Bytes(), 0644); err != nil {
		panic(err)
	}

	fmt.Println("File downloaded and saved successfully.")
}

func (d *Document) ReNameFile() {
	id, _ := primitive.ObjectIDFromHex("647ec47989598ac125402b95")
	if err := d.fileBucket.Rename(id, "mongodbTutorial.zip"); err != nil {
		panic(err)
	}
}

// Delete, bucket 데이터 다 지우는 Drop
