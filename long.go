package jabberwebby

import (
    "math/rand"
    "time"
)

func Long() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Pistol",
                        "SMG",
                        "Assault Rifle",
                        "Hunting Rifle",
                        "Sniper Rifle",
                        "Rocket Launcher",
                        "Crossbow",
                        "Long Bow",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
