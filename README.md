#  NewsFeed
NewsFeed is a API using Publish/Subscribe Model

# Running the API

go run *.go

# Testing 
You can either use 

CURL:
curl -H "Accept:application/json" http://localhost:8000/users

or
RestSoftware - https://insomnia.rest


# API Methods
Register Subscriber
 ``` POST localhost:8000/register 
  {
   "id": "3",
   "username": "User",
	}
 ```


Get Subscribers 
 ``` "/users" - Method("GET") ```

Delete Subscriber 
 ``` "/users/{username}" Method("DELETE") ```

Get User Details 
``` "/subscribe/{username}" Method("GET") ```

User Subscribe to Topic 
``` "/subscribe/{username}/{topic}" Methods("GET") ```

Get Topic News 
 ``` "/subscribed/{username}/topics" Method("GET") ```

Publish News and Topic 
 ``` "/publish/{topic}/{news}" Methods("POST") ```

