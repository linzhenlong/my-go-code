package sort


// 冒泡排序
func BubbleSort(arr []int,order string) []int {
	for i:=0;i<len(arr)-1;i++ {
		for j:=0;j<len(arr)-1-i;j++ {
			if order == "asc" {
				if arr[j] > arr[j+1] {
					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			} else {
				if arr[j] < arr[j+1] {
					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}
		}
	}
	return arr
}


