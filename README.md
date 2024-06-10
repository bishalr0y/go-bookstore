# Go Bookstore📚
---
A simple GoLang REST API with a structured folder structure. It makes use of Gin, a Golang-written HTTP web framework. To interact with the Postgres database docker container, it utilizes the use of GORM as the ORM.

## Project Structure🩻
---
![image](https://github.com/bishalr0y/go-bookstore/assets/56751927/487c3a6b-8c0d-401f-ba8e-dafe1ed0c320)

## How to run the project🪄
### Using Docker Compose 🐳
- Open the terminal and fire the command below
```
git clone <the URL of the GitHub repo>
```
- ``cd`` into the project directory and use docker-compose to run the application at port `9091`
```
cd go-bookstore
docker compose up -d
```
### Using Kubernetes ☸️
- Open the terminal and fire the command below
```
git clone <the URL of the GitHub repo>
```
- ``cd`` into the project directory, build the image of the application and push it to your dockerhub repo
```
git clone <the URL of the GitHub repo>
cd go-bookstore
docker build . -t <dockerhub_username>/go-bookstore:latest
docker push <dockerhub_username>/go-bookstore:latest
```
- After that, you need to create the secret to deploy your dockerhub credentials
```
kubectl create secret docker-registry dockerhub-cred \
--docker-username=DOCKER_USERNAME
--docker-password=DOCKER_PASSWORD
```
- Then, just deploy all of the Kubernetes manifest that is present in the ``kubernetes`` directory
```
kubectl deply -f kubernetes/
```
- If you are using ``minikube``, then to access the application locally in your browser use the ``kubectl`` command below
```
minikube service bookstore-service start
```

## API Endpoints👨‍💻
---
![image](https://github.com/bishalr0y/go-bookstore/assets/56751927/a3670e8e-b3fd-4883-9d1f-9282ae3ebe7d)

