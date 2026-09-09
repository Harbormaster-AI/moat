from django.urls import path
from manufacturingOnDjango.views import ProductionScheduleView

urlpatterns = [
    path('', ProductionScheduleView.index, name='index'),
	path('create', ProductionScheduleView.get, name='create'),
	path('get/<int:productionScheduleId>/', ProductionScheduleView.get, name='get'),
	path('save', ProductionScheduleView.save, name='save'),
	path('getAll', ProductionScheduleView.getAll, name='getAll'),
	path('delete/<int:productionScheduleId>/', ProductionScheduleView.delete, name='delete'),
	path('assignPlant/<int:productionScheduleId>/<int:PlantId>/', ProductionScheduleView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:productionScheduleId>/', ProductionScheduleView.unassignPlant, name='unassignPlant'),
	path('addWorkOrders/<int:productionScheduleId>/<WorkOrdersIds>/', ProductionScheduleView.addWorkOrders, name='addWorkOrders'),
	path('removeWorkOrders/<int:productionScheduleId>/<WorkOrdersIds>/', ProductionScheduleView.removeWorkOrders, name='removeWorkOrders'),
]
