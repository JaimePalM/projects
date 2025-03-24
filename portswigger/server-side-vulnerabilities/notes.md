---
title: Server-Side Vulnerabilities
description: Portswigger Server-Side Vulnerabilities
---

# Access Control

## Horizontal Privilege Escalation

Horizontal privilege escalation attacks may use similar types of exploit methods to vertical privilege escalation. For example, a user might access their own account page using the following URL:

`https://insecure-website.com/myaccount?id=123`

If an attacker modifies the id parameter value to that of another user, they might gain access to another user's account page, and the associated data and functions.

> Note: This is an example of an insecure direct object reference (**IDOR**) vulnerability. This type of vulnerability arises where user-controller parameter values are used to access resources or functions directly.

# Authentication

## Brute-forcing usernames
During auditing, check whether the website discloses potential usernames publicly. For example, are you able to access user profiles without logging in? Even if the actual content of the profiles is hidden, the name used in the profile is sometimes the same as the login username. You should also check HTTP responses to see if any email addresses are disclosed. Occasionally, responses contain email addresses of high-privileged users, such as administrators or IT support.

## Username enumeration
Username enumeration is when an attacker is able to observe changes in the website's behavior in order to identify whether a given username is valid.

Username enumeration typically occurs either on the login page, for example, when you enter a valid username but an incorrect password, or on registration forms when you enter a username that is already taken. This greatly reduces the time and effort required to brute-force a login because the attacker is able to quickly generate a shortlist of valid usernames.

## Bypassing two-factor authentication
At times, the implementation of two-factor authentication is flawed to the point where it can be bypassed entirely.

If the user is first prompted to enter a password, and then prompted to enter a verification code on a separate page, the user is effectively in a "logged in" state before they have entered the verification code. In this case, it is worth testing to see if you can directly skip to "logged-in only" pages after completing the first authentication step. Occasionally, you will find that a website doesn't actually check whether or not you completed the second step before loading the page.

# Server-Side Request Forgery

## What is SSRF?
Server-side request forgery is a web security vulnerability that allows an attacker to cause the server-side application to make requests to an unintended location.

In a typical SSRF attack, the attacker might cause the server to make a connection to internal-only services within the organization's infrastructure. In other cases, they may be able to force the server to connect to arbitrary external systems. This could leak sensitive data, such as authorization credentials.

## SSRF Attacks

In an SSRF attack against the server, the attacker causes the application to make an HTTP request back to the server that is hosting the application, via its loopback network interface. This typically involves supplying a URL with a hostname like 127.0.0.1 (a reserved IP address that points to the loopback adapter) or localhost (a commonly used name for the same adapter).

For example, imagine a shopping application that lets the user view whether an item is in stock in a particular store. To provide the stock information, the application must query various back-end REST APIs. It does this by passing the URL to the relevant back-end API endpoint via a front-end HTTP request. 

In some cases, the application server is able to interact with back-end systems that are not directly reachable by users. These systems often have non-routable private IP addresses. The back-end systems are normally protected by the network topology, so they often have a weaker security posture. In many cases, internal back-end systems contain sensitive functionality that can be accessed without authentication by anyone who is able to interact with the systems.

In the previous example, imagine there is an administrative interface at the back-end URL `https://192.168.0.68/admin`.

### Scanning through API URL

If the website uses an API for several purposes, it is possible to scan through the API URLs to see if any of the IP hosts are reachable. For example, if the API is reachable at `https://192.168.0.1:8080/checkStock?id=2`, then the website might be able to access the back-end API at some host at `https://192.168.0.X:8080/admin`. Scan by brute-forcing the IP addresses and once you discover the back-end API, you can then send admin API requests to the back-end API.

## File Upload Vulnerability

File upload vulnerabilities are when a web server allows users to upload files to its filesystem without sufficiently validating things like their name, type, contents, or size. A basic image upload function can be used to upload arbitrary and potentially dangerous files instead. This could even include server-side script files that enable remote code execution.

### Introduction

From a security perspective, the worst possible scenario is when a website allows you to upload server-side scripts, such as PHP, Java, or Python files, and is also configured to execute them as code. This makes it trivial to create your own web shell (command through HTTP requests) on the server.

For example, the following PHP one-line file could be used to read arbitrary files from the server's filesystem:
```php
<?php echo file_get_contents('/path/to/target/file'); ?>
```

Once uploaded, sending a request for this malicious file will return the target file's contents in the response.

A more versatile web shell may look something like this:
```php
<?php echo system($_GET['command']); ?>
```
This script enables you to pass an arbitrary system command via a query parameter as follows:
```http
GET /example/exploit.php?command=id HTTP/1.1
```

### Bypass Validation

There are two ways to send information to the server:
- `application/x-www-form-urlencoded`: the information is sent in the request body as name=value pairs.
- `multipart/form-data`: the information is sent in the request body as a multipart/form-data object. The data is splitted into fields and it can be a binary file (image, audio, video, etc.) or text.


Example of `multipart/form-data`:

```http
POST /images HTTP/1.1
    Host: normal-website.com
    Content-Length: 12345
    Content-Type: multipart/form-data; boundary=---------------------------012345678901234567890123456

    ---------------------------012345678901234567890123456
    Content-Disposition: form-data; name="image"; filename="example.jpg"
    Content-Type: image/jpeg

    [...binary content of example.jpg...]

    ---------------------------012345678901234567890123456
    Content-Disposition: form-data; name="description"

    This is an interesting description of my image.

    ---------------------------012345678901234567890123456
    Content-Disposition: form-data; name="username"

    wiener
    ---------------------------012345678901234567890123456--
```

One way that websites may attempt to validate file uploads is to check that this input-specific `Content-Type` header matches an expected MIME type. If the server is only expecting image files, for example, it may only allow types like `image/jpeg` and `image/png`. Problems can arise when the value of this header is implicitly trusted by the server. If no further validation is performed to check whether the contents of the file actually match the supposed MIME type, this defense can be easily bypassed.

Web pages can be configured to only allow image files to be uploaded. However, sometimes this validation is done on the client-side. In this case, the client-side validation is bypassed intercepting the request from the server and changing the `Content-Type` header to something acceptable, such as `image/jpeg` or `image/png`.

## OS Command Injection

It allows an attacker to execute operating system (OS) commands on the server that is running an application, and typically fully compromise the application and its data. Often, an attacker can leverage an OS command injection vulnerability to compromise other parts of the hosting infrastructure, and exploit trust relationships to pivot the attack to other systems within the organization.

In other words, the server passes client parameters to a server shell, and the attacker can inject a command as parameter into a web page, and the server executes the command as the user running the web server.

Some useful commands are:
| Purpose of command    | Linux     | Windows |
| ---                   | ---       | --- |
| Name of current user  | whoami    | whoami
| Operating system      | uname -a  | ver
| Network configuration | ifconfig  | ipconfig /all
| Network connections   | netstat -an | netstat -an
| Running processes     | ps -ef    |  asklist

In this example, a shopping application lets the user view whether an item is in stock in a particular store. This information is accessed via a URL:
```url
https://insecure-website.com/stockStatus?productID=381&storeID=29
```
To provide the stock information, the application must query various legacy systems. For historical reasons, the functionality is implemented by calling out to a shell command with the product and store IDs as arguments: `stockreport.pl 381 29`. This command outputs the stock status for the specified item, which is returned to the user.

The application implements no defenses against OS command injection, so an attacker can submit the following input to execute an arbitrary command: `& echo aiwefwlguh &`. If this input is submitted in the `productID` parameter, the command executed by the application is:
`stockreport.pl & echo aiwefwlguh & 29`

Placing the additional command separator `&` after the injected command is useful because it separates the injected command from whatever follows the injection point. This reduces the chance that what follows will prevent the injected command from executing.

Regarding the payload to bypass the arguments, there are lots of ways to do it. For example, the "&", "|", ";", "<", ">", and "&&" characters can be used to separate the command and the arguments. In this repository it can be found more: http://github.com/payloadbox/command-injection-payload-list 

## SQL Injection (SQLi)

As the name suggests, SQL injection is a vulnerability that allows an attacker to inject SQL commands into a database. The attacker can use this vulnerability to read, modify, or delete data in the database. The idea is the same as the command injection vulnerability, but the SQL commands are used instead.

Typically:
- The single quote character `'` and look for errors or other anomalies.
- Some SQL-specific syntax that evaluates to the base (original) value of the entry point, and to a different value, and look for systematic differences in the application responses.
- Boolean conditions such as `OR 1=1` and `OR 1=2`, and look for differences in the application's responses.
- Payloads designed to trigger time delays when executed within a SQL query, and look for differences in the time taken to respond.
- OAST (Open Application Security Testing) payloads designed to trigger an out-of-band network interaction when executed within a SQL query, and monitor any resulting interactions

Injection example:
The web shop uses the URL `https://insecure-website.com/products?category=Gifts` and internally makes the following SQL query: `SELECT * FROM products WHERE category = 'Gifts' AND released = 1` because they don't want to show the existing products which aren't released. If we input as category `Gifts' --` the SQL query becomes `SELECT * FROM products WHERE category = 'Gifts' --' AND released = 1` and because of `--` is a SQL comment, this way bypasses the released filter.

> Warning
> 
> Take care when injecting the condition `OR 1=1` into a SQL query. Even if it appears to be harmless in the context you're injecting into, it's common for applications to use data from a single request in multiple different queries. If your condition reaches an `UPDATE` or `DELETE` statement, for example, it can result in an accidental loss of data.

One example to bypass the `WHERE` condition, modify the URL, after the value add `+OR+1=1--`. The `+` is decoded as a space, so the SQL query becomes:
```sql
SELECT * FROM products WHERE category = 'Gifts' OR 1=1--' AND released = 1
```
 
We also can use this to bypass logins and validate as any user. Where username field is `administrator'--`
```sql
SELECT * FROM users WHERE username = 'administrator'--' AND password = ''
```