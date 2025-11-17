.PHONY init-workspace:
.ONESHELL:
init-workspace:
	go work init
	go work use string-interpolation
	go work use array-and-slice