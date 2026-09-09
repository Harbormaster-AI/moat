from django.urls import path
from healthcareOnDjango.views import InventoryItemView

urlpatterns = [
    path('', InventoryItemView.index, name='index'),
	path('create', InventoryItemView.get, name='create'),
	path('get/<int:inventoryItemId>/', InventoryItemView.get, name='get'),
	path('save', InventoryItemView.save, name='save'),
	path('getAll', InventoryItemView.getAll, name='getAll'),
	path('delete/<int:inventoryItemId>/', InventoryItemView.delete, name='delete'),
	path('assignFacility/<int:inventoryItemId>/<int:FacilityId>/', InventoryItemView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:inventoryItemId>/', InventoryItemView.unassignFacility, name='unassignFacility'),
	path('assignSupplier/<int:inventoryItemId>/<int:SupplierId>/', InventoryItemView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:inventoryItemId>/', InventoryItemView.unassignSupplier, name='unassignSupplier'),
]
