from django.urls import path
from manufacturingOnDjango.views import SupplierView

urlpatterns = [
    path('', SupplierView.index, name='index'),
	path('create', SupplierView.get, name='create'),
	path('get/<int:supplierId>/', SupplierView.get, name='get'),
	path('save', SupplierView.save, name='save'),
	path('getAll', SupplierView.getAll, name='getAll'),
	path('delete/<int:supplierId>/', SupplierView.delete, name='delete'),
	path('addEnterprises/<int:supplierId>/<EnterprisesIds>/', SupplierView.addEnterprises, name='addEnterprises'),
	path('removeEnterprises/<int:supplierId>/<EnterprisesIds>/', SupplierView.removeEnterprises, name='removeEnterprises'),
	path('addItems/<int:supplierId>/<ItemsIds>/', SupplierView.addItems, name='addItems'),
	path('removeItems/<int:supplierId>/<ItemsIds>/', SupplierView.removeItems, name='removeItems'),
	path('addPurchaseOrders/<int:supplierId>/<PurchaseOrdersIds>/', SupplierView.addPurchaseOrders, name='addPurchaseOrders'),
	path('removePurchaseOrders/<int:supplierId>/<PurchaseOrdersIds>/', SupplierView.removePurchaseOrders, name='removePurchaseOrders'),
]
