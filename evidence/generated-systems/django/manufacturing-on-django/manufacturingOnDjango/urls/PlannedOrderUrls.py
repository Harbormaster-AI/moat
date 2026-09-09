from django.urls import path
from manufacturingOnDjango.views import PlannedOrderView

urlpatterns = [
    path('', PlannedOrderView.index, name='index'),
	path('create', PlannedOrderView.get, name='create'),
	path('get/<int:plannedOrderId>/', PlannedOrderView.get, name='get'),
	path('save', PlannedOrderView.save, name='save'),
	path('getAll', PlannedOrderView.getAll, name='getAll'),
	path('delete/<int:plannedOrderId>/', PlannedOrderView.delete, name='delete'),
	path('assignMrpRun/<int:plannedOrderId>/<int:MrpRunId>/', PlannedOrderView.assignMrpRun, name='assignMrpRun'),
	path('unassignMrpRun/<int:plannedOrderId>/', PlannedOrderView.unassignMrpRun, name='unassignMrpRun'),
	path('assignItem/<int:plannedOrderId>/<int:ItemId>/', PlannedOrderView.assignItem, name='assignItem'),
	path('unassignItem/<int:plannedOrderId>/', PlannedOrderView.unassignItem, name='unassignItem'),
	path('assignPlant/<int:plannedOrderId>/<int:PlantId>/', PlannedOrderView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:plannedOrderId>/', PlannedOrderView.unassignPlant, name='unassignPlant'),
]
