package envar_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fabritsius/envar"
)

func ExampleFill() {
	// create couple environment variables
	createCoupleEnvs()

	type config struct {
		Name  string `env:"HERO"`
		Place string `env:"PLACE"`
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("You gotta do it for Grandpa, %s. You gotta put these seeds inside your %s.\n", cfg.Name, cfg.Place)
	// Output:
	// You gotta do it for Grandpa, Bob. You gotta put these seeds inside your backpack.
}

func createCoupleEnvs() {
	os.Setenv("HERO", "Bob")
	os.Setenv("PLACE", "backpack")
}

func TestDataTypeErrors(t *testing.T) {
	type config struct {
		Name string `default:"Rey"`
		Age  int    `default:"29"`
		Job  string `default:"peacemaking"`
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err == nil {
		t.Error("wrong data type error isn't handled")
	}
}

func TestDefaults(t *testing.T) {
	type config struct {
		Item  string `default:"Face"`
		Error string `default:"Can't feel"`
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err != nil {
		t.Error("defaults aren't set properly")
	}

	if cfg.Item != "Face" || cfg.Error != "Can't feel" {
		t.Error("defaults have incorrect values")
	}
}

func TestEmpty(t *testing.T) {
	type config struct {
		Day  string ``
		Task string ``
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err == nil {
		t.Error("nothing to set error doesn't work")
	}
}
