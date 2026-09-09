from django.urls import path
from ecommerceOnDjango.views import InventoryItemView

urlpatterns = [
    path('', InventoryItemView.index, name='index'),
	path('create', InventoryItemView.get, name='create'),
	path('get/<int:inventoryItemId>/', InventoryItemView.get, name='get'),
	path('save', InventoryItemView.save, name='save'),
	path('getAll', InventoryItemView.getAll, name='getAll'),
	path('delete/<int:inventoryItemId>/', InventoryItemView.delete, name='delete'),
	path('assignVariant/<int:inventoryItemId>/<int:VariantId>/', InventoryItemView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:inventoryItemId>/', InventoryItemView.unassignVariant, name='unassignVariant'),
	path('assignFulfillmentCenter/<int:inventoryItemId>/<int:FulfillmentCenterId>/', InventoryItemView.assignFulfillmentCenter, name='assignFulfillmentCenter'),
	path('unassignFulfillmentCenter/<int:inventoryItemId>/', InventoryItemView.unassignFulfillmentCenter, name='unassignFulfillmentCenter'),
]
