---
layout: documentation
title: 'Mobile Main Street · Digital Ocean Install Guide'
css_assets:
  - "/css/docs.css"
---

**1. Create a Droplet**

> ![Screenshot showing create droplet button](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-1.png)
> Screenshot showing create droplet button

The first thing you'll need to do is click "Create Droplet" to create your server.

**2. Droplet Hostname**

> ![Screenshot showing a droplet's hostname and size](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-2.png) 
> Screenshot showing a droplet's hostname and size

You'll need to provide your droplet with a name. This is just for your personal use. The naming scheme we use is "mms-community-name-v1.0". This way we can tell at a glance that this is a Mobile Main Street server for "Community Name" that's running version "1.0".

**3. Select Size**

You'll need to select a size for your server. For most users the $5/month option should be more than sufficient. You can always go back and increase the size later. 

**4. Select Region**

> ![Screenshot showing a droplet's region and available settings.](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-3.png)
> Screenshot showing a droplet's region and available settings.

Your server will be placed in the region you select. Visitors closer to this region will have your community's Mobile Main Street load *slightly* quicker. A good rule of thumb is to select the region you or your community is closest to. 

**5. Available Settings**

These are advanced settings, typically they can all be left unchecked. The only one worth consideration is "Enable Backups". If this is selected, Digital Ocean will keep regular backups of your entire server that you can revert to for an additional 20% per month.

**6. Select Image**

> ![Screenshot showing an image selected to be installed on a droplet.](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-4.png)
> Screenshot showing an image selected to be installed on a droplet.

Go to the applications tab and select "Docker 1.x.x on Ubuntu". It's ok if the version number is different from the screenshot.

**7. Click Create Droplet and Wait a Few Seconds**

Once you click create, Digital Ocean will start building your server. 

**8. Check Your Email**

After your server is online, you will receive an email with your root password. This password is used to login to your server, so treat it like any other password and keep it private. 

**9. Login to Your Server**

> ![Screenshot showing a droplet's IP address.](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-5.png)
> Screenshot showing a droplet's IP address.

In order to login to your server, you'll need the password from the email you just received, and the IP address from the top of your Droplet page (highlighted in the screenshot above).

**10. Open the terminal** 

Open your operating system's terminal application. 
* **Mac** - On Mac OSX the application is called "Terminal". 
* **Windows** - On Windows this application is the Command Prompt. You will need a program called SSH in order to connect to your server. You can go to [OpenSSH.org](http://www.openssh.com/) to install this application.

**11. Connect to the Server**

> ![Screenshot showing a prompt to confirm connecting to a droplet.](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-6.png)
> Screenshot showing a prompt to confirm connecting to a droplet.

Run the following command, replacing X.X.X.X with the IP address of your Droplet.

```
ssh root@X.X.X.X
```

You may see  a screen that looks like the screenshot above. Confirm you want connect by typing in `yes`.

**12. Change Your Password**

> ![Screenshot showing a prompt to change password.](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/install-do-7.png)
> Screenshot showing a prompt to change password.

After confirming your connection, you will have to change your password from the password in your email. When it asks for your `(current) UNIX password` enter the password that Digital Ocean emailed you. If you forget the new password you enter, you can always reset it by going back to the Access tab on the Droplet page.

**13. Follow General Install Instructions**

You have successfully set up your server! To install Mobile Main Street on your server, visit the [General Install Instructions](../General-Install-Guide) page and follow the steps. 