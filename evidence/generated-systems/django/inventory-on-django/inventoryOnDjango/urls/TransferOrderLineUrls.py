from django.urls import path
from inventoryOnDjango.views import TransferOrderLineView

urlpatterns = [
    path('', TransferOrderLineView.index, name='index'),
	path('create', TransferOrderLineView.get, name='create'),
	path('get/<int:transferOrderLineId>/', TransferOrderLineView.get, name='get'),
	path('save', TransferOrderLineView.save, name='save'),
	path('getAll', TransferOrderLineView.getAll, name='getAll'),
	path('delete/<int:transferOrderLineId>/', TransferOrderLineView.delete, name='delete'),
	path('assignTransferOrder/<int:transferOrderLineId>/<int:TransferOrderId>/', TransferOrderLineView.assignTransferOrder, name='assignTransferOrder'),
	path('unassignTransferOrder/<int:transferOrderLineId>/', TransferOrderLineView.unassignTransferOrder, name='unassignTransferOrder'),
	path('assignSku/<int:transferOrderLineId>/<int:SkuId>/', TransferOrderLineView.assignSku, name='assignSku'),
	path('unassignSku/<int:transferOrderLineId>/', TransferOrderLineView.unassignSku, name='unassignSku'),
	path('assignLot/<int:transferOrderLineId>/<int:LotId>/', TransferOrderLineView.assignLot, name='assignLot'),
	path('unassignLot/<int:transferOrderLineId>/', TransferOrderLineView.unassignLot, name='unassignLot'),
	path('assignFromLocation/<int:transferOrderLineId>/<int:FromLocationId>/', TransferOrderLineView.assignFromLocation, name='assignFromLocation'),
	path('unassignFromLocation/<int:transferOrderLineId>/', TransferOrderLineView.unassignFromLocation, name='unassignFromLocation'),
	path('assignToLocation/<int:transferOrderLineId>/<int:ToLocationId>/', TransferOrderLineView.assignToLocation, name='assignToLocation'),
	path('unassignToLocation/<int:transferOrderLineId>/', TransferOrderLineView.unassignToLocation, name='unassignToLocation'),
	path('addSerialNumbers/<int:transferOrderLineId>/<SerialNumbersIds>/', TransferOrderLineView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:transferOrderLineId>/<SerialNumbersIds>/', TransferOrderLineView.removeSerialNumbers, name='removeSerialNumbers'),
]
