package depth

import "hw2/expr"

// Depth returns the maximum number of AST nodes between the
// root of the tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	//fmt.Printf("%f\n",e)
	// TODO: implement this function
	switch i:= e.(type){
	case expr.Var:

	case expr.Literal:

	case expr.Unary:
		return Depth(i.X) + 1
		break
	case expr.Binary: 
		return Max(Depth(i.X),Depth(i.Y)) + 1
		break
	default: 
		panic("unrecognized character")
	}
	return 1
}
func Max( x , y uint) uint{
	if x < y{
		return y
	}else {
		return x
	}
}