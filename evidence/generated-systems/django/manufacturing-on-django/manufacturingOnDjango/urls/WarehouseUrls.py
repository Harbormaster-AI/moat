from django.urls import path
from manufacturingOnDjango.views import WarehouseView

urlpatterns = [
    path('', WarehouseView.index, name='index'),
	path('create', WarehouseView.get, name='create'),
	path('get/<int:warehouseId>/', WarehouseView.get, name='get'),
	path('save', WarehouseView.save, name='save'),
	path('getAll', WarehouseView.getAll, name='getAll'),
	path('delete/<int:warehouseId>/', WarehouseView.delete, name='delete'),
	path('assignPlant/<int:warehouseId>/<int:PlantId>/', WarehouseView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:warehouseId>/', WarehouseView.unassignPlant, name='unassignPlant'),
	path('addLocations/<int:warehouseId>/<LocationsIds>/', WarehouseView.addLocations, name='addLocations'),
	path('removeLocations/<int:warehouseId>/<LocationsIds>/', WarehouseView.removeLocations, name='removeLocations'),
	path('addInventoryItems/<int:warehouseId>/<InventoryItemsIds>/', WarehouseView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:warehouseId>/<InventoryItemsIds>/', WarehouseView.removeInventoryItems, name='removeInventoryItems'),
]
