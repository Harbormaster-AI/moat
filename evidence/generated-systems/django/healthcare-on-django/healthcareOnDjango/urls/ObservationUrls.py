from django.urls import path
from healthcareOnDjango.views import ObservationView

urlpatterns = [
    path('', ObservationView.index, name='index'),
	path('create', ObservationView.get, name='create'),
	path('get/<int:observationId>/', ObservationView.get, name='get'),
	path('save', ObservationView.save, name='save'),
	path('getAll', ObservationView.getAll, name='getAll'),
	path('delete/<int:observationId>/', ObservationView.delete, name='delete'),
	path('assignEncounter/<int:observationId>/<int:EncounterId>/', ObservationView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:observationId>/', ObservationView.unassignEncounter, name='unassignEncounter'),
	path('assignPatient/<int:observationId>/<int:PatientId>/', ObservationView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:observationId>/', ObservationView.unassignPatient, name='unassignPatient'),
	path('assignDevice/<int:observationId>/<int:DeviceId>/', ObservationView.assignDevice, name='assignDevice'),
	path('unassignDevice/<int:observationId>/', ObservationView.unassignDevice, name='unassignDevice'),
	path('assignLabResult/<int:observationId>/<int:LabResultId>/', ObservationView.assignLabResult, name='assignLabResult'),
	path('unassignLabResult/<int:observationId>/', ObservationView.unassignLabResult, name='unassignLabResult'),
]
