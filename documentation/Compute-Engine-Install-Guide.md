---
layout: documentation
title: 'Mobile Main Street · Compute Engine Install Guide'
css_assets:
  - "/css/docs.css"
---


# Compute Engine Install Guide
1. Visit https://console.developers.google.com and login.

![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-2.png)
2. Click on "Create Project" and fill out information.

![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-3.png)
3. Select the project that was created.

![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-5.png)

4. Click on "Billing & Settings" and click "Enable Billing".  Fill out all the requested information. 
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-6.png)

5. After getting back to the project overview, select "Storage" and "Cloud SQL".  Press "Create Instance".
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-10.png)

6. Give the instance a name and press "Save".
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-11.png)

7. Select the created instance and click on "Access Control".  Press "Request an IP Address".  Also enter in a password in the "Set Root Password" field.  Both of these values will be used later.  
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-14.png)

8. In the sidebar select "Compute" then "Compute Engine" then "VM Instances".  Press "Create an Instance". 
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-15.png)

9. Fill out the new instance form.  Check allow http, allow https, or both.  Select the region that would be closest to the internet traffic it will receive. For image select coreos-stable-***.  For external IP select "New static IP address...".  Press "Create".   
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-16.png)

10. After the instance has been built, from the SSH menu select "Open in Browser Window".  
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-17.png)

11. Follow the general instructions for setting up a machine.  They can be found in [General Install Guide](../General-Install-Guide)
![](https://raw.githubusercontent.com/SyntropyDev/mms-web/docs/images/ce-install-18.png)