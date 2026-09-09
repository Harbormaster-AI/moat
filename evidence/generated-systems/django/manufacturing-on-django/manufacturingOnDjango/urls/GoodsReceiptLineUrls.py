from django.urls import path
from manufacturingOnDjango.views import GoodsReceiptLineView

urlpatterns = [
    path('', GoodsReceiptLineView.index, name='index'),
	path('create', GoodsReceiptLineView.get, name='create'),
	path('get/<int:goodsReceiptLineId>/', GoodsReceiptLineView.get, name='get'),
	path('save', GoodsReceiptLineView.save, name='save'),
	path('getAll', GoodsReceiptLineView.getAll, name='getAll'),
	path('delete/<int:goodsReceiptLineId>/', GoodsReceiptLineView.delete, name='delete'),
	path('assignGoodsReceipt/<int:goodsReceiptLineId>/<int:GoodsReceiptId>/', GoodsReceiptLineView.assignGoodsReceipt, name='assignGoodsReceipt'),
	path('unassignGoodsReceipt/<int:goodsReceiptLineId>/', GoodsReceiptLineView.unassignGoodsReceipt, name='unassignGoodsReceipt'),
	path('assignItem/<int:goodsReceiptLineId>/<int:ItemId>/', GoodsReceiptLineView.assignItem, name='assignItem'),
	path('unassignItem/<int:goodsReceiptLineId>/', GoodsReceiptLineView.unassignItem, name='unassignItem'),
	path('assignInventoryTransaction/<int:goodsReceiptLineId>/<int:InventoryTransactionId>/', GoodsReceiptLineView.assignInventoryTransaction, name='assignInventoryTransaction'),
	path('unassignInventoryTransaction/<int:goodsReceiptLineId>/', GoodsReceiptLineView.unassignInventoryTransaction, name='unassignInventoryTransaction'),
]
