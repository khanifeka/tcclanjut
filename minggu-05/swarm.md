#Docker Swarm

Mode ini memungkinkan pengguna untuk me-deploy container pada multiple hosts atau node, menggunakan overlay network. Swarm mode merupakan bagian dari command line interface Docker yang memudahkan pengguna untuk memanage komponen container apabila sudah familiar dengan command – command yang ada di Docker.

Mode Docker Swarm memperkenalkan tiga konsep :

- **Node** : Node merupakan turunan dari Docker Engine yang terhubung ke Swarm. Node adalah manajer atau pekerja. Manajer menjadwalkan container mana yang akan dijalankan. Pekerja melaksanakan tugas. Secara default, Manajer juga adalah pekerja.

- **Services** : merupakan konsep tingkat tinggi yang berkaitan dengan kumpulan tugas yang akan dieksekusi oleh pekerja. Contoh Server HTTP yang berjalan sebagai Docker Container pada tiga node.

- **Load Balancing** : Load balancing digunakan untuk memproses permintaan di semua kontainer dalam Service.

1. Inisialisasi Mode Swarm

Untuk mengetahui perintah-perintah apa saja yang terdapat pada mode ini, gunakan perintah **docker swarm --help**

![](img/2-1.png)

**Pada node-1** jalankan perintah :

**docker swarm init**

![](img/2-1b.png)

Perintah init digunakan untuk menginisialisasikan node ini sebagai manager.

2. Join Cluster

Dengan mengaktifkan Mode Swarm, dimungkinkan untuk menambahkan node tambahan dan mengeluarkan perintah di semuanya. Jika node hilang, misalnya, karena crash, kontainer yang berjalan pada host tersebut akan secara otomatis dijadwal ulang ke node lain yang tersedia. 

Untuk menambahkan node2 kedalam cluster, dilakukan dengan mengarahkan host lain ke manajer cluster saat ini yaitu node-1.

**Pada node-2** jalankan perintah : 
*token=$(docker -H 172.17.0.76:2345 swarm join-token -q worker) && echo $token*

![](img/2-2a.png)

Perintah ini menanyakan kepada manajer apa token itu via swarm join-token.

*docker swarm join 172.17.0.76:2377 --token $join-token*

![](img/2-2b.png)

Perintah ini menginisialisasikan node-2 sebagai worker dan join ke cluster swarm.

*docker node ls* digunakan untuk menampilkan semua node yang ada dalam cluster.

3. Membuat Overlay Network

Overlay network akan berfungsi sebagai media komunikasi antar container yang berbeda host atau node. Untuk me-create overlay network, digunakan perintah :

**docker network create -d overlay skynet**

![](img/2-3.png)

4. Mendeploy Service

Secara default, Docker menggunakan model replikasi penyebaran untuk memutuskan kontainer mana yang harus dijalankan pada host mana. Pendekatan penyebaran memastikan bahwa kontainer digunakan di seluruh cluster secara merata. Ini berarti bahwa jika salah satu node dihapus dari cluster, instance akan sudah berjalan di node lain.

Service digunakan untuk menjalankan container lintas cluster. Dengan memperbarui service ini, Docker juga akan memperbarui kontainer yang diperlukan secara terkelola.

![](img/2-4a.png)

docker service create --name http --network skynet --replicas 2 -p 80:80 katacoda/docker-http-server

Perintah ini digunakan untuk mendefinisikan service dengan nama http yang dilampirkan dalam jaringan skynet dan menjalankan image katacoda / docker-http-server. Replikasi di buat 2 dan kedua container ini berjalan bersama-sama di port 80. 

Mengirim permintaan HTTP ke salah satu node di cluster akan memproses permintaan dengan salah satu kontainer di dalam cluster. Node yang menerima permintaan mungkin bukan node di mana kontainer merespons. Sebagai gantinya, Docker memuat permintaan keseimbangan di semua kontainer yang tersedia.

**docker service ls** digunakan untuk menampilkan service yang berjalan pada cluster.

**docker ps** digunakan untuk menampilkan list container. Gambar pertama adalah list container pada node-1 dan gambar ke 2 adalah list container pada node-2.

![](img/2-4b.png)

![](img/2-4c.png)

Jika terdapat permintaan HTTP ke port publik, permintaan itu akan diproses oleh kedua kontainer.

![](img/2-4d.png)


5. Inspect State

Untuk memeriksa service dan keadaan cluster serta aplikasi yang sudah berjalan, dapat dilakukan dengan perintah :

![](img/2-5a.png)

Untuk melihat detail dan configurasi service nya dengan perintah :

![](img/2-5b.png)

Pada setiap node, dapat diketahui apakah tugas atau task yang sedang dijalankan. SELF merujuk pada manajer node.

![](img/2-5c.png)

Dengan menggunakan ID sebuah node, dapat dilakukan query dari masing-masing host.

![](img/2-5d.png)

6. Scale Service

Swarm Mode memungkinkan untuk me-scale service sesuai dengan kebutuhan. Service dapat digunakan untuk mengukur berapa banyak tugas yang berjalan di seluruh cluster. Karena service ini memahami cara menjalankan kontainer dan kontainer mana yang sedang berjalan, ia dapat dengan mudah memulai, atau mengeluarkan, kontainer seperti yang diperlukan. 

![](img/2-6a.png)

Perintah **docker service scale http=5** ini akan meningkatkan skala layanan http untuk berjalan di lima kontainer.

![](img/2-6b.png)

Perintah **docker ps** digunakan untuk melihat node tambahan yang berjalan.


[**KEMBALI**](README.md)



