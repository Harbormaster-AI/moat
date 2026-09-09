from django.urls import path
from ecommerceOnDjango.views import ChannelView

urlpatterns = [
    path('', ChannelView.index, name='index'),
	path('create', ChannelView.get, name='create'),
	path('get/<int:channelId>/', ChannelView.get, name='get'),
	path('save', ChannelView.save, name='save'),
	path('getAll', ChannelView.getAll, name='getAll'),
	path('delete/<int:channelId>/', ChannelView.delete, name='delete'),
	path('assignMerchant/<int:channelId>/<int:MerchantId>/', ChannelView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:channelId>/', ChannelView.unassignMerchant, name='unassignMerchant'),
	path('addCatalogs/<int:channelId>/<CatalogsIds>/', ChannelView.addCatalogs, name='addCatalogs'),
	path('removeCatalogs/<int:channelId>/<CatalogsIds>/', ChannelView.removeCatalogs, name='removeCatalogs'),
	path('addPromotions/<int:channelId>/<PromotionsIds>/', ChannelView.addPromotions, name='addPromotions'),
	path('removePromotions/<int:channelId>/<PromotionsIds>/', ChannelView.removePromotions, name='removePromotions'),
	path('addShippingMethods/<int:channelId>/<ShippingMethodsIds>/', ChannelView.addShippingMethods, name='addShippingMethods'),
	path('removeShippingMethods/<int:channelId>/<ShippingMethodsIds>/', ChannelView.removeShippingMethods, name='removeShippingMethods'),
	path('addPaymentProviders/<int:channelId>/<PaymentProvidersIds>/', ChannelView.addPaymentProviders, name='addPaymentProviders'),
	path('removePaymentProviders/<int:channelId>/<PaymentProvidersIds>/', ChannelView.removePaymentProviders, name='removePaymentProviders'),
]
