package main

func truncate(name string, newLength uint64) {
	client, err := getClient("")
	if err != nil {
		fatal(err)
	}

	err = client.Truncate(name, newLength)
	if err != nil {
		fatal(err)
	}
}
