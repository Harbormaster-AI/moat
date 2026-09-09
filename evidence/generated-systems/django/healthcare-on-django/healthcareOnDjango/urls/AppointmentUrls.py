from django.urls import path
from healthcareOnDjango.views import AppointmentView

urlpatterns = [
    path('', AppointmentView.index, name='index'),
	path('create', AppointmentView.get, name='create'),
	path('get/<int:appointmentId>/', AppointmentView.get, name='get'),
	path('save', AppointmentView.save, name='save'),
	path('getAll', AppointmentView.getAll, name='getAll'),
	path('delete/<int:appointmentId>/', AppointmentView.delete, name='delete'),
	path('assignPatient/<int:appointmentId>/<int:PatientId>/', AppointmentView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:appointmentId>/', AppointmentView.unassignPatient, name='unassignPatient'),
	path('assignClinician/<int:appointmentId>/<int:ClinicianId>/', AppointmentView.assignClinician, name='assignClinician'),
	path('unassignClinician/<int:appointmentId>/', AppointmentView.unassignClinician, name='unassignClinician'),
	path('assignFacility/<int:appointmentId>/<int:FacilityId>/', AppointmentView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:appointmentId>/', AppointmentView.unassignFacility, name='unassignFacility'),
	path('assignEncounter/<int:appointmentId>/<int:EncounterId>/', AppointmentView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:appointmentId>/', AppointmentView.unassignEncounter, name='unassignEncounter'),
]
