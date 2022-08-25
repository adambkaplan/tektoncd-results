package log

import (
	"bufio"
	"errors"
	"path/filepath"
	"strings"

	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/tektoncd/results/pkg/apis/v1alpha2"
	"github.com/tektoncd/results/pkg/conf"
)

type S3LogStreamer struct {
	LogStreamer
	io.WriterTo
	io.ReaderFrom

	bufSize int
	conf    *conf.ConfigFile
	path string
}

func NewS3LogStreamer(trl *v1alpha2.TaskRunLog, bufSize int, conf *conf.ConfigFile, logDataDir string) LogStreamer {
	if trl.Status.File == nil {
		trl.Status.File = &v1alpha2.FileLogTypeStatus{
			Path: defaultFilePath(trl),
		}
	}
	return &S3LogStreamer{
		conf:    conf,
		bufSize: bufSize,
		path: filepath.Join(logDataDir, trl.Status.File.Path),
	}
}

func (*S3LogStreamer) Type() string {
	return string(v1alpha2.S3LogType)
}

func (ls *S3LogStreamer) WriteTo(w io.Writer) (n int64, err error) {
	client := ls.initConfig()

	outPut, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &ls.conf.S3_BUCKET_NAME,
		Key:    &ls.path,
	})
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	defer outPut.Body.Close()

	reader := bufio.NewReaderSize(outPut.Body, ls.bufSize)
	n, err = reader.WriteTo(w)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return
}

func (ls *S3LogStreamer) ReadFrom(r io.Reader) (n int64, err error) {
	client := ls.initConfig()

	prevContent, err := ls.ReadS3(client)
	if err != nil {
		var oe *types.NoSuchKey
		if errors.As(err, &oe) {
			prevContent = []byte{}
		} else {
			return 0, fmt.Errorf(err.Error())
		}
	}

	newBts, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &ls.conf.S3_BUCKET_NAME,
		Key:    &ls.path,

		Body: strings.NewReader(string(prevContent) + string(newBts)),
	})

	if err != nil {
		return 0, fmt.Errorf("failed to put object ot S3 %v", err.Error())
	}
	fmt.Println("OK. Data was sent to S3 successfully!")

	return int64(len(newBts)), nil
}

func (ls *S3LogStreamer) ReadS3(client *s3.Client) ([]byte, error) {
	outPut, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &ls.conf.S3_BUCKET_NAME,
		Key:    &ls.path,
	})
	if err != nil {
		return []byte{}, err
	}

	defer outPut.Body.Close()

	bts, err := io.ReadAll(outPut.Body)
	if err != nil {
		return []byte{}, err
	}

	return bts, nil
}

// todo does this init once?
func (ls *S3LogStreamer) initConfig() *s3.Client {
	credentialsOpt := config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(ls.conf.S3_ACCESS_KEY_ID, ls.conf.S3_SECRET_ACCESS_KEY, ""))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(ls.conf.S3_REGION), credentialsOpt)

	if len(ls.conf.S3_ENDPOINT) > 0 {
		customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == dynamodb.ServiceID && region == ls.conf.S3_REGION {
				return aws.Endpoint{
					//PartitionID:   "aws",
					URL:           ls.conf.S3_ENDPOINT,
					SigningRegion: ls.conf.S3_REGION,
				}, nil
			}
			// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver), credentialsOpt)
	}

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return s3.NewFromConfig(cfg)
}
