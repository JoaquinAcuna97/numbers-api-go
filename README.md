# NUMBERS API GO 
API that categorizes and stores numbers based on their divisibility

### Go website :shipit:
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/website.svg)](https://pkg.go.dev/golang.org/x/website)

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 4.0 Attributions license][cc4-by].*


### Objective
> [!TIP]
> Implement an API that categorizes and stores numbers based on their divisibility by
3 and 5, mimicking the logic described below:
* Numbers divisible by 3: Categorized as "Type 1"
* Numbers divisible by 5: Categorized as "Type 2"
* Numbers divisible by both 3 and 5: Categorized as "Type 3"
* Other numbers: Stored as their original value

### Constraints
> [!CAUTION]
> Utilize only one if statement within the core logic. Avoid else, else if, ternary
operators, or switch statements.

### Functional Requirements
> [!NOTE]
>  The API shall store numbers and their classifications in memory or any
other support you consider appropriate.

#### Classifications include:
+ Numbers divisible by 3: Categorized as "Type 1"
+ Numbers divisible by 5: Categorized as "Type 2"
+ Numbers divisible by both 3 and 5: Categorized as "Type 3"
+ Other numbers: Stored as their original value



#### Number Addition:

<details>
 <summary><code>POST</code> <code><b>localhost:8080/numbers/</b></code> <code>The API provides a mechanism to save a number into its collection.</code></summary>

##### Body

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | number    |  required | object (JSON)           | `{"number":25}`                                                       |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`                | `Created successfully`                                              |
> | `400`         | `application/json`                | `"message": "There was an error processing the request!"`           |
</details>

#### Number Retrieval:
<details>
 <summary><code>GET</code> <code><b>localhost:8080/numbers/{number}</b></code> <code>The API allows the retrieval of a specific number's classification. If
the number is not found, the API returns 404</code></summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|
> | number    |  required | object (JSON)   | `/25`  |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"Number": 25,"Category": "Type 2"}`                                       |
> | `404`         | `application/json`                | `"message": "number not found"`                                     |
> | `400`         | `application/json`                | `"message": "Can't convert this to an int!"`                        |
</details>

#### Collection Retrieval:
<details>
 <summary><code>GET</code> <code><b>localhost:8080/numbers</b></code> <code>The API enables the retrieval of all stored numbers and their classifications.</code></summary>

##### Parameters

> | name      |  type     | data type               | description                                                           |
> |-----------|-----------|-------------------------|-----------------------------------------------------------------------|


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`                | `[{"Number": 25,"Category": "Type 2"}]`                             |
</details>


### For developer Folks:





> [!IMPORTANT]
> Key information Developers need to know to run, mantain and evolve this project.
>
> go version go1.22.1
>
> How to run the project:
go mod init
go mod tidy
go get numbers
go run main.go 