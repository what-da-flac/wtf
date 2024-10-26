.PHONY: bootstrap diff deploy destroy docker-build golang-imports test test-html-output

# PROFILE is aws cli profile
PROFILE := "wtf"

bootstrap:
	cdk bootstrap --profile $(PROFILE) --region us-east-2

docker-build:
	docker build -t wtf-go-builder:$(GIT_TAG) --platform linux/amd64 -f dockerfiles/wtf-go-builder.Dockerfile .

diff:
	cdk --profile $(PROFILE) --region us-east-2 diff --all

deploy:
	cdk --profile $(PROFILE) --region us-east-2 deploy --all --require-approval never

destroy:
	cdk --profile $(PROFILE) --region us-east-2 --all destroy

golang-imports:
	goimports -w .

test:
	go test -count 1 -cover ./...

test-html-output:
	go test -coverprofile=c.out ./... && go tool cover -html=c.out && rm -f c.out