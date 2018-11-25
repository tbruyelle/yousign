package yousign

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestProcedure(t *testing.T) {
	// Create a procedure
	p, resp, err := client.Procedure.Create(&ProcedureRequest{
		Name:        String("Procedure 1"),
		Description: String("My first procedure"),
		Start:       Bool(false),
	})
	if err != nil {
		fatal(t, err, resp)
	}

	p, resp, err = client.Procedure.Get(*p.ID)
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("procedure created:", *p.ID)

	// Add a file for this procedure
	pdf, err := os.Open("testdata/minimal.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer pdf.Close()

	file, resp, err := client.File.Create(&FileRequest{
		Procedure: p.ID,
		Name:      String("myFile"),
		Content:   base64Encode(pdf),
	})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file created:", *file.ID)

	file, resp, err = client.File.Get(*file.ID)
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file found:", *file.ID)

	content, resp, err := client.File.Download(*file.ID)
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file content:", *content)

	// Add a member to the procudure
	member, resp, err := client.Member.Create(&MemberRequest{
		Procedure: p.ID,
		Firstname: String("John"),
		Lastname:  String("Doe"),
		Phone:     String("+33xxxxxxxxx"),
		Email:     String("jdoe@mail.com"),
	})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("member added:", *member.ID)

	member, resp, err = client.Member.Get(*member.ID)
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("member found:", *member.ID)

	// Add a file object to the member
	fileObject, resp, err := client.FileObject.Create(&FileObjectRequest{
		Member:   member.ID,
		File:     file.ID,
		Page:     Int(1),
		Position: String("164,61,263,123"),
	})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file object added:", *fileObject.ID)

	fileObject, resp, err = client.FileObject.Get(*fileObject.ID)
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file object found:", *fileObject.ID, *fileObject.File, *fileObject.Member.ID)

	// Start the procedure
	p, resp, err = client.Procedure.Update(*p.ID, &ProcedureRequest{Start: Bool(true)})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("procedure started:", *p.ID)
}

func TestProcedureFast(t *testing.T) {
	// Add a file for this procedure
	pdf, err := os.Open("testdata/minimal.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer pdf.Close()

	file, resp, err := client.File.Create(&FileRequest{
		Name:    String("myFile"),
		Content: base64Encode(pdf),
	})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("file created:", *file.ID)

	// Stamp the file with a server stamp
	stamp, err := os.Open("testdata/Company-Stamp.png")
	if err != nil {
		t.Fatal(err)
	}
	defer stamp.Close()

	serverStamp, resp, err := client.ServerStamp.Create(&ServerStampRequest{
		File:      file.ID,
		SignImage: base64Encode(stamp),
		FileObjects: []FileObjectRequest{
			{
				Page:     Int(1),
				Position: String("24,65,122,127"),
			},
		},
	})

	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("server stamp created:", *serverStamp.ID)

	for {
		serverStamp, resp, err = client.ServerStamp.Get(*serverStamp.ID)
		if err != nil {
			fatal(t, err, resp)
		}
		fmt.Println("server stamp found:", *serverStamp.ID, *serverStamp.Status)
		if *serverStamp.Status != "pending" {
			break
		}
		time.Sleep(time.Second)
	}

	// Create a procedure
	p, resp, err := client.Procedure.Create(&ProcedureRequest{
		Name:        String("Procedure 1"),
		Description: String("My first procedure"),
		Start:       Bool(true),
		Members: []MemberRequest{
			{
				Firstname: String("John"),
				Lastname:  String("Doe"),
				Phone:     String("+33xxxxxxxxx"),
				Email:     String("jdoe@mail.com"),
				FileObjects: []FileObjectRequest{
					{
						Page:     Int(1),
						Position: String("164,61,263,123"),
						File:     file.ID,
					},
				},
			},
		},
	})
	if err != nil {
		fatal(t, err, resp)
	}
	fmt.Println("Procedure created and started:", *p.ID)

	fmt.Println("Sign URL:", client.SignURL(*p.Members[0].ID))

}

func base64Encode(r io.Reader) *string {
	var b bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &b)
	_, err := io.Copy(encoder, r)
	if err != nil {
		panic(err)
	}
	encoder.Close()
	return String(b.String())
}
