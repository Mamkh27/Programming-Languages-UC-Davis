
(defun matrix-add (m1 m2)
	"Takes two lists and does matrix addition"
   (mapcar ;;mapcar takes a function and lists then creates a new list of values by calling the function provided
   	(lambda (seg1 seg2)
   		(mapcar #'+ seg1 seg2))m1 m2))

(defun matrix-multiply (m1 m2)
	"takes 2 lists and does matrix multiplication"
	(mapcar
		(lambda (row)
			(apply #'mapcar
			(lambda (&rest column) ;;rest signals we are working with a list. remaining arguements in the variable name that follows it
				(apply #'+ (mapcar #'* row column))) m2)) m1))

(defun matrix-transpose (m)
	"Takes a matrix and rotates it"
    (apply #'mapcar #'list m) ;;list takes automatic &rest args
)
