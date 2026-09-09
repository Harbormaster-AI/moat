from django.urls import path
from inventoryOnDjango.views import TransferOrderView

urlpatterns = [
    path('', TransferOrderView.index, name='index'),
	path('create', TransferOrderView.get, name='create'),
	path('get/<int:transferOrderId>/', TransferOrderView.get, name='get'),
	path('save', TransferOrderView.save, name='save'),
	path('getAll', TransferOrderView.getAll, name='getAll'),
	path('delete/<int:transferOrderId>/', TransferOrderView.delete, name='delete'),
	path('assignOriginWarehouse/<int:transferOrderId>/<int:OriginWarehouseId>/', TransferOrderView.assignOriginWarehouse, name='assignOriginWarehouse'),
	path('unassignOriginWarehouse/<int:transferOrderId>/', TransferOrderView.unassignOriginWarehouse, name='unassignOriginWarehouse'),
	path('assignDestinationWarehouse/<int:transferOrderId>/<int:DestinationWarehouseId>/', TransferOrderView.assignDestinationWarehouse, name='assignDestinationWarehouse'),
	path('unassignDestinationWarehouse/<int:transferOrderId>/', TransferOrderView.unassignDestinationWarehouse, name='unassignDestinationWarehouse'),
	path('addLines/<int:transferOrderId>/<LinesIds>/', TransferOrderView.addLines, name='addLines'),
	path('removeLines/<int:transferOrderId>/<LinesIds>/', TransferOrderView.removeLines, name='removeLines'),
	path('addTransactions/<int:transferOrderId>/<TransactionsIds>/', TransferOrderView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:transferOrderId>/<TransactionsIds>/', TransferOrderView.removeTransactions, name='removeTransactions'),
]
