package main

func main() {
	err := RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
