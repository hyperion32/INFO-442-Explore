import React from 'react'
import { GetCookie } from '../GetCookie'
import {
  Typography,
  Container,
  Checkbox,
  TextField,
  Grid,
  Button,
  TextareaAutosize,
  Dialog, DialogTitle,
  DialogContent
} from "@material-ui/core";
import Alert from '@material-ui/lab/Alert';
import { Redirect } from "react-router-dom";

export class NewComment extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      auth: '0',
      showError: false,
      errorMessage: '',
      commentContent: '',
      groupId: this.props.groupId,

    }
  }

  componentDidMount() {
    var auth = GetCookie("access_token")
    this.setState({
      auth: auth
    })
  }

  handleContentChange = (event) => {
    console.log(event.target.value)
    this.setState({
      commentContent: event.target.value
    })
  };


  clickSubmitHandler() {
    if (this.state.commentContent.length < 1) {
      this.setState({
        showError: true,
        errorMessage: 'Empty comment'
      })
    } else if (this.state.commentContent.length > 200) {
      this.setState({
        showError: true,
        errorMessage: 'comment must be under 200 characters'
      })
    } else {
      this.createGroupComment(this.state.auth, this.state.groupId, this.state.commentContent, 1)
    }
  }

  removeAlert() {
    this.setState({
      showError: false,
      errorMessage: ''
    })
  }

  render = () => {
    const ErrorAlert = () => {
      if (this.state.showError === true) {
        return (
          <Alert severity="error" onClose={() => this.removeAlert()} dismissible>
            {this.state.errorMessage}
          </Alert>
        )
      } else {
        return null
      }
    }

    if (this.state.auth.length > 1) {
      return (
        <div>
          <Button size="medium" color="primary" onClick={() => this.setState({ showModal: true })}>
            Create Comment</Button>
          <Dialog
            open={this.state.showModal}
            onClose={() => this.setState({ showModal: false })}
            aria-labelledby="Create Group"
            aria-describedby="simple-modal-description"
          > <DialogTitle id="form-dialog-title">Create Comment</DialogTitle>
            <DialogContent>
              <Container maxWidth="lg">
                <Typography
                  component="h2"
                  align="center"
                  variant="h2"
                  color="textPrimary"
                  gutterBottom
                >
                  Create Comment
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
                <Grid container spacing={3}>
                  <TextField required label="Comment content" fullWidth onChange={this.handleContentChange} />
                  <Grid item xs={12}>
                    <Button id="create" variant="dark" size="sm" onClick={() => this.clickSubmitHandler()}>Post Comment</Button>
                    <hr></hr>
                    <ErrorAlert></ErrorAlert>
                  </Grid>
                </Grid>
              </Container>
            </DialogContent>
          </Dialog >
        </div>
      )
    }

    if (this.state.auth === '') {
      return <Redirect to={{
        pathname: '/'
      }} />
    }

    return (
      <div>
        <h1 id="title">Loading...</h1>
      </div>
    )
  }

  createGroupComment = (auth, groupId, commentContent, replyId) => {
    setTimeout(() => {
      if (replyId > 0) {
        var body =
        {
          "replyId": {
            "Int64": replyId,
            "Valid": true
          },
          "commentContent": commentContent
        }
      } else {
        var body =
        {
          "commentContent": commentContent
        }
      }

      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/comments"
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

}

export default NewComment;