import { React, Component } from 'react';
import { Typography, Paper, Container, Button, TextareaAutosize } from '@material-ui/core';
import Alert from '@material-ui/lab/Alert';
import { Row, Col } from "react-bootstrap";
import { GetCookie, toJSDate, timeSince } from "../UtilityFunctions";
import BlogComments from "./BlogComments";

export class BlogPost extends Component {
  constructor(props) {
    super(props);
    this.state = {
      newComment: '',
      commentData: '',
      moreResults: true,
      page: 1,
      showSuccess: false
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    if (this.props.location.state) {
      this.getBlogComments(auth, this.props.location.state.data.groupId, this.props.location.state.blogPost.blogPostId, 1)
      this.setState({
        auth: auth,
        blogPost: this.props.location.state.blogPost,
        groupData: this.props.location.state.data
      })
    } else {
      var groupId = this.props.location.pathname.split("/", 5)[2];
      var blogId = this.props.location.pathname.split("/", 5)[4];
      this.getGroup(auth, groupId)
      this.getBlogPost(auth, groupId, blogId)
      this.getBlogComments(auth, groupId, blogId, 1)
    }
  }

  handleNewBlogCommentChange = (event) => {
    this.setState({
      newComment: event.target.value
    })
  };

  clickSubmitHandler() {
    if (this.state.newComment.length < 1) {
      this.setState({
        showError: true,
        errorMessage: 'Empty comment'
      })
    } else if (this.state.newComment.length > 100) {
      this.setState({
        showError: true,
        errorMessage: 'Comment must be under 100 characters'
      })
    } else {
      this.createBlogComment(this.state.auth, this.state.groupData.groupId, this.state.blogPost.blogPostId, this.state.newComment, 0)
    }
  }

  showSuccessHandler = () => {
    this.setState({
      showSuccess: true
    })
  }

  removeAlert() {
    this.setState({
      showError: false,
      errorMessage: "",
    });
  }

  removeSuccessAlert() {
    this.setState({ showSuccess: false });
  }


  render() {
    const ErrorAlert = () => {
      if (this.state.showError === true) {
        return (
          <Alert style={{ float: "right" }} severity="error" onClose={() => this.removeAlert()} dismissible>
            {this.state.errorMessage}
          </Alert>
        )
      } else {
        return null
      }
    }

    const SuccessAlert = () => {
      if (this.state.showSuccess === true) {
        return (
          <Alert
            severity="success"
            onClose={() => this.removeSuccessAlert()}
            dismissible="true"
          >
            Comment posted!
          </Alert>
        );
      } else {
        return null;
      }
    };

    if (this.state.blogPost && this.state.groupData) {
      return (<div>
        <Container maxWidth="md">
          <Typography component="h2" variant="h2" align="center" color="textPrimary" gutterBottom>
            {this.state.blogPost.postTitle}
          </Typography>
          < hr style={{ marginTop: "-1rem", backgroundColor: "#3399FF", width: "300px", height: "3px" }} />
          <Paper variant="elevation" style={{ padding: "5px", marginBottom: "20px" }}>
            <div style={{ width: "100%", marginBottom: "10px" }} >
              <div style={{
                minHeight: "350px", maxHeight: "450px", minWidth: "50%", borderRadius: "5px",
                maxWidth: "600px", marginLeft: "auto", marginRight: "auto", display: "block", backgroundColor: "#A09BCC"
              }}
              /></div>
            <Typography style={{ marginBottom: "25px" }} variant="subtitle1" align="center">
              Created by {this.state.blogPost.user.firstName} {this.state.blogPost.user.lastName} <time class="timeago" dateTime={toJSDate(this.state.blogPost.createdAt)} title={toJSDate(this.state.blogPost.createdAt)}>{timeSince(toJSDate(this.state.blogPost.createdAt))}</time> ago
              </Typography>
            <Typography variant="subtitle1" color="textPrimary">
              {this.state.blogPost.postContent}
            </Typography>
            <Row>
              <Col>
                <br></br>
                {this.state.auth !== '' ?
                  <div>
                    <Typography variant="subtitle1" color="textPrimary">
                      Leave a comment
                </Typography>
                    <TextareaAutosize style={{ width: "100%" }} label="top level comment" rowsMin={3} onChange={this.handleNewBlogCommentChange} />
                    <Button style={{ marginBottom: "15px" }} size="medium" color="primary" onClick={() => this.clickSubmitHandler()}>Create Comment</Button>
                    <ErrorAlert></ErrorAlert>
                    <SuccessAlert></SuccessAlert>
                    <hr
                      style={{
                        marginTop: "-10px",
                      }}
                    />
                  </div> : null}
                <Typography
                  component="h6"
                  variant="h6"
                  align="left"
                  color="textSecondary"
                >
                  Comments
                </Typography>
                <hr
                  style={{
                    marginTop: "5px",
                    width: "100%",
                    marginBottom: "2px"
                  }}
                />
              </Col>
            </Row>
            <Row>
              <Col>
                {this.state.commentData !== '' ? <BlogComments isAdmin={this.state.groupData.isAdmin} blogPost={this.state.blogPost} auth={this.state.auth} groupId={this.state.groupData.groupId} commentData={this.state.commentData} moreResults={this.state.moreResults} getBlogComments={this.getBlogComments} showSuccessHandler={this.showSuccessHandler} /> : null}
                {this.state.moreResults && this.state.commentData.length > 2 ? <Button size="large" color="primary" onClick={() => this.getBlogComments(this.state.auth, this.state.groupData.groupId, this.state.blogPost.blogPostId, this.state.page + 1)}>
                  Show more comments</Button> : null}
              </Col>
            </Row>
          </Paper>
        </Container>
      </div>
      );
    } else if (this.state.notFound) {
      return <p style={{ textAlign: "center", fontSize: "40px" }}>404 Page Not Found</p>
    }

    return null;
  }

  createBlogComment = (auth, groupId, blogId, commentContent, replyId) => {
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

      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog/" + blogId + "/comments"
      console.log(url)
      fetch(url, {
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': auth
        },
        body: JSON.stringify(body)
      })
        .then((response) => {
          if (response.status <= 201) {
            response.json().then((data) => {
              console.log(data)
              var commentData = this.state.commentData
              commentData.unshift(data)
              this.setState({
                commentData: commentData,
                showSuccess: true
              })
            })
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  getBlogComments = (auth, groupId, blogId, page) => {
    setTimeout(() => {
      console.log(page)
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog/" + blogId + "/comments?";
      if (page !== '') {
        url = url + "page=" + page;
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
              console.log(data)
              if (page === 1) {
                this.setState({
                  commentData: data,
                });
              } else {
                var newData = this.state.commentData.concat(data)
                console.log(newData.length)
                if (newData.length % 3 !== 0 || data.length === 0) {
                  this.setState({
                    moreResults: false,
                    commentData: newData,
                    page: this.state.page + 1
                  })
                } else {
                  this.setState({
                    commentData: newData,
                    page: this.state.page + 1
                  })
                }
              }
            });
          } else {
            console.log("failed :(")
          }
        })
    }, 0)
  }

  getGroup = (auth, groupId) => {
    setTimeout(() => {
      console.log(auth);
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId;
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
              groupData: data,
              auth: auth
            });
          });
        } else {
          console.log("failed :(");
        }
      });
    }, 0);
  };

  getBlogPost = (auth, groupId, blogId) => {
    setTimeout(() => {
      console.log(auth);
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + '/blog/' + blogId
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
              blogPost: data,
              auth: auth
            });
          });
        } else {
          console.log("failed :(");
          this.setState({
            notFound: true
          })
        }
      });
    }, 0);
  }
}



