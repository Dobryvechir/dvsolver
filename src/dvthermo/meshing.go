package dvthermo

import (
	"errors"
	"strconv"
)

//allowable letters:
//M m move, only at first is allowed: M10 -20
//L l - line : L30.4 56
//H h - x: H34
//V v - y: V35
func prepareLoopInfo(loops []BodyLoop) (info []*LoopInfo, genMin [MAX_DIMS]float64,genMax [MAX_DIMS]float64, err error) {
	n := len(loops)
	info = make([]*LoopInfo, n)
	for i := 0; i < n; i++ {
		parts, path, e := ConvertPathToCurveGeoPaths(loops[i].Parts)
		if e != nil {
			return nil, nil, e
		}
		m := len(parts)
		if m == 0 {
			return nil, nil, errors.New("Non-descriptive borders at " + strconv.Itoa(i+1))
		}
                var min, max [MAX_DIMS]float64
                for d:=0;d<MAX_DIMS;d++ {
                      min[d]=parts[0].Min[d]
                      max[d]=parts[0].Max[d]
                }
		for j := 1; j < m; j++ {
                     for d:=0;d<MAX_DIMS;d++ {
			if parts[j].Min[d] < min[d] {
				min[d] = parts[j].Min[d]
			}
			if parts[j].Max[d] > max[d] {
				max[d] = parts[j].Max[d]
			}
                    } 
		}
		if i == 0 {
                     for d:=0;d<MAX_DIMS;d++ {
                          genMax[d] = max[d]
                          genMin[d] = min[d]
                     }
		} else {
                     for d:=0;d<MAX_DIMS;d++ {
			if min[d] < genMin[d] {
				genMin[d] = min[d]
			}
			if max[d] > genMax[d] {
				genMax[d] = max[d]
			}
                     }
		}
		info[i] = &LoopInfo{
			Parts:  parts,
			Path:   path,
			Inside: loops[i].Inside,
			Min: min,
                        Max: max,
			L:      loops[i].L,
			C:      loops[i].C,
			P:      loops[i].P,
		}
		if info[i].L <= 0 {
			info[i].L = 20
		}
		if info[i].C <= 0 {
			info[i].C = 570
		}
		if info[i].P <= 0 {
			info[i].P = 7600
		}
	}
	return
}

func MakeBodyMeshing(shape *BodyShape) (body *BodyComposition, err error) {
	loops, min, max, e := prepareLoopInfo(shape.Loops)
	if e != nil {
		return nil, e
	}
	body = &BodyComposition{
		Shape: shape,
		Mesh:  &BodyMesh{Min: min, Max: max},
		Loops: loops,
	}
	err = MakeMeshingInDirections(body)
	return
}
