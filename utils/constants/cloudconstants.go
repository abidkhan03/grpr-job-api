package constants

import (
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    SerpAPIKey         string `envconfig:"API_KEY"`
    GCloudBucketName   string `envconfig:"GCLOUD_BUCKET_NAME"`
    GCloudProjectName  string `envconfig:"GCP_PROJECT"`
}

var CloudConfig Config

func init() {
    err := envconfig.Process("", &CloudConfig)
    if err != nil {
        panic(err)
    }
}

