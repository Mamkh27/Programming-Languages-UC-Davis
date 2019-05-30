(defun reachable (transition start final input)
	"Return true if the NFA accepts the input and returns false if it cannot be reachable"
    (setq counter (countnum input 1))
    (cond
    	((listp start) (setq next (funcall transition (car start) (car input))))
    	(t (setq next (funcall transition start (car input)))))

    (cond
    	((and (eql counter nil)(eql start final)) t)
    	((eql counter nil)nil)
    	((or 
    		(reachable transition (car next) final (cdr input))
    		(reachable transition (cdr next) final (cdr input))) t)
    	(t nil)))
(defun countnum(xs counter)
	"count the number of items in a list"
	(cond
		((null xs) nil)
		((null (cdr xs)) counter)
		(t(countnum(cdr xs)(+ counter 1)))))
