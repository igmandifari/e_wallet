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

Full Documentation :  https://github.com/igmandifari/ReadMeImage

1. Login
   - Login does not include email/password
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/EmptyEmailPasswordGo.png?raw=true)

   - Login invalid credentials / entered the wrong email or password
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/EmailFailedGo.png?raw=true)
     
   - Login successful
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/LoginSuccessGo.png?raw=true)
     
2. Logout
   - Logout Successfull
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/LogoutSuccessGo.png?raw=true)
     
     
4. Transfer
   - Insufficient / Less balance
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/LessBalance.png?raw=true)
     
   - Transfer to your own account
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/ReceiverNotFoundGo.png?raw=true)
     
   - Transfer to unregistered users
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/ReceiverNotFound.png?raw=true)

   - Transfer Without Login
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/TransferWithoutLogin.png?raw=true)
     
   - Successful transfer
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/TransferSuccess.png?raw=true)

     
     
   **Screenshot of the Database : PostgreSQL**
   - Transaction
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/Transactions.png?raw=true)

   - Account History
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/ActivityHistories.png?raw=true)
     
   - Account
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/ActivityHistories.png?raw=true)
     
   - Users
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/Accounts.png?raw=true)
     
   - Auths
     ![GitHub Logo](https://github.com/igmandifari/ReadMeImage/blob/main/Auths.png?raw=true)
     
