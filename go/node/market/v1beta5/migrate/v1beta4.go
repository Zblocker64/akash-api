package migrate

import (
	"github.com/akash-network/akash-api/go/node/market/v1beta4"

	dmigrate "pkg.akt.io/go/node/deployment/v1beta4/migrate"
	"pkg.akt.io/go/node/market/v1"
	"pkg.akt.io/go/node/market/v1/migrate"
	v1migrate "pkg.akt.io/go/node/market/v1/migrate"
	"pkg.akt.io/go/node/market/v1beta5"
)

func ResourcesOfferFromV1beta4(from v1beta4.ResourcesOffer) v1beta5.ResourcesOffer {
	res := make(v1beta5.ResourcesOffer, 0, len(from))

	return res
}

func BidFromV1beta4(from v1beta4.Bid) v1beta5.Bid {
	return v1beta5.Bid{
		ID:             migrate.BidIDFromV1beta4(from.BidID),
		State:          migrate.BidStateFromV1beta4(from.State),
		Price:          from.Price,
		CreatedAt:      from.CreatedAt,
		ResourcesOffer: ResourcesOfferFromV1beta4(from.ResourcesOffer),
	}
}

func OrderFromV1beta4(from v1beta4.Order) v1beta5.Order {
	return v1beta5.Order{
		ID:        v1migrate.OrderIDFromV1beta4(from.OrderID),
		State:     v1.OrderState(from.State),
		Spec:      dmigrate.GroupSpecFromV1Beta3(from.Spec),
		CreatedAt: from.CreatedAt,
	}
}
