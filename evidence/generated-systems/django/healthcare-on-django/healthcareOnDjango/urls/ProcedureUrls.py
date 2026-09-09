from django.urls import path
from healthcareOnDjango.views import ProcedureView

urlpatterns = [
    path('', ProcedureView.index, name='index'),
	path('create', ProcedureView.get, name='create'),
	path('get/<int:procedureId>/', ProcedureView.get, name='get'),
	path('save', ProcedureView.save, name='save'),
	path('getAll', ProcedureView.getAll, name='getAll'),
	path('delete/<int:procedureId>/', ProcedureView.delete, name='delete'),
	path('assignEncounter/<int:procedureId>/<int:EncounterId>/', ProcedureView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:procedureId>/', ProcedureView.unassignEncounter, name='unassignEncounter'),
	path('assignPerformer/<int:procedureId>/<int:PerformerId>/', ProcedureView.assignPerformer, name='assignPerformer'),
	path('unassignPerformer/<int:procedureId>/', ProcedureView.unassignPerformer, name='unassignPerformer'),
	path('assignProcedureOrder/<int:procedureId>/<int:ProcedureOrderId>/', ProcedureView.assignProcedureOrder, name='assignProcedureOrder'),
	path('unassignProcedureOrder/<int:procedureId>/', ProcedureView.unassignProcedureOrder, name='unassignProcedureOrder'),
]
