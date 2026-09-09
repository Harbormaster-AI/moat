from django.urls import path
from manufacturingOnDjango.views import RoutingView

urlpatterns = [
    path('', RoutingView.index, name='index'),
	path('create', RoutingView.get, name='create'),
	path('get/<int:routingId>/', RoutingView.get, name='get'),
	path('save', RoutingView.save, name='save'),
	path('getAll', RoutingView.getAll, name='getAll'),
	path('delete/<int:routingId>/', RoutingView.delete, name='delete'),
	path('assignItem/<int:routingId>/<int:ItemId>/', RoutingView.assignItem, name='assignItem'),
	path('unassignItem/<int:routingId>/', RoutingView.unassignItem, name='unassignItem'),
	path('addOperations/<int:routingId>/<OperationsIds>/', RoutingView.addOperations, name='addOperations'),
	path('removeOperations/<int:routingId>/<OperationsIds>/', RoutingView.removeOperations, name='removeOperations'),
]
