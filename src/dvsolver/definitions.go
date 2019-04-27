package dvsolver;

type Point struct {
        X     int
        Y     int
}

type Node struct {
 	N []int	
} 


type Body struct {
	Nodes []Node
        Points []Point
}
