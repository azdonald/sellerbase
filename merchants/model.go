package merchants

import (
	"github.com/azdonald/sellerbase/src/util"
)

// Merchant represents the sellers on the platform
type Merchant struct {
	util.Base
	BusinessName string `gorm:"size:100;not null;unique" json:"business_name"`
}
