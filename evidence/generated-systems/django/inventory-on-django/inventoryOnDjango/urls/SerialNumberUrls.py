from django.urls import path
from inventoryOnDjango.views import SerialNumberView

urlpatterns = [
    path('', SerialNumberView.index, name='index'),
	path('create', SerialNumberView.get, name='create'),
	path('get/<int:serialNumberId>/', SerialNumberView.get, name='get'),
	path('save', SerialNumberView.save, name='save'),
	path('getAll', SerialNumberView.getAll, name='getAll'),
	path('delete/<int:serialNumberId>/', SerialNumberView.delete, name='delete'),
	path('assignSku/<int:serialNumberId>/<int:SkuId>/', SerialNumberView.assignSku, name='assignSku'),
	path('unassignSku/<int:serialNumberId>/', SerialNumberView.unassignSku, name='unassignSku'),
	path('assignCurrentInventoryItem/<int:serialNumberId>/<int:CurrentInventoryItemId>/', SerialNumberView.assignCurrentInventoryItem, name='assignCurrentInventoryItem'),
	path('unassignCurrentInventoryItem/<int:serialNumberId>/', SerialNumberView.unassignCurrentInventoryItem, name='unassignCurrentInventoryItem'),
	path('assignLot/<int:serialNumberId>/<int:LotId>/', SerialNumberView.assignLot, name='assignLot'),
	path('unassignLot/<int:serialNumberId>/', SerialNumberView.unassignLot, name='unassignLot'),
]
