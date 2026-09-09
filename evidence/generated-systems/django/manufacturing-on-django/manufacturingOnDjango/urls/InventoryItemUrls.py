from django.urls import path
from manufacturingOnDjango.views import InventoryItemView

urlpatterns = [
    path('', InventoryItemView.index, name='index'),
	path('create', InventoryItemView.get, name='create'),
	path('get/<int:inventoryItemId>/', InventoryItemView.get, name='get'),
	path('save', InventoryItemView.save, name='save'),
	path('getAll', InventoryItemView.getAll, name='getAll'),
	path('delete/<int:inventoryItemId>/', InventoryItemView.delete, name='delete'),
	path('assignItem/<int:inventoryItemId>/<int:ItemId>/', InventoryItemView.assignItem, name='assignItem'),
	path('unassignItem/<int:inventoryItemId>/', InventoryItemView.unassignItem, name='unassignItem'),
	path('assignLocation/<int:inventoryItemId>/<int:LocationId>/', InventoryItemView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:inventoryItemId>/', InventoryItemView.unassignLocation, name='unassignLocation'),
]
