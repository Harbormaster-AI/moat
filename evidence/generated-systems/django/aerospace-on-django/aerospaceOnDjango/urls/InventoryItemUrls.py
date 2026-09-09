from django.urls import path
from aerospaceOnDjango.views import InventoryItemView

urlpatterns = [
    path('', InventoryItemView.index, name='index'),
	path('create', InventoryItemView.get, name='create'),
	path('get/<int:inventoryItemId>/', InventoryItemView.get, name='get'),
	path('save', InventoryItemView.save, name='save'),
	path('getAll', InventoryItemView.getAll, name='getAll'),
	path('delete/<int:inventoryItemId>/', InventoryItemView.delete, name='delete'),
	path('assignComponent/<int:inventoryItemId>/<int:ComponentId>/', InventoryItemView.assignComponent, name='assignComponent'),
	path('unassignComponent/<int:inventoryItemId>/', InventoryItemView.unassignComponent, name='unassignComponent'),
	path('assignWarehouse/<int:inventoryItemId>/<int:WarehouseId>/', InventoryItemView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:inventoryItemId>/', InventoryItemView.unassignWarehouse, name='unassignWarehouse'),
]
