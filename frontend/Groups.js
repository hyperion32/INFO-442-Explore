import { React, Component } from 'react'
import "./GroupPage";
import './SearchBar'

// shows all groups within a selected category
class Groups extends Component {

    // loads the groups to display for the user based on the selected category
    // inputs: category selected - comes from Categories and is passed in as a prop
    // outputs: list of groups to display
    groupLoader(selectedCategory) {
        return listOfGroups;
    }

    // searches for groups that match the entered search phrase
    // input: search string
    // output: list of categories that match the phrase
    groupSearch(search) {
        return listOfSearchedGroups;
    }

    // creates a new group on the backend - all information should be validated before the group is created
    // input: form information to create the group along with user's ID
    // output: a request to the server to create the new group
    createGroup(newGroupFormInfo, user) {
        validateInfo(newGroupFormInfo) // validates that all the user info entered is correct

    }

    // lets user delete a group only if the user is the admin
    // input: group ID for the group being deleted and user ID
    // output: a request to the backend to delete a group 
    deleteGroup(groupID, userID) {

    }

    // adds group to user's groups
    // input: user ID and group ID
    // output: request to server to add group to user's groups
    joinGroup(groupID, userID) {

    }

    // removes group to user's groups
    // input: user ID and group ID
    // output: request to server to delete group from user's groups
    leaveGroup(groupID, userID) {
        
    }

    // loads all the information associated with a group
    // input: group id
    // output: group info
    groupInfo(groupID) {
        return groupInfo;
    }

    // returns the view for the groups page
    // loads list of groups - navigates to a group page if clicked on
    // shows group info as a pop up
    render() {

    }

}