# jwt-go
Repo ini adalah sebuah API dengan menggunakan GO, GORM dan JWT. Terdapat dua buah level user yaitu admin dan user.  
`ADMIN` dapat melakukan CRUD sedangkan `USER` hanya bisa melakukan `create` dan `read` product yang telah dibuatnya .  

Untuk menjalankan ketik `go run main.go` diterminal
# ubah config database

Jangan lupa untuk menyesuaikan config database dengan mengubah data di bagian  
pkg/database/db.go  

# body request postman
Untuk body request, kita akan menggunakan menggunakan form. Terdapat 2 buah endpoint yaitu `/users/` dan `/products/`.  
  
pertama kita harus melakukan register di `/users/register` dengan body sebagai berikut:  
*note: `level` yang kita gunakan hanya `user` dan `admin`  
  
![alt text](https://i.ibb.co/f2GGSxf/Screenshot-64.png)  
  
Selanjutnya kita melakukan login di `/users/login` untuk mendapatkan `token` dengan body sebagai berikut:  
  
 ![alt text](https://i.ibb.co/HKFzxQv/Screenshot-3.png)  
  
Setelah itu kita dapat melakukan insert produk baru di  `/products` dengan body sebagai berikut:  
   
   ![alt text](https://i.ibb.co/vcVTKJL/Screenshot-1.png)  
   
Jangan lupa untuk mengatur `token` yang sudah kita dapatkan di bagian `Authorization` seperti pada gambar dibawah:  
  
   ![alt text](https://i.ibb.co/XS38fbM/Screenshot-2.png)  
  
  
 # test postman
Berikut ini adalah HTTP method beserta route-nya  
POST `localhost:8080/users/register` //register user  
POST `localhost:8080/users/login` //login user   
  
POST `localhost:8080/products/` //create product  
GET `localhost:8080/products/:id` //read product by id  
PUT `localhost:8080/products/:id` //update product by id  
DELETE `localhost:8080/products/:id` //delete product by id  
 
