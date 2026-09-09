from django.urls import path
from inventoryOnDjango.views import StockKeepingUnitView

urlpatterns = [
    path('', StockKeepingUnitView.index, name='index'),
	path('create', StockKeepingUnitView.get, name='create'),
	path('get/<int:stockKeepingUnitId>/', StockKeepingUnitView.get, name='get'),
	path('save', StockKeepingUnitView.save, name='save'),
	path('getAll', StockKeepingUnitView.getAll, name='getAll'),
	path('delete/<int:stockKeepingUnitId>/', StockKeepingUnitView.delete, name='delete'),
	path('addInventoryItems/<int:stockKeepingUnitId>/<InventoryItemsIds>/', StockKeepingUnitView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:stockKeepingUnitId>/<InventoryItemsIds>/', StockKeepingUnitView.removeInventoryItems, name='removeInventoryItems'),
	path('addUomConversions/<int:stockKeepingUnitId>/<UomConversionsIds>/', StockKeepingUnitView.addUomConversions, name='addUomConversions'),
	path('removeUomConversions/<int:stockKeepingUnitId>/<UomConversionsIds>/', StockKeepingUnitView.removeUomConversions, name='removeUomConversions'),
	path('addReplenishmentPolicies/<int:stockKeepingUnitId>/<ReplenishmentPoliciesIds>/', StockKeepingUnitView.addReplenishmentPolicies, name='addReplenishmentPolicies'),
	path('removeReplenishmentPolicies/<int:stockKeepingUnitId>/<ReplenishmentPoliciesIds>/', StockKeepingUnitView.removeReplenishmentPolicies, name='removeReplenishmentPolicies'),
	path('addLots/<int:stockKeepingUnitId>/<LotsIds>/', StockKeepingUnitView.addLots, name='addLots'),
	path('removeLots/<int:stockKeepingUnitId>/<LotsIds>/', StockKeepingUnitView.removeLots, name='removeLots'),
	path('addSerialNumbers/<int:stockKeepingUnitId>/<SerialNumbersIds>/', StockKeepingUnitView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:stockKeepingUnitId>/<SerialNumbersIds>/', StockKeepingUnitView.removeSerialNumbers, name='removeSerialNumbers'),
]
