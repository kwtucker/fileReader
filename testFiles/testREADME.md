# App-Store

##### Requirements
* NodeJS /w npm
* MySQL
* Gulp - (globally)
* nodemon - (globally)
* mocha - (globally)
* istanbul - (globally)                   
## Getting Started

(:1:)
##### Installing Dependencies
```
npm install
```

##### Sample `.env` file

```
DB_NAME=apiCRUD
DB_USER=
DB_PASS=
DB_HOST=localhost
DB_SCHEMA=mysql
DB_PORT=3306
PORT = 3000
NODE_DEBUG=DEV
```

##### Start Server Without Debug
```
gulp startup
```
or
```
nodemon src/server.js
```

##### Start Server Dev Debug
```
DEBUG=true gulp startup
```
or
```
DEBUG=true nodemon src/server.js
```
***

## Debug Usage
``
const util = require('path/to/lib/util');

// How to use util
util.debug(title, obj, status);
``
## Test Usage - In Terminal
If you don't already install mocha globally.
* This will give you command line access to mocha.
```
npm install -g mocha
```
Then from project root run:
```
mocha
```
If you want to create a coverage report run Istanbul
* After istanbul is install Globally
```
istanbul cover _mocha
```
 * Now open the index.html in: /coverage/lcov-report/src/index.html   
***      

## Endpoints and response

### CRUD for Users
* [Create a User](#post-create-user)
* [Display User based upon id](#user-with-id-of-1)
* [Display all Users](#get-all-users)
* [Update User based upon id](#update-user)
* [Delete User based upon id](#delete-user)

### CRUD for Apps
* [Create an App](#post-create-app)
* [Display App based upon id](#get-app-with-id-of-1)
* [Display all App from specific userId](#get-app-by-user-id)
* [Display all Apps](#get-all-apps)
* [Update User based upon id](#post-update-app)
* [Update App UserId](#post-update-app-userid)

### ArtAssets to Apps
* [Create artassets](#post-create-artassets)
* [Display all artassets upon app id](#get-all-artassets-upon-app-id)
* [Display one artassets from specific app id and artassetid](#get-display-one-artassets-from-specific-app-id-and-artassetid)
* [Update artassets based upon id](#post-update-artassets-based-upon-id)


(:2:)
#### CRUD for Users
--=-=-=-=-=-=-=-=-

##### POST Create User
```
/api/users/
```
###### Sample
```
{
  "name": "Kevin",
}
```

##### GET User with Id of 1
```
/api/users/1
```
###### Sample Response
```
{
 "id": 1,
 "name": "Kevin Tucker",
 "createdAt": "2016-06-08T01:11:52.000Z",
 "updatedAt": "2016-06-08T01:11:52.000Z"
}
```


##### GET All Users
```
/api/users/
```
###### Sample Response
```
[
 {
  "id": 1,
  "name": "Kevin",
  "createdAt": "2016-06-08T01:35:43.000Z",
  "updatedAt": "2016-06-08T01:35:43.000Z"
 },
 {
  "id": 2,
  "name": "Kevin",
  "createdAt": "2016-06-08T01:35:47.000Z",
  "updatedAt": "2016-06-08T01:35:47.000Z"
 },
 {
  "id": 3,
  "name": "Kevin",
  "createdAt": "2016-06-08T01:35:48.000Z",
  "updatedAt": "2016-06-08T01:35:48.000Z"
 }
]
```      

##### POST Update User
```
/api/users/1
```
###### Sample
```
{
 "name": "Kevin T"
}
```      

##### DELETE User
```
/api/users/1
```
###### Response
```
1
```

(:3:)


#### CRUD for Apps
-=-=-=-=-=-=-=-=-

##### POST Create App
```
/api/apps
```
###### Sample
```
{
  "title": "Best routes Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z"
}
```


##### GET App with Id of 1
```
/api/apps/1
```
###### Sample Response
```
{
 "id": 1,
 "title": "Best routes Ever",
 "description": "A fast paced side scrolling shooter",
 "releaseDate": "2016-06-15T22:29:20.000Z",
 "createdAt": "2016-06-08T01:12:36.000Z",
 "updatedAt": "2016-06-08T01:12:36.000Z",
 "userId": 1,
 "artAssets": [
  {
   "id": 1,
   "title": "Best router Ever",
   "srcLink": "Kevin",
   "createdAt": "2016-06-08T01:16:11.000Z",
   "updatedAt": "2016-06-08T01:16:11.000Z",
   "appId": 1
  }
 ]
}
```


##### Get App by user id
```
/api/users/1/apps
```
###### Sample Response
```
[
 {
  "id": 1,
  "title": "Best routes Ever Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z",
  "createdAt": "2016-06-08T01:35:54.000Z",
  "updatedAt": "2016-06-08T01:44:12.000Z",
  "userId": 1,
  "artAssets": [
   {
    "id": 1,
    "title": "Best router Ever",
    "srcLink": "Kevin",
    "createdAt": "2016-06-08T01:36:16.000Z",
    "updatedAt": "2016-06-08T01:36:16.000Z",
    "appId": 1
   },
   {
    "id": 2,
    "title": "Best router Ever",
    "srcLink": "Kevin",
    "createdAt": "2016-06-08T01:36:20.000Z",
    "updatedAt": "2016-06-08T01:36:20.000Z",
    "appId": 1
   },
   {
    "id": 3,
    "title": "Best router Ever",
    "srcLink": "Kevin",
    "createdAt": "2016-06-08T01:36:25.000Z",
    "updatedAt": "2016-06-08T01:36:25.000Z",
    "appId": 1
   }
  ]
 }
]
```



##### GET All Apps
```
/api/apps/
```
###### Sample Response
```
[
 {
  "id": 1,
  "title": "Best routessdfsdfsdfr Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z",
  "createdAt": "2016-06-08T01:35:54.000Z",
  "updatedAt": "2016-06-08T01:35:54.000Z",
  "userId": null
 },
 {
  "id": 2,
  "title": "Best routessdfsdfsdfr Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z",
  "createdAt": "2016-06-08T01:35:55.000Z",
  "updatedAt": "2016-06-08T01:35:55.000Z",
  "userId": null
 }
]
```


##### POST Update App
```
/api/apps/1
```
###### Sample
```
{
  "title": "Best routes Ever Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z",
  "userId": 1
}
```


##### POST Update App UserId
```
/api/apps/1
```
###### Sample
```
{
  "title": "Best routes Ever Ever",
  "description": "A fast paced side scrolling shooter",
  "releaseDate": "2016-06-15T22:29:20.000Z",
  "userId": 2
}
```



### ArtAssets to Apps

##### POST Create artassets
```
/api/apps/1/artassets/
```
###### Sample
```
{
  "title": "Best router Ever",
  "srcLink": "Kevin",
  "appId": 1
}
```


##### GET All artassets upon app id
```
/api/apps/1/artassets/
```
###### Sample
```
[
 {
  "id": 1,
  "title": "Best router Ever",
  "srcLink": "Kevin",
  "createdAt": "2016-06-08T01:36:16.000Z",
  "updatedAt": "2016-06-08T02:33:50.000Z",
  "appId": 1
 },
 {
  "id": 2,
  "title": "Best router Ever",
  "srcLink": "Kevin",
  "createdAt": "2016-06-08T01:36:20.000Z",
  "updatedAt": "2016-06-08T01:36:20.000Z",
  "appId": 1
 }
]
```


##### GET Display one artassets from specific app id and artassetid
```
/api/apps/1/artassets/2
```
###### Sample Response
```
{
 "id": 2,
 "title": "Best router Ever",
 "srcLink": "Kevin",
 "createdAt": "2016-06-08T01:36:20.000Z",
 "updatedAt": "2016-06-08T01:36:      Z",
 "appId": 1
}
```

##### POST Update artassets based upon id
```
/api/apps/1/artassets/2
```
###### Sample
```
{
 "id": 2,
 "title": "Best router Ever",
 "srcLink": "Kevin",
 "createdAt": "2016-06-08T01:36:20.000Z",
 "updatedAt": "2016-06-08T01:36:20.000Z",
 "appId": 2
}
```
(:4:)
