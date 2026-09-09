from django.urls import path
from advertisingOnDjango.views import InventorySourceView

urlpatterns = [
    path('', InventorySourceView.index, name='index'),
	path('create', InventorySourceView.get, name='create'),
	path('get/<int:inventorySourceId>/', InventorySourceView.get, name='get'),
	path('save', InventorySourceView.save, name='save'),
	path('getAll', InventorySourceView.getAll, name='getAll'),
	path('delete/<int:inventorySourceId>/', InventorySourceView.delete, name='delete'),
	path('assignPublisher/<int:inventorySourceId>/<int:PublisherId>/', InventorySourceView.assignPublisher, name='assignPublisher'),
	path('unassignPublisher/<int:inventorySourceId>/', InventorySourceView.unassignPublisher, name='unassignPublisher'),
	path('addAdSlots/<int:inventorySourceId>/<AdSlotsIds>/', InventorySourceView.addAdSlots, name='addAdSlots'),
	path('removeAdSlots/<int:inventorySourceId>/<AdSlotsIds>/', InventorySourceView.removeAdSlots, name='removeAdSlots'),
	path('addDeals/<int:inventorySourceId>/<DealsIds>/', InventorySourceView.addDeals, name='addDeals'),
	path('removeDeals/<int:inventorySourceId>/<DealsIds>/', InventorySourceView.removeDeals, name='removeDeals'),
]
