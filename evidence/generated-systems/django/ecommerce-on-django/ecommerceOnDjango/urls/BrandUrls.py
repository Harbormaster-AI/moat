from django.urls import path
from ecommerceOnDjango.views import BrandView

urlpatterns = [
    path('', BrandView.index, name='index'),
	path('create', BrandView.get, name='create'),
	path('get/<int:brandId>/', BrandView.get, name='get'),
	path('save', BrandView.save, name='save'),
	path('getAll', BrandView.getAll, name='getAll'),
	path('delete/<int:brandId>/', BrandView.delete, name='delete'),
	path('assignMerchant/<int:brandId>/<int:MerchantId>/', BrandView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:brandId>/', BrandView.unassignMerchant, name='unassignMerchant'),
	path('addProducts/<int:brandId>/<ProductsIds>/', BrandView.addProducts, name='addProducts'),
	path('removeProducts/<int:brandId>/<ProductsIds>/', BrandView.removeProducts, name='removeProducts'),
]
