---
title: API Testing
description: Port Swigger API Testing
---

## API Recon
To start API testing, you first need to find out as much information about the API as possible, to discover its attack surface.

To begin, you should identify API endpoints. These are locations where an API receives requests about a specific resource on its server. 

Once you have identified the endpoints, you need to determine how to interact with them. This enables you to construct valid HTTP requests to test the API. For example, you should find out information about the following:

- The input data the API processes, including both compulsory and optional parameters.
- The types of requests the API accepts, including supported HTTP methods and media formats.
- Rate limits and authentication mechanisms.

## API Documentation

API documentation is often publicly available, particularly if the API is intended for use by external developers. If this is the case, always start your recon by reviewing the documentation.

Even if API documentation isn't openly available, you may still be able to access it by browsing applications that use the API.

To do this, you can use Burp Scanner to crawl the API. You can also browse applications manually using Burp's browser. Look for endpoints that may refer to API documentation, for example:
```
/api
/swagger/index.html
/openapi.json
```
If you identify an endpoint for a resource, make sure to investigate the base path. For example, if you identify the resource endpoint /api/swagger/v1/users/123, then you should investigate the following paths:
```
/api/swagger/v1
/api/swagger
/api
```
You can also use a list of common paths to find documentation using Intruder.

Sometimes it's so simple like changing the HTTP method. For example, to get a user, you can use the following request:
```http
GET /api/user/wiener
```
And if it's easy maybe to delete a user, you can use the following request:
```http
DELETE /api/user/wiener
```

## Indentify API endpoints

While browsing the application, look for patterns that suggest API endpoints in the URL structure, such as /api/. Also look out for JavaScript files. These can contain references to API endpoints that you haven't triggered directly via the web browser. 

As you interact with the API endpoints, review error messages and other responses closely. Sometimes these include information that you can use to construct a valid HTTP request. As the previous example shows.

API endpoints often expect data in a specific format. They may therefore behave differently depending on the content type of the data provided in a request. Changing the content type may enable you to:
- Trigger errors that disclose useful information.
- Bypass flawed defenses.
- Take advantage of differences in processing logic. For example, an API may be secure when handling JSON data but susceptible to injection attacks when dealing with XML.

To change the content type, modify the `Content-Type` header, then reformat the request body accordingly. 

In summary, it's a trial and error process to find API endpoints. Moreover it's possible to use fuzzing to find API endpoints and identify the input data expected by the API.

## Mass Assignment Vulnerabilities

Mass assignment (also known as auto-blinding) can inadvertently create hidden parameters. It occurs when software frameworks automatically bind request parameters to fields on an internal object. 

Mass assignment vulnerabilities can be used to bypass authentication, CSRF protection, or other security measures. This is a common vulnerability in web applications and APIs. 

### Identifying hidden parameters

Since mass assignment creates parameters from object fields, you can often identify these hidden parameters by manually examining objects returned by the API.

For example, consider a `PATCH /api/users/` request, which enables users to update their username and email, and includes the following JSON:
```json
{
    "username": "wiener",
    "email": "wiener@example.com",
}
```
A concurrent `GET /api/users/123` request returns the following JSON:
```json
{
    "id": 123,
    "name": "John Doe",
    "email": "john@example.com",
    "isAdmin": "false"
}
```
This may indicate that the hidden `id` and `isAdmin` parameters are bound to the internal user object, alongside the updated username and email parameters.

Then you can modify the PATCH request to include the hidden parameters, with valid ones or invalid ones. If the application behaves differently, this may suggest that the invalid value impacts the query logic, but the valid value doesn't. This may indicate that the parameter can be successfully updated by the user. For example, sending with your personal user email update the `isAdmin` parameter to `true`.

In summary, try to send requests with several methods and observe carefully the request and responses body to identify hidden parameters. Then use these hidden parameters to trigger mass assignment vulnerabilities.

### Server-side Parameter Pollution (SSVP)

Some systems contain internal APIs that aren't directly accessible from the internet. Server-side parameter pollution occurs when a website embeds user input in a server-side request to an internal API without adequate encoding. This means that an attacker may be able to manipulate or inject parameters.

To test for server-side parameter pollution in the query string, place query syntax characters like `#`, `&`, and `=` in your input and observe how the application responds.

For example, in some application if you send the following request:
```http
GET /userSearch?name=peter&back=/home
```
The frontend will try to access with:
```http	
GET /users/search?name=peter&publicProfile=true
```
And if we comment everything in the request after the query string (`%23` is the `#` character in URL encoding):
```http	
GET /users/search?name=peter%23foo&publicProfile=true
```
The server can respond in two ways: return the information of the peter user or an error indicating that the user doesn't exist. The former is good because it indicates that the server is vulnerable to server-side parameter pollution. In other words, we could be able to comment the `publicProfile` parameter with the `#` character and bypass the validation.

We can use the `&` character to add a second parameter for example:
```http	
GET /users/search?name=peter%26foo=xyz&publicProfile=true
```
Review the response for clues about how the additional parameter is parsed. For example, if the response is unchanged this may indicate that the parameter was successfully injected but ignored by the application.

If we put all together, we can see that, in the example in "Identifying hidden parameters", we send the parameter with a public URL. However, here wee're talking about an non accessible API so we should use this technique to test server-side parameter pollution.

Finally, to confirm whether the application is vulnerable to server-side parameter pollution, you could try to override the original parameter. Do this by injecting a second parameter with the same name.

One last example:
```http
GET /users/search?name=peter&name=carlos&publicProfile=true
```
The internal API interprets two `name` parameters. The impact of this depends on how the application processes the second parameter. This varies across different web technologies. For example:
- PHP parses the last parameter only. This would result in a user search for `carlos`.
- ASP.NET combines both parameters. This would result in a user search for `peter,carlos`, which might result in an `Invalid username` error message.
- Node.js / express parses the first parameter only. This would result in a user search for `peter`, giving an unchanged result.

If you're able to override the original parameter, you may be able to conduct an exploit. For example, you could add `name=administrator` to the request. This may enable you to log in as the administrator user.

Lab: if we send a request with the `#` (`%23`) character and the server returns an error like `Field not specified` this means that the server has a parameter called `field` that is not specified in the request.

#### REST Paths

A REST path is a URI that is used to identify a resource on a web server. It is often used to identify a resource in a RESTful API. For example, the path `api/users/123` might be used to retrieve information about a user with ID 123. An attacker may be able to manipulate the path to trigger a vulnerability.

If we access the URL `GET /edit_profile.php?name=peter` and this trigger a server-side request `GET /api/private/users/peter`. We could inject in the name parameter something like `GET /edit_profile.php?name=peter%2f..%2fadmin` which triggers a server-side request `GET /api/private/users/peter/../admin`

#### Data Formats

An attacker could manipulate parameters which the server process for other structured data formats, such as a JSON or XML. This could be adding hidden parameters to the data structure, or changing the data structure itself. For example if we have a `POST` method which change the user name, we could try something like: 
```http
POST /myaccount
name=peter","access_level":"administrator
```
This, when the server formats it to a JSON structure will look like:
```json
{"name": "peter","access_level": "administrator"}
```

To prevent server-side parameter pollution, use an allowlist to define characters that don't need encoding, and make sure all other user input is encoded before it's included in a server-side request. You should also make sure that all input adheres to the expected format and structure.