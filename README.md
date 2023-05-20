### Report

This is a simple web service that checks if a given text contains any obscene words. It's built with Go and uses the Gin web framework.


### Usage

To check if a given text contains any obscene words, send a POST request to the `/report` endpoint with the text parameter in the request body. Here's an example using the `curl` command:

```
curl -X POST -d "text=puta" http://localhost:8080/report
```

This will return a JSON response indicating whether or not the text contains any obscene words:

```
{"message":"The text contains an obscene word"}
```

If the text does not contain any obscene words, the response will look like this:

```
{"message":"The text does not contain obscene words"}
```

### Customization

If you want to customize the list of obscene words that the service checks for, you can update the `wordsList` variable in the `main.go` file. 

```
var wordsList = []string{"puta", "coño", "joder", "cabrón", "maricón", "gilipollas", "hijo de puta", "zorra", "mierda"}
```

### Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or a pull request on GitHub. 

### License

This project is licensed under the MIT License - see the LICENSE file for details.