# Capture-OTP CLI

This CLI application is designed to capture an image from the webcam, perform OCR (Optical Character Recognition) on the captured image, and extract any numeric values with a length of four or more digits from the recognized text. The application utilizes the Go programming language and two external libraries: "gocv" for webcam access and image processing, and "gosseract" for OCR functionality.

## Usage

1. First, make sure you have Go installed on your system.

2. Install the required external libraries using `go get`:

```
go get github.com/otiai10/gosseract/v2
go get -u -d gocv.io/x/gocv
cd $GOPATH/src/gocv.io/x/gocv
make install
```

3. Build the application using `go build`:

```
go build -o capture-otp
```

4. Run the application from the command line:

```
./capture-otp [-show-window|-sw]
```

Optional flags:

- `-show-window`, `-sw`: Display a window showing the live video stream from the webcam. Press the "Space" key to capture an image.

## Build Information

The application uses two external Go packages: "gocv" for webcam access and image processing, and "gosseract" for OCR functionality. The build instructions for these libraries are as follows:

1. **gosseract**:

   - Repository: [github.com/otiai10/gosseract](https://github.com/otiai10/gosseract)
   - License: MIT License

2. **gocv**:

   - Repository: [gocv.io/x/gocv](https://gocv.io)
   - License: BSD 3-Clause License

Make sure to comply with the respective licenses when using these packages.

## License

This Capture-OTP CLI Application is distributed under the MIT License. Feel free to use, modify, and distribute it as per the terms of the MIT License. Please see the `LICENSE` file for more details.
