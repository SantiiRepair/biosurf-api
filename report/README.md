## Report

This is a simple service that checks if a given text contains any obscene words. It's built with Go and uses the Gin framework.


### Usage

To check if a given text contains any obscene words, send a POST request to the `/report` endpoint with the text parameter in the request body. Here's an example using the `typescript` command:

```typescript
import axios from 'axios'

interface Response {
  message: string
  obscene: boolean
}

const message = async (text: string): Promise<void> => {
  try {
    const res = await axios.post<Response>('http://localhost:8080/report', { text }) 
    if (res.data.obscene) {
      console.log(res.data.message)
    } else {
      console.log(res.data.message)
    }
  } catch (error) {
    console.error(error)
  }
}
```

- Response

```typescript
await message('Este tio es un gilipollas') 
// Output: The text contains an obscene word
```

Or

```typescript
await message('Este tio es mi amigo') 
// Output: The text does not contain obscene words
```

This will return a JSON response indicating whether or not the text contains any obscene words:

```json
{
  "response": {
    "message": "The text contains an obscene word",
    "obscene": true
  }
}
```

If the text does not contain any obscene words, the response will look like this:

```json
{
  "response": {
    "message": "The text does not contain obscene words",
    "obscene": false
  }
}
```

### Customization

If you want to customize the list of obscene words that the service checks for, you can update the `wordsList` variable in the `bad_words.go` file. 

```go
var wordsList = []string{
	"puta",
	"coño",
	"joder",
	"cabrón",
	"maricón",
	"gilipollas",
	"hijo de puta",
	"zorra",
	"mierda",
}
```

### Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or a pull request on GitHub. 