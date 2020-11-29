test_examples: test_twitter test_computer_parts

test_twitter:
	(cd examples/twitter && make setup_db)
	(cd examples/twitter && make build)

test_computer_parts:
	(cd examples/computer_parts && make setup_db)
	(cd examples/computer_parts && make build)
