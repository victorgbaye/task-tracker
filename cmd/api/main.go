package main

func main()  {
	api := &application{addr: ":8080"}

	api.run(api.mount())
}