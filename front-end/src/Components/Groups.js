import { React, Component } from 'react'
import SearchBar from "material-ui-search-bar";
import { Typography, Grid, Container, Button, Card, CardActions, CardContent, CardMedia } from '@material-ui/core';
import { Row } from 'react-bootstrap';
import { Link } from 'react-router-dom'
import { GetCookie } from "../UtilityFunctions";
import NewGroup from './NewGroup';

// shows all groups within a selected category
export class Groups extends Component {

  constructor(props) {
    super(props);
    this.state = {
      auth: '',
      pagesShown: 1,
      query: '',
      moreDataToLoad: true,
      search: false
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    this.setState({
      auth: auth
    })
    this.searchGroups(auth, this.props.location.state.categoryId, 1, '');
    window.addEventListener('scroll', this.loadMore);
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this.loadMore);
  }


  onSave(card) {
    console.log("CLICK CLICK", card);
    card.isSaved === true
      ? this.unsaveGroup(this.state.auth, card.groupId)
      : this.saveGroup(this.state.auth, card.groupId)
  }

  search(value) {
    this.setState({
      search: true
    })
    this.searchGroups(this.state.auth, this.props.location.state.categoryId, this.state.pagesShown, value)
  }

  // returns the view for the groups page
  // loads list of groups - navigates to a group page if clicked on
  // shows group info as a pop up
  render() {
    if (this.state.data) {


      return (<div>
        <Container maxWidth="md">
          <Typography component="h2" variant="h2" align="center" color="textPrimary" gutterBottom>
            {this.props.location.state.categoryName}
          </Typography>
          < hr style={{ marginTop: "-1rem", backgroundColor: "#3399FF", width: "200px", height: "3px" }} />
          <Row className="w-100"><Button size="medium" color="primary" >
            <Link to={{
              pathname: '/',
              state: {
                auth: this.state.auth,
              }
            }}>Back</Link></Button>


            <div style={{ margin: "auto", width: "60%" }}>
              <SearchBar
                value={this.state.query}
                onChange={(newValue) => this.search(newValue)}
                onRequestSearch={(newValue) => this.search(newValue)}
              />
            </div>
            <NewGroup auth={this.state.auth} categoryId={this.props.location.state.categoryId} />
          </Row>
        </Container>
        <Container style={{ padding: "3.5rem 0" }} maxWidth="md">
          <Grid container spacing={4}>
            {this.state.data !== undefined && this.state.data.length < 1
              ? <Typography gutterBottom variant="h5" component="h4">No groups matching the search query were found. Try creating one!</Typography>
              : null}
            {this.state.data != undefined && this.state.data.map((card) => (
              <Grid item key={card.groupId} xs={12} sm={6} md={4}>
                <Card style={{ height: '100%', display: 'flex', flexDirection: 'column', padding: "20px" }}>
                  <CardContent style={{ flexGrow: 1, padding: "20px" }}>
                    <Typography style={{ marginBottom: "25px" }} gutterBottom variant="h4" component="h4">
                      {card.groupName}</Typography>
                    <Typography style={{ marginBottom: "10px" }}>
                      {card.groupDescription.length > 74 ? card.groupDescription.substring(0, 74) + "..." : card.groupDescription}</Typography>
                  </CardContent>
                  <CardActions>
                    <Button><Link to={{
                      pathname: '/group/' + card.groupId,
                      state: {
                        auth: this.state.auth,
                        groupId: card.groupId
                      }
                    }}>View </Link></Button>
                    {this.state.auth !== '' ?
                      <Button size="small" onClick={() => this.onSave(card)}>
                        {card.isSaved === true ? "Unsave" : "Save"}</Button> : null}
                  </CardActions>
                </Card>
              </Grid>
            ))}
          </Grid>
        </Container>
      </div>
      );
    }

    return null;
  }

  loadMore = () => {
    if (window.innerHeight + document.documentElement.scrollTop === document.scrollingElement.scrollHeight && this.state.moreDataToLoad) {
      this.searchGroups(this.state.auth, this.props.location.state.categoryId, this.state.pagesShown + 1, this.state.query)
    }
  }

  createGroup = (auth, categoryId, groupName, groupDescription) => {
    var body =
    {
      "category": {
        "categoryId": categoryId
      },
      "groupName": groupName,
      "groupDescription": groupDescription
    }

    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups"
      fetch(url, {
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        },
        body: body
      })
        .then((response) => {
          if (response.status <= 201) {
            response.json().then((data) => {
              console.log(data)
            })
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  getGroup = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId
      fetch(url, {
        method: 'get',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        }
      })
        .then((response) => {
          if (response.status <= 201) {
            response.json().then((data) => {
              console.log(data)
            })
          } else {
            console.log("failed :(")
          }
        })
    }, 0)
  }

  saveGroup = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/save"
      fetch(url, {
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        }
      })
        .then((response) => {
          if (response.status <= 201) {
            console.log("success")
            this.updateSavedState(groupId)
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  unsaveGroup = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/save"
      fetch(url, {
        method: 'delete',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        }
      })
        .then((response) => {
          if (response.status <= 201) {
            console.log("success")
            this.updateSavedState(groupId)
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  updateSavedState = (groupId) => {
    var groups = this.state.data
    for (var i = 0; i < groups.length; i++) {
      if (groups[i].groupId === groupId) {
        groups[i].isSaved = !groups[i].isSaved
        break
      }
    }

    this.setState({
      data: groups
    })
  }

  searchGroups = (auth, categoryId, page, query) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/search?"
      if (categoryId !== '') {
        url = url + "category=" + categoryId
      }

      if (page !== '') {
        if (url.charAt(url.length - 1) !== '?' || url.charAt(url.length - 1) !== '&') {
          url = url + '&'
        }
        url = url + "page=" + page
      }

      if (query !== '') {
        if (url.charAt(url.length - 1) !== '?' || url.charAt(url.length - 1) !== '&') {
          url = url + '&'
        }
        url = url + "query=" + query
      }

      fetch(url, {
        method: 'get',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        }
      })
        .then((response) => {
          if (response.status <= 201) {
            response.json().then((data) => {
              console.log("SEARCH", data)
              if (!this.state.data || this.state.search) {
                this.setState({
                  data: data
                })
              } else {
                var newData = this.state.data.concat(data)
                if (data < 1) {
                  this.setState({
                    moreDataToLoad: false
                  })
                } else {
                  this.setState({
                    data: newData,
                    pagesShown: this.state.pagesShown + 1
                  })
                }
              }

            })
          } else {
            console.log("failed :(")
          }
        })
    }, 0)
  }
}