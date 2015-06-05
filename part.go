package jabberwebby

import (
    "math/rand"
    "time"
)

func Part() string{

    rand.Seed(time.Now().Unix())

    // slice
    weapons := []string{
                        "jejunum",
                        "esophagus",
                        "lips",
                        "eyelashes",
                        "pinky toe",
                        "shin",
                        "nipple",
                        "uvula",
                        "everything",
                        "vestigial tail",
                        "nether regions",
                        "left gonad",
                        "twig and berries",
                        "elbow",
                        "spleen",
                        "junk",
                        "butt cheek",
                        "ego",
                        "gullet",
                        "front toof",
                        "lips",
                        "backside",
                        "front butt",
                        "special purpose",
                        "money maker",
                        "pie hole",
                        "hoo hoo",
                        }

    // return random element
    return weapons[rand.Intn(len(weapons))]
}
