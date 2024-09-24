migrate-up:
	goose postgres "user=postgres dbname=lms_db password=.,.. sslmode=disable" up

migrate-down:
	goose postgres "user=postgres dbname=lms_db password=.,.. sslmode=disable" down

migrate-redo:
	goose postgres "user=postgres dbname=lms_db password=.,.. sslmode=disable" redo

migrate-status:
	goose postgres "user=postgres dbname=lms_db password=.,.. sslmode=disable" status
