import React from 'react'
import {
  Typography,
  Container,
  TextField,
  Grid,
  Select,
  MenuItem,
  InputLabel,
  FormControl,
  Button,
  TextareaAutosize,
  Dialog, DialogTitle,
  DialogContent
} from "@material-ui/core";
import Alert from '@material-ui/lab/Alert';

export class NewGroup extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      auth: '0',
      showError: false,
      errorMessage: '',
      groupName: '',
      groupDesc: '',
      categoryId: this.props.categoryId,
      showModal: false,
      buttonText: "Create Group",
      buttonColor: "default"
    }
  }

  componentDidMount() {
    this.getCategories(this.props.auth, '')
  }

  createGroup = (auth, categoryId, groupName, groupDescription) => {
    console.log(parseInt(categoryId, 10))

    var body =
    {
      "category": {
        "categoryId": parseInt(categoryId, 10)
      },
      "groupName": groupName,
      "groupDescription": groupDescription
    }

    console.log(body)
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/groups"
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
              window.location.href = '/group/' + data.groupId
            })
          } else {
            console.log("failed :(", response.status)
          }
        })
    }, 0)
  }

  getCategories = (auth, query) => {
    setTimeout(() => {
      var url = "https://groups.cahillaw.me/v1/categories?"
      if (query !== '') {
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
          if (response.status === 200) {
            response.json().then((data) => {
              this.setState({
                data: data
              })
            })
          } else {
            console.log("failed :(")
          }
        })
    }, 0)
  }

  handleChange = (event) => {
    console.log(event.target.value)
    this.setState({
      categoryId: event.target.value
    })
  };

  handleGroupNameChange = (event) => {
    console.log(event.target.value)
    this.setState({
      groupName: event.target.value
    })
  };

  handleGroupDescChange = (event) => {
    console.log(event.target.value)
    this.setState({
      groupDesc: event.target.value
    })
  };

  clickSubmitHandler() {
    if (this.state.buttonText == "Create Group") {
      this.setState({
        buttonText: "Confirm",
        buttonColor: "primary"
      })
    } else {
      if (this.state.groupName.length < 1) {
        this.setState({
          showError: true,
          errorMessage: 'Empty group name.'
        })
      } else if (this.state.groupName.length > 32) {
        this.setState({
          showError: true,
          errorMessage: 'Group name must be under 32 characters.'
        })
      } else if (this.state.categoryId < 1) {
        this.setState({
          showError: true,
          errorMessage: 'You must select a category.'
        })
      } else if (this.state.groupDesc.length > 240) {
        this.setState({
          showError: true,
          errorMessage: 'Group description must be under 240 characters.'
        })
      } else if (this.state.groupDesc.length < 1) {
        this.setState({
          showError: true,
          errorMessage: 'Empty group description.'
        })
      } else if (this.state.groupName.length > 1 && this.state.groupDesc.length < 1) {
        this.setState({
          showError: true,
          errorMessage: "Enter a group description to continue."
        })
      }
      else {
        this.createGroup(this.props.auth, this.state.categoryId, this.state.groupName, this.state.groupDesc)
      }
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
          <Alert severity="error" style={{ marginTop: "-22px", marginBottom: "-5px", float: "right" }} onClose={() => this.removeAlert()} dismissible="true">
            {this.state.errorMessage}
          </Alert>
        )
      } else {
        return null
      }
    }

    if (this.props.auth.length > 1 && this.state.data) {
      return (
        <div>
          <Button size="medium" color="primary" onClick={() => this.setState({ showModal: true })}>
            Create Group</Button>
          <Dialog
            open={this.state.showModal}
            onClose={() => this.setState({ showModal: false })}
            aria-labelledby="Create Group"
            aria-describedby="simple-modal-description"
          >
            <DialogContent>
              <Container maxWidth="lg">
                <Typography
                  component="h2"
                  align="center"
                  variant="h2"
                  color="textPrimary"
                  gutterBottom
                >
                  Create Group
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
              <Container style={{ padding: "1.5rem 0", marginTop: "-35px" }} maxWidth="md">
                <Grid container spacing={3}>
                  <Grid item xs={12} md={6}>
                    <TextField required label="Group name" fullWidth onChange={this.handleGroupNameChange} />
                  </Grid>
                  <Grid item xs={12} md={6}>
                    <FormControl style={{ width: "100%" }}>
                      <InputLabel id="demo-simple-select-label">Category</InputLabel>
                      <Select
                        labelId="demo-simple-select-label"
                        id=""
                        defaultValue={this.props.categoryId}
                        onChange={this.handleChange}
                      >
                        {this.state.data.categories.map((card) => (
                          <MenuItem value={card.categoryId}>{card.categoryName}</MenuItem>))}
                      </Select>
                    </FormControl>
                  </Grid>
                  <Grid item xs={12} md={6}>
                    <div style={{ marginBottom: "5px" }}>Group description</div>
                    <TextareaAutosize style={{ width: "100%" }} label="Group description" rowsMin={5} fullWidth onChange={this.handleGroupDescChange} />
                  </Grid>
                  <Grid item xs={12}>
                    <Button style={{ marginTop: "-18px" }} id="create" variant="outlined"
                      color={this.state.buttonColor} size="sm" onClick={() => this.clickSubmitHandler()}>{this.state.buttonText}</Button>
                    <ErrorAlert></ErrorAlert>
                  </Grid>
                </Grid>
              </Container>
            </DialogContent>
          </Dialog >
        </div>
      )
    }

    return null;
  }

}

export default NewGroup