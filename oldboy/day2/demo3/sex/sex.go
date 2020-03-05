package sex

import "time"

const (
    Female = 2
    Men = 1
)

func Print() string {
    time := time.Now().Unix()
    if (time % Female == 0) {
        return "Female"
    } else {
        return "Men"
    }
}
