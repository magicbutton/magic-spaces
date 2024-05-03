/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package space

import (
    "log"

    "github.com/magicbutton/magic-spaces/applogic"
    "github.com/magicbutton/magic-spaces/database"
    "github.com/magicbutton/magic-spaces/services/models/spacemodel"
    . "github.com/magicbutton/magic-spaces/utils"
)

func SpaceSearch(query string) (*Page[spacemodel.Space], error) {
    log.Println("Calling Spacesearch")

    return applogic.Search[database.Space, spacemodel.Space]("name", query, applogic.MapSpaceOutgoing)

}
