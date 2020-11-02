package merchants

import (
	"github.com/azdonald/sellerbase/src/util"
)

// Merchant represents the sellers on the platform
type Merchant struct {
	util.Base
	Name string `json:"name"`
}
