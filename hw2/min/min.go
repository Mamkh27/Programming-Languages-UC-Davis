package min

import "fmt"
// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	// TODO: implement this function.
	var low int = 0
	for i:=0;i<len(arr);i++{
		if(i == 0){
			low = arr[i]
		}else if(arr[i] < low){
			low = arr[i]
		}
	}
	fmt.Printf("%d\n",low)
	return low
}
