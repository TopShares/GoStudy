package main

func main()  {
	var s [] int
	for i := 0; i<10;i ++{
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}