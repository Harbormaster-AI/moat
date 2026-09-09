from django.urls import path
from manufacturingOnDjango.views import SalesOrderLineView

urlpatterns = [
    path('', SalesOrderLineView.index, name='index'),
	path('create', SalesOrderLineView.get, name='create'),
	path('get/<int:salesOrderLineId>/', SalesOrderLineView.get, name='get'),
	path('save', SalesOrderLineView.save, name='save'),
	path('getAll', SalesOrderLineView.getAll, name='getAll'),
	path('delete/<int:salesOrderLineId>/', SalesOrderLineView.delete, name='delete'),
	path('assignSalesOrder/<int:salesOrderLineId>/<int:SalesOrderId>/', SalesOrderLineView.assignSalesOrder, name='assignSalesOrder'),
	path('unassignSalesOrder/<int:salesOrderLineId>/', SalesOrderLineView.unassignSalesOrder, name='unassignSalesOrder'),
	path('assignItem/<int:salesOrderLineId>/<int:ItemId>/', SalesOrderLineView.assignItem, name='assignItem'),
	path('unassignItem/<int:salesOrderLineId>/', SalesOrderLineView.unassignItem, name='unassignItem'),
]
