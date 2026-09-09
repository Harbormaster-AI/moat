from django.urls import path
from healthcareOnDjango.views import ProcedureOrderView

urlpatterns = [
    path('', ProcedureOrderView.index, name='index'),
	path('create', ProcedureOrderView.get, name='create'),
	path('get/<int:procedureOrderId>/', ProcedureOrderView.get, name='get'),
	path('save', ProcedureOrderView.save, name='save'),
	path('getAll', ProcedureOrderView.getAll, name='getAll'),
	path('delete/<int:procedureOrderId>/', ProcedureOrderView.delete, name='delete'),
	path('assignOrder/<int:procedureOrderId>/<int:OrderId>/', ProcedureOrderView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:procedureOrderId>/', ProcedureOrderView.unassignOrder, name='unassignOrder'),
	path('assignFacility/<int:procedureOrderId>/<int:FacilityId>/', ProcedureOrderView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:procedureOrderId>/', ProcedureOrderView.unassignFacility, name='unassignFacility'),
	path('assignProcedure/<int:procedureOrderId>/<int:ProcedureId>/', ProcedureOrderView.assignProcedure, name='assignProcedure'),
	path('unassignProcedure/<int:procedureOrderId>/', ProcedureOrderView.unassignProcedure, name='unassignProcedure'),
]
