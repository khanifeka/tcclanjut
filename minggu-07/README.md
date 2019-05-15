# Kubectl

Menggunakan Kubectl untuk membuat dan menjalankan Deployment, Pengendali Replikasi, dan mengeksposnya melalui Service tanpa menulis definisi yaml. Sehingga container dapat dijalankan dengan cepat ke dalam kluster.

1. Memulai Cluster

	- minikube start : perintah untuk memulai minikube.

	- kubectl get nodes : digunakan untuk mengechek apakah node telah siap




	

2. Menjalankan Kubectl




# Membangun Container menggunakan YAML

1. Membuat Deployment

   - Menambahkan definisi object deployment pada file deployment.yaml
     Apllikasi yang dijalankan bernana webapp1 menggunakan docker image katacoda/docker-http-server yang berjalan pada port 80.
   


   - kubectl create -f deployment.yaml : digunakan untuk mendeploy file deployment.yaml ke dalam cluster



   - kubectl get deployment : digunakan untuk menunjukkan list semua deployment object



   - kubectl describe deployment webapp1 : digunakan untuk menunjukkan detail dari deployment


   