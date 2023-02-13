# back_togther

# API 
This code provides a set of API endpoints for retrieving user, forum, message forum, message and admin data. 

## Endpoints 
The following endpoints are available: 
* `/user` - Retrieves user data with the following input and output formats: 
	* Input Data: `{"login":""}` 
	* Output Data: `{"id":"","nom":"","emal":"","login":"","pwd":""}` 
* `/forum` - Retrieves forum data. 
* `/message_forum` - Retrieves message forum data. 
* `/message` - Retrieves message data.  
* `/admin` - Retrieves admin data.