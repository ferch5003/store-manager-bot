const stripHtml = (html: string): string => {
  const div = document.createElement("div");
  div.innerHTML = html;
  return div.textContent || div.innerText || "";
}

export const extractLinks = (html: string): string[] => {
  const text = stripHtml(html);
  const urlPattern = /(https?:\/\/\S+)/g;
  const matches = text.match(urlPattern);
  return matches || [];
}