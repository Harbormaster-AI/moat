from django.urls import path
from inventoryOnDjango.views import StockAdjustmentLineView

urlpatterns = [
    path('', StockAdjustmentLineView.index, name='index'),
	path('create', StockAdjustmentLineView.get, name='create'),
	path('get/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.get, name='get'),
	path('save', StockAdjustmentLineView.save, name='save'),
	path('getAll', StockAdjustmentLineView.getAll, name='getAll'),
	path('delete/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.delete, name='delete'),
	path('assignAdjustment/<int:stockAdjustmentLineId>/<int:AdjustmentId>/', StockAdjustmentLineView.assignAdjustment, name='assignAdjustment'),
	path('unassignAdjustment/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.unassignAdjustment, name='unassignAdjustment'),
	path('assignSku/<int:stockAdjustmentLineId>/<int:SkuId>/', StockAdjustmentLineView.assignSku, name='assignSku'),
	path('unassignSku/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.unassignSku, name='unassignSku'),
	path('assignLot/<int:stockAdjustmentLineId>/<int:LotId>/', StockAdjustmentLineView.assignLot, name='assignLot'),
	path('unassignLot/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.unassignLot, name='unassignLot'),
	path('assignLocation/<int:stockAdjustmentLineId>/<int:LocationId>/', StockAdjustmentLineView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:stockAdjustmentLineId>/', StockAdjustmentLineView.unassignLocation, name='unassignLocation'),
	path('addSerialNumbers/<int:stockAdjustmentLineId>/<SerialNumbersIds>/', StockAdjustmentLineView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:stockAdjustmentLineId>/<SerialNumbersIds>/', StockAdjustmentLineView.removeSerialNumbers, name='removeSerialNumbers'),
]
