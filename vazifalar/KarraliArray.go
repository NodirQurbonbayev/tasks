package vazifalar
import "fmt"
func KarraliArray(arr []int, k int) []int {
	for i:=k-1; i<len(arr); i+=k {
        fmt.Println(arr[i])
    } 
	return arr

}