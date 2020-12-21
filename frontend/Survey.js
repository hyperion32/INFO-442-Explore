import { React, Component } from 'react'

class Survey extends Component {

    // takes in a list of answered questions based on user's interest of a few of the categories
    // re-orders the offered categories according to the user's interest and returns list
    // input: list of asnwered questions by user along with user ID -- ID is null if user is not logged in
    // output: returns list of ordered categories
    computeSuggestedCategories(userAnswers, userID) {
        return listOfOrderedCategories;
    }

    // displays a survey with predetermined questions to compute with categories to suggest first
    // on completion of survey: if user if not logged in, refirects to sign up page after answering survey
    // if user is logged in, redirects to catgories page with the ordered categories
    render() {

    }
}