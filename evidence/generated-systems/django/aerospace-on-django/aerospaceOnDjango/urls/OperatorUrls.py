from django.urls import path
from aerospaceOnDjango.views import OperatorView

urlpatterns = [
    path('', OperatorView.index, name='index'),
	path('create', OperatorView.get, name='create'),
	path('get/<int:operatorId>/', OperatorView.get, name='get'),
	path('save', OperatorView.save, name='save'),
	path('getAll', OperatorView.getAll, name='getAll'),
	path('delete/<int:operatorId>/', OperatorView.delete, name='delete'),
	path('assignSalesRegion/<int:operatorId>/<int:SalesRegionId>/', OperatorView.assignSalesRegion, name='assignSalesRegion'),
	path('unassignSalesRegion/<int:operatorId>/', OperatorView.unassignSalesRegion, name='unassignSalesRegion'),
	path('addAircraftOrders/<int:operatorId>/<AircraftOrdersIds>/', OperatorView.addAircraftOrders, name='addAircraftOrders'),
	path('removeAircraftOrders/<int:operatorId>/<AircraftOrdersIds>/', OperatorView.removeAircraftOrders, name='removeAircraftOrders'),
	path('addOperatedAircraft/<int:operatorId>/<OperatedAircraftIds>/', OperatorView.addOperatedAircraft, name='addOperatedAircraft'),
	path('removeOperatedAircraft/<int:operatorId>/<OperatedAircraftIds>/', OperatorView.removeOperatedAircraft, name='removeOperatedAircraft'),
]
