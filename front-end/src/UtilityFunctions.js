//from: https://www.w3schools.com/js/js_cookies.asp
export function GetCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) === ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) === 0) {
        return c.substring(name.length, c.length);
      }
    }
    return ""
}

//https://stackoverflow.com/questions/8847109/formatting-the-date-time-with-javascript
export function toJSDate (dateTime) {
  if(dateTime === '0001-01-01T00:00:00Z') {
    var m = new Date() 
    m.setHours(m.getHours() + 8);
    m.setSeconds(m.getSeconds() - 1);
    return m;
  }

  dateTime = dateTime.replace('Z', ' ').replace('T', ' ')

  dateTime = dateTime.split(" ");//dateTime[0] = date, dateTime[1] = time
  
  var date = dateTime[0].split("-");
  var time = dateTime[1].split(":");
  
  //(year, month, day, hours, minutes, seconds, milliseconds)
  // mont is 0 indexed so date[1] - 1 corrected format
  return new Date(date[0], date[1]-1, date[2], time[0], time[1], time[2], 0);
}

//https://stackoverflow.com/questions/3177836/how-to-format-time-since-xxx-e-g-4-minutes-ago-similar-to-stack-exchange-site
export function timeSince(date) {
  var seconds = Math.floor((new Date() - date) / 1000) + 28800;

  var interval = seconds / 31536000;
  if (interval >= 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " year";
    } else {
      return Math.floor(interval) + " years";
    }
  }
  interval = seconds / 2592000;
  if (interval >= 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " month";
    } else {
      return Math.floor(interval) + " months";
    }
  }
  interval = seconds / 86400;
  if (interval >= 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " day";
    } else {
      return Math.floor(interval) + " days";
    }
  }
  interval = seconds / 3600;
  if (interval >= 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " hour";
    } else {
      return Math.floor(interval) + " hours";
    }
  }
  interval = seconds / 60;
  if (interval >= 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " minute";
    } else {
      return Math.floor(interval) + " minutes";
    }
  }

  interval = seconds
  if (Math.floor(interval) === 1) {
    return Math.floor(interval) + " second";
  } else {
    return Math.floor(interval) + " seconds";
  }
}