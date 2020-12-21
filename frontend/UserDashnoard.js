import { React, Component } from 'react'

// loads info through userInfo function from User.js passed in as props
// userInfo returns an object with the user's joined and saved groups, posts and 
// comments they're made, which groups they're an admin of, and their login info
class UserDashboard extends Component {

    // takes in the user info from the userInfo function from User.js 
    // loads users requesting to join groups that user is an admin of
    // input: userInfo object
    // output: list of groups (groupIDs) that user is an admin of with associated 
    //         list of members (userIDs) requesting to join the group
    loadMemberRequestInfo(userInformation) {
        return listOfRequests;
    }

    // takes in the user info from the userInfo function from User.js and reorganizes
    // it to be displayed for the dashboard
    // input: userInfo object 
    // output: none
    loadInfo(userInformation) {

    }

    // displays user's joined groups, saved groups, and groups they're admins of
    // also shows option to approve or reject group members for groups user is 
    // an admin for. 
    render() {

    }

}