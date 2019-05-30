package nfa
// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, act rune) []state

//var sheep int = 1
// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	c := make(chan bool)

	go ReachNfa(transitions,start,final,input,c)

	ans := <- c

	return ans
}

func ReachNfa(transitions TransitionFunction,start,final state,input []rune,c chan bool) bool{
	if (len(input) == 0){
		if(start == final){
			c <- true
			return true
		}else{
			return false
		}
	}
	for _, next := range transitions(start,input[0]){
		if ReachNfa(transitions,next,final,input[1:],c){
			return true
		}
	}
	c <- false
	return false
}