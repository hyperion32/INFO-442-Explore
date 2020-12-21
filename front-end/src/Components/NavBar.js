import { React, Component } from 'react'
import { Navbar, Nav, Button, Container } from 'react-bootstrap';
import { NavLink } from 'react-router-dom';
import { GetCookie } from '../UtilityFunctions'
import '../css/NavBar.css'
import Logo from '../images/logo.png'

export class NavBar extends Component {
  constructor(props) {
    super(props)

    this.state = {
      loggedIn: false
    }
    this.loggedInterval = ""
  }

  componentDidMount() {
    window.onload = this.updateLoggedState;
    document.onmousemove = this.updateLoggedState;
  }

  updateLoggedState = () => {
    var auth = GetCookie("access_token")
    if (auth !== '') {
      this.setState({
        loggedIn: true
      })
    } else {
      this.setState({
        loggedIn: false
      })
    }
  }

  logOut = () => {
    document.cookie = "access_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    this.setState({
      loggedIn: false
    })

    //change once hosted
    if (window.location.href != 'https://romantic-ride-acb817.netlify.app/') {
      window.location.href = '/'
    }
  }

  render() {
    return (
      <Navbar className="color-nav" variant="light">
        <Container className="navcontainer">

          <Navbar.Brand href='/'><img height="50" src={Logo} /></Navbar.Brand>

          <Nav className="ml-auto">
            <Nav.Link><NavLink to="/">Home</NavLink></Nav.Link>
            {this.state.loggedIn ? <Nav.Link><NavLink to="/dashboard">My Groups</NavLink></Nav.Link> : null}
            {this.state.loggedIn ?
              <Button id="logout" variant="link" onClick={() => this.logOut()}>Logout</Button> :
              <Nav.Link><NavLink to="/authenticate">Log In With UW Gmail</NavLink></Nav.Link>
            }
          </Nav>
        </Container>
      </Navbar>
    );

  }
}