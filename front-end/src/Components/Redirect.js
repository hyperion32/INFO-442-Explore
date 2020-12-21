import React from 'react'
import { Redirect } from 'react-router-dom'

class RedirectPage extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      access_token: '0',
    }
  }

  componentDidMount() {
    var url = this.props.location.search
    var res = url.split("=", 3)
    var res2 = res[1].split("&")
    if (res2[0]) {
      this.setCookie("access_token", res2[0], (1 / 24))
    }

    this.setState({
      access_token: res2[0]
    })
  }

  render = () => {
    if (this.state.access_token !== '0') {
      return <Redirect to={{
        pathname: '/',
        state: {
          access_token: this.state.access_token,
        }
      }} />
    }

    return (
      null
    )
  }

  //from: https://www.w3schools.com/js/js_cookies.asp
  setCookie(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    var expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
  }

}

export default RedirectPage