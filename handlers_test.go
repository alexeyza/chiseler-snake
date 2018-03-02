package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMoveHandler(t *testing.T) {
	test_cases := []struct {
		name                 string
		request_msg          string
		success_returned_msg string // optional: use either success or fail return values
		fail_returned_msg    string // optional
	}{
		{name: "(1) colide with larger snake", request_msg: game_collide_with_larger_snake, fail_returned_msg: "{\"move\":\"right\"}\n"},
		{name: "(2) colide with same-size snake", request_msg: game_collide_with_same_size_snake, fail_returned_msg: "{\"move\":\"down\"}\n"},
		{name: "(3) turn into dead end", request_msg: game_turn_into_dead_end, fail_returned_msg: "{\"move\":\"up\"}\n"},
		{name: "(4) path to tail blocked, can't find a path", request_msg: game_cant_find_path_to_tail, fail_returned_msg: "{\"move\":\"up\"}\n"},
		{name: "(5) path to tail blocked by another smaller snake", request_msg: game_no_path_to_tail, fail_returned_msg: "{\"move\":\"up\"}\n"},
		{name: "(6) killed himself even though near food and smaller snake", request_msg: game_killed_himself_though_food_nearby_and_smaller_enemy_snake, success_returned_msg: "{\"move\":\"right\"}\n"},
		{name: "(7) kill himself when smaller snake blocked the way, even though could eat it", request_msg: game_smaller_snake_blocking_way, success_returned_msg: "{\"move\":\"down\"}\n"},
	}

	for _, test_case := range test_cases {

		req, err := http.NewRequest("POST", "//142.104.68.152:9000/move", strings.NewReader(test_case.request_msg))
		if err != nil {
			t.Fatalf("request failed %v", err)
		}
		rec := httptest.NewRecorder()
		MoveHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected Status OK; got %v", res.Status)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}

		if test_case.fail_returned_msg != "" {
			// checks if the returned value would lead us to death
			if string(b) == test_case.fail_returned_msg {
				t.Errorf(test_case.name)
			}
		}
		if test_case.success_returned_msg != "" {
			// checks if the returned value would lead us to death
			if string(b) != test_case.success_returned_msg {
				t.Errorf(test_case.name)
			}
		}
	}
}

var game_collide_with_larger_snake string = `{
  "width": 20,
  "height": 20,
  "id": 5171,
  "turn": 32,
  "snakes": {
    "data": [
      {
        "id": "9f63bca2-eee9-4394-9d77-35c59146f0e5",
        "health": 96,
        "length": 11,
        "taunt": "\"You Fish!\"",
        "name": "Big Asp",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 12,
              "y": 7
            },
            {
              "object": "point",
              "x": 12,
              "y": 6
            },
            {
              "object": "point",
              "x": 12,
              "y": 5
            },
            {
              "object": "point",
              "x": 12,
              "y": 4
            },
            {
              "object": "point",
              "x": 12,
              "y": 3
            },
            {
              "object": "point",
              "x": 13,
              "y": 3
            },
            {
              "object": "point",
              "x": 13,
              "y": 2
            },
            {
              "object": "point",
              "x": 13,
              "y": 1
            },
            {
              "object": "point",
              "x": 14,
              "y": 1
            },
            {
              "object": "point",
              "x": 14,
              "y": 2
            },
            {
              "object": "point",
              "x": 14,
              "y": 3
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "c3e89743-fbdd-465b-8208-45ea0919b05a",
        "health": 68,
        "length": 3,
        "taunt": "\"<(O.o)>\"",
        "name": "test",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": 17
            },
            {
              "object": "point",
              "x": 2,
              "y": 16
            },
            {
              "object": "point",
              "x": 1,
              "y": 16
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "69c37835-e672-4b8a-9327-47cd204e8411",
        "health": 68,
        "length": 3,
        "taunt": "\"lil pump!\"",
        "name": "UVSD Snake Rev 6",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 16,
              "y": 1
            },
            {
              "object": "point",
              "x": 16,
              "y": 0
            },
            {
              "object": "point",
              "x": 17,
              "y": 0
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 98,
        "length": 9,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 11,
              "y": 8
            },
            {
              "object": "point",
              "x": 11,
              "y": 9
            },
            {
              "object": "point",
              "x": 11,
              "y": 10
            },
            {
              "object": "point",
              "x": 12,
              "y": 10
            },
            {
              "object": "point",
              "x": 13,
              "y": 10
            },
            {
              "object": "point",
              "x": 13,
              "y": 11
            },
            {
              "object": "point",
              "x": 13,
              "y": 12
            },
            {
              "object": "point",
              "x": 13,
              "y": 13
            },
            {
              "object": "point",
              "x": 13,
              "y": 14
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "b36c2a40-5ab0-48cf-b0a7-3cfc8e824324",
        "health": 99,
        "length": 9,
        "taunt": "\"The Sauce is loose!\"",
        "name": "MrSauce",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 6,
              "y": 7
            },
            {
              "object": "point",
              "x": 6,
              "y": 8
            },
            {
              "object": "point",
              "x": 7,
              "y": 8
            },
            {
              "object": "point",
              "x": 7,
              "y": 9
            },
            {
              "object": "point",
              "x": 7,
              "y": 10
            },
            {
              "object": "point",
              "x": 7,
              "y": 11
            },
            {
              "object": "point",
              "x": 7,
              "y": 12
            },
            {
              "object": "point",
              "x": 7,
              "y": 13
            },
            {
              "object": "point",
              "x": 7,
              "y": 14
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 0,
        "y": 3
      },
      {
        "object": "point",
        "x": 0,
        "y": 1
      },
      {
        "object": "point",
        "x": 7,
        "y": 0
      },
      {
        "object": "point",
        "x": 6,
        "y": 4
      },
      {
        "object": "point",
        "x": 1,
        "y": 6
      },
      {
        "object": "point",
        "x": 13,
        "y": 8
      },
      {
        "object": "point",
        "x": 16,
        "y": 17
      },
      {
        "object": "point",
        "x": 18,
        "y": 10
      },
      {
        "object": "point",
        "x": 6,
        "y": 18
      },
      {
        "object": "point",
        "x": 13,
        "y": 0
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 98,
    "length": 9,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 11,
          "y": 8
        },
        {
          "object": "point",
          "x": 11,
          "y": 9
        },
        {
          "object": "point",
          "x": 11,
          "y": 10
        },
        {
          "object": "point",
          "x": 12,
          "y": 10
        },
        {
          "object": "point",
          "x": 13,
          "y": 10
        },
        {
          "object": "point",
          "x": 13,
          "y": 11
        },
        {
          "object": "point",
          "x": 13,
          "y": 12
        },
        {
          "object": "point",
          "x": 13,
          "y": 13
        },
        {
          "object": "point",
          "x": 13,
          "y": 14
        }
      ],
      "object": "list"
    }
  }
}`

var game_collide_with_same_size_snake string = `{
  "width": 14,
  "height": 11,
  "id": 5522,
  "turn": 7,
  "snakes": {
    "data": [
      {
        "id": "64f03f9f-c332-4330-8a1e-8342c8d887c7",
        "health": 93,
        "length": 3,
        "taunt": "\"Boop the snoot!\"",
        "name": "LUL",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 10,
              "y": 4
            },
            {
              "object": "point",
              "x": 10,
              "y": 3
            },
            {
              "object": "point",
              "x": 11,
              "y": 3
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "9f63bca2-eee9-4394-9d77-35c59146f0e5",
        "health": 98,
        "length": 5,
        "taunt": "\"You Fish!\"",
        "name": "Big Asp",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 8,
              "y": 2
            },
            {
              "object": "point",
              "x": 7,
              "y": 2
            },
            {
              "object": "point",
              "x": 6,
              "y": 2
            },
            {
              "object": "point",
              "x": 5,
              "y": 2
            },
            {
              "object": "point",
              "x": 4,
              "y": 2
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "9d152d4b-36c0-462d-a504-d86c20e6785e",
        "health": 100,
        "length": 6,
        "taunt": "\"I Sure Hope This Works Out!\"",
        "name": "Naive Snake",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 3,
              "y": 7
            },
            {
              "object": "point",
              "x": 3,
              "y": 8
            },
            {
              "object": "point",
              "x": 3,
              "y": 9
            },
            {
              "object": "point",
              "x": 4,
              "y": 9
            },
            {
              "object": "point",
              "x": 5,
              "y": 9
            },
            {
              "object": "point",
              "x": 5,
              "y": 9
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 100,
        "length": 6,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 3,
              "y": 5
            },
            {
              "object": "point",
              "x": 3,
              "y": 4
            },
            {
              "object": "point",
              "x": 2,
              "y": 4
            },
            {
              "object": "point",
              "x": 2,
              "y": 3
            },
            {
              "object": "point",
              "x": 2,
              "y": 2
            },
            {
              "object": "point",
              "x": 2,
              "y": 2
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 9,
        "y": 2
      },
      {
        "object": "point",
        "x": 11,
        "y": 2
      },
      {
        "object": "point",
        "x": 4,
        "y": 6
      },
      {
        "object": "point",
        "x": 13,
        "y": 10
      },
      {
        "object": "point",
        "x": 8,
        "y": 4
      },
      {
        "object": "point",
        "x": 9,
        "y": 0
      },
      {
        "object": "point",
        "x": 8,
        "y": 1
      },
      {
        "object": "point",
        "x": 3,
        "y": 6
      },
      {
        "object": "point",
        "x": 0,
        "y": 8
      },
      {
        "object": "point",
        "x": 13,
        "y": 2
      },
      {
        "object": "point",
        "x": 0,
        "y": 10
      },
      {
        "object": "point",
        "x": 12,
        "y": 0
      },
      {
        "object": "point",
        "x": 1,
        "y": 10
      },
      {
        "object": "point",
        "x": 9,
        "y": 5
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 100,
    "length": 6,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 3,
          "y": 5
        },
        {
          "object": "point",
          "x": 3,
          "y": 4
        },
        {
          "object": "point",
          "x": 2,
          "y": 4
        },
        {
          "object": "point",
          "x": 2,
          "y": 3
        },
        {
          "object": "point",
          "x": 2,
          "y": 2
        },
        {
          "object": "point",
          "x": 2,
          "y": 2
        }
      ],
      "object": "list"
    }
  }
}`

var game_turn_into_dead_end string = `{
  "width": 15,
  "height": 19,
  "id": 5498,
  "turn": 91,
  "snakes": {
    "data": [
      {
        "id": "b36c2a40-5ab0-48cf-b0a7-3cfc8e824324",
        "health": 100,
        "length": 26,
        "taunt": "\"The Sauce is loose!\"",
        "name": "MrSauce",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": 14
            },
            {
              "object": "point",
              "x": 2,
              "y": 13
            },
            {
              "object": "point",
              "x": 1,
              "y": 13
            },
            {
              "object": "point",
              "x": 1,
              "y": 12
            },
            {
              "object": "point",
              "x": 2,
              "y": 12
            },
            {
              "object": "point",
              "x": 3,
              "y": 12
            },
            {
              "object": "point",
              "x": 4,
              "y": 12
            },
            {
              "object": "point",
              "x": 5,
              "y": 12
            },
            {
              "object": "point",
              "x": 6,
              "y": 12
            },
            {
              "object": "point",
              "x": 7,
              "y": 12
            },
            {
              "object": "point",
              "x": 8,
              "y": 12
            },
            {
              "object": "point",
              "x": 8,
              "y": 13
            },
            {
              "object": "point",
              "x": 9,
              "y": 13
            },
            {
              "object": "point",
              "x": 10,
              "y": 13
            },
            {
              "object": "point",
              "x": 11,
              "y": 13
            },
            {
              "object": "point",
              "x": 11,
              "y": 12
            },
            {
              "object": "point",
              "x": 11,
              "y": 11
            },
            {
              "object": "point",
              "x": 11,
              "y": 10
            },
            {
              "object": "point",
              "x": 10,
              "y": 10
            },
            {
              "object": "point",
              "x": 9,
              "y": 10
            },
            {
              "object": "point",
              "x": 8,
              "y": 10
            },
            {
              "object": "point",
              "x": 8,
              "y": 9
            },
            {
              "object": "point",
              "x": 9,
              "y": 9
            },
            {
              "object": "point",
              "x": 10,
              "y": 9
            },
            {
              "object": "point",
              "x": 11,
              "y": 9
            },
            {
              "object": "point",
              "x": 11,
              "y": 9
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 100,
        "length": 25,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 14,
              "y": 4
            },
            {
              "object": "point",
              "x": 13,
              "y": 4
            },
            {
              "object": "point",
              "x": 12,
              "y": 4
            },
            {
              "object": "point",
              "x": 11,
              "y": 4
            },
            {
              "object": "point",
              "x": 11,
              "y": 5
            },
            {
              "object": "point",
              "x": 10,
              "y": 5
            },
            {
              "object": "point",
              "x": 10,
              "y": 4
            },
            {
              "object": "point",
              "x": 10,
              "y": 3
            },
            {
              "object": "point",
              "x": 10,
              "y": 2
            },
            {
              "object": "point",
              "x": 10,
              "y": 1
            },
            {
              "object": "point",
              "x": 10,
              "y": 0
            },
            {
              "object": "point",
              "x": 9,
              "y": 0
            },
            {
              "object": "point",
              "x": 8,
              "y": 0
            },
            {
              "object": "point",
              "x": 7,
              "y": 0
            },
            {
              "object": "point",
              "x": 6,
              "y": 0
            },
            {
              "object": "point",
              "x": 5,
              "y": 0
            },
            {
              "object": "point",
              "x": 4,
              "y": 0
            },
            {
              "object": "point",
              "x": 3,
              "y": 0
            },
            {
              "object": "point",
              "x": 2,
              "y": 0
            },
            {
              "object": "point",
              "x": 2,
              "y": 1
            },
            {
              "object": "point",
              "x": 2,
              "y": 2
            },
            {
              "object": "point",
              "x": 2,
              "y": 3
            },
            {
              "object": "point",
              "x": 2,
              "y": 4
            },
            {
              "object": "point",
              "x": 3,
              "y": 4
            },
            {
              "object": "point",
              "x": 3,
              "y": 4
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 6,
        "y": 3
      },
      {
        "object": "point",
        "x": 3,
        "y": 16
      },
      {
        "object": "point",
        "x": 12,
        "y": 0
      },
      {
        "object": "point",
        "x": 2,
        "y": 16
      },
      {
        "object": "point",
        "x": 13,
        "y": 1
      },
      {
        "object": "point",
        "x": 3,
        "y": 2
      },
      {
        "object": "point",
        "x": 5,
        "y": 17
      },
      {
        "object": "point",
        "x": 4,
        "y": 18
      },
      {
        "object": "point",
        "x": 11,
        "y": 0
      },
      {
        "object": "point",
        "x": 13,
        "y": 11
      },
      {
        "object": "point",
        "x": 14,
        "y": 13
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [
      {
        "id": "64f03f9f-c332-4330-8a1e-8342c8d887c7",
        "health": 13,
        "length": 3,
        "taunt": "\"Boop the snoot!\"",
        "name": "LUL",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": 12
            },
            {
              "object": "point",
              "x": 2,
              "y": 11
            },
            {
              "object": "point",
              "x": 2,
              "y": 10
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 100,
    "length": 25,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 14,
          "y": 4
        },
        {
          "object": "point",
          "x": 13,
          "y": 4
        },
        {
          "object": "point",
          "x": 12,
          "y": 4
        },
        {
          "object": "point",
          "x": 11,
          "y": 4
        },
        {
          "object": "point",
          "x": 11,
          "y": 5
        },
        {
          "object": "point",
          "x": 10,
          "y": 5
        },
        {
          "object": "point",
          "x": 10,
          "y": 4
        },
        {
          "object": "point",
          "x": 10,
          "y": 3
        },
        {
          "object": "point",
          "x": 10,
          "y": 2
        },
        {
          "object": "point",
          "x": 10,
          "y": 1
        },
        {
          "object": "point",
          "x": 10,
          "y": 0
        },
        {
          "object": "point",
          "x": 9,
          "y": 0
        },
        {
          "object": "point",
          "x": 8,
          "y": 0
        },
        {
          "object": "point",
          "x": 7,
          "y": 0
        },
        {
          "object": "point",
          "x": 6,
          "y": 0
        },
        {
          "object": "point",
          "x": 5,
          "y": 0
        },
        {
          "object": "point",
          "x": 4,
          "y": 0
        },
        {
          "object": "point",
          "x": 3,
          "y": 0
        },
        {
          "object": "point",
          "x": 2,
          "y": 0
        },
        {
          "object": "point",
          "x": 2,
          "y": 1
        },
        {
          "object": "point",
          "x": 2,
          "y": 2
        },
        {
          "object": "point",
          "x": 2,
          "y": 3
        },
        {
          "object": "point",
          "x": 2,
          "y": 4
        },
        {
          "object": "point",
          "x": 3,
          "y": 4
        },
        {
          "object": "point",
          "x": 3,
          "y": 4
        }
      ],
      "object": "list"
    }
  }
}`

var game_cant_find_path_to_tail string = `{
  "width": 18,
  "height": 7,
  "id": 5479,
  "turn": 38,
  "snakes": {
    "data": [
      {
        "id": "b36c2a40-5ab0-48cf-b0a7-3cfc8e824324",
        "health": 94,
        "length": 10,
        "taunt": "\"The Sauce is loose!\"",
        "name": "MrSauce",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 14,
              "y": 2
            },
            {
              "object": "point",
              "x": 14,
              "y": 3
            },
            {
              "object": "point",
              "x": 14,
              "y": 4
            },
            {
              "object": "point",
              "x": 15,
              "y": 4
            },
            {
              "object": "point",
              "x": 15,
              "y": 5
            },
            {
              "object": "point",
              "x": 16,
              "y": 5
            },
            {
              "object": "point",
              "x": 17,
              "y": 5
            },
            {
              "object": "point",
              "x": 17,
              "y": 4
            },
            {
              "object": "point",
              "x": 16,
              "y": 4
            },
            {
              "object": "point",
              "x": 16,
              "y": 3
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 100,
        "length": 11,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 6,
              "y": 0
            },
            {
              "object": "point",
              "x": 7,
              "y": 0
            },
            {
              "object": "point",
              "x": 8,
              "y": 0
            },
            {
              "object": "point",
              "x": 9,
              "y": 0
            },
            {
              "object": "point",
              "x": 10,
              "y": 0
            },
            {
              "object": "point",
              "x": 10,
              "y": 1
            },
            {
              "object": "point",
              "x": 11,
              "y": 1
            },
            {
              "object": "point",
              "x": 12,
              "y": 1
            },
            {
              "object": "point",
              "x": 13,
              "y": 1
            },
            {
              "object": "point",
              "x": 14,
              "y": 1
            },
            {
              "object": "point",
              "x": 14,
              "y": 1
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 8,
        "y": 3
      },
      {
        "object": "point",
        "x": 16,
        "y": 2
      },
      {
        "object": "point",
        "x": 10,
        "y": 4
      },
      {
        "object": "point",
        "x": 10,
        "y": 5
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 100,
    "length": 11,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 6,
          "y": 0
        },
        {
          "object": "point",
          "x": 7,
          "y": 0
        },
        {
          "object": "point",
          "x": 8,
          "y": 0
        },
        {
          "object": "point",
          "x": 9,
          "y": 0
        },
        {
          "object": "point",
          "x": 10,
          "y": 0
        },
        {
          "object": "point",
          "x": 10,
          "y": 1
        },
        {
          "object": "point",
          "x": 11,
          "y": 1
        },
        {
          "object": "point",
          "x": 12,
          "y": 1
        },
        {
          "object": "point",
          "x": 13,
          "y": 1
        },
        {
          "object": "point",
          "x": 14,
          "y": 1
        },
        {
          "object": "point",
          "x": 14,
          "y": 1
        }
      ],
      "object": "list"
    }
  }
}`

var game_no_path_to_tail string = `{
  "width": 17,
  "height": 19,
  "id": 5468,
  "turn": 41,
  "snakes": {
    "data": [
      {
        "id": "c9db82f2-bdd1-482f-a4da-21bb5789210a",
        "health": 59,
        "length": 3,
        "taunt": "\"Hello\"",
        "name": "lw-testsnake (should not win)",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": 0
            },
            {
              "object": "point",
              "x": 2,
              "y": 1
            },
            {
              "object": "point",
              "x": 1,
              "y": 1
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 89,
        "length": 10,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 9,
              "y": 11
            },
            {
              "object": "point",
              "x": 9,
              "y": 10
            },
            {
              "object": "point",
              "x": 9,
              "y": 9
            },
            {
              "object": "point",
              "x": 9,
              "y": 8
            },
            {
              "object": "point",
              "x": 9,
              "y": 7
            },
            {
              "object": "point",
              "x": 8,
              "y": 7
            },
            {
              "object": "point",
              "x": 8,
              "y": 8
            },
            {
              "object": "point",
              "x": 7,
              "y": 8
            },
            {
              "object": "point",
              "x": 7,
              "y": 9
            },
            {
              "object": "point",
              "x": 7,
              "y": 10
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "2d8fe9b9-d467-4c8a-903e-912d87cb8492",
        "health": 59,
        "length": 3,
        "taunt": "\"SSSNK!\"",
        "name": "sssnk",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 4,
              "y": 10
            },
            {
              "object": "point",
              "x": 4,
              "y": 11
            },
            {
              "object": "point",
              "x": 4,
              "y": 12
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "e0323921-3386-4883-b531-e8ad277b370e",
        "health": 96,
        "length": 9,
        "taunt": "\"python!\"",
        "name": "ds",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 7,
              "y": 11
            },
            {
              "object": "point",
              "x": 6,
              "y": 11
            },
            {
              "object": "point",
              "x": 6,
              "y": 10
            },
            {
              "object": "point",
              "x": 6,
              "y": 9
            },
            {
              "object": "point",
              "x": 6,
              "y": 8
            },
            {
              "object": "point",
              "x": 5,
              "y": 8
            },
            {
              "object": "point",
              "x": 4,
              "y": 8
            },
            {
              "object": "point",
              "x": 4,
              "y": 7
            },
            {
              "object": "point",
              "x": 4,
              "y": 6
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "c3e89743-fbdd-465b-8208-45ea0919b05a",
        "health": 81,
        "length": 4,
        "taunt": "\"<(o,o)>\"",
        "name": "test",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 13,
              "y": 7
            },
            {
              "object": "point",
              "x": 12,
              "y": 7
            },
            {
              "object": "point",
              "x": 12,
              "y": 8
            },
            {
              "object": "point",
              "x": 11,
              "y": 8
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 9,
        "y": 16
      },
      {
        "object": "point",
        "x": 16,
        "y": 10
      },
      {
        "object": "point",
        "x": 16,
        "y": 6
      },
      {
        "object": "point",
        "x": 8,
        "y": 11
      },
      {
        "object": "point",
        "x": 6,
        "y": 15
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [
      {
        "id": "1d442390-4195-4f47-93ce-21a56b250f4b",
        "health": 99,
        "length": 6,
        "taunt": "TIMED OUT",
        "name": "Uter",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 15,
              "y": 1
            },
            {
              "object": "point",
              "x": 15,
              "y": 2
            },
            {
              "object": "point",
              "x": 16,
              "y": 2
            },
            {
              "object": "point",
              "x": 16,
              "y": 3
            },
            {
              "object": "point",
              "x": 15,
              "y": 3
            },
            {
              "object": "point",
              "x": 15,
              "y": 4
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 89,
    "length": 10,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 9,
          "y": 11
        },
        {
          "object": "point",
          "x": 9,
          "y": 10
        },
        {
          "object": "point",
          "x": 9,
          "y": 9
        },
        {
          "object": "point",
          "x": 9,
          "y": 8
        },
        {
          "object": "point",
          "x": 9,
          "y": 7
        },
        {
          "object": "point",
          "x": 8,
          "y": 7
        },
        {
          "object": "point",
          "x": 8,
          "y": 8
        },
        {
          "object": "point",
          "x": 7,
          "y": 8
        },
        {
          "object": "point",
          "x": 7,
          "y": 9
        },
        {
          "object": "point",
          "x": 7,
          "y": 10
        }
      ],
      "object": "list"
    }
  }
}`

//https://play.snakedown.com/app/replay/febe0f5b-b29e-45f9-9809-500261f65de9
var game_killed_himself_though_food_nearby_and_smaller_enemy_snake string = `{
  "width": 21,
  "height": 11,
  "id": 5789,
  "turn": 139,
  "snakes": {
    "data": [
      {
        "id": "f8fa5342-950d-4ab7-9341-2ec14e8afc90",
        "health": 96,
        "length": 10,
        "taunt": "\"noodly noodly\"",
        "name": "danger-noodle",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 14,
              "y": 2
            },
            {
              "object": "point",
              "x": 15,
              "y": 2
            },
            {
              "object": "point",
              "x": 15,
              "y": 1
            },
            {
              "object": "point",
              "x": 16,
              "y": 1
            },
            {
              "object": "point",
              "x": 16,
              "y": 2
            },
            {
              "object": "point",
              "x": 17,
              "y": 2
            },
            {
              "object": "point",
              "x": 18,
              "y": 2
            },
            {
              "object": "point",
              "x": 19,
              "y": 2
            },
            {
              "object": "point",
              "x": 19,
              "y": 3
            },
            {
              "object": "point",
              "x": 19,
              "y": 4
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 94,
        "length": 25,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 13,
              "y": 1
            },
            {
              "object": "point",
              "x": 13,
              "y": 0
            },
            {
              "object": "point",
              "x": 12,
              "y": 0
            },
            {
              "object": "point",
              "x": 12,
              "y": 1
            },
            {
              "object": "point",
              "x": 11,
              "y": 1
            },
            {
              "object": "point",
              "x": 10,
              "y": 1
            },
            {
              "object": "point",
              "x": 10,
              "y": 2
            },
            {
              "object": "point",
              "x": 11,
              "y": 2
            },
            {
              "object": "point",
              "x": 12,
              "y": 2
            },
            {
              "object": "point",
              "x": 12,
              "y": 3
            },
            {
              "object": "point",
              "x": 12,
              "y": 4
            },
            {
              "object": "point",
              "x": 13,
              "y": 4
            },
            {
              "object": "point",
              "x": 13,
              "y": 5
            },
            {
              "object": "point",
              "x": 13,
              "y": 6
            },
            {
              "object": "point",
              "x": 14,
              "y": 6
            },
            {
              "object": "point",
              "x": 14,
              "y": 5
            },
            {
              "object": "point",
              "x": 14,
              "y": 4
            },
            {
              "object": "point",
              "x": 15,
              "y": 4
            },
            {
              "object": "point",
              "x": 15,
              "y": 5
            },
            {
              "object": "point",
              "x": 16,
              "y": 5
            },
            {
              "object": "point",
              "x": 16,
              "y": 4
            },
            {
              "object": "point",
              "x": 16,
              "y": 3
            },
            {
              "object": "point",
              "x": 15,
              "y": 3
            },
            {
              "object": "point",
              "x": 14,
              "y": 3
            },
            {
              "object": "point",
              "x": 13,
              "y": 3
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "1beb8a5d-97d1-49ce-b4df-0769e5ba17a0",
        "health": 92,
        "length": 24,
        "taunt": "\"hiss!\"",
        "name": "0ff3b768-1a83-11e8-92e0-8c85901cffa7",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 14,
              "y": 8
            },
            {
              "object": "point",
              "x": 13,
              "y": 8
            },
            {
              "object": "point",
              "x": 12,
              "y": 8
            },
            {
              "object": "point",
              "x": 12,
              "y": 7
            },
            {
              "object": "point",
              "x": 11,
              "y": 7
            },
            {
              "object": "point",
              "x": 11,
              "y": 6
            },
            {
              "object": "point",
              "x": 10,
              "y": 6
            },
            {
              "object": "point",
              "x": 9,
              "y": 6
            },
            {
              "object": "point",
              "x": 9,
              "y": 7
            },
            {
              "object": "point",
              "x": 8,
              "y": 7
            },
            {
              "object": "point",
              "x": 7,
              "y": 7
            },
            {
              "object": "point",
              "x": 6,
              "y": 7
            },
            {
              "object": "point",
              "x": 5,
              "y": 7
            },
            {
              "object": "point",
              "x": 4,
              "y": 7
            },
            {
              "object": "point",
              "x": 3,
              "y": 7
            },
            {
              "object": "point",
              "x": 2,
              "y": 7
            },
            {
              "object": "point",
              "x": 1,
              "y": 7
            },
            {
              "object": "point",
              "x": 0,
              "y": 7
            },
            {
              "object": "point",
              "x": 0,
              "y": 6
            },
            {
              "object": "point",
              "x": 0,
              "y": 5
            },
            {
              "object": "point",
              "x": 1,
              "y": 5
            },
            {
              "object": "point",
              "x": 1,
              "y": 4
            },
            {
              "object": "point",
              "x": 1,
              "y": 3
            },
            {
              "object": "point",
              "x": 1,
              "y": 2
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 0,
        "y": 8
      },
      {
        "object": "point",
        "x": 2,
        "y": 10
      },
      {
        "object": "point",
        "x": 14,
        "y": 9
      },
      {
        "object": "point",
        "x": 14,
        "y": 1
      },
      {
        "object": "point",
        "x": 11,
        "y": 3
      },
      {
        "object": "point",
        "x": 2,
        "y": 9
      },
      {
        "object": "point",
        "x": 15,
        "y": 6
      },
      {
        "object": "point",
        "x": 4,
        "y": 6
      },
      {
        "object": "point",
        "x": 8,
        "y": 6
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [
      {
        "id": "9f63bca2-eee9-4394-9d77-35c59146f0e5",
        "health": 97,
        "length": 3,
        "taunt": "\"You Fish!\"",
        "name": "Big Asp",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 0,
              "y": 6
            },
            {
              "object": "point",
              "x": 0,
              "y": 5
            },
            {
              "object": "point",
              "x": 0,
              "y": 4
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "9d152d4b-36c0-462d-a504-d86c20e6785e",
        "health": 95,
        "length": 16,
        "taunt": "TIMED OUT",
        "name": "Naive Snake",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": -1
            },
            {
              "object": "point",
              "x": 2,
              "y": 0
            },
            {
              "object": "point",
              "x": 2,
              "y": 1
            },
            {
              "object": "point",
              "x": 3,
              "y": 1
            },
            {
              "object": "point",
              "x": 3,
              "y": 2
            },
            {
              "object": "point",
              "x": 3,
              "y": 3
            },
            {
              "object": "point",
              "x": 3,
              "y": 4
            },
            {
              "object": "point",
              "x": 2,
              "y": 4
            },
            {
              "object": "point",
              "x": 2,
              "y": 5
            },
            {
              "object": "point",
              "x": 1,
              "y": 5
            },
            {
              "object": "point",
              "x": 1,
              "y": 4
            },
            {
              "object": "point",
              "x": 0,
              "y": 4
            },
            {
              "object": "point",
              "x": 0,
              "y": 3
            },
            {
              "object": "point",
              "x": 1,
              "y": 3
            },
            {
              "object": "point",
              "x": 1,
              "y": 2
            },
            {
              "object": "point",
              "x": 1,
              "y": 1
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "d9979736-9e2e-4419-b156-e2f808bbf7c3",
        "health": 93,
        "length": 16,
        "taunt": "TIMED OUT",
        "name": "whelp aggro",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 20,
              "y": 3
            },
            {
              "object": "point",
              "x": 20,
              "y": 4
            },
            {
              "object": "point",
              "x": 20,
              "y": 5
            },
            {
              "object": "point",
              "x": 20,
              "y": 6
            },
            {
              "object": "point",
              "x": 19,
              "y": 6
            },
            {
              "object": "point",
              "x": 19,
              "y": 5
            },
            {
              "object": "point",
              "x": 18,
              "y": 5
            },
            {
              "object": "point",
              "x": 18,
              "y": 4
            },
            {
              "object": "point",
              "x": 19,
              "y": 4
            },
            {
              "object": "point",
              "x": 19,
              "y": 3
            },
            {
              "object": "point",
              "x": 20,
              "y": 3
            },
            {
              "object": "point",
              "x": 20,
              "y": 2
            },
            {
              "object": "point",
              "x": 20,
              "y": 1
            },
            {
              "object": "point",
              "x": 20,
              "y": 0
            },
            {
              "object": "point",
              "x": 19,
              "y": 0
            },
            {
              "object": "point",
              "x": 19,
              "y": 1
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 94,
    "length": 25,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 13,
          "y": 1
        },
        {
          "object": "point",
          "x": 13,
          "y": 0
        },
        {
          "object": "point",
          "x": 12,
          "y": 0
        },
        {
          "object": "point",
          "x": 12,
          "y": 1
        },
        {
          "object": "point",
          "x": 11,
          "y": 1
        },
        {
          "object": "point",
          "x": 10,
          "y": 1
        },
        {
          "object": "point",
          "x": 10,
          "y": 2
        },
        {
          "object": "point",
          "x": 11,
          "y": 2
        },
        {
          "object": "point",
          "x": 12,
          "y": 2
        },
        {
          "object": "point",
          "x": 12,
          "y": 3
        },
        {
          "object": "point",
          "x": 12,
          "y": 4
        },
        {
          "object": "point",
          "x": 13,
          "y": 4
        },
        {
          "object": "point",
          "x": 13,
          "y": 5
        },
        {
          "object": "point",
          "x": 13,
          "y": 6
        },
        {
          "object": "point",
          "x": 14,
          "y": 6
        },
        {
          "object": "point",
          "x": 14,
          "y": 5
        },
        {
          "object": "point",
          "x": 14,
          "y": 4
        },
        {
          "object": "point",
          "x": 15,
          "y": 4
        },
        {
          "object": "point",
          "x": 15,
          "y": 5
        },
        {
          "object": "point",
          "x": 16,
          "y": 5
        },
        {
          "object": "point",
          "x": 16,
          "y": 4
        },
        {
          "object": "point",
          "x": 16,
          "y": 3
        },
        {
          "object": "point",
          "x": 15,
          "y": 3
        },
        {
          "object": "point",
          "x": 14,
          "y": 3
        },
        {
          "object": "point",
          "x": 13,
          "y": 3
        }
      ],
      "object": "list"
    }
  }
}`

var game_smaller_snake_blocking_way string = `{
  "width": 10,
  "height": 14,
  "id": 5814,
  "turn": 24,
  "snakes": {
    "data": [
      {
        "id": "9d152d4b-36c0-462d-a504-d86c20e6785e",
        "health": 99,
        "length": 13,
        "taunt": "\"I Sure Hope This Works Out!\"",
        "name": "Naive Snake",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 2,
              "y": 10
            },
            {
              "object": "point",
              "x": 2,
              "y": 9
            },
            {
              "object": "point",
              "x": 1,
              "y": 9
            },
            {
              "object": "point",
              "x": 0,
              "y": 9
            },
            {
              "object": "point",
              "x": 0,
              "y": 8
            },
            {
              "object": "point",
              "x": 0,
              "y": 7
            },
            {
              "object": "point",
              "x": 1,
              "y": 7
            },
            {
              "object": "point",
              "x": 1,
              "y": 8
            },
            {
              "object": "point",
              "x": 2,
              "y": 8
            },
            {
              "object": "point",
              "x": 3,
              "y": 8
            },
            {
              "object": "point",
              "x": 3,
              "y": 7
            },
            {
              "object": "point",
              "x": 3,
              "y": 6
            },
            {
              "object": "point",
              "x": 3,
              "y": 5
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "f8fa5342-950d-4ab7-9341-2ec14e8afc90",
        "health": 97,
        "length": 7,
        "taunt": "\"noodly noodly\"",
        "name": "danger-noodle",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 7,
              "y": 13
            },
            {
              "object": "point",
              "x": 6,
              "y": 13
            },
            {
              "object": "point",
              "x": 6,
              "y": 12
            },
            {
              "object": "point",
              "x": 5,
              "y": 12
            },
            {
              "object": "point",
              "x": 5,
              "y": 11
            },
            {
              "object": "point",
              "x": 5,
              "y": 10
            },
            {
              "object": "point",
              "x": 4,
              "y": 10
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
        "health": 100,
        "length": 13,
        "taunt": "\"You've just been ERASED!!\"",
        "name": "chiseler",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 9,
              "y": 11
            },
            {
              "object": "point",
              "x": 9,
              "y": 10
            },
            {
              "object": "point",
              "x": 8,
              "y": 10
            },
            {
              "object": "point",
              "x": 8,
              "y": 11
            },
            {
              "object": "point",
              "x": 8,
              "y": 12
            },
            {
              "object": "point",
              "x": 7,
              "y": 12
            },
            {
              "object": "point",
              "x": 7,
              "y": 11
            },
            {
              "object": "point",
              "x": 6,
              "y": 11
            },
            {
              "object": "point",
              "x": 6,
              "y": 10
            },
            {
              "object": "point",
              "x": 7,
              "y": 10
            },
            {
              "object": "point",
              "x": 7,
              "y": 9
            },
            {
              "object": "point",
              "x": 7,
              "y": 8
            },
            {
              "object": "point",
              "x": 7,
              "y": 8
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "food": {
    "data": [
      {
        "object": "point",
        "x": 1,
        "y": 12
      },
      {
        "object": "point",
        "x": 9,
        "y": 8
      },
      {
        "object": "point",
        "x": 1,
        "y": 1
      },
      {
        "object": "point",
        "x": 7,
        "y": 4
      },
      {
        "object": "point",
        "x": 0,
        "y": 4
      },
      {
        "object": "point",
        "x": 2,
        "y": 2
      },
      {
        "object": "point",
        "x": 1,
        "y": 0
      },
      {
        "object": "point",
        "x": 0,
        "y": 1
      },
      {
        "object": "point",
        "x": 1,
        "y": 3
      },
      {
        "object": "point",
        "x": 4,
        "y": 1
      },
      {
        "object": "point",
        "x": 6,
        "y": 8
      },
      {
        "object": "point",
        "x": 6,
        "y": 1
      },
      {
        "object": "point",
        "x": 4,
        "y": 4
      }
    ],
    "object": "list"
  },
  "object": "world",
  "dead_snakes": {
    "data": [
      {
        "id": "1fc1c054-a38e-454a-9a96-79b5e2c43836",
        "health": 97,
        "length": 4,
        "taunt": "\"hello\"",
        "name": "Training Snake 10",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 4,
              "y": 1
            },
            {
              "object": "point",
              "x": 4,
              "y": 2
            },
            {
              "object": "point",
              "x": 4,
              "y": 3
            },
            {
              "object": "point",
              "x": 4,
              "y": 4
            }
          ],
          "object": "list"
        }
      },
      {
        "id": "9f63bca2-eee9-4394-9d77-35c59146f0e5",
        "health": 94,
        "length": 5,
        "taunt": "\"You Fish!\"",
        "name": "Big Asp",
        "object": "snake",
        "body": {
          "data": [
            {
              "object": "point",
              "x": 3,
              "y": 6
            },
            {
              "object": "point",
              "x": 2,
              "y": 6
            },
            {
              "object": "point",
              "x": 2,
              "y": 5
            },
            {
              "object": "point",
              "x": 2,
              "y": 4
            },
            {
              "object": "point",
              "x": 2,
              "y": 3
            }
          ],
          "object": "list"
        }
      }
    ],
    "object": "list"
  },
  "you": {
    "id": "0c75298e-0045-4f8b-9524-a552f881f08b",
    "health": 100,
    "length": 13,
    "taunt": "\"You've just been ERASED!!\"",
    "name": "chiseler",
    "object": "snake",
    "body": {
      "data": [
        {
          "object": "point",
          "x": 9,
          "y": 11
        },
        {
          "object": "point",
          "x": 9,
          "y": 10
        },
        {
          "object": "point",
          "x": 8,
          "y": 10
        },
        {
          "object": "point",
          "x": 8,
          "y": 11
        },
        {
          "object": "point",
          "x": 8,
          "y": 12
        },
        {
          "object": "point",
          "x": 7,
          "y": 12
        },
        {
          "object": "point",
          "x": 7,
          "y": 11
        },
        {
          "object": "point",
          "x": 6,
          "y": 11
        },
        {
          "object": "point",
          "x": 6,
          "y": 10
        },
        {
          "object": "point",
          "x": 7,
          "y": 10
        },
        {
          "object": "point",
          "x": 7,
          "y": 9
        },
        {
          "object": "point",
          "x": 7,
          "y": 8
        },
        {
          "object": "point",
          "x": 7,
          "y": 8
        }
      ],
      "object": "list"
    }
  }
}`
