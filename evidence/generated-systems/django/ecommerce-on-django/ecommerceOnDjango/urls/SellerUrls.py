from django.urls import path
from ecommerceOnDjango.views import SellerView

urlpatterns = [
    path('', SellerView.index, name='index'),
	path('create', SellerView.get, name='create'),
	path('get/<int:sellerId>/', SellerView.get, name='get'),
	path('save', SellerView.save, name='save'),
	path('getAll', SellerView.getAll, name='getAll'),
	path('delete/<int:sellerId>/', SellerView.delete, name='delete'),
	path('assignMerchant/<int:sellerId>/<int:MerchantId>/', SellerView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:sellerId>/', SellerView.unassignMerchant, name='unassignMerchant'),
	path('addProducts/<int:sellerId>/<ProductsIds>/', SellerView.addProducts, name='addProducts'),
	path('removeProducts/<int:sellerId>/<ProductsIds>/', SellerView.removeProducts, name='removeProducts'),
	path('addPayouts/<int:sellerId>/<PayoutsIds>/', SellerView.addPayouts, name='addPayouts'),
	path('removePayouts/<int:sellerId>/<PayoutsIds>/', SellerView.removePayouts, name='removePayouts'),
	path('addOrders/<int:sellerId>/<OrdersIds>/', SellerView.addOrders, name='addOrders'),
	path('removeOrders/<int:sellerId>/<OrdersIds>/', SellerView.removeOrders, name='removeOrders'),
]
