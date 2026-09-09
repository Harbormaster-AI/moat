from django.urls import path
from inventoryOnDjango.views import InboundShipmentView

urlpatterns = [
    path('', InboundShipmentView.index, name='index'),
	path('create', InboundShipmentView.get, name='create'),
	path('get/<int:inboundShipmentId>/', InboundShipmentView.get, name='get'),
	path('save', InboundShipmentView.save, name='save'),
	path('getAll', InboundShipmentView.getAll, name='getAll'),
	path('delete/<int:inboundShipmentId>/', InboundShipmentView.delete, name='delete'),
	path('assignWarehouse/<int:inboundShipmentId>/<int:WarehouseId>/', InboundShipmentView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:inboundShipmentId>/', InboundShipmentView.unassignWarehouse, name='unassignWarehouse'),
	path('addLines/<int:inboundShipmentId>/<LinesIds>/', InboundShipmentView.addLines, name='addLines'),
	path('removeLines/<int:inboundShipmentId>/<LinesIds>/', InboundShipmentView.removeLines, name='removeLines'),
	path('addTransactions/<int:inboundShipmentId>/<TransactionsIds>/', InboundShipmentView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:inboundShipmentId>/<TransactionsIds>/', InboundShipmentView.removeTransactions, name='removeTransactions'),
]
