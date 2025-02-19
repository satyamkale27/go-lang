package main

func intSeq() func() int { //  This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
	i := 0

	return func() int { // this is closure
		i++ //The closure (the returned function) can still access i, even though intSeq() has already finished.
		return i
	}
}

/*
Additional note:-
the returned function (the closure) doesn't have to execute when intSeq() finishes executing. Let me explain:

intSeq() finishes executing before the closure is ever called.So, intSeq() itself has already completed its execution and returned the closure.
The closure is a separate function that doesn't start running until you actually call it (with nextInt()).
The closure remembers i, so every time it's called, it can access and modify i even though intSeq() is already done executing.


 nextInt := intSeq()
fmt.Println(nextInt())  now the closure is called in main func and it executes


*/
