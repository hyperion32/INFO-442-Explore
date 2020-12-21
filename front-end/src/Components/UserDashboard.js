import { React, Component } from "react";
import { Typography, Grid, Container, Button, Card, CardActions, CardContent, CardMedia, Dialog, DialogContent } from "@material-ui/core";
import { Link } from "react-router-dom";
import { GetCookie } from "../UtilityFunctions";

// loads info through userInfo function from User.js passed in as props
// userInfo returns an object with the user's joined and saved groups, posts and
// comments they're made, which groups they're an admin of, and their login info
export class UserDashboard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showModal: false,
      showLeaveModal: false
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    this.getAdminGroups(auth);
  }

  getAdminGroups = (auth) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/admin";

      fetch(url, {
        method: "get",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth,
        },
      }).then((response) => {
        if (response.status <= 201) {
          response.json().then((data) => {
            this.setState({
              data: data,
              auth: auth
            });
            console.log(data);
          });
        } else {
          console.log("failed :(");
        }
      });
    }, 0);
  };

  // displays user's joined groups, saved groups, and groups they're admins of
  // also shows option to approve or reject group members for groups user is
  // an admin for.
  render() {
    if (this.state.data) {
      console.log(this.state.data.adminGroups.length)
      return (
        <div>
          {<Dialog
            open={this.state.showModal}
            onClose={() => this.setState({ showModal: false })}
            aria-labelledby="Delete Group"
            aria-describedby="simple-modal-description"
          >
            <DialogContent>
              <Container maxWidth="lg">
                <Typography component="h5" align="center" variant="h5" color="textPrimary" gutterBottom>
                  Are you sure you want to delete this group? All contents, including members and blog posts, will be deleted.
              </Typography>
                <Button style={{ marginLeft: "auto" }} id="delete" variant="dark" size="sm" onClick={() => this.clickSubmitHandler()}>Confirm</Button>
              </Container>
            </DialogContent>
          </Dialog >}
          {<Dialog
            open={this.state.showLeaveModal}
            onClose={() => this.setState({ showLeaveModal: false })}
            aria-labelledby="Leave Group"
            aria-describedby="simple-modal-description"
          >
            <DialogContent>
              <Container maxWidth="lg">
                <Typography component="h5" align="center" variant="h5" color="textPrimary" gutterBottom>
                  Are you sure you want to leave this group?
              </Typography>
                <Button style={{ marginLeft: "auto" }} id="delete" variant="dark" size="sm" onClick={() => this.clickLeaveHandler()}>Confirm</Button>
              </Container>
            </DialogContent>
          </Dialog >}
          <Container maxWidth="lg">
            <Typography
              component="h2"
              align="center"
              variant="h2"
              color="textPrimary"
              gutterBottom
            >
              Dashboard
            </Typography>
            <hr
              style={{
                marginTop: "-1rem",
                backgroundColor: "#3399FF",
                width: "200px",
                height: "3px",
              }}
            />
          </Container>
          <Container style={{ padding: "1.5rem 0" }} maxWidth="md">
            <Typography
              component="h5"
              align="left"
              variant="h5"
              color="textPrimary"
              style={{ paddingBottom: "10px" }}
            >
              Admin Groups
            </Typography>
            {this.state.data.adminGroups.length > 0 ? <Grid container spacing={4}>{this.state.data.adminGroups.map((card) => (
              <Grid item key={card} xs={6} sm={4} md={4}>
                <Card style={{ height: '100%', display: 'flex', flexDirection: 'column', padding: "20px" }}>
                  <CardContent style={{ flexGrow: 1, padding: "20px" }}>
                    <Typography style={{ marginBottom: "25px" }} gutterBottom variant="h4" component="h4">
                      {card.groupName}</Typography>
                    <Typography style={{ marginBottom: "10px" }}>
                      {card.groupDescription.length > 74 ? card.groupDescription.substring(0, 74) + "..." : card.groupDescription}</Typography>
                  </CardContent>
                  <CardActions>
                    <Button>
                      <Link
                        to={{
                          pathname: "/group/" + card.groupId,
                          state: {
                            groupId: card.groupId
                          },
                        }}
                      >
                        View{" "}
                      </Link>
                    </Button>
                    <Button size="medium" onClick={() => this.handleDelete(card.groupId)}>
                      Delete</Button>
                  </CardActions>
                </Card>
              </Grid>
            ))}</Grid> :
              <Grid xs={6} sm={4} md={4}>
                <Container style={{ width: "300%", margin: "auto", padding: "10px" }}>
                  <Typography gutterBottom variant="h5" component="h2" style={{ textAlign: "center", fontSize: "20px" }}>
                    You have 0 created groups.
                      </Typography>
                </Container>
              </Grid>}
          </Container>
          <hr
            style={{
              margin: "1rem auto",
              backgroundColor: "#3399FF",
              width: "550px",
              height: "3px",
            }}
          />
          <Container style={{ padding: "1.5rem 0" }} maxWidth="md">
            <Typography
              component="h5"
              align="left"
              variant="h5"
              color="textPrimary"
              style={{ paddingBottom: "10px" }}
            >
              Joined Groups
            </Typography>
            {/* End hero unit */}
            {this.state.data.joinedGroups.length > 0 ? <Grid container spacing={4}>{this.state.data.joinedGroups.map((card) => (
              <Grid item key={card} xs={6} sm={4} md={4}>
                <Card style={{ height: '100%', display: 'flex', flexDirection: 'column', padding: "20px" }}>
                  <CardContent style={{ flexGrow: 1, padding: "20px" }}>
                    <Typography style={{ marginBottom: "25px" }} gutterBottom variant="h4" component="h4">
                      {card.groupName}</Typography>
                    <Typography style={{ marginBottom: "10px" }}>
                      {card.groupDescription.length > 74 ? card.groupDescription.substring(0, 74) + "..." : card.groupDescription}</Typography>
                  </CardContent>
                  <CardActions>
                    <Button>
                      <Link
                        to={{
                          pathname: "/group/" + card.groupId,
                          state: {
                            groupId: card.groupId
                          },
                        }}
                      >
                        View
                        </Link>
                    </Button>
                    <Button size="medium" onClick={() => this.handleLeave(card.groupId)}>
                      Leave</Button>
                  </CardActions>
                </Card>
              </Grid>
            ))}</Grid> :
              <Grid xs={6} sm={4} md={4}>
                <Container style={{ width: "300%", margin: "auto", padding: "10px" }}>
                  <Typography gutterBottom variant="h5" component="h2" style={{ textAlign: "center", fontSize: "20px" }}>
                    You have 0 joined groups.
                      </Typography>
                </Container>
              </Grid>}
          </Container>
          <hr
            style={{
              margin: "1rem auto",
              backgroundColor: "#3399FF",
              width: "550px",
              height: "3px",
            }}
          />
          <Container style={{ padding: "1.5rem 0" }} maxWidth="md">
            <Typography
              component="h5"
              align="left"
              variant="h5"
              color="textPrimary"
              style={{ paddingBottom: "10px" }}
            >
              Saved Groups
            </Typography>
            {/* End hero unit */}
            {this.state.data.savedGroups.length > 0 ? <Grid container spacing={4}>{this.state.data.savedGroups.map((card) => (
              <Grid item key={card} xs={6} sm={4} md={4}>
                <Card style={{ height: '100%', display: 'flex', flexDirection: 'column', padding: "20px" }}>
                  <CardContent style={{ flexGrow: 1, padding: "20px" }}>
                    <Typography style={{ marginBottom: "25px" }} gutterBottom variant="h4" component="h4">
                      {card.groupName}</Typography>
                    <Typography style={{ marginBottom: "10px" }}>
                      {card.groupDescription.length > 74 ? card.groupDescription.substring(0, 74) + "..." : card.groupDescription}</Typography>
                  </CardContent>
                  <CardActions>
                    <Button>
                      <Link
                        to={{
                          pathname: "/group/" + card.groupId,
                          state: {
                            groupId: card.groupId
                          },
                        }}
                      >
                        View{" "}
                      </Link>
                    </Button>
                  </CardActions>
                </Card>
              </Grid>
            ))}</Grid> :
              <Grid xs={6} sm={4} md={4}>
                <Container style={{ width: "300%", margin: "auto", padding: "10px" }}>
                  <Typography gutterBottom variant="h5" component="h2" style={{ textAlign: "center", fontSize: "20px" }}>
                    You have 0 saved groups.
                      </Typography>
                </Container>
              </Grid>}
          </Container>
        </div>
      );
    }
    return (
      <div>
        <Container maxWidth="lg">
          <Typography
            component="h2"
            align="center"
            variant="h2"
            color="textPrimary"
            gutterBottom
          >
            Dashboard
          </Typography>
          <hr
            style={{
              marginTop: "-1rem",
              backgroundColor: "#3399FF",
              width: "200px",
              height: "3px",
            }}
          />
        </Container>
        <div>Loading...</div>
      </div>
    );
  }

  leaveGroup = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/requests"
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
            this.removeFromJoinedGroups(groupId)
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  removeFromJoinedGroups = (groupId) => {
    var data = this.state.data
    for (var i = 0; i < data.joinedGroups.length; i++) {
      if (data.joinedGroups[i].groupId === groupId) {
        data.joinedGroups.splice(i, 1)
        break
      }
    }

    this.setState({
      data: data
    })
  }

  handleDelete(groupId) {
    this.setState({
      showModal: true,
      groupId: groupId
    })
  }

  clickSubmitHandler() {
    this.deleteGroup(this.state.auth, this.state.groupId)
    this.setState({
      showModal: false,
      groupId: ''
    })
  }

  handleLeave(leaveGroupId) {
    this.setState({
      showLeaveModal: true,
      leaveGroupId: leaveGroupId
    })
  }

  clickLeaveHandler() {
    this.leaveGroup(this.state.auth, this.state.leaveGroupId)
    this.setState({
      showLeaveModal: false,
      leaveGroupId: ''
    })
  }

  deleteGroup = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId
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
            this.removeFromAdminGroups(groupId)
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  removeFromAdminGroups = (groupId) => {
    var data = this.state.data
    for (var i = 0; i < data.adminGroups.length; i++) {
      if (data.adminGroups[i].groupId === groupId) {
        data.adminGroups.splice(i, 1)
        break
      }
    }

    this.setState({
      data: data
    })
  }
}
