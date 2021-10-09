# Insta-API
 HTTP JSON API that can be used to perform basic tasks of Creating User, Creating Post, Getting User & Post details through their Unique ID. It is developed using Go-Lang and MongoDB as database

## Brief Overview of API
The API provides following features
 1.  **Create User**: 
  * A POST Request is used, data is when entered in JSON body gets added to MongoDB (which is used for storage). Also, the data is added at the ``` https://localhost:9000/users``` . <br>

   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/createUser.png)
  * The Passwords are stored in their SHA256  equivalent so that they can't be reverse engineered.
  * Output: 
   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/createUserOutput.png)
  * Every user is assigned with their Unique ID.

 2.  **Create Posts**:
 * The Posts are created using POST Method
 * The URL is ```http://localhost:9000/posts/``` 
 * The request body is in JSON format.
 * The TimeStamp is automatically updated using ```time.Now()``` function
 * CreatePosts Function:
   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/createPost.png)
 * Creating a post using Postman :
   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/createPostOutput.png)
   
 3.  **Get User Details by ID**:
 * The user details can be fetched from database by sending a GET Request
 * The url is ```https://localhost:9000/users/<UserID>```
 * The Code for Fetching User Details: 
 * the Get User Function: 
   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/getUser.png)
 * The Output of Fetching userdetails by its ID in Postman:
   ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/createPostOutput.png)
   
 4.  **Get Posts Details by ID**:
 * The Posts Details can be from MongoDB using the Unique ID of Posts
 * The GET request is to be sent along with id in URL
 * The URL for fetching posts details: ```http://localhost:9000/posts/<id>```
 * The Get Post Function:


 ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/getPost.png)

**NOTE** I was unable to create a function for fetching Posts according to User ID, so I created an extra function for deleting users.

 5.  **Delete User Information**:
 * This function deletes the User Records from the Database.
 * It uses DELETE to perform the operation
 * The ID of user is passed in URL for deletion: ```http://localhost:9000/user/<id>```
  ![](https://github.com/IamLucif3r/Insta-API/blob/main/docs/delete.png)
  
<hr>

This repository is maintained by : [IamLucif3r](https://www.linkedin.com/in/anmolsinghyadav/)
