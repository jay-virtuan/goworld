package operators

func All(val1,val2 int) (int,int,int,float32) {

	  // use := creating new variables and assign
      addv,subv,mulv,divev := all(val2,val1)

      return addv,subv,mulv,divev
}


func all(val1, val2 int) (add,sub,mul int,dive float32)  {

	// use = only when not need to create new variables
	add = Add(val2,val1)
    sub = Sub(val2,val1)
    mul = Mal(val2,val1)
    dive = Devide(val1,val2)

    return // implicitly return above variables since it's declared in the function signature
}