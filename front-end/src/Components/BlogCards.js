import { React, Component } from "react";
import {
  Typography,
  Grid,
  Container,
  Button,
  Card,
  CardContent,
  CardMedia,
  Hidden,
  Dialog,
  DialogContent
} from "@material-ui/core";
import { Link } from 'react-router-dom'
import RemoveCircleOutlineIcon from '@material-ui/icons/RemoveCircleOutline';
import { Row, Col } from "react-bootstrap";
import IconButton from '@material-ui/core/IconButton';
import { GetCookie, toJSDate, timeSince } from "../UtilityFunctions";
import { NewBlog } from "./NewBlog";


// shows all blog posts as cards
export class BlogCards extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showBlog: false,
      showModal: false,
      groupCommentId: '',
      blogId: ''
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    this.getBlogPosts(auth, this.props.groupId)
    console.log(this.props.data)
  }

  getBlogPosts = (auth, groupId) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog"

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
              this.setState({
                blogs: data
              })
              console.log(data)
            })
          } else {
            console.log("failed :(")
          }
        })
    }, 0)
  }

  handleDelete(blogId) {
    this.setState({
      showModal: true,
      blogId: blogId
    })
  }

  clickSubmitHandler() {
    this.deleteBlogPost(this.props.auth, this.props.groupId, this.state.blogId)
    this.setState({
      showModal: false,
      blogId: ''
    })
  }


  // loads the group info about the group
  // gets group name as prop from Groups
  // shows edit and accept options if user is admin
  render() {
    if (this.state.blogs) {
      return (
        <div>
          {<Dialog
            open={this.state.showModal}
            onClose={() => this.setState({ showModal: false })}
            aria-labelledby="Delete Blog Post"
            aria-describedby="simple-modal-description"
          >
            <DialogContent>
              <Container maxWidth="lg">
                <Typography component="h5" align="center" variant="h5" color="textPrimary" gutterBottom>
                  Are you sure you want to delete this blog post?
              </Typography>
                <Button style={{ marginLeft: "auto" }} id="delete" variant="dark" size="sm" onClick={() => this.clickSubmitHandler()}>Confirm</Button>
              </Container>
            </DialogContent>
          </Dialog >}
          <Container style={{ padding: "1rem 0" }} maxWidth="md">
            <NewBlog
              groupId={this.props.groupId}
              data={this.props.data}
            />
            {this.state.blogs.length > 0 ?
              <Grid container spacing={4}>
                {this.state.blogs.map((blogPost) => (
                  <Grid item xs={12} md={12}>
                    <Card style={{ display: "flex", padding: "5px" }}>
                      <Hidden xsDown>
                        <div style={{ width: 160, backgroundColor: "#EE9D94", borderRadius: "5px" }}>
                        </div>
                      </Hidden>
                      <div style={{ flex: 1 }}>
                        <CardContent>
                          <Row>
                            <Col xs={10}>
                              <Typography component="h3" variant="h5">
                                {blogPost.postTitle}
                              </Typography>
                              <Typography style={{ margin: "2px 0 10px 0" }} variant="subtitle1" color="textSecondary">
                                Created by {blogPost.user.firstName} {blogPost.user.lastName} <time class="timeago" dateTime={toJSDate(blogPost.createdAt)} title={toJSDate(blogPost.createdAt)}>{timeSince(toJSDate(blogPost.createdAt))}</time> ago
                            </Typography>
                              <Typography variant="subtitle1" paragraph>
                                {blogPost.postContent.substring(0, 115) + "..."}
                              </Typography>
                              <Button ><Link to={{
                                pathname: '/group/' + this.props.groupId + '/blog/' + blogPost.blogPostId,
                                state: {
                                  groupId: this.props.groupId,
                                  data: this.props.data,
                                  blogPost: blogPost,
                                }
                              }}>Continue Reading...</Link>
                              </Button>
                            </Col>
                            <Col xs={2}>
                              {/* if isAdmin or wrote the comment? Simplifief to if admin */}
                              {this.props.isAdmin
                                ? <IconButton onClick={() => this.handleDelete(blogPost.blogPostId)} aria-label="delete" variant="contained" size="small" style={{ padding: "0", marginRight: "0px", marginLeft: "60px" }}><RemoveCircleOutlineIcon style={{ fontSize: "15px" }} /></IconButton>
                                : null}
                            </Col></Row>
                        </CardContent>
                      </div>
                    </Card>
                  </Grid>
                ))}
              </Grid> : <div style={{ padding: "40px", marginLeft: "10px", marginTop: "15px", marginBottom: "10px" }}><bold>No blog posts yet...</bold></div>}
          </Container>
        </div>
      );
    }
    return null;
  }

  deleteBlogPost = (auth, groupId, blogId) => {
    setTimeout(() => {
      var url =
        "https://groups.cahillaw.me/v1/groups/" + groupId + "/blog/" + blogId;
      fetch(url, {
        method: "delete",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth,
        },
      }).then((response) => {
        if (response.status <= 201) {
          console.log("success");
          this.removeFromBlogCards(blogId)
        } else {
          console.log("failed :(", response.status);
        }
      });
    }, 0);
  }

  removeFromBlogCards = (blogPostId) => {
    var blogs = this.state.blogs
    for (var i = 0; i < blogs.length; i++) {
      if (blogs[i].blogPostId === blogPostId) {
        blogs.splice(i, 1)
        break
      }
    }

    this.setState({
      blogs: blogs
    })
  }
}
