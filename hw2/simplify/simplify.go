//INCOMPLETE
package simplify

import "hw2/expr"
import "fmt"

var count int
// Depth should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	//fmt.Println("Top:")
	//fmt.Println(e)

	switch i:= e.(type){
	case expr.Var:
		temp,cow := convert(i,env)
		//check if float # is 0/nil. I ran into issues converting env variables b/c it would turn into 0 after
		if(cow <= 0){
			return e
		}else{
			return temp
		}
	case expr.Literal:
		temp,_ := convert(i,env)
		return temp
	case expr.Unary:
		i.X = Simplify(i.X,env)
		//Evaluate Unary, convert to float64 then convert to string type then to expr type 
		temp,cow := convert(i,env)

		check := fmt.Sprintf("%T",temp.(expr.Unary).X)
		if(check == "expr.Literal" && cow < 0){
			return -temp.(expr.Unary).X.(expr.Literal)
		}
		//check if float # is 0/nil
		if(cow <= 0){
			return e
		}else{
			return temp
		}
		break
	case expr.Binary: 
		//fmt.Printf("%T,%T\n",i.X,i.Y)
		//fmt.Println(i.X , i.Y)
		count = count + 1
		i.X = Simplify(i.X,env)
		count = count - 1
		//checks if multiplied by 1
		if(zoo(i.Y,env) && i.Op != 45 && i.Op != 43){
			return i.X
		}
		if(zoo(i.X,env) && i.Op != 45 && i.Op != 43){
			return i.Y
		}
		//Evaluate Binary, convert to float64 then convert to string type then to expr type 
		temp,_ := convert(i,env)
		


		if(count == 0){
			check := fmt.Sprintf("%T,%T",i.X,i.Y)
			if(check == "expr.Literal,expr.Literal"){
				return temp
			}
			return i
		}else{
			return temp
		}

		break
	default: 
		panic("unrecognized character")
	}
	return e
	panic("TODO: implement this!")
}

func zoo(rhs expr.Expr, env expr.Env) bool{
	if(rhs.Eval(env) == 1){
		return true
	}
	return false
}
func convert(i expr.Expr, env expr.Env) (expr.Expr, float64){
		cow:= i.Eval(env)
		pig:= fmt.Sprintf("%f",cow)
		temp,_:=expr.Parse(pig)
		return temp,cow
}