**The scope of this application includes:**
1. Login
2. Logout
3. Transcation (Transfer between customers)
4. I Already add collections (postman) in this repo for testing API

**How to run app:**
1. add the .env file with the contents according to yours, im using postgresql
   - contents of the .env file :
      DB_URL=
      JWT_SECRET=
      DURATION_EXPIRE=
2. run migration : export DB_URL=postgresql://yourHost/go_ewallet?sslmode=disable && make migrateup
   
**Below is a screenshot with the steps and test case:**
1. Login
   - Login does not include email/password
   - Login invalid credentials / entered the wrong email or password
   - Login successful
     
3. Logout
   - Logout Successfull
     
5. Transfer
   - Insufficient / Less balance
   - Transfer to your own account
   - Transfer to unregistered users
   - Successful transfer
