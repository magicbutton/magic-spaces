    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-spaces/services/endpoints/spacetype"
                    "github.com/magicbutton/magic-spaces/services/models/spacetypemodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestSpaceTypeupdate(t *testing.T) {
                                // transformer v1
            object := spacetypemodel.SpaceType{}
         
            result,err := spacetype.SpaceTypeUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
