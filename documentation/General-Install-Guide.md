---
layout: documentation
title: 'Mobile Main Street · General Install Guide'
css_assets:
  - "/css/docs.css"
---

## Create Server
Create a server to host the Mobile Main Street application.  If you need help with this step please reference the [Digital Ocean](../Digital-Ocean-Install-Guide) or [Compute Engine](../Compute-Engine-Install-Guide) install guides.  Choose an option that comes with Docker preinstalled such as Ubuntu with Docker or CoreOS.  

## Create MySQL Instance
Create a MySQL instance for the server.  The instance can be local or located at another IP address.   Once created record the following information for use in the MySQL URL:
- IP Address (ip)
- User (user)
- Password (password)
- Database Name (databaseName)

MySQL URL Format: user:password@tcp(ip)/databaseName

## Create a Twitter App
Create a Twitter App to be able to use the Twitter API.  Visit https://apps.twitter.com/app/new login and fill out the form.
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/twitter-1.png)

Select the Application and select "Keys and Access Tokens".  Record the "Consumer Key" and "Consumer Secret".
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/twitter-2.png)

## Create a Facebook App
Create a Facebook App to be able to use the Facebook Graph API.  Visit https://developers.facebook.com/quickstarts/?platform=web login and select "Skip and Create App ID".  Then fill out the form shown.
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/facebook-1.png)

Select the created application and hit "show" next to the "App Secret".  Record the "App ID" and "App Secret".
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/facebook-2.png)

## Create a Mailgun Account
Visit https://mailgun.com/signup and fill out the requested information.  
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/mailgun-1.png)

After creating the account record the "API Key", "Public API Key", and the first of the "Mailgun Subdomains".  These are the same values that are blacked out in the screenshot.  
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/mailgun-2.png)

## Install on the Server
SSH into the server and run the following command:
```bash
docker pull loganjspears/mms-api
sudo docker run -p 80:8080 \
-e twitterApiKey={{twitterApiKey}} \
-e twitterApiSecret={{twitterApiSecret}} \
-e facebookApiID={{facebookApiID}} \
-e facebookAppSecret={{facebookAppSecret}} \
-e mailgunDomain={{mailgunDomain}} \
-e mailgunPublicApiKey={{mailgunPublicApiKey}} \
-e mailgunPrivateApiKey={{mailgunPrivateApiKey}} \
-e mysql={{mysqlURL}} \
-t loganjspears/mms-api
```

All values in double brackets "{{" "}}" represent the urls and api keys retrieved in the previous sections.  For the mysqlURL it should escape the parenthesis shown below: 
user:password@tcp\(ip\)/databaseName

This command pulls the mobile main street docker repo and runs it with the environmental variables given.  80:8080 represents the port mapping and to use SSL it would be changed to 443:8080.    