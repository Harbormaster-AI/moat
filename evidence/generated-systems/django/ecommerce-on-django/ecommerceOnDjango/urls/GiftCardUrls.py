from django.urls import path
from ecommerceOnDjango.views import GiftCardView

urlpatterns = [
    path('', GiftCardView.index, name='index'),
	path('create', GiftCardView.get, name='create'),
	path('get/<int:giftCardId>/', GiftCardView.get, name='get'),
	path('save', GiftCardView.save, name='save'),
	path('getAll', GiftCardView.getAll, name='getAll'),
	path('delete/<int:giftCardId>/', GiftCardView.delete, name='delete'),
	path('assignCustomer/<int:giftCardId>/<int:CustomerId>/', GiftCardView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:giftCardId>/', GiftCardView.unassignCustomer, name='unassignCustomer'),
	path('assignIssuedOrder/<int:giftCardId>/<int:IssuedOrderId>/', GiftCardView.assignIssuedOrder, name='assignIssuedOrder'),
	path('unassignIssuedOrder/<int:giftCardId>/', GiftCardView.unassignIssuedOrder, name='unassignIssuedOrder'),
	path('addRedemptions/<int:giftCardId>/<RedemptionsIds>/', GiftCardView.addRedemptions, name='addRedemptions'),
	path('removeRedemptions/<int:giftCardId>/<RedemptionsIds>/', GiftCardView.removeRedemptions, name='removeRedemptions'),
]
