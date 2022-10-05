# GoLang-K8s-Ingress
## Deployment GoLang app on Google Cloud Kubernetes.

docker build -t mertcakmak2/go-gke .

docker push mertcakmak2/go-gke 

kubectl apply -f deployment.yaml

kubectl apply -f ingress.yaml


![Screenshot_7](https://user-images.githubusercontent.com/21373505/194007931-8f4d762b-6b2d-4317-b54c-f201d65d1c11.png)


![Screenshot_5](https://user-images.githubusercontent.com/21373505/194007925-615c3a96-cbc8-4348-ab52-a5fb56890079.png)


![Screenshot_4](https://user-images.githubusercontent.com/21373505/194007920-7db50be9-35af-4c99-8fa1-d62c31e60555.png)
