from django.urls import path
from inventoryOnDjango.views import LotView

urlpatterns = [
    path('', LotView.index, name='index'),
	path('create', LotView.get, name='create'),
	path('get/<int:lotId>/', LotView.get, name='get'),
	path('save', LotView.save, name='save'),
	path('getAll', LotView.getAll, name='getAll'),
	path('delete/<int:lotId>/', LotView.delete, name='delete'),
	path('assignSku/<int:lotId>/<int:SkuId>/', LotView.assignSku, name='assignSku'),
	path('unassignSku/<int:lotId>/', LotView.unassignSku, name='unassignSku'),
	path('addInventoryItems/<int:lotId>/<InventoryItemsIds>/', LotView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:lotId>/<InventoryItemsIds>/', LotView.removeInventoryItems, name='removeInventoryItems'),
]
