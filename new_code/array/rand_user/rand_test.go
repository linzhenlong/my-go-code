package rand_user

import (
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T)  {
	var users [4]string
	users[0] = "杨骏"
	users[1] = "燕小伟"
	users[2] = "林振龙"
	users[3] = "陈静欢"
	t.Log(users)

	rand.Seed(time.Now().UnixNano())
	t.Log(time.Now().UnixNano())
	num := rand.Intn(len(users))
	t.Log(num)
	t.Log(users[num])
}