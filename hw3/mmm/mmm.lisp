(defun minim (xs)
	"recursive function to calculate minimum value of a list of numbers"
	(cond
		((null xs) nil)
		((null (cdr xs)) (car xs))
		((<(car xs)(car(cdr xs)))
			(minim (cons(car xs)
				(cdr(cdr xs)))))
		(t(minim (cdr xs)))))

(defun mean(xs)
	"recursive funtion to calculate the mean value of a list of numbers"
	(setq q(apply '+ xs))
	(setq w(countnum xs 1))
	(/ q w))
(defun countnum(xs counter)
	"count the number of #'s in a list"
	(cond
		((null xs) nil)
		((null (cdr xs)) counter)
		(t(countnum(cdr xs)(+ counter 1)))))
(defun maxim(xs) 
	"recursive function to calculate maximum value of a list of numbers"
	(cond
		((null xs) nil)
		((null(cdr xs))(car xs))
		((>(car xs)(car(cdr xs)))
			(maxim (cons(car xs)
				(cdr(cdr xs)))))
		(t(maxim(cdr xs)))))

(defun min-mean-max (xs)
    (setq a(minim xs))
    (setq b(mean xs))
    (setq c(maxim xs))
    (cons a(cons b(cons c nil))))

