from django.urls import path
from manufacturingOnDjango.views import GoodsReceiptView

urlpatterns = [
    path('', GoodsReceiptView.index, name='index'),
	path('create', GoodsReceiptView.get, name='create'),
	path('get/<int:goodsReceiptId>/', GoodsReceiptView.get, name='get'),
	path('save', GoodsReceiptView.save, name='save'),
	path('getAll', GoodsReceiptView.getAll, name='getAll'),
	path('delete/<int:goodsReceiptId>/', GoodsReceiptView.delete, name='delete'),
	path('assignPurchaseOrder/<int:goodsReceiptId>/<int:PurchaseOrderId>/', GoodsReceiptView.assignPurchaseOrder, name='assignPurchaseOrder'),
	path('unassignPurchaseOrder/<int:goodsReceiptId>/', GoodsReceiptView.unassignPurchaseOrder, name='unassignPurchaseOrder'),
	path('assignWarehouse/<int:goodsReceiptId>/<int:WarehouseId>/', GoodsReceiptView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:goodsReceiptId>/', GoodsReceiptView.unassignWarehouse, name='unassignWarehouse'),
	path('addLines/<int:goodsReceiptId>/<LinesIds>/', GoodsReceiptView.addLines, name='addLines'),
	path('removeLines/<int:goodsReceiptId>/<LinesIds>/', GoodsReceiptView.removeLines, name='removeLines'),
]
