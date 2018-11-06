## Motivation

Here we will lean to upload multi part data to disk in golang


* [uploadfile.go](https://github.com/saiumesh535/chi-http/blob/master/utils/uploadfiles.go) has Method which helps in uploading files. **do not use this example for uploading large file**

## Example
curl -X POST \
  http://localhost:4321/upload/somefile \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: 76970671-0355-41a2-bc5a-cb1e019180c3' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F 'resume=@C:\Users\SaiUmesh\Downloads\build_14007_step_120_container_0.txt'