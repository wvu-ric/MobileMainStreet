---
layout: documentation
title: 'Mobile Main Street · Documentation'
css_assets:
  - "/css/docs.css"
---

# What is it?
Mobile Main Street is a web-based publishing tool built for sharing real-time information from multiple members within a community. Unlike other platforms, we publish directly from community member's Facebook pages, Twitter feeds, or XML feeds. This means there's no need for storytellers and community members to login to post updates – it's all done automatically.

<a name="how-it-works"></a> 

# How does it work?

Once you set up your Mobile Main Street instance you add in the community members and their Twitter or Facebook user names. Mobile Main Street pulls profile information from the community member's social profiles and then adds them to the category you specify. 

Mobile Main Street then checks your community member's social feeds every few minutes for updates and automatically publishes these updates to the appropriate category on your page. 

As the community leader you have the ability to add and remove people from your community, mute individual posts, and open and close registration. Most importantly, however, the community leader can also invite community members to create a Mobile Main Street account so they can mute any of their own posts and update their own profile information. 

# Install Guide
Mobile Main Street, for the time being, is self-hosted. This means it needs to be installed on a server in order to work. Don't worry if you've never done this before, it's easier than it sounds. 

The first step is to set up the hosting environment.  If you are unsure how to set a hosting environment then follow the [Digital Ocean](Digital-Ocean-Install-Guide) or [Compute Engine](Compute-Engine-Install-Guide) installation guides. Once the hosting environment is set up, follow the [General Install Guide](General-Install-Guide).

* [Installing on Digital Ocean](Digital-Ocean-Install-Guide)
* [Installing on Compute Engine](Compute-Engine-Install-Guide)
* [General Install Guide](General-Install-Guide)

# Quick Start 
Getting started with Mobile Main Street is easy. You simply need to provide information about your community (name, location, and  website), add in categories within your community, and finally add the Twitter or Facebook usernames of the individuals you want in each category. 

1. Define the community
2. Add categories
3. Add storytellers
4. Invite storytellers (optional)

## API Documentation
We're currently expanding the set of documentation publicly available. Documentation will be made available through this Wiki, or via the project website. In the interim please refer to the Google Doc below for our preliminary version of the API. 

[API Models and Routes](https://docs.google.com/document/d/1c76EKzKif4Gk1BEG0i7x-B-Tf1wnnxHzzhCJTuEVASI/edit?usp=sharing)

### Community JSON API
A public RESTful JSON API used to access available members, stories and categories inside of a community. 

#### Routes

* Community
* Member
* Category
* Story

### Administration JSON API
A RESTful JSON API used to access administrative functions related to adding members to the community, user management, category, story and community management. All the community's settings may be accessed through this API.

#### Overview
* Authentication
* Permissions

#### Routes
* Community
* Member
* Story
* Category
* Feed
* User

## Contributing Code
Interested in contributing code to Mobile Main Street? Great! Just submit a pull request and we'll take care of the rest. Please be detailed about what feature you're improving or bug you're fixing, this helps us get the request approved more quickly. 

### Suggested Starting Points
Want to contribute, but don't know where to start? We'll keep this list update with potential starting points. Currently, we're excited about the following:

* Expanded unit testing
* Integration tests and test server
* Better theming support
* Additional adapters to connect different social networks (Google+, LinkedIn, etc.)
* Schema.org compatible profile pages (See [http://schema.org/Organization](http://schema.org/Organization))
* Reducing load and render times on mobile devices
* Suggestions or improvements to story ranking

### Filing Bug Reports
If you find a bug please report it using our [Github Issues page](https://github.com/SyntropyDev/mms-web/issues). Please be detailed and describe steps to reproduce the issue. Screenshots of issues where appropriate are **highly recommended**. 
