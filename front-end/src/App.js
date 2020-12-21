import React, { Component } from 'react';
import './App.css';
import { NavBar } from './Components/NavBar';
import { Categories } from "./Components/Categories";
import LogIn from "./Components/LogIn";
import { Container } from '@material-ui/core';
import { Route, Switch } from 'react-router-dom'
import { UserDashboard } from "./Components/UserDashboard";
import { Groups } from './Components/Groups';
import { BlogCards } from "./Components/BlogCards";
import { BlogPost } from './Components/BlogPost';
import { Group } from './Components/Group';
import Authenticate from './Components/Authenticate'
import RedirectPage from './Components/Redirect'
import NewGroup from './Components/NewGroup'

export default class App extends Component {
  render() {
    return (
        <main>
           <Container maxWidth={false}>
          <>
            <NavBar />
            <Switch>
              <Route exact path='/' component={Categories} />
              <Route path='/login' component={LogIn} />
              <Route path="/dashboard" component={UserDashboard} />
              <Route path="/groups" component={Groups} />
              <Route path="/grouppage" component={BlogCards} />
              <Route path="/blog" component={BlogPost} />
              <Route exact path="/group/:groupId" component={Group} />
              <Route exact path="/group/:groupId/blog/:blogId" component={BlogPost}></Route>
              <Route exact path = "/authenticate" component={Authenticate}/>
              <Route path = "/redirect" component={RedirectPage} />
              <Route path = "/create" component={NewGroup} />
            </Switch>
          </>
          </Container>
        </main>
    )
  }
}

