run:
	dev_appserver.py app.yaml

deploy:
	gcloud app --project=${PROJECT_ID} deploy
