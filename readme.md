# Triển khai Database và Web Services trên Render  

## I. Deploy Database  
1. Đăng nhập [Render](https://render.com/) bằng GitHub.  
2. Vào **Dashboard** → **New** → **PostgreSQL**.  
3. Đặt tên cho database.  
4. Lưu lại các thông tin sau:  
   - Hostname  
   - Port  
   - Database  
   - Username  
   - Password  

## II. Deploy Web Services  
1. Trong `main.go`, thay đoạn mã:  
   ```go
   err := godotenv.Load()
   if err != nil {
       log.Fatal("Error loading .env file")
   }
bằng
if err := godotenv.Load(); err != nil {
    log.Println(".env file not found — using environment variables from Render")
}
2. Quay lại Dashboard → New → Web Services.

3. Environment: Chọn Docker (nhớ chọn gói miễn phí để tránh mất phí).

4. Do Render không sử dụng trực tiếp file .env, nên cần thêm biến môi trường thủ công:

	Copy nội dung file .env và thay các thông tin kết nối database bằng những giá trị đã deploy.