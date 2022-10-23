# avitoTechInternship
## Description
Микросервис для работы с балансом пользователей

## Installation
First step clone the git repository
```bash
git clone https://github.com/rro1699/avitoTechInternship.git
```

After that go to the folder avitoTechInternship
```bash
cd avitoTechInternship
```

After that run the following command
```bash
docker-compose up 
```

After the containers are deployed, the microservice is ready to user.

## Request/response examples

The orders and users tables are initially empty. The servs table looks like this:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197414561-4e165e12-1531-4c5d-b3ea-426e8971309d.png)<br> *Pic 1. Table servs.* <br><br>

### The method of accruing funds to the balance has the following url: http://localhost:10000/user/accrual <br><br>
#### Request:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197414771-598df33c-9d98-4707-aede-649585477bfa.png)<br> *Pic 2. Accrual request.* <br><br>
#### Response:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197414814-8ff22909-e478-4929-bc84-fad1ed3d5432.png)<br> *Pic 3. Accrual response.* <br><br>

### The method of reserving funds from the main balance in a separate account has the following url: http://localhost:10000/user/reserv<br><br>
#### Request:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415218-6438ca75-bcfa-4d2d-9ad0-2cd15a60b77a.png)<br> *Pic 4. Reserving request.* <br><br>
#### Response:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415258-45b514d3-b48d-4849-ba24-eccb5ad4749c.png)<br> *Pic 5. Reserving successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415290-82c9978f-4970-4c6c-a9b9-f42ce4bcbd2a.png)<br> *Pic 6. Reserving badly response.* <br><br>

### Revenue recognition method has the following url: http://localhost:10000/user/recogn <br><br>
#### Request:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415409-9efc8c0f-b004-426e-8baa-975a0b7cc164.png)<br> *Pic 7. Recognition request.* <br><br>
#### Response:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415443-71f37eb1-d8dc-4543-9cf7-78bd40e0ce3d.png)<br> *Pic 8. Recognition successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415476-80bc0d6c-b5de-4fb2-b3e2-d9b8626481cb.png)<br> *Pic 9. Recognition badly response 1.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415494-3c02c3d5-1085-4124-b6ed-7bd1fb0b1d6c.png)<br> *Pic 10. Recognition badly response 2.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415785-2b941117-f0fb-4dc8-8bba-76a8a92138ce.png)<br> *Pic 11. Recognition badly response 3.* <br><br>


### User balance receipt method has the following url: http://localhost:10000/user/balance <br><br>
#### Request:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197414865-8150de09-6084-490b-ab54-fd69714ed5c1.png)<br> *Pic 12. Get balance request.* <br><br>
#### Response:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415549-a92f2754-be74-4a76-a379-e5e8f89274ad.png)<br> *Pic 13. Get balance successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415575-118cc5be-680b-49b6-8204-43d6d9166dab.png)<br> *Pic 14. Reserving badly response.* <br><br>

### Method to get monthly report has the following url: http://localhost:10000/user/report <br><br>
#### Request:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415645-a67c6e2a-0386-49a7-a185-c699b1f99744.png)<br> *Pic 15. Get report request.* <br><br>
#### Response:<br><br>
![image](https://user-images.githubusercontent.com/79422421/197415670-66e2d1b1-a75e-490a-8dd6-c9cc7fa8c789.png)<br> *Pic 16. Get report successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415695-74dfdb5e-d0eb-4c9d-8dbb-ee7999ab8c30.png)<br> *Pic 17. Get report badly response 1.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415716-7b385af7-9da1-4d99-8283-9ffe368e157c.png)<br> *Pic 18. Get report badly response 2.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415730-070253bf-8ef7-4982-a4ce-88e4701153d3.png)<br> *Pic 19. Get report badly response 3.* <br><br>

#### Example report
![image](https://user-images.githubusercontent.com/79422421/197416253-865c8d94-04fd-489f-af61-04239f0c5ed9.png)<br> *Pic 20. Report example.* <br><br>










