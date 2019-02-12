package Basic

import "fmt"

func main() {
	m := map[string]string{
		"name":    "Go",
		"quality": "notbad",
		"star":    "five",
	}
	m2 := make(map[string]int) // m == empty map
	var m3 map[string]int      // m2 == nil
	fmt.Println(m, m2, m3)
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	if fName, ok := m["fName"]; ok {
		fmt.Println(fName)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println("Deleting values")
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
