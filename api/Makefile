GCP_APP_NAME := api
GCP_PROJECT_NAME := activities-298512
TAG := eu.gcr.io/$(GCP_PROJECT_NAME)/api
REGION := europe-west1

.PHONY: build
build:
	gcloud builds submit --tag $(TAG)

.PHONY: deploy
deploy:
	gcloud run deploy $(GCP_APP_NAME) --image $(TAG) --project $(GCP_PROJECT_NAME) --region $(REGION) --allow-unauthenticated
