# Family Tree

Simple app to manage people and their kinships and build their Family Tree.

The intention behind this project is to further understand a few things about building web applications in Go, like how to manage routing without third party packages and to test out different database management styles.

## Dev

### Setup

simply run on your terminal
```sh
make setup
```

### Run

run either

```sh
make dev
```

or

```sh
wtc
```

to Run the hotreload version of the project which then runs on port 3000

## Docker version

simply
```sh
docker compose up --build
```
and the api will be available on port 3000

## Basic Api Usage

Endpoints:
- <details>
  <summary>person</summary>
  `localhos:3000/person`

  ### Methods
   - GET: Fetches a person by Id or people Name like the following
     - `localhost:3000/person?id=1` returns
`
[
    {
        "id": 1,
        "name": "Arthur"
    }
]
`
     - `localhost:3000/person?name=Arthur` returns
`
[
    {
        "id": 1,
        "name": "Arthur"
    },
    {
        "id": 3,
        "name": "Arthur"
    }
]
`

    - POST: Registers a new person given a json containing the person's name, like so:
      - `localhos:3000/person` with `{"name": "Arthur"}` as body returns `{"id": 1, "name": "Arthur"}
    - DELETE: Deletes a person given an id, like so:
      - `localhost:3000/person?id=1` returns status 204 NoContent
    - PUT: Edits a person's Name given the request body and the id provided in the URL:
      - `localhost:3000/person?id=1` with `{"name": "Arthur"}` as body returns `{"id": 1, "name": "Arthur"}` having updated the object.
</details>

- <details>
  <summary>kinship</summary>
  `localhos:3000/kinship`

  ### Methods
   - GET: Fetches a person's Kinships by their id:
     - `localhost:3000/kinship?id=1` responds with:
       - `[
	{
		"parent": 2,
		"child": 1
	},
	{
		"parent": 3,
		"child": 1
	}
]`

    - POST: Registers a new kinship given the ids provided in the body
      - `localhos:3000/kinship` with `{"parent": 2,"child": 1}` returns 204 No Content if the kinship is allowed
    - DELETE: Deletes a kinship given the ids provided in the body
      - `localhost:3000/kinship` with `{"parent": 2,"child": 1}` as body returns status 204 NoContent if the kinship exists

</details>

- <details>
  <summary>tree</summary>
  `localhos:3000/kinship`

  ### Methods
   - GET: Fetches a person's direct ascendants and descendants by their id:
     - `localhost:3000/tree?id=1` responds with:
       - `[
	{
		"ID": 1,
		"name": "Arthur",
		"relationships": [
			{
				"person_id": 2,
				"name": "Geraldo",
				"kind": "parent"
			},
			{
				"person_id": 3,
				"name": "Valeria",
				"kind": "parent"
			}
		]
	},
	{
		"ID": 2,
		"name": "Geraldo",
		"relationships": []
	},
	{
		"ID": 3,
		"name": "Valeria",
		"relationships": []
	}
]`

</details>

The API does not allow for incestuous kinships to be registered, so if the parent of a child shares a parent with the child's other parent the kinship isn't registered, likewise if the child shares a child with the parent's other children.

## Future Steps

 - Implement a simple SQL database using the package [sqlx](https://pkg.go.dev/github.com/jmoiron/sqlx) as a means to test it and see how it differs from [gorp](https://pkg.go.dev/github.com/go-gorp/gorp/v3)
   - Includes full refactor of the repository part of the stack.
 - Unit testing, made easy with generated code and mockgen (currently not really feasable for the repositories).
 - Further exploring relationships in the graph structure created in the API (seeking cousins and other types of familial relationships).