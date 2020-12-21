# Problem

COVID-19, quarantine, online-schooling, and online work will forever shape the way students interact with others. Instead of being able to interact with others in a physical environment, students are now forced to interact purely through virtual means which can be very difficult when it comes to meeting new people. Social interaction is a pivotal aspect of the college experience as according to Hurst, B., Wallace, R., & Nixon, S. B. (2013). Finding other students with similar interests is a much larger challenge now than ever before as students are quarantined and taking classes from behind a screen. How might we enable students to make meaningful connections with one another in a virtual environment?

# Solution

## Homescreen & Categories

We envisioned a web application that will allow users to browse categories of interest, join groups, and post updates to groups. This type of platform fosters social interaction with one another and hosts an avenue of making new friends and exploring new interests. We will leverage a card layout for the content on the pages, sizable to the screen being used by the user.

This service is for UW students, but could be expanded to other institutions in the future. This platform is meant for broad and specific groups, (i.e. "Studying" or "IMA pick-up basketball 5-7pm"). The overarching goal is to facilitate meaningful connections with other students.

On the *Categories* page, there will be a set amount of predetermined categories that we believe encompass a wide range of topics. The user will be able to scroll down the page to view more categories.

If the user is signed in, the user can click on the “recommended” button to show them recommended categories based on preferences they selected when creating the account. If the user is not logged in, the recommended button will alert the user to log in to receive category recommendations.

<img src="/public/img/explore.png" width="700">

## Log In

On the *Categories* page, there are many interactions that a user can make. To start with, *Log in*. Clicking the log in button on the *Categories* page will navigate the user to the log in page, prompting them to enter an email and password. The user can then log in, and be routed to *Categories*. If the user doesn’t have an account, they can create an account by filling out a few questions. To log out of their account, the user simply needs to hit the log out button.

<img src="/public/img/login.png" width="700">

If the user types in incorrect login information, there will be an error that shows up that explains the email and/or password is incorrect.

<img src="/public/img/incorrectlogin.png" width="700">

## New Account Survey

The questions the user will answer allows the web application to generate recommended interest groups. The process flow for filling out the questions are as follows: enter your name > choose at least 3 interests > hit ok > route back to *Explore*. While the user fills out the questions, a progress bar will move closer to 100% as the user completes the survey.

<img src="/public/img/surveyname.png" width="700">

<img src="/public/img/emailinput.png" width="700" padding="2">

<img src="/public/img/interestselector.png" width="700" padding="2">

Some errors that may occur during this process flow are: “what is your name?“ is mandatory and won’t let the user click next until there are letters in the text area, given interests will change color once selected, the “OK” button won’t be activated until 3 interests have been selected. Users can always return to *Explore* by hitting the logo on the top left of each page, or the *Explore* button in the header. This route of navigation is available on all pages.

## Groups

From *Explore*, a user can view a category which will bring them to the underlying groups in that category as seen below. The groups listed will be related to the category that was clicked on.

<img src="/public/img/grouplist.png" width="700">
 
## Search Bar

The search bar can be utilized to filter out the groups in real time. 

<img src="/public/img/searchgroup.png" width="700">

## Create New Group

If a group doesn’t exist after a user searches for it, the user will be prompted to create a new group if they would like.

<img src="/public/img/groupnotexist.png" width="700">

A form will appear asking for group name, category, and search tags. From there, the group will then be created. The group creator will then become the group admin. In order to create a group, the user MUST be logged into an account. If the user is not logged in, it will prompt them to log in or create an account.

<img src="/public/img/newgroupdetails.png" width="700">

## View Group: Not Joined

When viewing a group as the user that hasn’t joined yet, they will see the option to join the group. The back button will route the user to the group list page. The page will show group details.

<img src="/public/img/groupoverview.png" width="700">

Users can only join or post within a group if they have an account. Clicking either of these buttons will redirect to the log in page. Once joined, the join button will change to pending as the group admin will admit the user to the group. And after the user is accepted into a group, the pending button will change to a “leave group” button if they decide they want to leave the group. Leaving a group will prompt the user with a confirmation message so that the user does not accidentally leave the group.

Users can “star” a group which will highlight that certain group in the *My Groups* screen below. 

<img src="/public/img/pendingjoin.png" width="700">

## View Group: Joined

When viewing a group as a user that has already joined, the recent posts in that group will appear in chronological order. Users can comment on posts they choose by clicking on the comment button. This will give them the option to comment and view other comments. Users can then click on “Load more comments” to view more comments if there are more than three comments.

<img src="/public/img/groupblog.png" width="700">

<img src="/public/img/loadmorecomments.png" width="700" padding="2">

## Create Post

When creating a post, a title and description are required. An image is optional. The user can hit the back button to return to the recent posts within the group.

<img src="/public/img/createpost.png" width="700">

When hitting submit post, the user will be prompted with a confirmation message in which they can cancel and continue editing their post, or confirm the post to the group.

<img src="/public/img/postconfirmation.png" width="700">

## My Groups

My Groups will show the groups that the user has joined, allowing for quick access. By clicking on the “Admin Groups” button, it will display groups that the user is an admin for.

<img src="/public/img/mygroups.png" width="700">

## Admin User

If the user is an admin, they will be able to see pending join requests in the group admin screen. Instead of being asked to join on the group screen, the admin will have a different view where they can edit the group description and manage member requests to join that specific group.

<img src="/public/img/groupadminview.png" width="700">

<img src="/public/img/adminrequests.png" width="700" padding="2">

The admin user will be prompted a confirmation message if they would like to admit the new user to the group.

<img src="/public/img/adminadmit.png" width="700">
