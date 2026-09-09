from django.urls import path
from inventoryOnDjango.views import QuarantineView

urlpatterns = [
    path('', QuarantineView.index, name='index'),
	path('create', QuarantineView.get, name='create'),
	path('get/<int:quarantineId>/', QuarantineView.get, name='get'),
	path('save', QuarantineView.save, name='save'),
	path('getAll', QuarantineView.getAll, name='getAll'),
	path('delete/<int:quarantineId>/', QuarantineView.delete, name='delete'),
	path('assignWarehouse/<int:quarantineId>/<int:WarehouseId>/', QuarantineView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:quarantineId>/', QuarantineView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLot/<int:quarantineId>/<int:LotId>/', QuarantineView.assignLot, name='assignLot'),
	path('unassignLot/<int:quarantineId>/', QuarantineView.unassignLot, name='unassignLot'),
	path('addItems/<int:quarantineId>/<ItemsIds>/', QuarantineView.addItems, name='addItems'),
	path('removeItems/<int:quarantineId>/<ItemsIds>/', QuarantineView.removeItems, name='removeItems'),
	path('addSerialNumbers/<int:quarantineId>/<SerialNumbersIds>/', QuarantineView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:quarantineId>/<SerialNumbersIds>/', QuarantineView.removeSerialNumbers, name='removeSerialNumbers'),
]
