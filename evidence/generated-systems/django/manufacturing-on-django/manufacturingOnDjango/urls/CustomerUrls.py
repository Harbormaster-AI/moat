from django.urls import path
from manufacturingOnDjango.views import CustomerView

urlpatterns = [
    path('', CustomerView.index, name='index'),
	path('create', CustomerView.get, name='create'),
	path('get/<int:customerId>/', CustomerView.get, name='get'),
	path('save', CustomerView.save, name='save'),
	path('getAll', CustomerView.getAll, name='getAll'),
	path('delete/<int:customerId>/', CustomerView.delete, name='delete'),
	path('addEnterprises/<int:customerId>/<EnterprisesIds>/', CustomerView.addEnterprises, name='addEnterprises'),
	path('removeEnterprises/<int:customerId>/<EnterprisesIds>/', CustomerView.removeEnterprises, name='removeEnterprises'),
	path('addSalesOrders/<int:customerId>/<SalesOrdersIds>/', CustomerView.addSalesOrders, name='addSalesOrders'),
	path('removeSalesOrders/<int:customerId>/<SalesOrdersIds>/', CustomerView.removeSalesOrders, name='removeSalesOrders'),
]
