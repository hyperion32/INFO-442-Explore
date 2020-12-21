# Team Explore

## Architecture Description

### CategoryModel

This component is a model that stores all of the categories that the user will be able to scroll through, and the links to underlying groups in those categories.

The model resides only on the server.

* The **CategoryController** and **SearchController** communicate with the model. They communicate:
* The **CategoryController** can ask the **CategoryModel** to load the categories.
* The **SearchController** can ask the **CategoryModel** for matching category search results.

### GroupModel

This component is a model that stores all of the group information that the user will be able to scroll through.

The model resides only on the server.

The **GroupTileController**, **SearchController**, **NewGroupController**, **GroupController**, **CategoryController**, and **BlogController** communicate with the model. They communicate:
* The **GroupTileController** can ask the **GroupModel** for subsequent groups in a category as the user scrolls.
* The **SearchController** can ask the **GroupModel** for matching group search results.
* The **NewGroupController** can ask the **GroupModel** to update with a new group.
* The **GroupController** can ask the **GroupModel** for the group information: current amount of users, description.
* The **GroupController** can ask the **GroupModel** for the joined users of the group.
* The **CategoryController** can ask the **CategoryModel** for the groups in a category.
* The **BlogController** can ask the **GroupModel** for the group info and users.

### UserModel

This component is a model that stores all of the user information so that the user will have a personalized experience every time they log in to the application.

The model resides only on the server.

The **AdminController**, **GroupController**, **PostContoller**, **SurveyController**, **LoginController**, and **SavedGroupController** communicate with the model. They communicate:
* The **AdminController** can ask the **UserModel** for the admin groups and admin user information.
* The **GroupController** can ask the **UserModel** for groups the user has joined or saved.
* The **GroupController** can ask the **UserModel** for joined users of a group.
* The **PostController** can ask the **UserModel** for matching a post with a user.
* The **SurveyController** can ask the **UserModel** for creation of a new user based on survey results.
* The **LoginController** can ask the **UserModel** for user login information, if it exists.
* The **LoginController** can ask the **UserModel** to create a new user.
* The **SavedGroupController** can ask the **UserModel** for the user’s saved groups.

### CommentModel

This component is a model that stores comment information such as the time stamp, comment text, and the user that commented so that the users will be able to interact with comments on blog posts.

The model resides only on the server.

The **CommentController** communicates with the model. They communicate:
* The **CommentController** can ask the **CommentModel** for creation of a new comment.

### BlogModel

This component is a model that stores blog information such as the blog posts in a specific group and who created the blog post so that the users will be able to view all the blog posts within a group.

The model resides only on the server.

The **BlogController** communicates with the model. They communicate:
* The **BlogController** can ask the **BlogModel** for the blog posts in a group.
* The **BlogController** can ask the **BlogModel** for tnew blog creation.

### MemberRequestModel

This component is a model that stores all member request information in a specific group.

The model resides only on the server.

* The **MemberGroupController** communicates with the model. They communicate:
* The **MemberGroupController** can ask the **MemberRequestModel** for the member requests in a group.

### GroupTileController

This component is a controller that updates models and receives alerts from views about groups.

The **GroupTileView** component communicates with **GroupTileController** by:
* The **GroupTileView** sends user input to **GroupTileController** when the user clicks to view a group within a category.
* The **GroupTileView** sends user input to **GroupTileController** when the user clicks to join a group within a category.
* The **GroupTileView** sends user input to **GroupTileController** when the user clicks to leave a group within a category.

### SearchController

This component is a controller that updates models and receives alerts from views about the user searching for items.

The **SearchView** component communicates with **SearchController** by:
* The **SearchView** sends user input to **SearchController** when the user clicks in the search bar and begins typing a search query for categories.
* The **SearchView** sends user input to **SearchController** when the user clicks in the search bar and begins typing a search query for groups.

### CommentController

This component is a controller that updates models and receives alerts from views about user’s commenting on blog posts.

The **CommentView** and **NewCommentView** component communicate with **CommentController** by:
* The **CommentView** sends user input to **CommentController** when the user clicks to view comments.
* The **NewCommentView** sends user input to **CommentController** when the user clicks to create a new comment on a post.

### CategoryController

This component is a controller that updates models and receives alerts from views about the user’s home page.

The **CategoryView** component communicates with **CategoryController** by:
* The **CategoryView** sends user input to **CategoryController** when the user scrolls through the homepage categories in alphabetical order.

### BlogController

This component is a controller that updates models and receives alerts from views about blog posts.

The **BlogView** component communicates with **BlogController** by:
* The **BlogView** sends user input to **BlogController** when the user scrolls through the blog posts of a specific group.
* The **BlogView** sends user input to **BlogController** when the user clicks to create a new blog post in a group.
* The **BlogView** sends user input to **BlogController** when the user clicks to submit a new blog post in a group.

### AdminController

This component is a controller that updates models such as the **GroupModel** and **UserModel** and receives alerts from views about the user’s admin page.

The **AdminView** and **PopUpView** components communicate with **AdminController** by:
* The **AdminView** sends user input to **AdminController** when the user clicks on the Admin Panel to view their admin groups and pending requests.
* The **AdminView** sends user input to **AdminController** when the user clicks on the Admin Panel to manage their admin groups.
* The **PopUpView** sends user input to **AdminController** when an admin user clicks to confirm a join request.

### NewGroupController

This component is a controller that updates models and receives alerts from views about the new groups being created.

The **NewGroupView** and **PopUpView** components communicate with **NewGroupController** by:
* The **NewGroupView** sends user input to **NewGroupController** when the user clicks to create a new group.
* The **PopUpView** sends user input to **NewGroupController** when a user clicks to create a new group.

### SurveyController

This component is a controller that updates models and receives alerts from views about the new user survey.

The **SurveyView** component communicates with **SurveyController** by:
* The **SurveyView** sends user input to **SurveyController** when the user doesn’t have an account and fills out the new user survey through multiple steps.
* The **SurveyView** sends user input to **SurveyController** as the user completes the survey.

### LoginController

This component is a controller that updates models and receives alerts from views about the user login process.

The **LoginView** component communicates with **LoginController** by:
* The **LoginView** sends user input to **LoginController** when the user navigates through logs in and enters their email and password information.
* The **LoginView** sends user input to **LoginController** if the user doesn’t have an account, prompting them to sign up.

### GroupController

This component is a controller that updates models and receives alerts from views about the group information.

The **GroupView** and **PopUpView** components communicate with **GroupController** by:
* The **GroupView** sends user input to **GroupController** when the user clicks on a group within a category.
* The **PopUpView** sends user input to **GroupController** when the user clicks leave a group.
* The **PopUpView** sends user input to **GroupController** when an admin user clicks to delete a group.
* The **PopUpView** sends user input to **GroupController** when a user clicks to join a group at maximum capacity.

### SavedGroupController

This component is a controller that updates models and receives alerts from views about the saved groups.

The **SavedGroupView** component communicates with **SavedGroupController** by:
* The **SavedGroupView** sends user input to **SavedGroupController** when the user clicks to save or unsave a group.

### MemberGroupController

This component is a controller that updates models and receives alerts from views about member requests.

* The **MemberGroupView** component communicates with **MemberGroupController** by:
* The **MemberGroupView** sends user input to **MemberGroupController** when the user clicks to join a group.
* The **MemberGroupView** sends user input to **MemberGroupController** when the admin user accepts or declines the request.

### GroupTileView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with groups within a category.

The **GroupModel** and **UserModel** components communicate with **GroupTileView** by:
* The **GroupModel** updates **GroupTileView** with the groups within a certain category.
* The **UserModel** updates **GroupTileView** with the groups that the user has joined, giving access to the posts and comments.

### SearchView

This component is a view that grabs model data to update itself and alerts controllers of events when the user provides search queries.

The **GroupModel** and **CategoryModel** components communicate with **SearchView** by:
* The **GroupModel** updates **SearchView** with the groups defined by the user’s search query.
* The **CategoryModel** updates **SearchView** with the categories groups defined by the user’s search query.

### CommentView

This component is a view that grabs model data to update itself and alerts controllers of events when the user comments on blog posts.

The **CommentModel component communicates with **CommentView** by:
* The **CommentModel** updates **CommentView** with the comments on a specific post within a group.
* The **CommentModel** updates **CommentView** with the replies to comments on a specific post..

### NewCommentView

This component is a view that grabs model data to update itself and alerts controllers of events when the user comments on blog posts.

The **CommentModel** component communicates with **NewCommentView** by:
* The **CommentModel** updates **NewCommentView** with the new comment creation.

### CategoryView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with the home page.

The **CategoryModel** component communicates with **CategoryView** by:
* The **CategoryModel** updates **CategoryView** with the categories to be displayed to the user.

### BlogView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with blog posts.

The **BlogModel** component communicates with **BlogView** by:
* The **BlogModel** updates **BlogView** with the blogs that are supposed to be displayed to the user if they’ve joined the group.
* The **BlogModel** updates **BlogView** with the information for creating a new blog post (title, description, image).

### UserView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with their profile.

The **UserModel** component communicates with **UserView** by:
* The **UserModel** updates **UserView** with the admin group information pertaining to a specific user.
* The **UserModel** updates **UserView** with the joined groups.

### SurveyView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with the new account survey.

The **UserModel** component communicates with **SurveyView** by: 
* The **UserModel** updates **SurveyView** with the survey information pertaining to a specific user, creating the new user when the survey is completed.
* The **UserModel** updates **SurveyView** with the survey progression as the user completes items.
* The **UserModel** updates **SurveyView** with navigation for the users to move freely between survey questions, saving completed data.

### LoginView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with the login screen.

The **UserModel** component communicates with **LoginView** by: 
* The **UserModel** updates **LoginView** with the login information for returning users.
* The **UserModel** updates **LoginView** with a sign in prompt if an account doesn’t exist.

### NewGroupView

This component is a view that grabs model data to update itself and alerts controllers of events when the user interacts with the New Group Page

The **UserModel** component communicates with **NewGroupView** by: 
* The **UserModel** updates **NewGroupView** with the information for creating a new group (group name, group category, group description).

### PopUpView

This component is a view that grabs models data to update itself and alerts controllers of events when the user interacts with confirmation commands.

The **GroupModel** and **UserModel** communicate with **PopUpView** by:
* The **GroupModel** updates **PopUpView** with a confirmation if a user tries to join a group that’s already at maximum capacity.
* The **GroupModel** updates **PopUpView** with a confirmation if a user would like to create a new group.
* The **GroupModel** updates **PopUpView** with a confirmation if a user would like to leave a group.
* The **UserModel** updates **PopUpView** with a post confirmation when the user tries to create a new post.
* The **UserModel** updates **PopUpView** with a confirmation message when an admin user tries to admit a new user to a group.
* The **UserModel** updates **PopUpView** with a confirmation message when when an admin user clicks to confirm a join request.

### GroupView

This component is a view that grabs model data to update itself and alerts controllers of events when the user creates a new blog post.

The **GroupModel** component communicates with **GroupView** by: 
* The **GroupModel** updates **GroupView** with the information of the group that the user clicked on, including the description and current amount of users.

### SavedGroupView

This component is a view that grabs model data to update itself and alerts controllers of events when the user creates a new blog post.

The **UserModel** component communicates with **SavedGroupView** by: 
* The **UserModel** updates **SavedGroupView** with the information if a user has saved a group or not, indicating this to the user.

### MemberGroupView

This component is a view that grabs model data to update itself and alerts controllers of events when the user requests to join a group

* The **MemberRequestModel** component communicate with **MemberGroupView** by:
* The **MemberRequestModel** updates **MemberGroupView** with the status of the user’s request.

### AdminView

This component is a view that grabs model data to update itself and alerts controllers of events when the user requests to join a group

* The **GroupModel** component communicate with **AdminView** by:
* The **GroupModel** updates **AdminView** with the pending member request tiles and admin group tiles.