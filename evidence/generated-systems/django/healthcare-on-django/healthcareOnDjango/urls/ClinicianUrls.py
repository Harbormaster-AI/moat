from django.urls import path
from healthcareOnDjango.views import ClinicianView

urlpatterns = [
    path('', ClinicianView.index, name='index'),
	path('create', ClinicianView.get, name='create'),
	path('get/<int:clinicianId>/', ClinicianView.get, name='get'),
	path('save', ClinicianView.save, name='save'),
	path('getAll', ClinicianView.getAll, name='getAll'),
	path('delete/<int:clinicianId>/', ClinicianView.delete, name='delete'),
	path('addCareTeams/<int:clinicianId>/<CareTeamsIds>/', ClinicianView.addCareTeams, name='addCareTeams'),
	path('removeCareTeams/<int:clinicianId>/<CareTeamsIds>/', ClinicianView.removeCareTeams, name='removeCareTeams'),
	path('addAppointments/<int:clinicianId>/<AppointmentsIds>/', ClinicianView.addAppointments, name='addAppointments'),
	path('removeAppointments/<int:clinicianId>/<AppointmentsIds>/', ClinicianView.removeAppointments, name='removeAppointments'),
	path('addEncounters/<int:clinicianId>/<EncountersIds>/', ClinicianView.addEncounters, name='addEncounters'),
	path('removeEncounters/<int:clinicianId>/<EncountersIds>/', ClinicianView.removeEncounters, name='removeEncounters'),
	path('addProcedures/<int:clinicianId>/<ProceduresIds>/', ClinicianView.addProcedures, name='addProcedures'),
	path('removeProcedures/<int:clinicianId>/<ProceduresIds>/', ClinicianView.removeProcedures, name='removeProcedures'),
	path('addImagingReports/<int:clinicianId>/<ImagingReportsIds>/', ClinicianView.addImagingReports, name='addImagingReports'),
	path('removeImagingReports/<int:clinicianId>/<ImagingReportsIds>/', ClinicianView.removeImagingReports, name='removeImagingReports'),
]
