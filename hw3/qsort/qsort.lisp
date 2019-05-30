
(defun pivot (n xs)
  "Function to write the left and right of pivot"
  (cond 
  		((null xs) '(nil nil))
  		(t
  		(myappend2
  			(list<= n xs)
  			(list> n xs)))))

(defun quicksort (xs)
  	"Recursive function for doing a quick sort on a list"
  	(cond
	    ((null xs) nil)
	    (t
	    (myappend (myappend
	        (quicksort (list<= (car xs) (cdr xs)))
	        (cons (car xs) nil))
	        (quicksort (list> (car xs) (cdr xs)))))))

(defun list<= (a b)
	"Recursively sort numbers lesser than the pivot"
	(cond
		((or (null a) (null b)) nil)
		((<= a (car b)) (list<= a (cdr b)))
		(t (cons (car b) (list<= a (cdr b))))))

(defun list> (a b)
	"Recursively sort numbers greater than the pivot"
	(cond
		((or (null a) (null b)) nil)
		((> a (car b)) (list> a (cdr b)))
		(t (cons (car b) (list> a (cdr b))))))

(defun myappend(a b)
	"Append a b for quicksort"
	(cond 
		((null a) b)
	(t (cons (car a)(myappend(cdr a) b)))))

(defun myappend2(a b)
	"Append a b for pivot"
	(cons a (cons b nil)))