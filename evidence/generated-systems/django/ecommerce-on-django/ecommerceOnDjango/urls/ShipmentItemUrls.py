from django.urls import path
from ecommerceOnDjango.views import ShipmentItemView

urlpatterns = [
    path('', ShipmentItemView.index, name='index'),
	path('create', ShipmentItemView.get, name='create'),
	path('get/<int:shipmentItemId>/', ShipmentItemView.get, name='get'),
	path('save', ShipmentItemView.save, name='save'),
	path('getAll', ShipmentItemView.getAll, name='getAll'),
	path('delete/<int:shipmentItemId>/', ShipmentItemView.delete, name='delete'),
	path('assignShipment/<int:shipmentItemId>/<int:ShipmentId>/', ShipmentItemView.assignShipment, name='assignShipment'),
	path('unassignShipment/<int:shipmentItemId>/', ShipmentItemView.unassignShipment, name='unassignShipment'),
	path('assignOrderLine/<int:shipmentItemId>/<int:OrderLineId>/', ShipmentItemView.assignOrderLine, name='assignOrderLine'),
	path('unassignOrderLine/<int:shipmentItemId>/', ShipmentItemView.unassignOrderLine, name='unassignOrderLine'),
]
