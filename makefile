test tests t:
	go test -count=1000000 -race

fast-test fast-tests fast-t:
	go test -count=1 -race