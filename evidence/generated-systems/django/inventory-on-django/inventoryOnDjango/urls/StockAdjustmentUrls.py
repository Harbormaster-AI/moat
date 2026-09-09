from django.urls import path
from inventoryOnDjango.views import StockAdjustmentView

urlpatterns = [
    path('', StockAdjustmentView.index, name='index'),
	path('create', StockAdjustmentView.get, name='create'),
	path('get/<int:stockAdjustmentId>/', StockAdjustmentView.get, name='get'),
	path('save', StockAdjustmentView.save, name='save'),
	path('getAll', StockAdjustmentView.getAll, name='getAll'),
	path('delete/<int:stockAdjustmentId>/', StockAdjustmentView.delete, name='delete'),
	path('assignWarehouse/<int:stockAdjustmentId>/<int:WarehouseId>/', StockAdjustmentView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:stockAdjustmentId>/', StockAdjustmentView.unassignWarehouse, name='unassignWarehouse'),
	path('addLines/<int:stockAdjustmentId>/<LinesIds>/', StockAdjustmentView.addLines, name='addLines'),
	path('removeLines/<int:stockAdjustmentId>/<LinesIds>/', StockAdjustmentView.removeLines, name='removeLines'),
	path('addTransactions/<int:stockAdjustmentId>/<TransactionsIds>/', StockAdjustmentView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:stockAdjustmentId>/<TransactionsIds>/', StockAdjustmentView.removeTransactions, name='removeTransactions'),
]
