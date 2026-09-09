from django.urls import path
from aerospaceOnDjango.views import AirworthinessDirectiveView

urlpatterns = [
    path('', AirworthinessDirectiveView.index, name='index'),
	path('create', AirworthinessDirectiveView.get, name='create'),
	path('get/<int:airworthinessDirectiveId>/', AirworthinessDirectiveView.get, name='get'),
	path('save', AirworthinessDirectiveView.save, name='save'),
	path('getAll', AirworthinessDirectiveView.getAll, name='getAll'),
	path('delete/<int:airworthinessDirectiveId>/', AirworthinessDirectiveView.delete, name='delete'),
	path('addWorkOrders/<int:airworthinessDirectiveId>/<WorkOrdersIds>/', AirworthinessDirectiveView.addWorkOrders, name='addWorkOrders'),
	path('removeWorkOrders/<int:airworthinessDirectiveId>/<WorkOrdersIds>/', AirworthinessDirectiveView.removeWorkOrders, name='removeWorkOrders'),
]
