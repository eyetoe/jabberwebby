package jabberwebby

import (
    "math/rand"
    "time"
)

func Grenade() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Explosive",
                        "Flash Bang",
                        "Tear Gas",
                        "Bouncy",
                        "Sticky",
                        "Incindiary",
                        "Cryo",
                        "Tesla",
                        "Molotov",
                        "Cluster",
                        "Shrapnel",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
