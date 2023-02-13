# back_togther

# API 
This code provides a set of API endpoints for retrieving user, forum, message forum, message and admin data. 

## Endpoints 
The following endpoints are available: 
* `/user` - Retrieves user data with the following input and output formats: 
	* Input Data: `{"login":""}` 
	* Output Data: `{"id":"","nom":"","emal":"","login":"","pwd":""}` 
* `/forum` - Retrieves forum data. 
	* Input Data: `{"id":""}` 
	* Output Data: `{"id":"","description":""}` 
* `/message_forum` - Retrieves message forum data. 
	* Input Data: `{"id":""}` 
	* Output Data: `{"id":"","titre":"","contenu":"","date":"","id_frm":"","id_user":""}` 
* `/message` - Retrieves message data.  
	* Input Data: `{"id":""}` 
	* Output Data: `{"id":"","obj_msg":"","contenu":""}` 
* `/admin` - Retrieves admin data.
	* Input Data: `{"id":""}` 
	* Output Data: `{"id":"","login":"","pwd":""}` 