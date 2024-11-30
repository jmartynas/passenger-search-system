run:
	go build
	env $$(cat environment) ./main
