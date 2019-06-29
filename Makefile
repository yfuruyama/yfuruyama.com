run:
	go run main.go

deploy:
	gcloud app --project=${PROJECT_ID} deploy
