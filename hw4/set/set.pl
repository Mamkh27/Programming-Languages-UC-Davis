isUnion([Head|Tail], Set2, Union) :- 
	member(Head,Set2),
	!,
	isUnion(Tail,Set2,Union).
isUnion([Head|Tail], Set2, [Head|Union]) :-
	isUnion(Tail,Set2,Union).
isUnion([],Union1,Union1).

isIntersection([],Head,[]).
isIntersection([Head|Tail], Set2, [Head|Intersection]) :-
	member(Head,Set2),
	!,
	isIntersection(Tail,Set2,Intersection).
isIntersection([Head|Tail], Set2, Intersection) :- 
	isIntersection(Tail,Set2,Intersection).

isEqual([],_).
isEqual([Head|Tail], Set2) :-
	member(Head,Set2),
	isEqual(Tail,Set2).

powerSet([], []).
powerSet([_|Tail], P) :- 
	powerSet(Tail,P).
powerSet([Head|Tail], [Head|P]) :- 
	powerSet(Tail,P).

