from django.urls import path
from healthcareOnDjango.views import CareTeamView

urlpatterns = [
    path('', CareTeamView.index, name='index'),
	path('create', CareTeamView.get, name='create'),
	path('get/<int:careTeamId>/', CareTeamView.get, name='get'),
	path('save', CareTeamView.save, name='save'),
	path('getAll', CareTeamView.getAll, name='getAll'),
	path('delete/<int:careTeamId>/', CareTeamView.delete, name='delete'),
	path('assignDepartment/<int:careTeamId>/<int:DepartmentId>/', CareTeamView.assignDepartment, name='assignDepartment'),
	path('unassignDepartment/<int:careTeamId>/', CareTeamView.unassignDepartment, name='unassignDepartment'),
	path('addClinicians/<int:careTeamId>/<CliniciansIds>/', CareTeamView.addClinicians, name='addClinicians'),
	path('removeClinicians/<int:careTeamId>/<CliniciansIds>/', CareTeamView.removeClinicians, name='removeClinicians'),
	path('addPatients/<int:careTeamId>/<PatientsIds>/', CareTeamView.addPatients, name='addPatients'),
	path('removePatients/<int:careTeamId>/<PatientsIds>/', CareTeamView.removePatients, name='removePatients'),
]
