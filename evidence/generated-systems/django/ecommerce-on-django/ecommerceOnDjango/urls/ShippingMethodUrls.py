from django.urls import path
from ecommerceOnDjango.views import ShippingMethodView

urlpatterns = [
    path('', ShippingMethodView.index, name='index'),
	path('create', ShippingMethodView.get, name='create'),
	path('get/<int:shippingMethodId>/', ShippingMethodView.get, name='get'),
	path('save', ShippingMethodView.save, name='save'),
	path('getAll', ShippingMethodView.getAll, name='getAll'),
	path('delete/<int:shippingMethodId>/', ShippingMethodView.delete, name='delete'),
	path('assignCarrierService/<int:shippingMethodId>/<int:CarrierServiceId>/', ShippingMethodView.assignCarrierService, name='assignCarrierService'),
	path('unassignCarrierService/<int:shippingMethodId>/', ShippingMethodView.unassignCarrierService, name='unassignCarrierService'),
	path('addChannels/<int:shippingMethodId>/<ChannelsIds>/', ShippingMethodView.addChannels, name='addChannels'),
	path('removeChannels/<int:shippingMethodId>/<ChannelsIds>/', ShippingMethodView.removeChannels, name='removeChannels'),
]
