from django.urls import path
from ecommerceOnDjango.views import CustomerAddressView

urlpatterns = [
    path('', CustomerAddressView.index, name='index'),
	path('create', CustomerAddressView.get, name='create'),
	path('get/<int:customerAddressId>/', CustomerAddressView.get, name='get'),
	path('save', CustomerAddressView.save, name='save'),
	path('getAll', CustomerAddressView.getAll, name='getAll'),
	path('delete/<int:customerAddressId>/', CustomerAddressView.delete, name='delete'),
	path('assignCustomer/<int:customerAddressId>/<int:CustomerId>/', CustomerAddressView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:customerAddressId>/', CustomerAddressView.unassignCustomer, name='unassignCustomer'),
]
