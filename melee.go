package jabberwebby

import (
    "math/rand"
    "time"
)

func Melee() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "Sword",
                        "Sledgehammer",
                        "Chainsaw",
                        "Shillelagh",
                        "Welding Torch",
                        "Ice Pick",
                        "Number 2 Pencil",
                        "Crowbar",
                        "Bar stool",
                        "Wrench",
                        "Old Boot",
                        "Broom",
                        "Rusty Nail",
                        "Feather Duster",
                        "Stun Stick",
                        "Sausage Link",
                        "Fart",
                        "Poop-on-a-stick",
                        "Spatula",
                        "Dead frog",
                        "Dirty Spoon",
                        "Pokey Stick",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
