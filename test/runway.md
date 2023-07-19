Berikut merupakan runway untuk test blue green deployment dengan skema change:

1. Deploy app versi blue

`kubectl apply -k test/example/app/deploy/blue/.`

2. Deploy service

`kubectl apply -f test/example/app/deploy/app-service.yaml`

3. Deploy k6

`kubectl apply -k test/example/app/deploy/k6/.`

4. Exec ke k6 dan jalanin scriptnya

`kubectl exec -it k6 sh`

5. Jalanin k6 test nya 

`k6 run /config/script.js --vus 1 --duration 10m --rps 10 --no-connection-reuse`

6. Build aplikasinya

`make build`

7. Jalanin aplikasinya

`./main run -c=config.test-app.json`

8. Mulai sync data

`./main start change-ddl`

Deploy versi green

`kubectl apply -k test/example/app/deploy/green/.`

Kalo dah mulai blue green 

`bash run.sh`
