from django.urls import path
from manufacturingOnDjango.views import SalesOrderView

urlpatterns = [
    path('', SalesOrderView.index, name='index'),
	path('create', SalesOrderView.get, name='create'),
	path('get/<int:salesOrderId>/', SalesOrderView.get, name='get'),
	path('save', SalesOrderView.save, name='save'),
	path('getAll', SalesOrderView.getAll, name='getAll'),
	path('delete/<int:salesOrderId>/', SalesOrderView.delete, name='delete'),
	path('assignCustomer/<int:salesOrderId>/<int:CustomerId>/', SalesOrderView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:salesOrderId>/', SalesOrderView.unassignCustomer, name='unassignCustomer'),
	path('assignPlant/<int:salesOrderId>/<int:PlantId>/', SalesOrderView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:salesOrderId>/', SalesOrderView.unassignPlant, name='unassignPlant'),
	path('addLines/<int:salesOrderId>/<LinesIds>/', SalesOrderView.addLines, name='addLines'),
	path('removeLines/<int:salesOrderId>/<LinesIds>/', SalesOrderView.removeLines, name='removeLines'),
	path('addWorkOrders/<int:salesOrderId>/<WorkOrdersIds>/', SalesOrderView.addWorkOrders, name='addWorkOrders'),
	path('removeWorkOrders/<int:salesOrderId>/<WorkOrdersIds>/', SalesOrderView.removeWorkOrders, name='removeWorkOrders'),
]
