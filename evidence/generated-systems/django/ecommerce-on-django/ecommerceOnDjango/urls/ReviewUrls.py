from django.urls import path
from ecommerceOnDjango.views import ReviewView

urlpatterns = [
    path('', ReviewView.index, name='index'),
	path('create', ReviewView.get, name='create'),
	path('get/<int:reviewId>/', ReviewView.get, name='get'),
	path('save', ReviewView.save, name='save'),
	path('getAll', ReviewView.getAll, name='getAll'),
	path('delete/<int:reviewId>/', ReviewView.delete, name='delete'),
	path('assignProduct/<int:reviewId>/<int:ProductId>/', ReviewView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:reviewId>/', ReviewView.unassignProduct, name='unassignProduct'),
	path('assignCustomer/<int:reviewId>/<int:CustomerId>/', ReviewView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:reviewId>/', ReviewView.unassignCustomer, name='unassignCustomer'),
	path('assignOrder/<int:reviewId>/<int:OrderId>/', ReviewView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:reviewId>/', ReviewView.unassignOrder, name='unassignOrder'),
]
