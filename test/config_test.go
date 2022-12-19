package config_test

import (
	"fmt"
	"go-blog/internal/config"
	"testing"
)

func TestPgConfig(t *testing.T) {
	fmt.Printf("pgconfig:%+v", config.PgConfig)
}
