# Go Errors

Like in any where in software development, if you have an operation that may return an error, YOU SHOULD ALWAYS RETURN AN ERROR.

This makes your code able to handle the possible outcome of an error

When handling errors in golang, make sure to handle it in the top most caller of the culprit

## Struct For APIs

Say you have an API, it is ideal to have your app in layers do that you can handle your errors more gracefully.

### HTTP Layer
This should be the top most layer of your app, so users hit this layer, basically the layer the client interacts with directly, and all errors should be pricipitated here for proper relating back to the client
### Service Layer
This layer should handle the business logic of the application and all its errors should be returned to the HTTP layer where it can be properly handled, wether it would be logged or related to the user in the most appropriate way

### Repository/Storage Layer
This is the communication layer where the backend talks with the db, this is where the app makes calls to the db, and db failure or similar errors can be returned to the caller, ie the service layer and the errors are then returned to the HTTP layer where it would be properly shaped for logging or reporting to the client.

## API Security
When handling error in a backend application, make sure to becareful the errors that are returned to the client, in applications that errors may not be too expensive, crash fast and fix would be a way to go. 
That would mean crash as soon as you hit an error and log the error for monitoring or report the error to the client.
Internal errors like db failures should not be logged to the client, rather they should get generic messages like 500 internal error, and NEVER return the stack trace of a call