from django.urls import path
from ecommerceOnDjango.views import CarrierServiceView

urlpatterns = [
    path('', CarrierServiceView.index, name='index'),
	path('create', CarrierServiceView.get, name='create'),
	path('get/<int:carrierServiceId>/', CarrierServiceView.get, name='get'),
	path('save', CarrierServiceView.save, name='save'),
	path('getAll', CarrierServiceView.getAll, name='getAll'),
	path('delete/<int:carrierServiceId>/', CarrierServiceView.delete, name='delete'),
	path('addShippingMethods/<int:carrierServiceId>/<ShippingMethodsIds>/', CarrierServiceView.addShippingMethods, name='addShippingMethods'),
	path('removeShippingMethods/<int:carrierServiceId>/<ShippingMethodsIds>/', CarrierServiceView.removeShippingMethods, name='removeShippingMethods'),
]
