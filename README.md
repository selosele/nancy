# nancy

Cloudinary 파일 업로드 API

## 기술스택

- Golang 1.21.1
- WSL2 Ubuntu 22.04 LTS

## TO-DO

- Postman에서 유효한 토큰 없이 호출되면 안됨
- 이미지뿐만 아니라 동영상 파일 등 구분없이 업로드할 수 있도록 개선
  - 파일 URL 형식: `https://res.cloudinary.com/<cloud_name>/<resource_type>/<delivery_type>[/<transformations>]/<version>/<public_id>`
  - `resource_type` 예) image
  - `delivery_type` 예) upload
  - `version`       예) v1684887827
  - `public_id`     예) y2ofmteyrfkvg8jldraz.jpg
