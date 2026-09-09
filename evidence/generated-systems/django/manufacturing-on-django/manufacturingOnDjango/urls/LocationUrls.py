from django.urls import path
from manufacturingOnDjango.views import LocationView

urlpatterns = [
    path('', LocationView.index, name='index'),
	path('create', LocationView.get, name='create'),
	path('get/<int:locationId>/', LocationView.get, name='get'),
	path('save', LocationView.save, name='save'),
	path('getAll', LocationView.getAll, name='getAll'),
	path('delete/<int:locationId>/', LocationView.delete, name='delete'),
	path('assignWarehouse/<int:locationId>/<int:WarehouseId>/', LocationView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:locationId>/', LocationView.unassignWarehouse, name='unassignWarehouse'),
	path('addInventoryItems/<int:locationId>/<InventoryItemsIds>/', LocationView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:locationId>/<InventoryItemsIds>/', LocationView.removeInventoryItems, name='removeInventoryItems'),
]
