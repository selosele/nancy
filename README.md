# nancy

Cloudinary 파일 API를 간편하게 활용하는 애플리케이션

## 작업기간
- 2023.09.30. ~ 2023.10.07.

## 기술스택 및 개발환경

- `Go` - 1.21.1
- `WSL2` - Ubuntu 22.04 LTS

## etc.

- 파일 URL 형식: `https://res.cloudinary.com/<cloud_name>/<resource_type>/<delivery_type>[/<transformations>]/<version>/<public_id>`
  - `resource_type` 예) image, video
  - `delivery_type` 예) upload
  - `version`       예) v1684887827
  - `public_id`     예) y2ofmteyrfkvg8jldraz.jpg
