test tests t:
	go test -count=100000 -race

fast-test fast-tests fast-t:
	go test -count=1 -race