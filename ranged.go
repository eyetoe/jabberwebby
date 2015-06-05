package jabberwebby

import (
    "math/rand"
    "time"
)

func Ranged() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Flame Thrower",
                        "Shotgun",
                        "Double Barrel Shotgun",
                        "Gernade Launcher",
                        "Throwing Knife",
                        "Tazer Stun Gun",
                        "Short Bow",
                        "Crossbow Pistol",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
