package main

const configPath = "config.json"

func main() {
	cli, err := newCli(configPath)

	if err != nil {
		panic(err)
	}

	if err := cli.run(); err != nil {
		panic(err)
	}

	web, err := newWeb(configPath)

	if err != nil {
		panic(err)
	}

	if err := web.run(); err != nil {
		panic(err)
	}
}
