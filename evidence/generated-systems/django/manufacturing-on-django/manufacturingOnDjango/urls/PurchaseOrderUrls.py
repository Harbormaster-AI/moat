from django.urls import path
from manufacturingOnDjango.views import PurchaseOrderView

urlpatterns = [
    path('', PurchaseOrderView.index, name='index'),
	path('create', PurchaseOrderView.get, name='create'),
	path('get/<int:purchaseOrderId>/', PurchaseOrderView.get, name='get'),
	path('save', PurchaseOrderView.save, name='save'),
	path('getAll', PurchaseOrderView.getAll, name='getAll'),
	path('delete/<int:purchaseOrderId>/', PurchaseOrderView.delete, name='delete'),
	path('assignSupplier/<int:purchaseOrderId>/<int:SupplierId>/', PurchaseOrderView.assignSupplier, name='assignSupplier'),
	path('unassignSupplier/<int:purchaseOrderId>/', PurchaseOrderView.unassignSupplier, name='unassignSupplier'),
	path('assignPlant/<int:purchaseOrderId>/<int:PlantId>/', PurchaseOrderView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:purchaseOrderId>/', PurchaseOrderView.unassignPlant, name='unassignPlant'),
	path('addLines/<int:purchaseOrderId>/<LinesIds>/', PurchaseOrderView.addLines, name='addLines'),
	path('removeLines/<int:purchaseOrderId>/<LinesIds>/', PurchaseOrderView.removeLines, name='removeLines'),
	path('addGoodsReceipts/<int:purchaseOrderId>/<GoodsReceiptsIds>/', PurchaseOrderView.addGoodsReceipts, name='addGoodsReceipts'),
	path('removeGoodsReceipts/<int:purchaseOrderId>/<GoodsReceiptsIds>/', PurchaseOrderView.removeGoodsReceipts, name='removeGoodsReceipts'),
]
