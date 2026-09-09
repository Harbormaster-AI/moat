from django.urls import path
from healthcareOnDjango.views import AdmissionView

urlpatterns = [
    path('', AdmissionView.index, name='index'),
	path('create', AdmissionView.get, name='create'),
	path('get/<int:admissionId>/', AdmissionView.get, name='get'),
	path('save', AdmissionView.save, name='save'),
	path('getAll', AdmissionView.getAll, name='getAll'),
	path('delete/<int:admissionId>/', AdmissionView.delete, name='delete'),
	path('assignEncounter/<int:admissionId>/<int:EncounterId>/', AdmissionView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:admissionId>/', AdmissionView.unassignEncounter, name='unassignEncounter'),
	path('assignFacility/<int:admissionId>/<int:FacilityId>/', AdmissionView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:admissionId>/', AdmissionView.unassignFacility, name='unassignFacility'),
]
