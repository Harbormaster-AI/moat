from django.urls import path
from healthcareOnDjango.views import MedicalDeviceView

urlpatterns = [
    path('', MedicalDeviceView.index, name='index'),
	path('create', MedicalDeviceView.get, name='create'),
	path('get/<int:medicalDeviceId>/', MedicalDeviceView.get, name='get'),
	path('save', MedicalDeviceView.save, name='save'),
	path('getAll', MedicalDeviceView.getAll, name='getAll'),
	path('delete/<int:medicalDeviceId>/', MedicalDeviceView.delete, name='delete'),
	path('assignPatient/<int:medicalDeviceId>/<int:PatientId>/', MedicalDeviceView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:medicalDeviceId>/', MedicalDeviceView.unassignPatient, name='unassignPatient'),
	path('addObservations/<int:medicalDeviceId>/<ObservationsIds>/', MedicalDeviceView.addObservations, name='addObservations'),
	path('removeObservations/<int:medicalDeviceId>/<ObservationsIds>/', MedicalDeviceView.removeObservations, name='removeObservations'),
	path('addSoftwareUpdates/<int:medicalDeviceId>/<SoftwareUpdatesIds>/', MedicalDeviceView.addSoftwareUpdates, name='addSoftwareUpdates'),
	path('removeSoftwareUpdates/<int:medicalDeviceId>/<SoftwareUpdatesIds>/', MedicalDeviceView.removeSoftwareUpdates, name='removeSoftwareUpdates'),
]
