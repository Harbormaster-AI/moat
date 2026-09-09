from django.urls import path
from manufacturingOnDjango.views import WorkOrderView

urlpatterns = [
    path('', WorkOrderView.index, name='index'),
	path('create', WorkOrderView.get, name='create'),
	path('get/<int:workOrderId>/', WorkOrderView.get, name='get'),
	path('save', WorkOrderView.save, name='save'),
	path('getAll', WorkOrderView.getAll, name='getAll'),
	path('delete/<int:workOrderId>/', WorkOrderView.delete, name='delete'),
	path('assignItem/<int:workOrderId>/<int:ItemId>/', WorkOrderView.assignItem, name='assignItem'),
	path('unassignItem/<int:workOrderId>/', WorkOrderView.unassignItem, name='unassignItem'),
	path('assignPlant/<int:workOrderId>/<int:PlantId>/', WorkOrderView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:workOrderId>/', WorkOrderView.unassignPlant, name='unassignPlant'),
	path('assignRouting/<int:workOrderId>/<int:RoutingId>/', WorkOrderView.assignRouting, name='assignRouting'),
	path('unassignRouting/<int:workOrderId>/', WorkOrderView.unassignRouting, name='unassignRouting'),
	path('assignBom/<int:workOrderId>/<int:BomId>/', WorkOrderView.assignBom, name='assignBom'),
	path('unassignBom/<int:workOrderId>/', WorkOrderView.unassignBom, name='unassignBom'),
	path('assignProductionSchedule/<int:workOrderId>/<int:ProductionScheduleId>/', WorkOrderView.assignProductionSchedule, name='assignProductionSchedule'),
	path('unassignProductionSchedule/<int:workOrderId>/', WorkOrderView.unassignProductionSchedule, name='unassignProductionSchedule'),
	path('assignSalesOrder/<int:workOrderId>/<int:SalesOrderId>/', WorkOrderView.assignSalesOrder, name='assignSalesOrder'),
	path('unassignSalesOrder/<int:workOrderId>/', WorkOrderView.unassignSalesOrder, name='unassignSalesOrder'),
]
