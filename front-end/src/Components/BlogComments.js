import React, { useState } from "react";
import { Typography, Container, Button, Dialog, DialogContent, CardContent, TextareaAutosize } from "@material-ui/core";
import { toJSDate, timeSince } from "../UtilityFunctions";
import RemoveCircleOutlineIcon from '@material-ui/icons/RemoveCircleOutline';
import IconButton from '@material-ui/core/IconButton';
import ReplyIcon from '@material-ui/icons/Reply';
import { Row, Col } from "react-bootstrap";
import Alert from '@material-ui/lab/Alert';

export class BlogComments extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: "",
      showModal: false
    };
  }

  handleDelete(blogCommentId) {
    this.setState({
      showModal: true,
      blogCommentId: blogCommentId
    })
  }

  clickSubmitHandler() {
    this.deleteBlogComment(this.props.auth, this.props.groupId, this.props.blogPost.blogPostId, this.state.blogCommentId)
    this.setState({
      showModal: false,
      blogCommentId: ''
    })
    setTimeout(() => {
      this.props.getBlogComments(this.props.auth, this.props.groupId, this.props.blogPost.blogPostId, 1)
    }, 1000)
  }

  render() {
    const Comment = (props) => {
      const [showReply, setShowReply] = useState(false);
      const [replyComment, setReplyComment] = useState('');
      const [showError, setShowError] = useState(false);

      const handleReplySubmit = () => {
        if (replyComment.length < 1) {
          setShowError(true)
        } else if (replyComment.length > 100) {
          setShowError(true)
        } else {
          createBlogComment(this.props.auth, this.props.groupId, this.props.blogPost.blogPostId, replyComment, props.card.blogCommentId)
        }
      }

      const ErrorAlert = () => {
        if (showError === true) {
          return (
            <Alert style={{ float: "right" }} severity="error" onClose={() => setShowError(false)} dismissible="true">
              Comment length must be between 1 and 100
            </Alert>
          )
        } else {
          return null
        }
      }

      const createBlogComment = (auth, groupId, blogId, commentContent, replyId) => {
        setTimeout(() => {
          var body
          if (replyId > 0) {
            body =
            {
              "replyId": {
                "Int64": replyId,
                "Valid": true
              },
              "commentContent": commentContent
            }
          } else {
            body =
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
                  setShowReply(false)
                  this.props.getBlogComments(this.props.auth, this.props.groupId, this.props.blogPost.blogPostId, 1)
                  this.props.showSuccessHandler()
                })
              } else {
                console.log("failed :(", response.status)
              }
            })
        }, 0)
      }

      return (
        <div>
          <CardContent key={props.card.blogCommentId}
            style={{ padding: "5px", paddingLeft: props.indent, borderBottom: "0.5px solid #ebebeb" }}
          >
            <Row>
              <Col xs={10}>
                {!props.card.deleted ? <Typography
                  component="p"
                  variant="p"
                  style={{ fontSize: "15px" }}
                >
                  {props.card.commentContent}
                </Typography> : <Typography
                  component="p"
                  variant="p"
                  style={{ fontSize: "15px", color: "#e34949" }}
                >
                    Comment deleted
                  </Typography>}
                <div style={{ float: "left", fontSize: "13px", fontWeight: "500" }}>{props.card.user.firstName} {props.card.user.lastName} </div>
                <div style={{ float: "left", marginLeft: "8px", fontSize: "13px" }}>posted <time class="timeago" dateTime={toJSDate(props.card.createdAt)} title={toJSDate(props.card.createdAt)}>{timeSince(toJSDate(props.card.createdAt))}</time> ago</div>
                <br></br>
              </Col>
              <Col xs={2}>
                {this.props.isAdmin && !props.card.deleted
                  ? <IconButton onClick={() => this.handleDelete(props.card.blogCommentId)} aria-label="delete" variant="contained" size="small" style={{ float: "right", marginTop: "15px" }}><RemoveCircleOutlineIcon style={{ fontSize: "15px" }} /></IconButton>
                  : null}
                {this.props.auth !== '' && !props.card.deleted ? <IconButton onClick={() => setShowReply(true)} style={{ float: "right" }}>
                  <ReplyIcon />
                </IconButton> : null}
              </Col>
            </Row>
            <Row>
              <Col>
                {showReply ? <div>
                  <TextareaAutosize style={{ width: "100%" }} label="top level comment" rowsMin={2} onChange={(event) => setReplyComment(event.target.value)} />
                  <Button style={{ marginBottom: "15px" }} size="small" color="primary" onClick={() => handleReplySubmit()}>Create Comment</Button>
                  <ErrorAlert></ErrorAlert>
                </div> : null}
              </Col>
            </Row>
          </CardContent>
          {props.card.children ?
            <div>
              {props.card.children.map((children) => (
                <Comment card={children} indent={props.indent + 20}></Comment>
              ))}
            </div> : null}
        </div>
      )
    }
    return (
      <div>
        {<Dialog
          open={this.state.showModal}
          onClose={() => this.setState({ showModal: false })}
          aria-labelledby="Delete Comment"
          aria-describedby="simple-modal-description"
        >
          <DialogContent>
            <Container maxWidth="lg">
              <Typography component="h5" align="center" variant="h5" color="textPrimary" gutterBottom>
                Are you sure you want to delete this comment?
              </Typography>
              <Button style={{ marginLeft: "auto" }} id="delete" variant="dark" size="sm" onClick={() => this.clickSubmitHandler()}>Confirm</Button>
            </Container>
          </DialogContent>
        </Dialog >}
        {this.props.commentData && this.props.commentData.length > 0 ? (
          <div>
            {this.props.commentData.map((card) => (
              <Comment card={card} getGroupComments={this.props.getGroupComments} showSuccessHandler={this.props.showSuccessHandler} indent={5}></Comment>
            ))}
          </div>
        ) : (
            <Container
              style={{
                Width: "100%",
                Height: "20x",
              }}
              maxWidth="md"
            >
              <bold>No comments yet...</bold>
            </Container>
          )}
      </div>
    );
  }

  deleteBlogComment = (auth, groupId, blogId, commentId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog/" + blogId + "/comments/" + commentId;
      fetch(url, {
        method: "delete",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth,
        },
      }).then((response) => {
        if (response.status <= 201) {
          console.log("success");
          this.setAsDeleted(commentId);
        } else {
          console.log("failed :(", response.status);
        }
      });
    }, 0);
  };

  setAsDeleted = (blogCommentId) => {
    var commentData = this.props.commentData
    for (var i = 0; i < commentData.length; i++) {
      if (commentData[i].blogCommentId === blogCommentId) {
        commentData[i].deleted = true
        break
      }
    }

    this.setState({
      commentData: commentData
    })
  }

}

export default BlogComments;
