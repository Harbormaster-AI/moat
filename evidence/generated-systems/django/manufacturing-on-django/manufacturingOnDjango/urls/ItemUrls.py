from django.urls import path
from manufacturingOnDjango.views import ItemView

urlpatterns = [
    path('', ItemView.index, name='index'),
	path('create', ItemView.get, name='create'),
	path('get/<int:itemId>/', ItemView.get, name='get'),
	path('save', ItemView.save, name='save'),
	path('getAll', ItemView.getAll, name='getAll'),
	path('delete/<int:itemId>/', ItemView.delete, name='delete'),
	path('assignBusinessUnit/<int:itemId>/<int:BusinessUnitId>/', ItemView.assignBusinessUnit, name='assignBusinessUnit'),
	path('unassignBusinessUnit/<int:itemId>/', ItemView.unassignBusinessUnit, name='unassignBusinessUnit'),
	path('addBoms/<int:itemId>/<BomsIds>/', ItemView.addBoms, name='addBoms'),
	path('removeBoms/<int:itemId>/<BomsIds>/', ItemView.removeBoms, name='removeBoms'),
	path('addRoutings/<int:itemId>/<RoutingsIds>/', ItemView.addRoutings, name='addRoutings'),
	path('removeRoutings/<int:itemId>/<RoutingsIds>/', ItemView.removeRoutings, name='removeRoutings'),
	path('addSuppliers/<int:itemId>/<SuppliersIds>/', ItemView.addSuppliers, name='addSuppliers'),
	path('removeSuppliers/<int:itemId>/<SuppliersIds>/', ItemView.removeSuppliers, name='removeSuppliers'),
	path('addQualitySpecifications/<int:itemId>/<QualitySpecificationsIds>/', ItemView.addQualitySpecifications, name='addQualitySpecifications'),
	path('removeQualitySpecifications/<int:itemId>/<QualitySpecificationsIds>/', ItemView.removeQualitySpecifications, name='removeQualitySpecifications'),
	path('addInventoryItems/<int:itemId>/<InventoryItemsIds>/', ItemView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:itemId>/<InventoryItemsIds>/', ItemView.removeInventoryItems, name='removeInventoryItems'),
]
