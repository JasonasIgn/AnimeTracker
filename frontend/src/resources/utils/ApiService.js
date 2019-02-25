import config from "../../config";

export function request(path, type = "GET", data = null) {
  const apiPath = `${config.apiUrl}/${path}`;
  return fetch(apiPath, {
    method: type,
    data: data
  });
}
