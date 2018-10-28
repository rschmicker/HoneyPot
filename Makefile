all:
	glide update
	glide install
	go install .