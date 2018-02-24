# CHISEL Battle Snake 2018

This is the CHISEL group's engineered snake for the [BattleSnake programming competition](http://battlesnake.io). We chose to write it in Go. 

The game server API documentation can be found at [https://stembolthq.github.io/battle_snake/](https://stembolthq.github.io/battle_snake/). 

![Battlesnake game example](static/snakes.gif "BattleSnake game")


### Running the BattleSnake locally

1) [Fork this repo](https://github.com/alexeyza/battlesnake-go/fork).

2) Clone repo to your development environment:
``` bash
git clone git@github.com:USERNAME/battlesnake-go.git $GOPATH/github.com/USERNAME/battlesnake-go
cd $GOPATH/github.com/USERNAME/battlesnake-go
```

3) Compile the battlesnake-go server.
``` bash
go build
```
This will create a `battlesnake-go` executable.

4) Run the server.
``` bash
./battlesnake-go
```

5) Test the client in your browser: [http://127.0.0.1:9000](http://127.0.0.1:9000). I recommend using [Insomnia](https://insomnia.rest/) or [Postman](https://www.getpostman.com/) for testing.

Example start game request:
``` json
{
  "width": 20,
  "height": 20,
  "game_id": "b1dadee8-a112-4e0e-afa2-2845cd1f21aa"
}
:ok
```

Example move request:
``` json
{
  "you": "2c4d4d70-8cca-48e0-ac9d-03ecafca0c98",
  "width": 2,
  "turn": 0,
  "snakes": [
    {
      "taunt": "git gud",
      "name": "my-snake",
      "id": "2c4d4d70-8cca-48e0-ac9d-03ecafca0c98",
      "health_points": 93,
      "coords": [
        [
          0,
          0
        ],
        [
          0,
          0
        ],
        [
          0,
          0
        ]
      ]
    },
    {
      "taunt": "gotta go fast",
      "name": "other-snake",
      "id": "c35dcf26-7f48-492c-b7b5-94ae78fbc713",
      "health_points": 50,
      "coords": [
        [
          1,
          0
        ],
        [
          1,
          0
        ],
        [
          1,
          0
        ]
      ]
    }
  ],
  "height": 2,
  "game_id": "a2facef2-b031-44ba-a36c-0859c389ef96",
  "food": [
    [
      1,
      1
    ]
  ],
  "dead_snakes": [
    {
      "taunt": "gotta go fast",
      "name": "other-snake",
      "id": "83fdf2b9-c8d0-44f4-acb2-0c506139079e",
      "health_points": 50,
      "coords": [
        [
          5,
          0
        ],
        [
          5,
          0
        ],
        [
          5,
          0
        ]
      ]
    }
  ]
}
:ok
```


### Deploying to Heroku

1) Create a new Go Heroku app using Go buildpack.
``` bash
heroku create [APP_NAME]
```

2) Add a buildpack for Go.
``` bash
heroku buildpacks:set heroku/go
```

3) Push code to Heroku servers. Make sure you have a `vendor/vendor.json`, otherwise Heroku will fail building.
``` bash
git push heroku master
```

4) Open Heroku app in browser.
``` bash
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

5) View/stream server logs.
``` bash
heroku logs --tail
```


### Running Your Own Game Server (With Docker)

1) Install Docker. For Ubuntu follow the instructions described here: https://docs.docker.com/engine/installation/linux/ubuntu/#install-docker/

2) Install the game server (this should also run the game server for you).
``` bash
docker run -it -p 4000:4000 stembolt/battle_snake
```

To stop/start the game server use:
``` bash
docker start vibrant_kowalevski

docker stop vibrant_kowalevski
```
Where `vibrant_kowalevski` is the name of my local game server.

3) Visit http://localhost:4000
NOTE: Docker runs on a virtual LAN so when you add a snake to the game you cannot use `http://localhost:9000`, use your internal IP instead (also remove trailing `/`).


### Acknowledgments

Our code is based on [Alexey's naive attempt](https://github.com/alexeyza/battlesnake-go) at the BattleSnake competition in 2017. [SendWithUs](https://www.sendwithus.com/) and [Stembolt](https://stembolt.com/) have put much effort and created an awesome game server for the competition - Thank you!