from django.urls import path
from aerospaceOnDjango.views import BuildScheduleView

urlpatterns = [
    path('', BuildScheduleView.index, name='index'),
	path('create', BuildScheduleView.get, name='create'),
	path('get/<int:buildScheduleId>/', BuildScheduleView.get, name='get'),
	path('save', BuildScheduleView.save, name='save'),
	path('getAll', BuildScheduleView.getAll, name='getAll'),
	path('delete/<int:buildScheduleId>/', BuildScheduleView.delete, name='delete'),
	path('addProductionOrders/<int:buildScheduleId>/<ProductionOrdersIds>/', BuildScheduleView.addProductionOrders, name='addProductionOrders'),
	path('removeProductionOrders/<int:buildScheduleId>/<ProductionOrdersIds>/', BuildScheduleView.removeProductionOrders, name='removeProductionOrders'),
]
