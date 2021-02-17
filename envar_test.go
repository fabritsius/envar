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
		Age  uint   `default:"29"`
		Job  int    `default:"peacemaking"`
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

func TestInts(t *testing.T) {
	type config struct {
		Port  int   `default:"3000"`
		Zero  int   `default:"0"`
		Three int8  `default:"3"`
		Four  int16 `default:"4"`
		Five  int32 `default:"5"`
		Six   int64 `default:"6"`
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err != nil {
		t.Error("int values aren't set properly:", err)
	}

	if cfg.Port != 3000 || cfg.Zero != 0 {
		t.Error("int values have incorrect values")
	}

	type badConfig struct {
		Name int `default:"Benedict"`
	}
	badcfg := badConfig{}
	if err := envar.Fill(&badcfg); err == nil {
		t.Error("non-int values are parsed as int")
	}
}

func TestBools(t *testing.T) {
	type config struct {
		Allow bool `default:"true"`
		Hero  bool `default:"false"`
	}

	cfg := config{}
	if err := envar.Fill(&cfg); err != nil {
		t.Error("int values aren't set properly")
	}

	if cfg.Allow != true || cfg.Hero != false {
		t.Error("bool values have incorrect values")
	}

	type badConfig struct {
		Number bool `default:"6"`
	}
	badcfg := badConfig{}
	if err := envar.Fill(&badcfg); err == nil {
		t.Error("non-bool values are parsed as bool")
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
