function checkSignIn() {
  requestData(serverURL + "APIs?id=getUser", null, responseLoad, "GET");
}

function responseLoad(responseText) {
  alert(responseText);
}