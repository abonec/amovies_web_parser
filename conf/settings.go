package conf

import (
  "os"
  "fmt"
  "amovies_parser/models"
)

var (
  DOWNLOAD_DIR = ""
  TEMP_PREFIX = ".part"
  DOWNLOADS_FILE = "/tmp/downloads.json"
  DOWNLOADS = &models.Downloads{}
)

func init() {
  DOWNLOAD_DIR = setParam(DOWNLOAD_DIR, "DIR")
  TEMP_PREFIX = setParam(TEMP_PREFIX, "TEMP_PREFIX")
  DOWNLOADS_FILE = setParam(DOWNLOADS_FILE, "DOWNLOADS_FILE")

  DOWNLOADS.RestoreFile(DOWNLOADS_FILE)
  fmt.Println("Download dir is: ", DOWNLOAD_DIR)
}

func setParam(value, env string) string{
  user_set := os.Getenv(env)
  if user_set != "" {
    return user_set
  }
  return value
}
