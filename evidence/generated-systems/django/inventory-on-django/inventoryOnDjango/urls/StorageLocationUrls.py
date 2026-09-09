from django.urls import path
from inventoryOnDjango.views import StorageLocationView

urlpatterns = [
    path('', StorageLocationView.index, name='index'),
	path('create', StorageLocationView.get, name='create'),
	path('get/<int:storageLocationId>/', StorageLocationView.get, name='get'),
	path('save', StorageLocationView.save, name='save'),
	path('getAll', StorageLocationView.getAll, name='getAll'),
	path('delete/<int:storageLocationId>/', StorageLocationView.delete, name='delete'),
	path('assignWarehouse/<int:storageLocationId>/<int:WarehouseId>/', StorageLocationView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:storageLocationId>/', StorageLocationView.unassignWarehouse, name='unassignWarehouse'),
	path('assignParentLocation/<int:storageLocationId>/<int:ParentLocationId>/', StorageLocationView.assignParentLocation, name='assignParentLocation'),
	path('unassignParentLocation/<int:storageLocationId>/', StorageLocationView.unassignParentLocation, name='unassignParentLocation'),
	path('addChildLocations/<int:storageLocationId>/<ChildLocationsIds>/', StorageLocationView.addChildLocations, name='addChildLocations'),
	path('removeChildLocations/<int:storageLocationId>/<ChildLocationsIds>/', StorageLocationView.removeChildLocations, name='removeChildLocations'),
	path('addInventoryItems/<int:storageLocationId>/<InventoryItemsIds>/', StorageLocationView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:storageLocationId>/<InventoryItemsIds>/', StorageLocationView.removeInventoryItems, name='removeInventoryItems'),
]
