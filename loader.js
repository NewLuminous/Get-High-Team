function checkSignIn() {
  requestData(serverURL + "APIs?id=getUser", null, responseLoad, "GET");
}

function responseLoad(responseText) {
  //alert(responseText);
  let signIn = document.getElementById("signin");
  let signUp = document.getElementById("signup");
  let signOut = document.getElementById("signout");
  let welcome = document.getElementById("welcome");
  if (responseText.length > 0) {
    signIn.style.display = "none";
    signUp.style.display = "none";
    signOut.style.display = "block";
    welcome.style.display = "block";
    welcome.innerHTML = "Welcome " + responseText;
  }
  else {
    signIn.style.display = "block";
    signUp.style.display = "block";
    signOut.style.display = "none";
    welcome.style.display = "none";
  }
}