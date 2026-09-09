from django.urls import path
from governanceOnDjango.views import AttestationView

urlpatterns = [
    path('', AttestationView.index, name='index'),
	path('create', AttestationView.get, name='create'),
	path('get/<int:attestationId>/', AttestationView.get, name='get'),
	path('save', AttestationView.save, name='save'),
	path('getAll', AttestationView.getAll, name='getAll'),
	path('delete/<int:attestationId>/', AttestationView.delete, name='delete'),
	path('assignControl/<int:attestationId>/<int:ControlId>/', AttestationView.assignControl, name='assignControl'),
	path('unassignControl/<int:attestationId>/', AttestationView.unassignControl, name='unassignControl'),
	path('assignPolicy/<int:attestationId>/<int:PolicyId>/', AttestationView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:attestationId>/', AttestationView.unassignPolicy, name='unassignPolicy'),
	path('assignComplianceProgram/<int:attestationId>/<int:ComplianceProgramId>/', AttestationView.assignComplianceProgram, name='assignComplianceProgram'),
	path('unassignComplianceProgram/<int:attestationId>/', AttestationView.unassignComplianceProgram, name='unassignComplianceProgram'),
]
