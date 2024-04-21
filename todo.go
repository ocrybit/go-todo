package main

import ( "fmt"; "bufio"; "os"; "strings" )

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter command: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	words := strings.Split(strings.TrimSpace(input), ",")
	cmd := words[0]
	
	fmt.Println("Command:", cmd)
}
