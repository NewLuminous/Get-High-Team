const serverURL = "http://192.168.43.186:1234/";

function requestData(url, dbParam, callback, httpMethod = "POST") {
  let xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      callback(this.responseText);  
    }
  };
  xmlhttp.open(httpMethod, url, true);
  if (httpMethod == "POST") {
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xmlhttp.send(dbParam);
  }
  else if (httpMethod == "GET") {
    xmlhttp.send();
    //alert(url);
  }
}