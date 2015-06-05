package jabberwebby

import (
    "math/rand"
    "time"
)

func Lastname() string{

    rand.Seed(time.Now().Unix())

    // slice
    names := []string{
                       "Jammypants",
                       "Waperson",
                       "Dandypants",
                       "Montoya",
                       "Hoobastank",
                       "Dinkski",
                       "Wooshblatz",
                       "Slime",
                       "Maximus",
                       "Cashew",
                       "Malarkey",
                       "Blandsten",
                       "Palendrome",
                       "Fooperson",
                       "Happyslap",
                       "Farkus",
                       "Banzai",
                       "Koala-nose",
                       "Slothface",
                       "Wilikins",
                       "Wilson",
                       "Stark",
                       "Octopus",
                       "Yaya",
                       "Lamarr",
                       "Smallberries",
                       "Slangwhanger",
                       "Frankenstein",
                       "Bigbootay",
                       "Blunderbuss",
                       "Hodor",
                       "Troglodyte",
                       "Comeuppance",
                       "Googlethat",
                       "Nevermind",
                       "Skedaddle",
                       "Mugwump",
                       "Canoodle",
                       "Fishface",
                       "Shenanigans",
                       "Fists-o-fury",
                       "La-di-da",
                       "Baba Ghanoush",
                       "Not-in-the-face",
                        }

    // return random element
    return names[rand.Intn(len(names))]
}
