from django.urls import path
from inventoryOnDjango.views import ReservationView

urlpatterns = [
    path('', ReservationView.index, name='index'),
	path('create', ReservationView.get, name='create'),
	path('get/<int:reservationId>/', ReservationView.get, name='get'),
	path('save', ReservationView.save, name='save'),
	path('getAll', ReservationView.getAll, name='getAll'),
	path('delete/<int:reservationId>/', ReservationView.delete, name='delete'),
	path('assignSku/<int:reservationId>/<int:SkuId>/', ReservationView.assignSku, name='assignSku'),
	path('unassignSku/<int:reservationId>/', ReservationView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:reservationId>/<int:WarehouseId>/', ReservationView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:reservationId>/', ReservationView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLocation/<int:reservationId>/<int:LocationId>/', ReservationView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:reservationId>/', ReservationView.unassignLocation, name='unassignLocation'),
	path('assignInventoryItem/<int:reservationId>/<int:InventoryItemId>/', ReservationView.assignInventoryItem, name='assignInventoryItem'),
	path('unassignInventoryItem/<int:reservationId>/', ReservationView.unassignInventoryItem, name='unassignInventoryItem'),
	path('assignLot/<int:reservationId>/<int:LotId>/', ReservationView.assignLot, name='assignLot'),
	path('unassignLot/<int:reservationId>/', ReservationView.unassignLot, name='unassignLot'),
	path('assignDemandSignal/<int:reservationId>/<int:DemandSignalId>/', ReservationView.assignDemandSignal, name='assignDemandSignal'),
	path('unassignDemandSignal/<int:reservationId>/', ReservationView.unassignDemandSignal, name='unassignDemandSignal'),
	path('addSerialNumbers/<int:reservationId>/<SerialNumbersIds>/', ReservationView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:reservationId>/<SerialNumbersIds>/', ReservationView.removeSerialNumbers, name='removeSerialNumbers'),
]
