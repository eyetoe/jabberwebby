package jabberwebby

import (
    "math/rand"
    "time"
)

func Monster() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Skeletal Racoon",
                        "Humplefrump",
                        "Calamari",
                        "Unicorn on the cob",
                        "Jabberwocky",
                        "Odd Sock",
                        "Skeevy Skivvies",
                        "Bandersnatch",
                        "12ft Duck",
                        "Tinysaurus Rex",
                        "Slithering Slime",
                        "Owl Badger",
                        "Mole Rat",
                        "Earthworm",
                        "Tardigrade",
                        "Kangashrew",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
