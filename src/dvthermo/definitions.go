package dvthermo

const MAX_DIMS = 2

type BodyMesh struct {
	NumDim         int
	Dims           [MAX_DIMS]int
	Min            [MAX_DIMS]float64
	Max            [MAX_DIMS]float64
	Val            []float64
	PointCondition []int
}

const (
	POINT_OUT = iota
	POINT_MID
	POINT_CONSTANT
	POINT_NO_FLOW
	POINT_HTC
)

type PointCondition struct {
	pointKind [MAX_DIMS]int
	thermo    float64
	boundary  [MAX_DIMS]float64
}

type ConditionPart struct {
	Condition   string  `json:"condition"`
	Description string  `json:"description"`
	Boundary    float64 `json:"boundary"`
}

type LoopInfo struct {
	Parts  []*CurveGeoPath
	Path   string
	Inside bool
	Min    [MAX_DIMS]float64
	Max    [MAX_DIMS]float64
	L      float64
	C      float64
	P      float64
}

type BodyLoop struct {
	Inside bool            `json:"inside"`
	Parts  []ConditionPart `json:"parts"`
	L      float64         `json:"L"`
	C      float64         `json:"C"`
	P      float64         `json:"P"`
}

type BodyShape struct {
	Axes     string     `json:"axes"`
	MeshDims []int      `json:"dims"`
	Loops    []BodyLoop `json:"loops"`
}

type ManageInfo struct {
	CurrentStep   int
	AllSteps      int
	StepTime      float64
	AfterTimeStep func(*ManageInfo) error
	Extra         interface{}
	Body          *BodyComposition
}

type BodyComposition struct {
	Shape *BodyShape
	Mesh  *BodyMesh
	Loops []*LoopInfo
}

const (
	CURVE_MOVE = 0
	CURVE_LINE = 1
	LINE_H     = 1
	LINE_V     = 2
	LINE_HV    = 3
)

type CurveGeoPoint struct {
	Kind           int
	SubKind        int
	X              [MAX_DIMS]float64
	Condition      int
	ConditionValue float64
}

type CurveGeoPath struct {
	Min    [MAX_DIMS]float64
	Max    [MAX_DIMS]float64
	Points []*CurveGeoPoint
}
