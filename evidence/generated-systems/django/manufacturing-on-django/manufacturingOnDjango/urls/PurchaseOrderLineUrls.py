from django.urls import path
from manufacturingOnDjango.views import PurchaseOrderLineView

urlpatterns = [
    path('', PurchaseOrderLineView.index, name='index'),
	path('create', PurchaseOrderLineView.get, name='create'),
	path('get/<int:purchaseOrderLineId>/', PurchaseOrderLineView.get, name='get'),
	path('save', PurchaseOrderLineView.save, name='save'),
	path('getAll', PurchaseOrderLineView.getAll, name='getAll'),
	path('delete/<int:purchaseOrderLineId>/', PurchaseOrderLineView.delete, name='delete'),
	path('assignPurchaseOrder/<int:purchaseOrderLineId>/<int:PurchaseOrderId>/', PurchaseOrderLineView.assignPurchaseOrder, name='assignPurchaseOrder'),
	path('unassignPurchaseOrder/<int:purchaseOrderLineId>/', PurchaseOrderLineView.unassignPurchaseOrder, name='unassignPurchaseOrder'),
	path('assignItem/<int:purchaseOrderLineId>/<int:ItemId>/', PurchaseOrderLineView.assignItem, name='assignItem'),
	path('unassignItem/<int:purchaseOrderLineId>/', PurchaseOrderLineView.unassignItem, name='unassignItem'),
]
