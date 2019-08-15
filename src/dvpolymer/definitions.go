package main

type PolymerConfig struct {
	Radius           float64     `json:"radius"`
	SteelL           float64     `json:"steelL"`
	SteelC           float64     `json:"steelC"`
	SteelP           float64     `json:"steelP"`
	PolymerL         float64     `json:"polymerL"`
	PolymerC         float64     `json:"polymerC"`
	PolymerP         float64     `json:"polymerP"`
	QuenchingT       float64     `json:"quenchingT"`
	Mesh             int         `json:"mesh"`
	TimeStep         float64     `json:"timeStep"`
	TimeShow         float64     `json:"timeShow"`
	Points           []float64   `json:"points"`
	GraphTimeStep    float64     `json:"graphTimeStep"`
	GraphXStep       float64     `json:"graphXStep"`
	StartTemperature float64     `json:"startTemperature"`
	Htc              float64     `json:"htc"`
	Layer            [][]float64 `json:"layer"`
}

type FarbeTabelle struct {
	Colors []string `json:"colors"`
}

type RechnenInfo struct {
	Config            *PolymerConfig
	Colors            []string
	MaxDepth          float64
	CurrentDepthIndex int
	WholeSize         float64
	WholeTime         float64
	SvgTime           int
	SvgAllTime        int
	SvgTBase          int
	SvgTStep          int
	SvgColors         int
	SvgXSize          int
	SvgTable          [][]int
	CsvTime           int
	CsvAllTime        int
	CsvLen            int
	CsvTable          []float64
	LastSteelPoint    int
}
