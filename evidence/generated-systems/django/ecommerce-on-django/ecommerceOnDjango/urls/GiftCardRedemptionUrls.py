from django.urls import path
from ecommerceOnDjango.views import GiftCardRedemptionView

urlpatterns = [
    path('', GiftCardRedemptionView.index, name='index'),
	path('create', GiftCardRedemptionView.get, name='create'),
	path('get/<int:giftCardRedemptionId>/', GiftCardRedemptionView.get, name='get'),
	path('save', GiftCardRedemptionView.save, name='save'),
	path('getAll', GiftCardRedemptionView.getAll, name='getAll'),
	path('delete/<int:giftCardRedemptionId>/', GiftCardRedemptionView.delete, name='delete'),
	path('assignGiftCard/<int:giftCardRedemptionId>/<int:GiftCardId>/', GiftCardRedemptionView.assignGiftCard, name='assignGiftCard'),
	path('unassignGiftCard/<int:giftCardRedemptionId>/', GiftCardRedemptionView.unassignGiftCard, name='unassignGiftCard'),
	path('assignOrder/<int:giftCardRedemptionId>/<int:OrderId>/', GiftCardRedemptionView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:giftCardRedemptionId>/', GiftCardRedemptionView.unassignOrder, name='unassignOrder'),
]
