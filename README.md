## Introduction

폴더 단위로 구성된 image files를 pdf로 변환하는 앱

### Examples

- input/
    - a/
        - 1.png
        - 2.png
    - b/
        - 3.jpg
        - 4.png

위와 같은 구조의 image files를 아래와 같은 pdf files로 변환한다.

- input/
    - ...
- output/
    - a.pdf
    - b.pdf


## Getting Started

### Build

```sh
$ go build
```

build가 완료되면 `pdfconv.exe` 파일이 생성됨.

### Run App

```sh
$ ./pdfconv.exe <input_dir_path> <output_dir_path>
```

이때 `input_dir_path`의 directory 구조는 위 `Examples` 처럼 2-level layer로 구성되어 있어야 함.

