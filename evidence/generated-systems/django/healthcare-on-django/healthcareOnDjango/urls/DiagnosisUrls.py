from django.urls import path
from healthcareOnDjango.views import DiagnosisView

urlpatterns = [
    path('', DiagnosisView.index, name='index'),
	path('create', DiagnosisView.get, name='create'),
	path('get/<int:diagnosisId>/', DiagnosisView.get, name='get'),
	path('save', DiagnosisView.save, name='save'),
	path('getAll', DiagnosisView.getAll, name='getAll'),
	path('delete/<int:diagnosisId>/', DiagnosisView.delete, name='delete'),
	path('assignEncounter/<int:diagnosisId>/<int:EncounterId>/', DiagnosisView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:diagnosisId>/', DiagnosisView.unassignEncounter, name='unassignEncounter'),
	path('assignPatient/<int:diagnosisId>/<int:PatientId>/', DiagnosisView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:diagnosisId>/', DiagnosisView.unassignPatient, name='unassignPatient'),
]
