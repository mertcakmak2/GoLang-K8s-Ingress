docker build -t mertcakmak2/go-gke .

docker push mertcakmak2/go-gke 

kubectl apply -f deployment.yaml

kubectl apply -f ingress.yaml