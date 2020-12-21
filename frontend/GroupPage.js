import { React, Component } from 'react'

// shows a singular selected group page with posts
class GroupPage extends Component {

    // loads all the blog posts made in that group
    // input: takes in the ID of the group - is passed this in as a prop from Groups
    // output: returns list of posts made on the group
    postLoader(groupID) {
        return listOfPosts;
    }

    // creates a new blog post
    // input: form content for the new blog post along with the user ID
    // output: a request to the backend to create a new blog post 
    createPost(newPostFormContent, userID) {
        return createdPost;
    }

    // lets user delete blog post only if the user made the post or is the admin
    // input: post ID for the post being deleted and user ID
    // output: a request to the backend to delete a blog post 
    deletePost(postID, userID) {
    }

    // loads all the comments made in that post
    // input: takes in the ID of the post 
    // output: returns list of comments made on the post
    commentsLoader(postID) {
        return listOfComments;
    }

    // creates a new comment
    // input: form content for the new comment along with the user ID
    // output: a request to the backend to create a new comment
    createComment(newCommentFormContent, userID) {
        return createdComment;
    }

    // lets user delete comment only if the user made the comment or is the admin
    // input: post ID + comment ID for the post being deleted and user ID
    // output: a request to the backend to delete a comment 
    deleteComment(postID, commentID, userID) {
    }

    // loads the group info about the group
    // gets group name as prop from Groups
    // shows edit and accept options if user is admin
    render() {

    }
}