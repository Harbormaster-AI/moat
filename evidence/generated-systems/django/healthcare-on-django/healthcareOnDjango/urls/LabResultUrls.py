from django.urls import path
from healthcareOnDjango.views import LabResultView

urlpatterns = [
    path('', LabResultView.index, name='index'),
	path('create', LabResultView.get, name='create'),
	path('get/<int:labResultId>/', LabResultView.get, name='get'),
	path('save', LabResultView.save, name='save'),
	path('getAll', LabResultView.getAll, name='getAll'),
	path('delete/<int:labResultId>/', LabResultView.delete, name='delete'),
	path('assignLaboratoryOrder/<int:labResultId>/<int:LaboratoryOrderId>/', LabResultView.assignLaboratoryOrder, name='assignLaboratoryOrder'),
	path('unassignLaboratoryOrder/<int:labResultId>/', LabResultView.unassignLaboratoryOrder, name='unassignLaboratoryOrder'),
	path('assignLaboratory/<int:labResultId>/<int:LaboratoryId>/', LabResultView.assignLaboratory, name='assignLaboratory'),
	path('unassignLaboratory/<int:labResultId>/', LabResultView.unassignLaboratory, name='unassignLaboratory'),
	path('addObservations/<int:labResultId>/<ObservationsIds>/', LabResultView.addObservations, name='addObservations'),
	path('removeObservations/<int:labResultId>/<ObservationsIds>/', LabResultView.removeObservations, name='removeObservations'),
]
