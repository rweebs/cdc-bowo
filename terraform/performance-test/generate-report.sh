python3 -m pip install requirements.txt
POSTGRES_HOST=$(aws rds describe-db-instances --db-instance-identifier cdc-secondary --output text --region us-west-2 --query 'DBInstances[0].Endpoint.Address')
POSTGRES_HOST=$POSTGRES_HOST python3 -m generate-report.py