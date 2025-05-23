# CSRF

Cross-site request forgery (CSRF) is a security vulnerability that allows an attacker to trick a user into performing an action on their behalf.

## How does CSRF work?

For a CSRF attack to be possible, three key conditions must be in place:

- **A relevant action.** There is an action within the application that the attacker has a reason to induce. This might be a privileged action (such as modifying permissions for other users) or any action on user-specific data (such as changing the user's own password).
- **Cookie-based session handling.** Performing the action involves issuing one or more HTTP requests, and the application relies solely on session cookies to identify the user who has made the requests. There is no other mechanism in place for tracking sessions or validating user requests.
- **No unpredictable request parameters.** The requests that perform the action do not contain any parameters whose values the attacker cannot determine or guess. For example, when causing a user to change their password, the function is not vulnerable if an attacker needs to know the value of the existing password.

For example, if for changing the email address of a user, the HTTP request is like this:
```http
POST /email/change HTTP/1.1
Host: vulnerable-website.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 30
Cookie: session=yvthwsztyeQkAPzeQ5gHgTvlyxHfsAfE

email=wiener@normal-user.com
```

The three key conditions are met:
- The action is to change the email address of a user (then change passwords, etc.).
- The application relies on cookies for session handling (no tokens).
- The request does not contain any parameters whose values the attacker can not guess.

To know the session cookie, the attacker can build a web page in which the user enter their email and send the form with the session cookie (assuming `SameSite=None`). It will be redirected to the vulnerable web page, but the attacker will see the session cookie.

> Although CSRF is normally described in relation to cookie-based session handling, it also arises in other contexts where the application automatically adds some user credentials to requests, such as HTTP Basic authentication and certificate-based authentication.


## How to construct a CSRF attack?

Manually creating the HTML needed for a CSRF exploit can be cumbersome, particularly where the desired request contains a large number of parameters, or there are other quirks in the request. The easiest way to construct a CSRF exploit is using the CSRF PoC generator. There are several generators online which receive the HTTP request and generate the necessary HTML.

Lab: we try to change our email address, capture the HTTP request and generate the HTML with the CSRF PoC generator.

TODO: I don't know why but it doesn't work.

The deliver mechanisms for CSRF attacks are essentially the same ass for reflected XSS. The previous lab we redirect the victim to an attacker web server. However, it can be done placing a image or comment in the vulnerable web page. For example:
```html
<img src="https://vulnerable-website.com/email/change?email=pwned@evil-user.net">
```
This can be done in simple CSRF in which the whole GET request can be fully self-contained with a single URL.

## Defense against CSRF

The most common defenses you'll encounter are as follows:

- **CSRF tokens** - A CSRF token is a unique, secret, and unpredictable value that is generated by the server-side application and shared with the client. When attempting to perform a sensitive action, such as submitting a form, the client must include the correct CSRF token in the request. This makes it very difficult for an attacker to construct a valid request on behalf of the victim.
- **SameSite cookies** - SameSite is a browser security mechanism that determines when a website's cookies are included in requests originating from other websites. As requests to perform sensitive actions typically require an authenticated session cookie, the appropriate SameSite restrictions may prevent an attacker from triggering these actions cross-site. Since 2021, Chrome enforces `Lax` SameSite restrictions by default which implies that the cookie will only be sent with requests to the same origin.
- **Referer-based validation** - Some applications make use of the HTTP Referer header to attempt to defend against CSRF attacks, normally by verifying that the request originated from the application's own domain. This is generally less effective than CSRF token validation.

A common way to share CSRF tokens with the client is to include them as a hidden parameter in an HTML form, for example:
```html
<form name="change-email-form" action="/my-account/change-email" method="POST">
    <label>Email</label>
    <input required type="email" name="email" value="example@normal-website.com">
    <input required type="hidden" name="csrf" value="50FaWgdOhi9M9wyna8taR1k3ODOR8d6u">
    <button class='button' type='submit'> Update email </button>
</form>
```
Submitting this form results in the following request:
```http
POST /my-account/change-email HTTP/1.1
Host: normal-website.com
Content-Length: 70
Content-Type: application/x-www-form-urlencoded

csrf=50FaWgdOhi9M9wyna8taR1k3ODOR8d6u&email=example@normal-website.com
```

## Common flaws in CSRF token validation

Some applications correctly validate the token when the request uses the POST method but skip the validation when the GET method is used. In this situation, the attacker can switch to the GET method to bypass the validation and deliver a CSRF attack.

> CSRF tokens don't have to be sent as hidden parameters in a POST request. Some applications place CSRF tokens in HTTP headers, for example. The way in which tokens are transmitted has a significant impact on the security of a mechanism as a whole.

Some applications correctly validate the token when it is present but skip the validation if the token is omitted. In this situation, the attacker can remove the entire parameter containing the token (not just its value) to bypass the validation and deliver a CSRF attack.

Some applications do not validate that the token belongs to the same session as the user who is making the request. Instead, the application maintains a global pool of tokens that it has issued and accepts any token that appears in this pool. In this situation, the attacker can log in to the application using their own account, obtain a valid token, and then feed that token to the victim user in their CSRF attack. However, this token may be only valid for a limited period of time or one-use. So the attacker could inspect the form before sending it and maybe in the HTML code the token is not hidden from the user.



### Summary

- Change method (e.g. use GET instead of POST)
- Remove token (the case of not sending the token is not covered)
- Reuse token (server only validate a pool of tokens, not ownership)
  
## SameSite Bypass

## Referer-based Bypass