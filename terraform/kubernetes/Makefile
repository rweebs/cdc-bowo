.PHONY: run
init: 
	tfenv use min-required
	terraform init
	tflint --init
	pre-commit run -a
run:
	terraform apply -auto-approve