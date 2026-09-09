from django.urls import path
from ecommerceOnDjango.views import CustomerView

urlpatterns = [
    path('', CustomerView.index, name='index'),
	path('create', CustomerView.get, name='create'),
	path('get/<int:customerId>/', CustomerView.get, name='get'),
	path('save', CustomerView.save, name='save'),
	path('getAll', CustomerView.getAll, name='getAll'),
	path('delete/<int:customerId>/', CustomerView.delete, name='delete'),
	path('addAddresses/<int:customerId>/<AddressesIds>/', CustomerView.addAddresses, name='addAddresses'),
	path('removeAddresses/<int:customerId>/<AddressesIds>/', CustomerView.removeAddresses, name='removeAddresses'),
	path('addCarts/<int:customerId>/<CartsIds>/', CustomerView.addCarts, name='addCarts'),
	path('removeCarts/<int:customerId>/<CartsIds>/', CustomerView.removeCarts, name='removeCarts'),
	path('addOrders/<int:customerId>/<OrdersIds>/', CustomerView.addOrders, name='addOrders'),
	path('removeOrders/<int:customerId>/<OrdersIds>/', CustomerView.removeOrders, name='removeOrders'),
	path('addPayments/<int:customerId>/<PaymentsIds>/', CustomerView.addPayments, name='addPayments'),
	path('removePayments/<int:customerId>/<PaymentsIds>/', CustomerView.removePayments, name='removePayments'),
	path('addReviews/<int:customerId>/<ReviewsIds>/', CustomerView.addReviews, name='addReviews'),
	path('removeReviews/<int:customerId>/<ReviewsIds>/', CustomerView.removeReviews, name='removeReviews'),
	path('addWishlists/<int:customerId>/<WishlistsIds>/', CustomerView.addWishlists, name='addWishlists'),
	path('removeWishlists/<int:customerId>/<WishlistsIds>/', CustomerView.removeWishlists, name='removeWishlists'),
	path('addSubscriptions/<int:customerId>/<SubscriptionsIds>/', CustomerView.addSubscriptions, name='addSubscriptions'),
	path('removeSubscriptions/<int:customerId>/<SubscriptionsIds>/', CustomerView.removeSubscriptions, name='removeSubscriptions'),
	path('addCouponRedemptions/<int:customerId>/<CouponRedemptionsIds>/', CustomerView.addCouponRedemptions, name='addCouponRedemptions'),
	path('removeCouponRedemptions/<int:customerId>/<CouponRedemptionsIds>/', CustomerView.removeCouponRedemptions, name='removeCouponRedemptions'),
	path('addGiftCards/<int:customerId>/<GiftCardsIds>/', CustomerView.addGiftCards, name='addGiftCards'),
	path('removeGiftCards/<int:customerId>/<GiftCardsIds>/', CustomerView.removeGiftCards, name='removeGiftCards'),
]
