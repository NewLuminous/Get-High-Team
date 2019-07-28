function getInfo(info) {
  let posLeft = location.href.indexOf(info);
  if (posLeft < 0) return "1";
  let posRight = location.href.indexOf("&", posLeft);
  if (posRight < 0) return location.href.substring(posLeft + info.length + 1);
  return location.href.subtring(posLeft + info.length + 1, posRight);
}