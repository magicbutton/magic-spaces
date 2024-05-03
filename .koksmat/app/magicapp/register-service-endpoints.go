/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-spaces/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
    root.AddEndpoint("person", micro.HandlerFunc(services.HandlePersonRequests))
        root.AddEndpoint("spacetype", micro.HandlerFunc(services.HandleSpaceTypeRequests))
        root.AddEndpoint("space", micro.HandlerFunc(services.HandleSpaceRequests))
    }
