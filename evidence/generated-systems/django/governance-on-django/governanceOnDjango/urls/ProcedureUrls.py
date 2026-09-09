from django.urls import path
from governanceOnDjango.views import ProcedureView

urlpatterns = [
    path('', ProcedureView.index, name='index'),
	path('create', ProcedureView.get, name='create'),
	path('get/<int:procedureId>/', ProcedureView.get, name='get'),
	path('save', ProcedureView.save, name='save'),
	path('getAll', ProcedureView.getAll, name='getAll'),
	path('delete/<int:procedureId>/', ProcedureView.delete, name='delete'),
	path('assignPolicy/<int:procedureId>/<int:PolicyId>/', ProcedureView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:procedureId>/', ProcedureView.unassignPolicy, name='unassignPolicy'),
	path('addControls/<int:procedureId>/<ControlsIds>/', ProcedureView.addControls, name='addControls'),
	path('removeControls/<int:procedureId>/<ControlsIds>/', ProcedureView.removeControls, name='removeControls'),
]
