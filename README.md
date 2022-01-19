# Yousign API client written in Go

This library doesn't cover all the API, so contributions are welcome.

That said, the main parts are covered so you can create procedures, manage files,
add server stamps... etc. You can check the test files to have an overview.

### Usage

```go
client := NewClient("YOUR-API-KEY")

// Add a file for this procedure
pdf, err := os.Open("mydocument.pdf")
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

// Use this link to sign the document for a particular member of the
// procedure
fmt.Println("Sign URL:", client.SignURL(*p.Members[0].ID))
```
