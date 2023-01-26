## 1. Register User
```
curl --location --request POST 'localhost:8080/user/register' \
--header 'Content-Type: text/plain' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InVzZXIxMjMiLCJpc3MiOiJkYXZpZGFsZXhhbmRlci5jb20iLCJleHAiOjE2NzQ3NTA3NTJ9.xAxBKutA8ugVhkmxEI_xXH2ILw9C6GL_tiurE53CkiM; username=user123' \
--data-raw '{
    "name": "david alexander",
    "username": "user123",
    "password": "user123"
}'
```

## 2. Login
```
curl --location --request POST 'localhost:8080/user/login' \
--header 'Content-Type: text/plain' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InVzZXIxMjMiLCJpc3MiOiJkYXZpZGFsZXhhbmRlci5jb20iLCJleHAiOjE2NzQ3NTA5NDB9.n0BRJsb18Bay6EEOCon6hFE4gcT6ibAfiqe3969AvSg; username=user123' \
--data-raw '{
    "username": "user123",
    "password": "user123"
}'
```

# 3. Upload Image
```
curl --location --request POST 'http://localhost:8080/file/upload' \
--header 'Authorization: Bearer test' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InVzZXIxMjMiLCJpc3MiOiJkYXZpZGFsZXhhbmRlci5jb20iLCJleHAiOjE2NzQ3NTA5NDB9.n0BRJsb18Bay6EEOCon6hFE4gcT6ibAfiqe3969AvSg; username=user123' \
--form 'file=@"/C:/Users/ASUS VIVO BOOK/OneDrive/Pictures/david.jpg"'
```

# 4. Get All Files
```
curl --location --request GET 'http://localhost:8080/file/' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InVzZXIxMjMiLCJpc3MiOiJkYXZpZGFsZXhhbmRlci5jb20iLCJleHAiOjE2NzQ3NTA1OTd9.Hoopu-kDVvVVMUj_zlsN46j-rfcwrq4ni0WWaSvJWbY; username=user123'
```

# 5. Download File
```
curl --location --request GET 'http://localhost:8080/ex/file/download' \
--header 'Content-Type: text/plain' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InVzZXIxMjMiLCJpc3MiOiJkYXZpZGFsZXhhbmRlci5jb20iLCJleHAiOjE2NzQ3NTEyODJ9.wqv8V3hoZTpfXgRauJ491HB3HFKSCp9SXH9tQgsrVoU; username=user123' \
--data-raw '{
    "path": "IMAGE-8cc2d13835fa4ec7800fb74904ca4d51"
}'
```