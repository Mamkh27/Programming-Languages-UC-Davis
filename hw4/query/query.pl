memberlist(L1,L2) :- member(A,L1),member(A,L2).
memberlist(A,L1,L2):-member(A,L1),member(A,L2).
year_1953_1996_novels(Book) :-
    novel(Book,1953);
    novel(Book,1996).

period_1800_1900_novels(Book) :-
    novel(Book,Z),Z=<1900,Z>=1800.

lotr_fans(Fan) :-
   fan(Fan,Z), member(the_lord_of_the_rings,Z).

author_names(Author) :-
    author(Author,Z),fan(chandler,X),memberlist(X,Z).

fans_names(Fan) :-
    fan(Fan,Z),author(brandon_sanderson,X),memberlist(X,Z).

mutual_novels(Book) :-
    novel(Book,_),fan(phoebe,X),fan(ross,Y),fan(monica,Z),(memberlist(Book,X,Y);memberlist(Book,Y,Z);memberlist(Book,X,Z)).