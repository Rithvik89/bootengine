package store

func StringifyS3() string {

	s3 := `
	// For S3 init.
	// https://github.com/Rithvik89/pkg-store/blob/main/pkg/object_store/aws_s3/aws_s3.go
		 
	func initS3 (AWSAccessID, AWSSecret, Region, Bucket string) *S3.S3{
		s3,err := S3.New(AWSAccessID, AWSSecret, Region, Bucket)
		if err != nil {
			panic(err)
		}
		return s3
	}`
	return s3
}
