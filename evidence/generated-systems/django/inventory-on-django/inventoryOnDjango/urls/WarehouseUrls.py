from django.urls import path
from inventoryOnDjango.views import WarehouseView

urlpatterns = [
    path('', WarehouseView.index, name='index'),
	path('create', WarehouseView.get, name='create'),
	path('get/<int:warehouseId>/', WarehouseView.get, name='get'),
	path('save', WarehouseView.save, name='save'),
	path('getAll', WarehouseView.getAll, name='getAll'),
	path('delete/<int:warehouseId>/', WarehouseView.delete, name='delete'),
	path('addStorageLocations/<int:warehouseId>/<StorageLocationsIds>/', WarehouseView.addStorageLocations, name='addStorageLocations'),
	path('removeStorageLocations/<int:warehouseId>/<StorageLocationsIds>/', WarehouseView.removeStorageLocations, name='removeStorageLocations'),
	path('addInventoryItems/<int:warehouseId>/<InventoryItemsIds>/', WarehouseView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:warehouseId>/<InventoryItemsIds>/', WarehouseView.removeInventoryItems, name='removeInventoryItems'),
	path('addInboundShipments/<int:warehouseId>/<InboundShipmentsIds>/', WarehouseView.addInboundShipments, name='addInboundShipments'),
	path('removeInboundShipments/<int:warehouseId>/<InboundShipmentsIds>/', WarehouseView.removeInboundShipments, name='removeInboundShipments'),
	path('addOutboundAllocations/<int:warehouseId>/<OutboundAllocationsIds>/', WarehouseView.addOutboundAllocations, name='addOutboundAllocations'),
	path('removeOutboundAllocations/<int:warehouseId>/<OutboundAllocationsIds>/', WarehouseView.removeOutboundAllocations, name='removeOutboundAllocations'),
	path('addOriginTransfers/<int:warehouseId>/<OriginTransfersIds>/', WarehouseView.addOriginTransfers, name='addOriginTransfers'),
	path('removeOriginTransfers/<int:warehouseId>/<OriginTransfersIds>/', WarehouseView.removeOriginTransfers, name='removeOriginTransfers'),
	path('addDestinationTransfers/<int:warehouseId>/<DestinationTransfersIds>/', WarehouseView.addDestinationTransfers, name='addDestinationTransfers'),
	path('removeDestinationTransfers/<int:warehouseId>/<DestinationTransfersIds>/', WarehouseView.removeDestinationTransfers, name='removeDestinationTransfers'),
	path('addCycleCounts/<int:warehouseId>/<CycleCountsIds>/', WarehouseView.addCycleCounts, name='addCycleCounts'),
	path('removeCycleCounts/<int:warehouseId>/<CycleCountsIds>/', WarehouseView.removeCycleCounts, name='removeCycleCounts'),
]
