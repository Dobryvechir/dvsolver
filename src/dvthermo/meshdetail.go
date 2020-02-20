package dvthermo

func makeMeshingInVertical(x []float64, body *BodyComposition) (res float64, pointer int) {
	//TODO: 3 dimensional
	//Here we implement 2 directional
        loops:=body.Loops
	cond := POINT_OUT
        res = body.Mesh.Max[1]
        n:=len(loops)
        for i:=0; i< n; i++ {
           loop:=loops[i]
           parts:=loop.Parts
           m:=len(parts)
           for j:=0;j<m;j++ {
/********************************************************************************************
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

*********************************************************************************************/
                todo!!!!!!!!!!!!!!!!!!!!!!!!!
           }
        }
        return
}

func MakeMeshingInDirections(body *BodyComposition) error {
	dim := body.Mesh.NumDim
	last := body.Mesh.Dims[dim-1]
	n := 1
	m := 1
	for i := 0; i < dim; i++ {
		m = n
		n *= body.Mesh.Dims[i]
	}
	currentPoint := make([]float64, dims)
	body.Mesh.Val = make([]float64, n)
	kinds := make([]int, n)
	body.Mesh.PointCondition = make([]int, n)
	fb := body.Mesh.Min[dim-1]
	fk := (body.Mesh.Max[dim-1] - fb) / (last - 1)
	for j := 0; j < m; j++ {
		vm := j
		for i := dim - 1; i >= 0; i-- {
			p := body.Mesh.Dims[i]
			k := vm % p
			vm = vm / p
			currentPoint[i] = (body.Mesh.Max[i]-body.Mesh.Min[i])*k/(p-1) + body.Mesh.Min[i]
		}
		k := 0
		s := kinds[j*last:]
		for k < last {
			currentPoint[dim-1] = fk*k + fb
			val, kind := makeMeshingInVertical(currentPoint, body)
			s[k] = kind
			k++
			lastK := int((val-fb)/fk + 0.5)
			if lastK > last {
				lastK = last
			}
			for ; k < lastK; k++ {
				s[k] = kind
			}
		}
	}
	return nil
}
