package config

import (
	"math/rand"
	"os"
	"order-service/exception"
	"time"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	Timezone() *time.Location
	RangeIn(low, hi int) int
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config *configImpl) Timezone() *time.Location {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return loc
}

func (config *configImpl) RangeIn(low, hi int) int {
	rand.Seed(time.Now().Unix())
	return low + rand.Intn(hi-low)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
