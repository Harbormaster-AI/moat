from django.urls import path
from ecommerceOnDjango.views import ShipmentView

urlpatterns = [
    path('', ShipmentView.index, name='index'),
	path('create', ShipmentView.get, name='create'),
	path('get/<int:shipmentId>/', ShipmentView.get, name='get'),
	path('save', ShipmentView.save, name='save'),
	path('getAll', ShipmentView.getAll, name='getAll'),
	path('delete/<int:shipmentId>/', ShipmentView.delete, name='delete'),
	path('assignOrder/<int:shipmentId>/<int:OrderId>/', ShipmentView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:shipmentId>/', ShipmentView.unassignOrder, name='unassignOrder'),
	path('assignFulfillmentCenter/<int:shipmentId>/<int:FulfillmentCenterId>/', ShipmentView.assignFulfillmentCenter, name='assignFulfillmentCenter'),
	path('unassignFulfillmentCenter/<int:shipmentId>/', ShipmentView.unassignFulfillmentCenter, name='unassignFulfillmentCenter'),
	path('addShipmentItems/<int:shipmentId>/<ShipmentItemsIds>/', ShipmentView.addShipmentItems, name='addShipmentItems'),
	path('removeShipmentItems/<int:shipmentId>/<ShipmentItemsIds>/', ShipmentView.removeShipmentItems, name='removeShipmentItems'),
]
