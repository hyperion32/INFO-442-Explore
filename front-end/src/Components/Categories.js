import { React, Component } from "react";
import {
  Typography,
  Grid,
  Container,
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import SearchBar from "material-ui-search-bar";
import { GetCookie } from "../UtilityFunctions";

export class Categories extends Component {
  constructor(props) {
    super(props);
    this.state = {
      categoryId: 0,
      query: "",
      numMaxResults: 9
    };
  }

  componentDidMount() {
    var auth = GetCookie("access_token");
    this.setState({
      auth: auth,
    });
    this.getCategories(auth, "");
    window.addEventListener('scroll', this.loadMore);

  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this.loadMore);
  }

  render() {
    return (
      <div>
        <Container maxWidth="md">
          <Typography
            component="h2"
            variant="h2"
            align="center"
            color="textPrimary"
            gutterBottom
          >
            Categories
          </Typography>
          <hr
            style={{
              marginTop: "-1rem",
              backgroundColor: "#3399FF",
              width: "200px",
              height: "3px",
            }}
          />
          <SearchBar
            value={this.state.query}
            onChange={(newValue) => this.getCategories(this.state.auth, newValue)}
            onRequestSearch={(newValue) =>
              this.getCategories(this.state.auth, newValue)
            }
          />
        </Container>
        <Container style={{ padding: "3.5rem 0" }} maxWidth="md">
          <Grid container spacing={4}>
            {this.state.data !== undefined && this.state.data.categories.length < 1
              ? <Typography gutterBottom variant="h5" component="h4">No categories matching the search query were found.</Typography>
              : null}
            {this.state.data !== undefined &&
              this.state.data.categories.slice(0, this.state.numMaxResults).map((card) => (

                <Grid item key={card.categoryId} xs={12} sm={6} md={4}>
                  <Card
                    style={{
                      height: "100%",
                      display: "flex",
                      flexDirection: "column",
                    }}
                  >
                    <CardMedia
                      style={{ paddingTop: "40.25%", backgroundColor: this.getColor(card.categoryId) }}
                      // image="https://source.unsplash.com/random"
                      title="Category image"
                    />
                    <CardContent style={{ flexGrow: 1 }}>
                      <Typography gutterBottom variant="h5" component="h2">
                        {card.categoryName}
                      </Typography>
                    </CardContent>
                    <CardActions>
                      <Button>
                        <Link
                          to={{
                            pathname: "/groups",
                            state: {
                              categoryId: card.categoryId,
                              categoryName: card.categoryName,
                            },
                          }}
                        >
                          View{" "}
                        </Link>
                      </Button>
                    </CardActions>
                  </Card>
                </Grid>
              ))}
          </Grid>
          <br></br>
          <Container style={{ display: "flex", flexWrap: "wrap", justifyContent: "center", marginTop: "10px" }}>
            {this.state.data !== undefined && this.state.data.categories.length > 9 && this.state.numMaxResults < this.state.data.categories.length ?
              <Button size="large" color="primary" onClick={() => this.setState({ numMaxResults: this.state.numMaxResults + 9 })}>
                Show more results</Button> : null}
          </Container>
        </Container>
      </div>
    );
  }

  loadMore = () => {
    if (window.innerHeight + document.documentElement.scrollTop === document.scrollingElement.scrollHeight) {
      this.setState({ numMaxResults: this.state.numMaxResults + 9 })
    }
  }

  getCategories = (auth, query) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/categories?";
      if (query !== "") {
        url = url + "query=" + query;
      }
      fetch(url, {
        method: "get",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth,
        },
      }).then((response) => {
        if (response.status === 200) {
          response.json().then((data) => {
            console.log(data);
            this.setState({
              data: data,
              numMaxResults: 9
            });
          });
        } else {
          console.log("failed :(");
        }
      });
    }, 0);
  };

  getColor(categoryId) {
    if (categoryId % 10 == 0) {
      return "#EE9D94"
    } else if (categoryId % 10 == 1) {
      return "#9DBBE3"
    } else if (categoryId % 10 == 2) {
      return "#BFE3DF"
    } else if (categoryId % 10 == 3) {
      return "#A09BCC"
    } else if (categoryId % 10 == 4) {
      return "#FFF5CC"
    } else if (categoryId % 10 == 5) {
      return "#90CCDE"
    } else if (categoryId % 10 == 6) {
      return "teal"
    } else if (categoryId % 10 == 7) {
      return "#D3F1D2"
    } else if (categoryId % 10 == 8) {
      return "#F3C6D1"
    } else {
      return "#FFD9BD"
    }
  }

}
