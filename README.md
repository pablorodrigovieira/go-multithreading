# Go Multithreading CEP API Challenge

## 📌 Environment Setup

To run this project, create a `.env` file in the **server folder** with the following content:

```env
BRASIL_API_URL=
VIACEP_API_URL=
```

## 📌 Running the Server

1. Navigate to the **server folder**:

```bash
go run main.go
```

The server will start on http://localhost:8000
Endpoint available: /cep/{cep}

## 📌 Description

In this challenge, you will apply what we’ve learned about **multithreading** and **APIs** to fetch the fastest response between two different APIs.

### ✅ Requirements

- Make two requests **simultaneously** to the following APIs (using a `cep` value):
  - `https://brasilapi.com.br/api/cep/v1/01153000 + cep`
  - `http://viacep.com.br/ws/" + cep + "/json/`
- Accept the API that returns the **fastest response** and discard the slower one.
- Display the result on the **command line** with the address data and which API provided it.
- Limit the response time to **1 second**. If exceeded, show a **timeout error**.
