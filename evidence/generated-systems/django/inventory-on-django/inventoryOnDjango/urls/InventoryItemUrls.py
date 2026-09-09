from django.urls import path
from inventoryOnDjango.views import InventoryItemView

urlpatterns = [
    path('', InventoryItemView.index, name='index'),
	path('create', InventoryItemView.get, name='create'),
	path('get/<int:inventoryItemId>/', InventoryItemView.get, name='get'),
	path('save', InventoryItemView.save, name='save'),
	path('getAll', InventoryItemView.getAll, name='getAll'),
	path('delete/<int:inventoryItemId>/', InventoryItemView.delete, name='delete'),
	path('assignSku/<int:inventoryItemId>/<int:SkuId>/', InventoryItemView.assignSku, name='assignSku'),
	path('unassignSku/<int:inventoryItemId>/', InventoryItemView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:inventoryItemId>/<int:WarehouseId>/', InventoryItemView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:inventoryItemId>/', InventoryItemView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLocation/<int:inventoryItemId>/<int:LocationId>/', InventoryItemView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:inventoryItemId>/', InventoryItemView.unassignLocation, name='unassignLocation'),
	path('assignLot/<int:inventoryItemId>/<int:LotId>/', InventoryItemView.assignLot, name='assignLot'),
	path('unassignLot/<int:inventoryItemId>/', InventoryItemView.unassignLot, name='unassignLot'),
	path('addSerialNumbers/<int:inventoryItemId>/<SerialNumbersIds>/', InventoryItemView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:inventoryItemId>/<SerialNumbersIds>/', InventoryItemView.removeSerialNumbers, name='removeSerialNumbers'),
	path('addTransactions/<int:inventoryItemId>/<TransactionsIds>/', InventoryItemView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:inventoryItemId>/<TransactionsIds>/', InventoryItemView.removeTransactions, name='removeTransactions'),
	path('addReservations/<int:inventoryItemId>/<ReservationsIds>/', InventoryItemView.addReservations, name='addReservations'),
	path('removeReservations/<int:inventoryItemId>/<ReservationsIds>/', InventoryItemView.removeReservations, name='removeReservations'),
]
