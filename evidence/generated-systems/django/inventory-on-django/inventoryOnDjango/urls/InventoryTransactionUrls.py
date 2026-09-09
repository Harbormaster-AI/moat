from django.urls import path
from inventoryOnDjango.views import InventoryTransactionView

urlpatterns = [
    path('', InventoryTransactionView.index, name='index'),
	path('create', InventoryTransactionView.get, name='create'),
	path('get/<int:inventoryTransactionId>/', InventoryTransactionView.get, name='get'),
	path('save', InventoryTransactionView.save, name='save'),
	path('getAll', InventoryTransactionView.getAll, name='getAll'),
	path('delete/<int:inventoryTransactionId>/', InventoryTransactionView.delete, name='delete'),
	path('assignSku/<int:inventoryTransactionId>/<int:SkuId>/', InventoryTransactionView.assignSku, name='assignSku'),
	path('unassignSku/<int:inventoryTransactionId>/', InventoryTransactionView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:inventoryTransactionId>/<int:WarehouseId>/', InventoryTransactionView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:inventoryTransactionId>/', InventoryTransactionView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLocation/<int:inventoryTransactionId>/<int:LocationId>/', InventoryTransactionView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:inventoryTransactionId>/', InventoryTransactionView.unassignLocation, name='unassignLocation'),
	path('assignLot/<int:inventoryTransactionId>/<int:LotId>/', InventoryTransactionView.assignLot, name='assignLot'),
	path('unassignLot/<int:inventoryTransactionId>/', InventoryTransactionView.unassignLot, name='unassignLot'),
	path('assignRelatedReservation/<int:inventoryTransactionId>/<int:RelatedReservationId>/', InventoryTransactionView.assignRelatedReservation, name='assignRelatedReservation'),
	path('unassignRelatedReservation/<int:inventoryTransactionId>/', InventoryTransactionView.unassignRelatedReservation, name='unassignRelatedReservation'),
	path('assignTransferOrder/<int:inventoryTransactionId>/<int:TransferOrderId>/', InventoryTransactionView.assignTransferOrder, name='assignTransferOrder'),
	path('unassignTransferOrder/<int:inventoryTransactionId>/', InventoryTransactionView.unassignTransferOrder, name='unassignTransferOrder'),
	path('assignAdjustment/<int:inventoryTransactionId>/<int:AdjustmentId>/', InventoryTransactionView.assignAdjustment, name='assignAdjustment'),
	path('unassignAdjustment/<int:inventoryTransactionId>/', InventoryTransactionView.unassignAdjustment, name='unassignAdjustment'),
	path('assignCycleCount/<int:inventoryTransactionId>/<int:CycleCountId>/', InventoryTransactionView.assignCycleCount, name='assignCycleCount'),
	path('unassignCycleCount/<int:inventoryTransactionId>/', InventoryTransactionView.unassignCycleCount, name='unassignCycleCount'),
	path('addSerialNumbers/<int:inventoryTransactionId>/<SerialNumbersIds>/', InventoryTransactionView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:inventoryTransactionId>/<SerialNumbersIds>/', InventoryTransactionView.removeSerialNumbers, name='removeSerialNumbers'),
]
