package main

import ()

func tabelleVorbereiten(conf *PolymerConfig, colors []string) (*ManageInfo, error) {
	cols := len(Colors)
	wholeTime := float64(0)
	maxDepth := float64(0)
	n := len(conf.Limits)
	for i := 0; i < n; i++ {
		wholeTime = conf.Limits[i][0]
		if conf.Limits[i][1] > maxDepth {
			maxDepth = conf.Limits[i][1]
		}
	}
	if wholeTime < conf.TimeStep {
		conf.TimeStep = wholeTime
	}
	if wholeTime < conf.TimeShow {
		conf.TimeShow = wholeTime
	}
	if wholeTime < conf.GraphTimeStep {
		conf.GraphTimeStep = wholeTime
	}
	wholeSize := conf.Radius + maxDepth
	svgAllTime := int(wholeTime/conf.GraphTimeStep+0.01) + 1
	svgTStep := int((conf.StartTemperature-conf.QuenchingT)/cols + 0.01)
	svgT := int(conf.QuenchingT + 0.01)
	firstCol := cols - 1
	if svgT > conf.StartTemperature {
		svgT = int(conf.StartTemperature + 0.01)
		svgTStep = -svgTStep
		firstCol = 0
	}
	svgTable := make([][]int, svgAllTime)
	svgXSize := int(wholeSize/conf.GraphXStep) + 1
	csvLen := len(conf.Points)
	csvAllTime := int(wholeTime/conf.TimeShow+0.5) + 1
	csvTable := make([]float64, csvAllTime*csvLen)
	for i := 0; i < csvLen; i++ {
		if conf.Points[i] <= 100 {
			csvTable[i] = conf.StartTemperature
		} else {
			csvTable[i] = conf.QuenchingT
		}
	}
	extraInfo := &RechnenInfo{Config: conf,
		Colors:            colors,
		SvgColors:         cols,
		MaxDepth:          maxDepth,
		WholeSize:         wholeSize,
		WholeTime:         wholeTime,
		SvgTable:          svgTable,
		SvgTime:           0,
		SvgTBase:          svgT,
		SvgTStep:          svgTStep,
		SvgAllTime:        svgAllTime,
		SvgXSize:          svgXSize,
		CurrentDepthIndex: 0,
		CsvTime:           1,
		CsvLen:            csvLen,
		CsvAllTime:        csvAllTime,
		CsvTable:          csvTable,
	}
	rad := fmt.Sprintf("%f", conf.Radius)
	dep := fmt.Sprintf("%f", maxDepth)
	loop1 := dvthermo.BodyLoop{
		Inside: true,
		Parts: []ConditionPart{
			ConditionPart{
				Condition:   "NOFLOW",
				Description: "h" + rad,
			},
			ConditionPart{
				Condition:   "MID",
				Description: "v0.5",
				Boundary:    conf.Htc,
			},
			ConditionPart{
				Condition:   "HTC",
				Description: "v0.001",
				Boundary:    conf.Htc,
			},
			ConditionPart{
				Condition:   "NOFLOW",
				Description: "H0z",
			},
		},
		L: conf.SteelL,
		C: conf.SteelC,
		P: conf.SteelP,
	}
	loop2 := dvthermo.BodyLoop{
		Inside: true,
		Parts: []ConditionPart{
			ConditionPart{
				Condition:   "NOFLOW",
				Description: "M" + rad + " 0h" + dep,
			},
			ConditionPart{
				Condition:   "CONSTANT",
				Description: "v0.501",
				Boundary:    conf.QuenchingT,
			},
			ConditionPart{
				Condition:   "NOFLOW",
				Description: "h-" + dep + "z",
			},
		},
		L: conf.SteelL,
		C: conf.SteelC,
		P: conf.SteelP,
	}
	shape := &dvthermo.BodyShape{
		Axes:     "Z",
		MeshDims: []int{conf.Mesh},
		Loops:    []dvthermo.BodyLoop{loop1, loop2},
	}
	body, err := dvthermo.MakeBodyMeshing(shape)
	manager := &dvthermo.ManageInfo{
		CurrentStep:   0,
		AllSteps:      int(wholeTime/conf.TimeStep+0.9) + 1,
		StepTime:      conf.TimeStep,
		AfterTimeStep: intercessor,
		Extra:         extraInfo,
		Body:          body,
	}
	extraInfo.LastSteelPoint, _, _ = manager.GetPointDimensions(conf.Radius, 0)
	return manager, err
}
