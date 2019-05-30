reachable(StartState,FinalState,[]):- StartState == FinalState.
reachable(StartState, FinalState, [A|B]) :- transition(StartState,A,T),member(U,T),reachable(U,FinalState,B).
