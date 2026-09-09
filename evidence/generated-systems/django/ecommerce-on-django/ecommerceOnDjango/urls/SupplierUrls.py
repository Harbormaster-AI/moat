from django.urls import path
from ecommerceOnDjango.views import SupplierView

urlpatterns = [
    path('', SupplierView.index, name='index'),
	path('create', SupplierView.get, name='create'),
	path('get/<int:supplierId>/', SupplierView.get, name='get'),
	path('save', SupplierView.save, name='save'),
	path('getAll', SupplierView.getAll, name='getAll'),
	path('delete/<int:supplierId>/', SupplierView.delete, name='delete'),
	path('assignMerchant/<int:supplierId>/<int:MerchantId>/', SupplierView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:supplierId>/', SupplierView.unassignMerchant, name='unassignMerchant'),
	path('addProducts/<int:supplierId>/<ProductsIds>/', SupplierView.addProducts, name='addProducts'),
	path('removeProducts/<int:supplierId>/<ProductsIds>/', SupplierView.removeProducts, name='removeProducts'),
	path('addFulfillmentCenters/<int:supplierId>/<FulfillmentCentersIds>/', SupplierView.addFulfillmentCenters, name='addFulfillmentCenters'),
	path('removeFulfillmentCenters/<int:supplierId>/<FulfillmentCentersIds>/', SupplierView.removeFulfillmentCenters, name='removeFulfillmentCenters'),
]
