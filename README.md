[![Go Tests](https://github.com/schufeli/minute-word/actions/workflows/test.yml/badge.svg)](https://github.com/schufeli/minute-word/actions/workflows/test.yml)
[![Docker Image CI](https://github.com/schufeli/minute-word/actions/workflows/docker-image.yml/badge.svg)](https://github.com/schufeli/minute-word/actions/workflows/docker-image.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/schufeli/minute-word/badge)](https://www.codefactor.io/repository/github/schufeli/minute-word)
![License](https://img.shields.io/github/license/schufeli/minute-word?label=License)

# Word of the Minute - Microservice

During my service in the Swiss Armed Forces, some friends and I played a game that we called "Wort der Minute" (Word of the Minute). One would open a German-language dictionary on any random site and read the first word with its description to the rest of the group. As we didn't have a dictionary on hand all the time, I decided to write a small microservice for that. Minute-word is a Go microservice that retrieves a random word from the German language. The response contains a word with its grammar, an explanation, and some usage examples.

## ⚠️ Disclaimer

*Please be aware that this microservice crawls and extracts data from the website https://dwds.de/. Before using this service, it is your responsibility to obtain the necessary permissions from the owner of the dwds.de website. Failure to do so may result in legal action against you. I will not be held liable for any unauthorized use of this service. By using this service, you agree to assume full responsibility for obtaining the necessary permissions and to indemnify me from any legal repercussions that may result from your use of the service.*

## 📦 Deployment

There is another Docker image available for you to use for easier deployment. Below you will find the `Docker Command`, and here you can even find a pre-written [Docker-Compose](docker-compose.yml) file. Please keep in mind that you have to create a `.env` file , as detailed below. 

*The service will listen on port 8000, but feel free to map it to your liking or use case.*

### Environment File

The first and most important thing is to create a new .env file and mount it to your container (how will be explained later). The file is needed to load the API_KEY, which you can compose as you like, into the environment. I strongly recommend that you use some sort of randomly generated string and then hash it with MD5 or SHA1. Example of a .env file: [.env.example](.env.example)

```shell
API_KEY=<Your API_KEY>
```

### Docker Command

Run the following command in the directory containing your `docker-compose.yml` and `.env` file. 

```shell
 docker run --env-file .env <-p "8000:8000"> schufeli/minute-word
```

### Docker Compose

To run the container using [docker-compose](https://docs.docker.com/compose/install/), run the following command in the directory containing your `docker-compose.yml` and `.env` file. 

```shell
docker compose up <-d>
```

## 🚀 How to use

### Requesting a word

When you have the service up and running, you can start using it. For simplicity, I will use `http://localhost:8000` in the following examples. Be aware that this may differ from your deployment (port and URL).

Before you request a random word from the service, you will need to specify the header for your API key: `X-API-KEY`

### Request Header

| Name      | Value                                            |
| --------- | ------------------------------------------------ |
| X-API-KEY | The api key is defined in your environment file. |

After you set the above-explained header, you can send a normal get request to your endpoint url.

### Response

Below you will find an example of a JSON response from the minute-word service.

```json
{
    "word": "Hydrosphäre, die",
    "grammar": "Substantiv (Femininum) · Genitiv Singular: Hydrosphäre · Nominativ Plural: Hydrosphären",
    "explanation": "Wasserhülle der Erde",
    "example": "die Erforschung der Hydrosphäre durch Tauchboote",
    "usages": [
        "Denn Atmosphäre und Hydrosphäre hängen in sehr komplexer Art zusammen.",
        "Eine Beeinflussung durch die Bewegungen der Hydrosphäre aber ist wahrscheinlich.",
        "Falsch sei eine Beseitigung durch Versenkung in Ozeanen angesichts der möglichen Vergiftung der Hydrosphäre.",
        "Die Schnee‑ und Eisregionen der Erde bilden neben der Atmosphäre, der Hydrosphäre (Wasserhülle) und Lithosphäre (Gesteinshülle) ein eigenes System, nämlich das der Kryosphäre.",
        "Es verfügt über Teilsysteme namens Atmosphäre, Hydrosphäre, Geo‑, Kryo‑ und Biosphäre, die alle hochkomplex sind und zudem in zeitlich und räumlich variierender Wechselwirkung stehen."
    ]
}
```

**Note: Sometimes the grammar, explanation, example, or usage in the response can be empty. This means you will need to deal with this in your consuming client example.**


