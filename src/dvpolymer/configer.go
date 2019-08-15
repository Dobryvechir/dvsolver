package main

import (
	"encoding/json"
	"errors"
)

func validateConfig(conf *PolymerConfig) error {
	if conf.Radius <= 0 {
		return errors.New("radius must be positive")
	}
	if conf.SteelL <= 0 {
		return errors.New("steelL must be positive")
	}
	if conf.SteelC <= 0 {
		return errors.New("steelC must be positive")
	}
	if conf.SteelP <= 0 {
		return errors.New("SteelP must be positive")
	}
	if conf.PolymerL <= 0 {
		return errors.New("polymerL must be positive")
	}
	if conf.PolymerC <= 0 {
		return errors.New("polymerC must be positive")
	}
	if conf.PolymerP <= 0 {
		return errors.New("polymerP must be positive")
	}
	if conf.QuenchingT < 0 {
		return errors.New("quenchingT must not be negative")
	}
	if conf.Mesh < 3 {
		return errors.New("mesh must be at least 3")
	}
	if conf.TimeStep <= 0 {
		return errors.New("timeStep must be positive")
	}
	if conf.TimeShow < conf.TimeStep {
		return errors.New("timeShow must be greater than or equal to timeStep")
	}
	p := len(conf.Points)
	if p == 0 {
		return errors.New("points must be specified")
	}
	for i := 0; i < p; i++ {
		if conf.Points[i] < 0 {
			return errors.New("points must be not negative")
		}
		if conf.Points[i] >= 200 {
			return errors.New("points must be less than 200")
		}
	}
	if conf.GraphTimeStep <= 0 {
		return errors.New("graphTimeStep must be positive")
	}
	if conf.GraphXStep <= 0 {
		return errors.New("graphXStep must be positive")
	}
	if conf.StartTemperature <= 0 {
		return errors.New("startTemperature must be positive")
	}
	if conf.Htc <= 0 {
		return errors.New("htc must be positive")
	}
	l := len(conf.Layer)
	if l <= 1 {
		return errors.New("layer must have at least 2 entries")
	}
	for i := 0; i < l; i++ {
		c := conf.Layer[i]
		if len(c) != 2 {
			return errors.New("each entry in the layer must have 2 numbers: time in seconds and polymer layer depth in meters")
		}
		if i == 0 {
			if c[0] != 0 {
				return errors.New("first number of each entry in the layer is time in seconds and at first it must be zero")
			}
		} else {
			if c[0] <= conf.Layer[i-1][0] {
				return errors.New("first number of each entry in the layer is time in seconds and it must be in ascential order")
			}
		}
		if c[1] < 0 {
			return errors.New("second number of each entry in the layer is polymer layer depth in meters and it must not be negative")
		}
	}
	return nil
}

func ReadConfig() (conf *PolymerConfig, err error) {
	conf = &PolymerConfig{}
	var data []byte
	data, err = ioutil.ReadFile("polymer.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, conf)
	if err == nil {
		err = validateConfig(conf)
	}
	return
}

func validateColors(conf *FarbeTabelle) error {
    c:=conf.Colors
    n:=len(c)
    if n<2 {
	return errors.New("At least 2 colors must be present")
    }
    for i:=0;i<n;i++ {
         d:=c[i]
         if len(d)!=7 || d[0]!='#' {
              return errors.New("Each color must have format #xxxxxx")
         }
         for j:=1;j<=6;j++ {
             k:=d[j]
             if !(k>='0' && k<='9' || k>='a' && k<='f' || k>='A' && k<='F') {
                return errors.New("Each color must have format #09afAF")
             }
         }
    }
    return nil
}

func ReadColors() (conf *FarbeTabelle, err error) {
	conf = &FarbeTabelle{}
	var data []byte
	data, err = ioutil.ReadFile("colors.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, conf)
	if err == nil {
		err = validateColors(conf)
	}
	return
}
