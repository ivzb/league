package main

const configPath = "config.json"

func main() {
	app, err := newApp(configPath)

	if err != nil {
		panic(err)
	}

	if err := app.run(); err != nil {
		panic(err)
	}
}
