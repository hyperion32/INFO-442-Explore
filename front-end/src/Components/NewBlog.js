import React from "react";
import { GetCookie } from "../UtilityFunctions";
import {
  Typography,
  Container,
  Checkbox,
  TextField,
  FormControlLabel,
  Grid,
  Select,
  MenuItem,
  InputLabel,
  FormControl,
  Button,
  TextareaAutosize,
  Dialog,
  DialogTitle,
  DialogContent,
} from "@material-ui/core";
import Alert from "@material-ui/lab/Alert";

export class NewBlog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      auth: "0",
      showError: false,
      errorMessage: "",
      postTitle: "",
      postContent: "",
      blogPost: '',
      buttonText: "Create Post",
      buttonColor: "default"
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    this.setState({
      auth: auth,
    });
  }

  handlePostTitleChange = (event) => {
    this.setState({
      postTitle: event.target.value,
    });
  };

  handlePostContentChange = (event) => {
    this.setState({
      postContent: event.target.value,
    });
  };

  clickSubmitHandler() {
    if (this.state.buttonText == "Create Post") {
      this.setState({
        buttonText: "Confirm",
        buttonColor: "primary"
      })
    } else {
      if (this.state.postTitle.length < 1) {
        this.setState({
          showError: true,
          errorMessage: "Empty post title.",
        });
      } else if (this.state.postTitle.length > 50) {
        this.setState({
          showError: true,
          errorMessage: "Post title must be under 50 characters.",
        });
      } else if (this.state.postContent.length > 1000) {
        this.setState({
          showError: true,
          errorMessage: "Post must be under 1000 characters.",
        });
      } else {
        this.createBlogPost(
          this.state.auth,
          this.props.groupId,
          this.state.postTitle,
          this.state.postContent
        );
        console.log(this.props.groupId)
        console.log(this.props.data.isJoined)
        console.log(this.props.data.isAdmin)
        this.setState({
          showModal: false,
          showSuccess: true,
        });
        console.log("check post is created");
      }
    }
  }

  removeAlert() {
    this.setState({
      showError: false,
      errorMessage: "",
    });
  }

  render = () => {
    const ErrorAlert = () => {
      if (this.state.showError === true) {
        return (
          <Alert
            severity="error"
            onClose={() => this.removeAlert()}
            dismissible
          >
            {this.state.errorMessage}
          </Alert>
        );
      } else {
        return null;
      }
    };

    if (this.props.data != '') {
      return (
        <div>
          {this.props.data.joinedStatus === "Joined" || this.props.data.isAdmin ? (
            <div>
              <Button style={{ left: "45%" }}
                size="medium"
                color="primary"
                onClick={() => this.setState({ showModal: true })}
              >
                Create New Post
              </Button>
              {/* <SuccessAlert></SuccessAlert> */}
              <Dialog
                open={this.state.showModal}
                onClose={() => this.setState({ showModal: false, buttonText: "Create Post", buttonColor: "default" })}
                aria-labelledby="Create Post"
                aria-describedby="simple-modal-description"
              >
                {" "}
                <DialogTitle id="form-dialog-title">Create Post</DialogTitle>
                <DialogContent>
                  <Container maxWidth="lg">
                    <Typography
                      component="h3"
                      align="center"
                      variant="h3"
                      color="textPrimary"
                      gutterBottom
                    >
                      Create Post
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
                      <TextField
                        required
                        label="Post Title"
                        fullWidth
                        onChange={this.handlePostTitleChange}
                      />
                      <div style={{ margin: "10px 0px 3px 0px" }}>
                        Post Content
                      </div>
                      <TextareaAutosize
                        style={{ width: "100%" }}
                        label="Post Content"
                        rowsMin={5}
                        fullWidth
                        onChange={this.handlePostContentChange}
                      />
                      <div style={{ width: "100%", margin: "10px auto 0px auto" }}>
                        <Button
                          id="create"
                          variant="outlined"
                          color={this.state.buttonColor}
                          size="sm"
                          onClick={() => this.clickSubmitHandler()}
                        >
                          {this.state.buttonText}
                        </Button>
                      </div>
                      <hr></hr>
                      <ErrorAlert></ErrorAlert>
                    </Grid>
                  </Container>
                </DialogContent>
              </Dialog>
            </div>
          ) : null}
        </div>
      );
    } else {
      return null
    }
    //s
  };

  createBlogPost = (auth, groupId, postTitle, postContent) => {
    var body = {
      postTitle: postTitle,
      postContent: postContent,
    };

    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog";
      fetch(url, {
        method: "post",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth,
        },
        body: JSON.stringify(body),
      }).then((response) => {
        if (response.status <= 201) {
          response.json().then((data) => {
            console.log(data);
            window.location.href = '/group/' + groupId + '/blog/' + data.blogPostId
          });
        } else {
          console.log("failed :(", response.status);
        }
      });
    }, 0);
  };
}
