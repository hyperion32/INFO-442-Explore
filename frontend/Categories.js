import { React, Component } from 'react'
import './Groups'
import './Survey'
import './SearchBar'

// shows all categories
class Categories extends Compoonent {

    // loads the categories to display for the user
    // inputs: none
    // outputs: list of categories to display
    categoryLoader() {
        return listOfCategories;
    }


    // loads the categories to display that match the search phrase
    // inputs: search string
    // outputs: list of categories to display that start with the same search phrase
    categorySearch(search) {
        return listOfSearchCategories;
    }

    
    // returns the view for the categories page
    // loads list of categories - navigates to the categories group page if clicked on
    render() {

    }
}