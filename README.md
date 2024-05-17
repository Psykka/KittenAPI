# KittenAPI

This is a simple API that returns a random kitten image. It is built using Go and the Gin framework.

## Endpoints

The API has the following endpoints:
- `/` - status endpoint with total number of images
- `/img` - random image
- `/gif` - random gif
- `/mp4` - random mp4

## Running the API

To run the API, you need to have Go installed on your machine. You can run the following command to start the API:

```sh
go run main.go
```

This will start the API on port 8080. You can access the API by going to `http://localhost:8080`.

## Contributing

### Images

The images are stored in the `images` directory. If you would like to add more images to the API, feel free to add them to this directory.
The images should be in the with the following format: `<number>.<ext>` where `<number>` is a unique number between 1 and 10000 and `<ext>` is the file extension (e.g. jpg, gif, mp4).

### Code

If you would like to contribute to this project, feel free to fork the repository and submit a pull request. I will review the changes and merge them if they are appropriate.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
