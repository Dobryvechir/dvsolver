package dvthermo

import (
	"errors"
	"strconv"
)

var pathTokens map[string]int = map[string]int{
	"m": 2,
	"M": 2,
	"l": 2,
	"L": 2,
	"h": 1,
	"H": 1,
	"v": 1,
	"V": 1,
	"z": 0,
	"Z": 0,
}

func ReadPathToken(path string, pos int) (c byte, x []float64, npos int, err error) {
	n := len(path)
	c = ' '
	for pos < n && path[pos] <= 32 {
		pos++
	}
	npos = pos
	if npos >= n {
		return
	}
	s := path[npos : npos+1]
	c = path[npos]
	npos++
	m, ok := pathTokens[s]
	if !ok {
		err = errors.New("Unexpected character " + s)
		return
	}
	for npos < n && path[npos] <= 32 {
		npos++
	}
	if npos >= n {
		if m != 0 {
			err = errors.New("Expected numbers after " + s)
		}
		return
	}
	if _, ok = pathTokens[path[npos:npos+1]]; ok {
		if m != 0 {
			err = errors.New("Expected numbers after " + s)
		}
		return
	}
	d := path[npos]
	if !(d >= '0' && d <= '9' || d == '+' || d == '-' || d == '.') {
		err = errors.New("Unexpected character " + path[npos:npos+1])
		return
	}
	if m == 0 {
		err = errors.New(s + " must not be followed by numbers")
		return
	}
	x = make([]float64, m)
	for i := 0; i < m; i++ {
		pos = npos
		for npos < n {
			d = path[npos]
			if d >= '0' && d <= '9' || d == '+' || d == '-' || d == '.' || d == 'e' {
				npos++
			}
		}
		x[i] = path[pos:npos]
		for npos < n && path[npos] <= 32 {
			npos++
		}
		if npos >= n {
			if i+1 != m {
				err = errors.New(s + " requires " + strconv.Itoa(m) + " numbers but end of line is met")
			}
			return
		}
		d = path[npos]
		if d >= '0' && d <= '9' || d == '+' || d == '-' || d == '.' {
			if i+1 < m {
				continue
			}
			err = errors.New(s + " requires only " + m + " numbers")
			return
		}
		if _, ok = pathTokens[path[npos:npos+1]]; ok {
			err = errors.New(s + " requires " + strconv.Itoa(m) + " numbers")
		} else {
			err = errors.New("Unexpected character " + path[npos:npos+1])
		}
		break
	}
	return
}

func ConvertPathToCurveGeoPaths(parts []ConditionPart) ([]*CurveGeoPath, string, error) {
	path := ""
	fullPath := ""
	partNo := -1
	partAmount := len(parts)
	x := 0
	y := 0
	px := x
	py := y
	pos := 0
	condition := 0
	var condValue float64
	var currentGeoPath *CurveGeoPath
	var u []float64
	var c byte = '0'
	r := make([]*CurveGeoPath, 0, 7)
CurveLoop:
	for partNo <= partAmount {
		if path == "" {
			partNo++
			pos = 0
			if partNo == partAmount {
				if currentGeoPath == nil {
					break
				}
				path = "z"
				fullPath += path
			} else {
				path = strings.TrimSpace(parts[partNo].Description)
				if path == "" {
					continue
				}
				fullPath += path
				condValue = parts[partNo].Boundary
				switch strings.TrimSpace(parts[partNo].Condition) {
				case "NOFLOW":
					condition = POINT_NO_FLOW
				case "CONSTANT":
					condition = POINT_CONSTANT
				case "ALPHA":
					condition = POINT_HTC
				case "MID":
					condition = POINT_MID
				default:
					return nil, "", errors.New("Condition may be only NOFLOW, CONSTANT, or ALPHA")
				}
			}
		}
		kind := -1
		subKind = 0
		c, u, pos, err = ReadPathToken(path, pos)
		if err != nil {
			return nil, "", err
		}
		switch c {
		case 'l':
			if u[0] != 0 {
				subKind = LINE_H
				if u[1] != 0 {
					subKind = LINE_HV
				}
			} else if u[1] != 0 {
				subKind = LINE_V
			}
			if subKind != 0 {
				x += u[0]
				y += u[1]
				kind = CURVE_LINE
			}
		case 'L':
			if u[0] != x {
				subKind = LINE_H
				if u[1] != y {
					subKind = LINE_HV
				}
			} else if u[1] != y {
				subKind = LINE_V
			}
			if subKind != 0 {
				x = u[0]
				y = u[1]
				kind = CURVE_LINE
			}
		case 'h':
			if u[0] != 0 {
				x += u[0]
				subKind = LINE_H
				kind = CURVE_LINE
			}
		case 'H':
			if u[0] != x {
				x = u[0]
				subKind = LINE_H
				kind = CURVE_LINE
			}
		case 'v':
			if u[0] != 0 {
				y += u[0]
				subKind = LINE_V
				kind = CURVE_LINE
			}
		case 'V':
			if u[0] != y {
				y = u[0]
				subKind = LINE_V
				kind = CURVE_LINE
			}
		case 'm':
			kind = CURVE_MOVE
			x += u[0]
			y += u[1]
			px = x
			py = y
		case 'M':
			kind = CURVE_MOVE
			x = u[0]
			y = u[1]
			px = x
			py = y
		case 'z', 'Z':
			if x != px {
				subKind = LINE_H
				if y != py {
					subKind = LINE_HV
				}
			} else if y != py {
				subKind = LINE_V
			}
			if subKind == 0 {
				currentGeoPath = nil
			} else {
				kind = CURVE_LINE
				c = 'z'
				x = px
				y = py
			}
		default:
			path = ""
			continue CurveLoop
		}
		switch kind {
		case CURVE_MOVE:
			if currentGeoPath == nil {
				continue CurveLoop
			} else {
				return nil, "", errors.New("It is expected to place z before m (move)")
			}
		case CURVE_LINE:
			if currentGeoPath == nil {
				currentGeoPath := &CurveGeoPath{
					Points: make([]*CurveGeoPoint, 0, 4),
				}
				CurveGeoPath.Min[0] = x
				CurveGeoPath.Min[1] = y
				CurveGeoPath.Max[0] = x
				CurveGeoPath.Max[1] = y
				r = append(r, currentGeoPath)
			}
			if x > currentGeoPath.Max[0] {
				currentGeoPath.Max[0] = x
			}
			if y > currentGeoPath.Max[1] {
				currentGeoPath.Max[1] = y
			}
			if x < currentGeoPath.Min[0] {
				currentGeoPath.Min[0] = x
			}
			if y < currentGeoPath.Min[1] {
				currentGeoPath.Min[1] = y
			}
			curveGeoPoint := &CurveGeoPoint{
				Kind:           kind,
				SubKind:        subKind,
				Condition:      condition,
				ConditionValue: condValue,
			}
			curveGeoPoint.X[0] = x
			curveGeoPoint.X[1] = y
			currentGeoPath.Points = append(currentGeoPath.Points, curveGeoPoint)
			if c == 'z' {
				currentGeoPath = nil
			}
		default:
			continue CurveLoop
		}
	}
	return r, fullPath, nil
}
