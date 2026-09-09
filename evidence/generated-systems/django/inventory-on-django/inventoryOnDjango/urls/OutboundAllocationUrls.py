from django.urls import path
from inventoryOnDjango.views import OutboundAllocationView

urlpatterns = [
    path('', OutboundAllocationView.index, name='index'),
	path('create', OutboundAllocationView.get, name='create'),
	path('get/<int:outboundAllocationId>/', OutboundAllocationView.get, name='get'),
	path('save', OutboundAllocationView.save, name='save'),
	path('getAll', OutboundAllocationView.getAll, name='getAll'),
	path('delete/<int:outboundAllocationId>/', OutboundAllocationView.delete, name='delete'),
	path('assignWarehouse/<int:outboundAllocationId>/<int:WarehouseId>/', OutboundAllocationView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:outboundAllocationId>/', OutboundAllocationView.unassignWarehouse, name='unassignWarehouse'),
	path('assignSku/<int:outboundAllocationId>/<int:SkuId>/', OutboundAllocationView.assignSku, name='assignSku'),
	path('unassignSku/<int:outboundAllocationId>/', OutboundAllocationView.unassignSku, name='unassignSku'),
	path('assignInventoryItem/<int:outboundAllocationId>/<int:InventoryItemId>/', OutboundAllocationView.assignInventoryItem, name='assignInventoryItem'),
	path('unassignInventoryItem/<int:outboundAllocationId>/', OutboundAllocationView.unassignInventoryItem, name='unassignInventoryItem'),
	path('assignReservation/<int:outboundAllocationId>/<int:ReservationId>/', OutboundAllocationView.assignReservation, name='assignReservation'),
	path('unassignReservation/<int:outboundAllocationId>/', OutboundAllocationView.unassignReservation, name='unassignReservation'),
	path('assignLot/<int:outboundAllocationId>/<int:LotId>/', OutboundAllocationView.assignLot, name='assignLot'),
	path('unassignLot/<int:outboundAllocationId>/', OutboundAllocationView.unassignLot, name='unassignLot'),
	path('assignSourceLocation/<int:outboundAllocationId>/<int:SourceLocationId>/', OutboundAllocationView.assignSourceLocation, name='assignSourceLocation'),
	path('unassignSourceLocation/<int:outboundAllocationId>/', OutboundAllocationView.unassignSourceLocation, name='unassignSourceLocation'),
	path('addSerialNumbers/<int:outboundAllocationId>/<SerialNumbersIds>/', OutboundAllocationView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:outboundAllocationId>/<SerialNumbersIds>/', OutboundAllocationView.removeSerialNumbers, name='removeSerialNumbers'),
]
