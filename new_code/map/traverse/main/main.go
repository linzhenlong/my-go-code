package main

import "fmt"

func main()  {
	
	var map1 = make(map[string]string)

	for i:=0;i<10;i++ {
		key := fmt.Sprintf("map_key_%d",i)
		map1[key] = fmt.Sprintf("map_value_%d",i)
	}

	// 遍历
	for key,val := range map1 {
		fmt.Printf("map key:%s,value:%s\n",key,val)
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$")

	map2 := make(map[string]map[string]string)

	for i:=0;i<10;i++ {
		key := fmt.Sprintf("map_key_%d",i)
		map2[key] = make(map[string]string)
		map2[key]["name"] = fmt.Sprintf("key-%d-name-%d",i,i)
		map2[key]["sex"] = fmt.Sprintf("key-%d-sex-%d",i,i)
	}

	for key1,val :=range map2{

		fmt.Printf("map,key=%s \n",key1)
		for key2,val2 := range val {
			fmt.Printf("map,key=%s,value=%s \t",key2,val2)
		}
		fmt.Println()

	}
}
