let pages = document.getElementsByClassName("pages");
  
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
    
    function activatePage(num) {
      if (num == 0) num = getCurrentPage();
      leaveCurrentPage();
      pages[getPageIndex(num)].id = "active";
      num = getCurrentPage();
    }
  
    function checkPageNum() {
      const maxPageNum = 20;
      let lastPageNum = getLastPageNum();
      if (maxPageNum < pages.length - 2)  {
        for (let i = 1; i <= pages.length - 2; ++i)
          if (i >= maxPageNum + 1) pages[i].style.display = "none";
            else pages[i].style.display = "block";
      }
      else if (maxPageNum < lastPageNum) {
        for (let i = pages.length - 2; i >= 1; --i) pages[i].innerHTML = maxPageNum - pages.length + 2 + i;
      }
      activatePage(0);
      return maxPageNum;
    }
  
    function changePage(pageEle) {
      checkPageNum();
      activatePage(parseInt(pageEle.innerHTML));
    }
    
    function goToPrevPage() {
      let maxPageNum = checkPageNum();
      let currentPageNum = getCurrentPage();
      if (currentPageNum == 1) return;
      let firstPageNum = getFirstPageNum();
      if (currentPageNum != firstPageNum) activatePage(currentPageNum - 1);
      else {
        for (let i = 1; i <= pages.length - 2; ++i) pages[i].innerHTML = parseInt(pages[i].innerHTML) - 1;
      }
    }
    
    function goToNextPage() {
      let maxPageNum = checkPageNum();
      let currentPageNum = getCurrentPage();
      if (currentPageNum == maxPageNum) return;
      let lastPageNum = getLastPageNum();
      if (currentPageNum != lastPageNum) activatePage(currentPageNum + 1);
      else {
        for (let i = 1; i <= pages.length - 2; ++i) pages[i].innerHTML = parseInt(pages[i].innerHTML) + 1;
      }
    }