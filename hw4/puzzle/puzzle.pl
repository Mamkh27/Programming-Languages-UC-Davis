%states which side of the river the obj is on
state(_,_,_,_).

%returns true or false if A,B are different sides of the river
opposite(left,right).
opposite(right,left).

move(State1,State2,M):-
	safe(State1),
	A = arc(_,State1,State2),
	A,
	\+ member(A,M),
	!.

move(State1,State2,M):-
	safe(State1),
	A1 = arc(take(_,_,_),State1,StateCheck),
	\+ member(A1,M),
	A1,
	move(StateCheck,State2,[A1|M]).

%Arc checks if move to X to Y is safe

arc(take(wolf, F1, F2), state(F1, F1, G, C), state(F2, F2, G, C)) :-
	take(wolf, F1, F2),
	safe(state(F2, F2, G, C)).

arc(take(goat, F1, F2), state(F1, W, F1, C), state(F2, W, F2, C)) :-
	take(goat, F1, F2),
	safe(state(F2, W, F2, C)).

arc(take(cabbage, F1, F2), state(F1, W, G, F1), state(F2, W, G, F2)) :-
	take(cabbage, F1, F2),
	safe(state(F2, W, G, F2)).

arc(take(farmer,F1,F2),state(F1,W,G,C),state(F2,W,G,C)):-
	take(farmer,F1,F2),
	safe(state(F2,W,G,C)).

safe(state(F,_,G,_)) :- F = G,!. 
%If both Farmer and goat are on the same side of river then SAFE

safe(state(_,W,G,C)):- opposite(W,G),opposite(G,C).
%checks if wolf/goat and goat/cabbage are on the opposite sides of the river , then SAFE

%not provable safe(A)
unsafe(A):- \+ safe(A).

%takes obj to other side of the river.
take(_,A,B):- opposite(A,B).

solve(F1, W1, G1, C1, F2, W2, G2, C2) :- 
	move(state(F1,W1,G1,C1),state(F2,W2,G2,C2),[]),!.
	
