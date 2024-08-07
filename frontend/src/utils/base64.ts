export const isBase64UrlImage = async (base64String: string): Promise<boolean> => {
  let image = new Image();
  image.src = base64String;
  return await (new Promise((resolve) => {
    image.onload = function() {
      if (image.height === 0 || image.width === 0) {
        resolve(false);
        return;
      }
      resolve(true);
    }
    image.onerror = () => {
      resolve(false);
    }
  }))
}