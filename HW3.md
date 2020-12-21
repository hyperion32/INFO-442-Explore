# Team Explore
## Requirements List

### Home Page

1.  **COMPLETE** Users must be able to view and type characters into the search bar on the home page.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks into the search bar 4. User types any character into the search bar 5. The search bar must display the characters typed by the user.

2.  **COMPLETE** The system must display fixed categories on the Explore page, as defined by our team, so that the user can browse through categories. 

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User must be able to view pre defined categories on the home page 4. There must not be a create category button.

3.  **REVISED** Users must be able to click on any category on the Explore page and then be routed to that category’s page once the user has logged in. If the user doesn't have an account, they can still view the category's page, but won't be able to join the group.

- Verification Steps: 1. User accesses the application 2. User clicks on any category they'd like 3. User should be able to view the category's list of groups, even if they aren't logged in.

- Reason for REVISED: We revised this requirement to allow for non-logged in users to have access to the category's page. We made this decision as it gives users the opportunity to explore the pages before they've logged in, affording the chance to determine if making an account will provide them with meaningful connections.

4.  **COMPLETE** The system must show all categories, and the user must be able to scroll down to see the rest if they do not all fit on the page.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User must be able to view pre defined categories on the home page 4. User must be able to freely scroll up and down the list of categories.

5.  **COMPLETE** The system must display the categories in alphabetical order.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. Reading from left to right, all categories should be in alphabetical order.

6.  **IMPOSSIBLE** The system must display a recommended section of categories to the user at below the search bar.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User must be able to view a section at the top of the page listing recommended categories.

- Reasoning for IMPOSSIBLE: We determined that this requirement doesn't fit the needs of our MVP to be beneficial to the users on our application. Having the ability to see recommended groups if a useful feature for efficiency, which could be implemented later on down the line, but isn't a feature that will afford users to make meaningful connections with each other right away. See the New Account Survey requirements for more information on why we made this decision.

### Viewing Groups in a Category

7. **COMPLETE** Users must be able to see all the groups in a certain category once they click on a category.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks into a category 4. User should be able to see the existing groups in the category. No groups will show if none have been created.

8. **COMPLETE** Users should be able to scroll down to view more groups if the results do not all fit on one screen.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks into a group 4. User must be able to freely scroll up and down the list of groups within a category if they all don't fit on the screen.

9. **COMPLETE** The groups displayed to the user must be in alphabetical order, following a left to right pattern so it is easy to find the exact group they are looking for.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. Reading from left to right, all groups withing the clicked on category should be in alphabetical order.

10. **REVISED** Users should be able to create a new group if none of the search results were what they were looking for by clicking on a “create group” button. If the user is logged in, it will redirect to the create new group page, otherwise there will not be a create group button on the page.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User searches or scrolls for a group 5. If the group they're looking for doesn't exist, they can click the "create group" button 6. If the user isn't logged in, the user won't have access to the create group button.

- Reasoning for REVISED: We revised this requirement due to the permissions as a non-logged in user. We decided that a non-logged in user shouldn't that they ability to see the create group button. The create group button will only be displayed in the case that a user is logged in to their UW account.

11. **REVISED** The system might display saved groups under the My Groups page.

- Verification Steps: 1. User accesses the application 2. User navigates to the My Groups by clicking it in the navigation bar 3. Users should be able to see saved groups at the bottom of the page. There will be no saved groups if the user hasn't saved any groups since making their account.

- Reasoning for REVISED: We revised this requirement to make the home page more elegant and less cluttered. We determined it was best to display the saved groups under the My Groups tab on the navigation bar as My Groups contains information about Admin Groups, Joined Groups, and now Saved Groups. It made sense from a user standpoint to make this decision to move the location of Saved Groups.

### Search Bar

12. **COMPLETE** Users must be able to type into the search bar while the system updates search results as the user types.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks into the search bar 4. User enters letters into the search bar 5. Search results should update as the user types their search query.

13. **COMPLETE** The system must query the term inputted against groups if on the group page, displaying the matching or partially matching results to the user for each new character inputted.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User types a query into the groups search bar 5. Matching results of the search term will then display.

14. **COMPLETE** The system must query the term inputted against categories if on the category page, displaying the matching or partially matching results to the user for each new character inputted.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User types a query into the categories search bar 5. Matching results of the search term will then display.

15. **COMPLETE** The system should display search results for groups in alphabetical order.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User types a query into the groups search bar 5. Matching results of the search term will then display in alphabetical order.

16. **COMPLETE** The system should display search results for categories in alphabetical order.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User types a query into the categories search bar 5. Matching results of the search term will then display in alphabetical order.

### Navigation Bar

17. **COMPLETE** The system must display a nav bar at the top of all pages to allow for easy navigation.

- Verification Steps: 1. User accesses the application 2. User mouses towards the top of the application 3. User should be able to easily navigate to Home, My Groups (when logged in), and Log in with UW Gmail/Log Out.

18. **COMPLETE** The system must have a logo on the nav bar, redirecting the user to the home page when clicked.

- Verification Steps: 1. User accesses the application 2. User navigates to the top left of the page. 3. User should be able to see the Explore logo 4. If the user clicks the Explore logo, they should be routed to the home page.

19. **REVISED** The system must have a My Groups tab on the nav bar, redirecting the user to the My Groups page when clicked and the user is logged in, otherwise the user won't have access to the My Groups tab.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click My Groups if the user is logged in. 3. If the user isn't logged in, they won't have access to the My Groups tab until they log in.

Reason for REVISED: We revised this requirement because we made the decision to not display the My Groups tab if the user isn't logged in. Once the user logs in, they will then have access to the My Groups tab. We made this decision as it made sense from a permissions standpoint to limit the functionality of the application until the user logs in to their account.

20. **COMPLETE** The system must have a Log In tab on the nav bar, redirecting the user to the Log In page when clicked. The system must change the Log In element to Log Out once a user has logged in.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In, taking them to the log in screen. 3. Once the user logs in, the element should change to Log Out, vice versa.

### Log in IMPOSSIBLE

- Reasoning for IMPOSSIBLE: We made the decision to delete this feature from our application as we chose to use Google for identity management purposes. We chose to do this as we wanted to restrict the application to only UW email users (at least for now). This decision makes the log in process easier and more elegant for users when they access our application and sign into their UW Gmail account.

21. **IMPOSSIBLE** Users must be able to type into the email and password text box.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. On the Log In page, the user should be able to type characters into both the email and password text box.

22. **IMPOSSIBLE** The system should show dots/stars instead of letters when typing in the password text box.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. On the Log In page, the user should be able to type characters into the password text box. 4. When typing the characters into the password box, they should be presented as stars for privacy purposes.

23. **IMPOSSIBLE** The system should notify the user if the email or password is incorrect and prompt them to re-enter username and password.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. On the Log In page, the user should be able to type characters into both the email and password. 4. If the email or password are incorrect, a prompt should present telling the user to try again.

24. **IMPOSSIBLE** The system should log the user in if the email and password match an existing account redirecting them to the home(explore) page after a successful authentication.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. On the Log In page, the user should be able to type characters into both the email and password text box. 4. If an account exists and the credentials are correct, the user should be routed back to the home page.

25. **IMPOSSIBLE** The system should show a sign-up button, directing the user to the new account survey page if clicked.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account.

### New Account Survey IMPOSSIBLE

- Reasoning for IMPOSSIBLE: We made the decision to delete this feature from our application because it doesn't contribute to making meaningful connections amongst users. The result of this survey would have been recommended categories for the new user that would be displayed at the top of the home page. We concluded that recommended sections wouldn't contribute to the experience of finding new categories as we have ~50 predefined categories that are easily accessible. We diverted the time from making this feature to ensuring that the home page experience was great for our users.

26. **IMPOSSIBLE** The new account survey is made up of multiple questions, with one question on each page of the survey.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. The survey should have multiple questions, one on each page of the survey.

27. **IMPOSSIBLE** The system should show back and next buttons, allowing the user to move between questions freely, saving inputted information so it doesn’t get lost.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. On the new account survey, the user should be able to navigate back and forth with the arrow buttons. 

28. **IMPOSSIBLE** The system must be able to calculate the progression percentage as users complete survey items and display in a progression bar on each survey page.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. On the survey, there should be a progress bar at the top of the screen that shows the user's progress as they complete the survey.

39. **IMPOSSIBLE** Users must be able to type into the name field to enter a name when on the survey name page.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. On the first page, the user should be able to click into the text back and enter their name, using only letters 6. If a number or symbol is used, an error will tell the user to not use numbers or symbols.

30. **IMPOSSIBLE** Users must be able to type their email into the email field when on the survey email page.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. On the second page, the user should be able to click into the text back and enter their email, using any character they'd like.

31. **IMPOSSIBLE** Users should be able to select up to 3 interests from a group of categories on the survey interest page, highlighting the selected categories in an accessible color when selected.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. On the third page, the user should be able to select up to three of the suggested category interest bubbles 6. When a bubble is selected, it should change to a highlighted color 7. If the user selects more than three interests, an error will be displayed.

32. **IMPOSSIBLE** The system shouldn’t let the user move to the next question if they have not entered information correctly; name must be >0 characters, email should be a valid email address, and the user must select 3 categories for the interest selection, graying out the next button until this is satisfied.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. When filling out the survey, the system won't allow the user to continue unless: name must be >0 characters, email should be a valid email address, and the user must select 3 categories for the interest selection 6. The next button on each page will be grayed out until each page's requirement is met.

33. **IMPOSSIBLE** The system must be able to generate recommended categories on the home page based on the user’s selected interests, matching their selected interests to existing categories and displaying at the top of the categories search results on the home page.

- Verification Steps: 1. User accesses the application 2. In the navbar, the user should be able to click Log In to be routed to the log in page. 3. If the user doesn't have an account, the user should be able to click the sign up button to create an account 4. The user will then be routed to the new account survey 5. After the last page of selecting interests, the system should generate suggested categories 6. The suggested categories should be displayed at the top of the home page list of categories.

### Create New Group

34. **COMPLETE** The new group page should have 3 inputs, a group name, a group category and a group description.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. The element should allow the user to enter a group name, select the desired category, and enter a group description.

35. **COMPLETE** Users must be able to type into the group name and group description text inputs.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. The user should be able to click into the group name and group description box and type into the boxes.

36. **COMPLETE** The system might only allow numbers and letters to be typed into each text box.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. The user might be able to click into the group name and group description box and only type letters and numbers into the boxes.

37. **COMPLETE** Users must be able to open the category dropdown menu to choose the correlating category.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. The user should be able to select a desired category from the category dropdown box 6. The user should be able to scroll up and down the category dropdown list.

38. **COMPLETE** The system must allow the user to create a group even if the group name is already the name of an existing group as group names are not unique.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. If the user enters the exact name of an existing group in the current category, the system should allow the group to be created as groups name are not unique.

39. **COMPLETE** The system must limit the number of characters per text box to 32 characters for group name and 240 for description.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying a modal that will allow the user to create a new group 5. The user should be able to type into the group name and group description boxes 6. If the user types more than 32 characters into the group name box and tries to submit, an error will display 7. If the user types more than 240 characters into the group description box and tries to submit, an error will display.

40. **REVISED** The system must give an error if the user tries to create a new group with missing group information.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on the create group button, displaying an element that will allow the user to create a new group 5. The user should be given an error if they try to submit the new group when there's missing group information (name, category, description).

- Reasoning for REVISED: We decided to revise this requirement to better match the workflow of the modal. Instead of greying out the "create group" button, we implemented an error that will be thrown if the user tries to create the group with missing information. By doing this, we are keeping modal buttons consistent in format and are following the same workflow for throwing errors when forms aren't correctly filled out.

41. **IMPOSSIBLE** Users must be able to create an account if the form contains all of the correct information.

- Reasoning for IMPOSSIBLE: This requirement is a typo on my (Greyson) part and shouldn't be considered part of our requirements list. It does not relate to creating a new group.

### View Group

42. **COMPlETED** The system must display the group name, description and category to the user, as defined by the group admin.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. The user should be able to review the following information on the group page: group name, group description, and category the group is in.

43. **IMPOSSIBLE** The group administrator’s email should be displayed if the user is in the group.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. If the user is a member of the group, they should be able to see the group admin's email in the group details section.

- Reasoning for IMPOSSIBLE: We determined that this requirement was impossible because the user's email is not returned from the Google authentication method we used, unless we choose to always return the email from the Google authentication method. We decided against the latter as we believe it could be a privacy issue if the user's email is accessible at all times in the application. If the students want to share their contact information, they can discuss that in comments (reply feature shoutout Andy), find each other through social media or the student directory.

44. **IMPOSSIBLE** The system must allow the user to route back to the previous page by clicking the back arrow.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. The user should be able to route back to the previous page by hitting the back arrow.

- Reasoning for IMPOSSIBLE: We decided that this requirement isn't necessary because users can just use their browser's back arrow to swipe of their touch pad.

45. **COMPLETE** If the user is logged in and not in the group, the system must allow the user to request to join the group by clicking the ask to join button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. If the user is logged into their account and has not joined the group they clicked, the user should be able to request to join the group.

46. **COMPLETE** The system must change the join status to pending if they requested to join. The join button will now say pending and will be greyed out.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. If the user is logged into their account and has not joined the group they clicked, the user should be able to request to join the group 6. The user's join status should change to pending and be greyed out once the user has requested to join the group.

47. **COMPLETE** The system must limit the number of group members to 10, displaying an error message if a user requests to join and the member count is already 10.

- Verification Steps (will be tricky as a group needs to have 10 people in it): 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. The number of members in the group should be displayed to the user 6. If the user tries to join a group with 10 members, an error will be displayed.

48. **COMPLETE** The user must be able to leave any group that they have joined by clicking on the “leave group” button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. If the user is logged into their account and is a member of the group, the user should be able to click the leave group button which will redact the user's membership from the group 6. If the user isn't a member of the group, the user must request to join, be admitted be the group admin, then can click to leave the group.

49. **REVISED** The system should allow users to leave comments by clicking on the comment button, and each comment should have a max character count error that will be thrown if the user creates a comment over the character limit.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. The user should be able to post a comment to the group by clicking the comment button 6. If the user reaches the max character count for a comment, an error will be displayed after the user tries to create the comment.

- Reasoning for REVISED: We revised this requirement to match how we implemented the error to be thrown if the character count is at the limit. Instead of throwing the error while the user is typing over the limit, we implemented the error to be thrown if the user tries to create a comment over the character limit. This change was made according to our development team and implementation methods that were chosen to be used.

50. **REVISED** The system should display comments below the group information and show these comments in chronological order with the most recent comment being at the top.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User should be able to view comments on the group in chronological order.

- Reasoning for REVISED: We revised this requirement's verbiage to match the location of where it's located in the application--group tab--and not on blog posts.

51. **REVISED** Users must be able to view and reply to comments on groups in a thread system below the original comment.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User should be able to view comments on the group 6. User should be able to reply to comments by hitting the reply button.

- Reasoning for REVISED: We revised this requirement's verbiage to match the location of where it's located in the application--group tab--and not on blog posts.

52. **REVISED** The system should only display the three most recent comments unless “load more comments” is pressed.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User should be able to view the three most recent comments on the group 7. If the user presses the load more comments button, the user should be able to see all the comments on the blog post.

- Reasoning for REVISED: We revised this requirement's verbiage to match the location of where it's located in the application--group tab--and not on blog posts.

53. **COMPLETE** Users should be able to save the group so they can return to the page at a later date (through the my groups page) by clicking on the save button. Once saved, you can click it again to unsave.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User should be able to save a group by clicking the save group button 6. The save button will then change to unsaved in which the user can click again to unsave the group.

### View Group: Blog Tab

54. **COMPLETE** The system must display blog posts in the group in chronological order (recent first).

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User should be able to view blog posts in chronological order.

55. **COMPLETE** For each blog post, the system should display the post title, author, description, and may show an image if the post has one.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. Blog posts should display the title, author, description (consolidated version), and may show an image.

56. **REVISED** Users should be routed to the “create post” screen if they click on the “create new post” button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User should be able to click on create new post to make a new blog post.

- Reasoning for REVISION: Changed the verbiage of buttons to accurately match the application.

57. **REVISED** The system should only allow users to post if they have already joined the group by only displaying the create new post button to group members.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User should only be able to see the create new post button if they are logged in and have already joined the group.

- Reasoning for REVISION: Changed the verbiage of buttons to accurately match the application.

58. **COMPLETE** Users must be able to comment on any post in the group by clicking the comment button (comments work the same was as in the view group tab).

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User should be able to click into a blog post 7. User should be able to create a comment on a blog post by hitting the comment button.

### Create Blog Post

59. **IMPOSSIBLE** The system must allow the user to route back to the previous page by clicking the back arrow.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User should be able to click the back arrow to route back to the previous page.

- Reasoning for IMPOSSIBLE: We decided that this requirement isn't necessary because users can just use their browser's back arrow to swipe of their touch pad.

60. **COMPLETE** The system must allow the user to enter a title for their post, including letters and numbers with a maximum of 50 characters.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. User should be able to enter a post title 8. If the user has a title more than 50 characters, an error will appear.

61. **COMPLETE** The system must allow the user to enter a description (body text) for their post. Including letters and numbers with a maximum of 1,000 characters.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. User should be able to enter a post description containing any characters 8. If the user has a description more than 1000 characters, an error will appear.

62. **IMPOSSIBLE** The system might allow the user to upload an image for their post of file type jpeg or png.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. User may be able to attach an image of file type jpeg or png to their post.

- Reasoning for IMPOSSIBLE: We decided that this requirement is impossible because it doesn't contribute to making meaningful connections amongst users. We determined that an in-depth description for a blog post would be sufficient for users to make meaningful connections with each other. If a conversation would like to be taken further, perhaps the users could coordinate and share their contact information with each other where they could then transfer photos of their blog post, if they'd like.

63. **COMPLETE** The system must allow the user to submit their post by hitting the create post. 

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. User fills out the new post details 8. User can submit post by hitting the create post button.

64. **COMPLETE** The system must display an error if the title or description are empty, not allowing the user to post until satisfied.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. If the user doesn't fill out all post details, an error will appear.

65. **COMPLETE** The system must display a confirmation, ensuring that the user wants to post to the group or not by either clicking the post or cancel button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on a category 4. User clicks on a group 5. User clicks on blog posts 6. User clicks on create new post 7. User fills out the new post details 8. User can submit post by hitting the create post button 9. A confirmation message will appear, confirming if the user wants to post.

### My Groups

66. **COMPLETE** The system must display every group that the user has joined.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that the user has joined.

67. **COMPLETE** For each displayed group, there should be a “leave group” button within the tile that allows users to leave groups that they have previously joined.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that the user has joined 5. User should be able to click the leave group button to leave a group they have joined.

68. **COMPLETE** When hitting the leave button, the system must display a confirmation, ensuring that the user wants to leave the group by clicking the confirm or cancel button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that the user has joined 5. User should be able to click the leave group button to leave a group they have joined 6. A confirmation popup should appear to confirm this action.

69. **REVISED** The system must show the groups that the user is an admin for if they click on the “My Groups” tab in the nav bar. 

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that the user is an admin for (groups they've created).

- Reason for REVISED: We decided to revise this requirement to properly match the workflow that we've implemented into our application. To access Admin Groups, the user just needs to click into the My Groups tab in the nav panel. Their Admin groups will be displayed at the top of this page.

70. **COMPLETE** The system must display the groups that they have joined in chronological order.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that the user has joined 5. Joined groups should be displayed in chronological order, from left to right.

71. **COMPLETE** The system must route the user to the specific group page when clicked on.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to click on any group to be routed to that group's page.

### Admin User

72. **REVISED** The system must allow the user to view pending requests for groups the user has created.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. User routes to the groups page 6. User clicks Member Requests 7. User should see the pending user requests if another user has requested membership into the group.

- Reasoning for REVISED: We decided to revise this requirement to properly match the workflow that we've implemented into our application. To access pending requests, the Admin user just needs to click into a group they are an admin for, then view the Member Requests tab in the group.

73. **REVISED** For each group the user is an admin for, the system must display the name of each user pending.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. User routes to the groups page 6. User clicks Member Requests 7. User should see the pending user requests and their details if another user has requested membership into the group.

- Reasoning for REVISED: We decided to revise this requirement to properly match the workflow that we've implemented into our application. To view pending request information, the Admin user just needs to click into a group they are an admin for, then view the Member Requests tab in the group. The pending requests should display the requested user's details.

74. **COMPLETE** The system must allow the Admin user to click on a pending request to make the decision to admit them.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. User routes to the groups page 6. User clicks Member Requests 7. User should see the pending user requests 8. User can then make a decision.

75. **COMPLETE** The system must display a confirmation when the Admin user clicks on a pending request, allowing them to click confirm.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. User routes to the groups page 6. User clicks Member Requests 7. User should see the pending user requests 8. User can then make a decision 9. A confirmation should pop up when the user makes their decision.

76. **COMPLETE** The system must allow the pending user to access the group if they are admitted by setting the user as a group member or cancel the request if they are denied.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. The user should be able to click on the pending request to make an admission decision 6. If admitted, the user that requested admission should be able to enter the group they requested membership for.

77. **COMPLETE** The system must be able to recognize when a user is an admin for a group (a group they created), giving them the ability to delete the group if desired from the “Your Admin Groups” tab by hitting the delete button.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that they have created 5. User should be able to hit the delete button on a group tile to delete the group.

78. **REVISED** The system must display a delete confirmation popup, allowing the user to confirm the deletion the group.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User should be able to view the groups that they have created 5. User should be able to hit the delete button on a group tile to delete the group 6. A confirmation message will popup to confirm the user's decision.

- Reasoning for REVISED: We revised this requirement to match how we implemented the confirmation message. We decided that a popup confirmation would be best for this action as deleting a group will delete all of its information. So, we want the user to adequately think about this decision by showing them a confirmation popup.

79. **COMPLETE** The system must remove the pending request from the Admin user’s view after a decision has been made.

- Verification Steps: 1. User accesses the application 2. User navigates to the home page by clicking either Explore or Home 3. User clicks on My Groups 4. User clicks into a group under "Admin Groups" -- no groups will exist in this section if the user hasn't created any groups 5. The user should be able to click on the pending request to make an admission decision 6. Once a decision has been made, the pending request should be removed from the pending requests section.