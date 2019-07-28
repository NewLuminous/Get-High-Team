let pages = document.getElementsByClassName("pages");
let previews = document.getElementsByClassName("hostelPreview");
let maxPageNum = -1;
  
function getFirstPageNum() {
  return parseInt(pages[1].innerHTML);
}
    
    function getLastPageNum() {
      for (let i = pages.length - 2; i >= 1; --i)
        if (pages[i].style.display != "none") return parseInt(pages[i].innerHTML);
    }
    
    function getPageIndex(num) {
      for (let i = 1; i <= pages.length - 2; ++i)
        if (pages[i].innerHTML == num) {
          return i;
        }
      let firstPageNum = getFirstPageNum();
      if (num < firstPageNum) return 1;
      let lastPageNum = getLastPageNum();
      if (num > lastPageNum) return pages.length - 2;
    }
    
    function getCurrentPage() {
      let currentPage = document.getElementById("active");
      return parseInt(currentPage.innerHTML);
    }
    
    function leaveCurrentPage() {
      let currentPage = document.getElementById("active");
      currentPage.id = "";
      return parseInt(currentPage.innerHTML);
    }

function getInfo(info) {
  let posLeft = location.href.indexOf(info);
  if (posLeft < 0) return "1";
  let posRight = location.href.indexOf("&", posLeft);
  if (posRight < 0) return location.href.substring(posLeft + info.length + 1);
  return location.href.subtring(posLeft + info.length + 1, posRight);
}
    
function activatePage(num) {
  if (num == 0) num = getCurrentPage();
  leaveCurrentPage();
  pages[getPageIndex(num)].id = "active";
  num = getCurrentPage();
  let pageObj = {
    From: (num - 1) * 10,
    To: num * 10 - 1
  };
  let pageRequest = getInfo("page");
  let oldNum = parseInt(pageRequest);
  if (oldNum != num) location.href = location.pathname + "?page=" + num; 
  requestData(serverURL + "APIs/getIndexPost", JSON.stringify(pageObj), responsePage);
}

function responsePage(responseText) {
  let hostels = JSON.parse(responseText);
  for (let i = 0; i < hostels.Data.length; ++i) {
    let contents = previews[i].getElementsByTagName("*");
    //contents[0].innerHTML = hostels.Data[i].Image;
    contents[1].innerHTML = hostels.Data[i].Title;
    contents[2].innerHTML = hostels.Data[i].Price + " VND/month";
    contents[3].innerHTML = "Area: " + hostels.Data[i].Area + " m2";
    contents[4].innerHTML = "Address: " + hostels.Data[i].Address;
    contents[5].innerHTML = "Last update: " + hostels.Data[i].Date;
    //contents[6].innerHTML = hostels.Data[i].Id;
    contents[6].innerHTML = 7;
    previews[i].style.display = "block";
  }
  for (let i = hostels.Data.length; i <= 10; ++i) {
    previews[i].style.display = "none";
  }
}
  
    function checkPageNum() {
      requestData(serverURL + "APIs?id=getNumberOfPosts", null, responsePost, "GET");
    }
    
  function updatePage() {
      let lastPageNum = getLastPageNum();
      if (maxPageNum < pages.length - 2)  {
        for (let i = 1; i <= pages.length - 2; ++i)
          if (i >= maxPageNum + 1) pages[i].style.display = "none";
            else pages[i].style.display = "block";
      }
      else if (maxPageNum < lastPageNum) {
        for (let i = pages.length - 2; i >= 1; --i) pages[i].innerHTML = maxPageNum - pages.length + 2 + i;
      }
      let pageRequest = getInfo("page");
      let oldNum = parseInt(pageRequest);
      activatePage(oldNum);
  }
    
  function responsePost(responseText) {
    let x = parseInt(JSON.parse(responseText).Data);
    maxPageNum = parseInt(x / 10);
    if (x % 10 > 0) ++maxPageNum;
    updatePage();
  }
  
    function changePage(pageEle) {
      checkPageNum();
      activatePage(parseInt(pageEle.innerHTML));
    }
    
    function goToPrevPage() {
      checkPageNum();
      let currentPageNum = getCurrentPage();
      if (currentPageNum == 1) return;
      let firstPageNum = getFirstPageNum();
      if (currentPageNum != firstPageNum) activatePage(currentPageNum - 1);
      else {
        for (let i = 1; i <= pages.length - 2; ++i) pages[i].innerHTML = parseInt(pages[i].innerHTML) - 1;
      }
    }
    
    function goToNextPage() {
      checkPageNum();
      let currentPageNum = getCurrentPage();
      if (currentPageNum == maxPageNum) return;
      let lastPageNum = getLastPageNum();
      if (currentPageNum != lastPageNum) activatePage(currentPageNum + 1);
      else {
        for (let i = 1; i <= pages.length - 2; ++i) pages[i].innerHTML = parseInt(pages[i].innerHTML) + 1;
      }
    }