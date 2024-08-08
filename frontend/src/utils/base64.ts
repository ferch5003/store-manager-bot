export const isBase64UrlImage =  (base64String: string): boolean => {
  // Regular expression to check if a string is Base64-encoded and is an image
  const base64Regex = /^data:image\/(png|jpg|jpeg|gif|bmp|webp);base64,/;

  if (!base64Regex.test(base64String)) {
    return false;
  }

  const base64Data = base64String.replace(base64Regex, '');

  try {
    atob(base64Data);
    return true;
  } catch (e) {
    return false;
  }
}