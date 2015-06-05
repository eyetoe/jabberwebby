package jabberwebby

import (
    "math/rand"
    "time"
)

func Firstname() string{

    rand.Seed(time.Now().Unix())

    // slice
    names := []string{
                       "Dink",
                       "Yolanda",
                       "Lemonjello",
                       "Orangello",
                       "Grog",
                       "Dudley",
                       "Foop",
                       "Sly",
                       "Jenkins",
                       "Bent",
                       "Dorcus",
                       "Elvis",
                       "Fugnatious",
                       "Halcyon",
                       "Melvin",
                       "Woosh",
                       "Bark",
                       "Derp",
                       "Gern",
                       "Biff",
                       "Nemesis",
                       "Slim",
                       "Meh",
                       "Felicity",
                       "Walnut",
                       "Walrus",
                       "Dingo",
                       "Yahoo",
                       "Elixir",
                       "Cantankerous",
                       "Gonzo",
                       "Ethereal",
                       "Inigo",
                       "Vice",
                       "Gollum",
                       "Hodor",
                       "Hedley",
                       "Flummox",
                       "Codswallop",
                       "Doozy",
                       "Smeagol",
                       "Poodle",
                       "Doofus",
                       "Scut",
                       "Flip",
                       "Hero",
                       "Bart",
                       "Sherbert",
                       "Spatula",
                       "Fecalpap",
                       "Smorgasbord",
                       "Silly",
                       "Buckaroo",
                       "Ted",
                       "Disco",
                       "Burn",
                       "Horatio",
                       "Tingle",
                       "Splash",
                       "Johnny",
                       "Tuna",
                       "Nirvana",
                       "Wap",
                        }

    // return random element
    return names[rand.Intn(len(names))]
}
