API Documentation
Users
Get All Users
Mengambil daftar semua pengguna yang ada dalam database.

URL

/users

Method

GET

Response

Code: 200
Content:

json
Copy code
[
  {
    "ID": 1,
    "Name": "John Doe",
    "Email": "john.doe@example.com",
    "Age": 30
  },
  {
    "ID": 2,
    "Name": "Jane Smith",
    "Email": "jane.smith@example.com",
    "Age": 25
  }
]
Create User
Membuat pengguna baru dengan data yang diberikan.

URL

/users

Method

POST

Request Body

json
Copy code
{
  "Name": "John Doe",
  "Email": "john.doe@example.com",
  "Age": 30
}
Response

Code: 201
Content:

json
Copy code
{
  "ID": 3,
  "Name": "John Doe",
  "Email": "john.doe@example.com",
  "Age": 30
}
Update User
Memperbarui data pengguna berdasarkan ID yang diberikan.

URL

/users/:id

Method

PUT

URL Params

Required:

bash
Copy code
id=[integer]
Request Body

json
Copy code
{
  "Name": "John Doe (Updated)",
  "Email": "john.doe@example.com",
  "Age": 32
}
Response

Code: 200
Content:

json
Copy code
{
  "ID": 3,
  "Name": "John Doe (Updated)",
  "Email": "john.doe@example.com",
  "Age": 32
}
Delete User
Menghapus pengguna berdasarkan ID yang diberikan.

URL

/users/:id

Method

DELETE

URL Params

Required:

bash
Copy code
id=[integer]
Response

Code: 200
Content:

json
Copy code
{
  "message": "User deleted successfully"
}
Errors
User Not Found
Code: 404
Content:
json
Copy code
{
  "error": "User not found"
}
Internal Server Error
Code: 500
Content:
json
Copy code
{
  "error": "Internal Server Error"
}
Catatan:

Pastikan Anda telah menyesuaikan data yang ada dalam contoh dengan data aktual dari pengguna.
Selain operasi CRUD (Create, Read, Update, Delete), Anda juga dapat menambahkan endpoint lain yang sesuai dengan kebutuhan aplikasi Anda.
