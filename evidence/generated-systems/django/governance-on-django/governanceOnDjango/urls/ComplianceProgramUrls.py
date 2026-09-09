from django.urls import path
from governanceOnDjango.views import ComplianceProgramView

urlpatterns = [
    path('', ComplianceProgramView.index, name='index'),
	path('create', ComplianceProgramView.get, name='create'),
	path('get/<int:complianceProgramId>/', ComplianceProgramView.get, name='get'),
	path('save', ComplianceProgramView.save, name='save'),
	path('getAll', ComplianceProgramView.getAll, name='getAll'),
	path('delete/<int:complianceProgramId>/', ComplianceProgramView.delete, name='delete'),
	path('assignOrganization/<int:complianceProgramId>/<int:OrganizationId>/', ComplianceProgramView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:complianceProgramId>/', ComplianceProgramView.unassignOrganization, name='unassignOrganization'),
	path('addRequirements/<int:complianceProgramId>/<RequirementsIds>/', ComplianceProgramView.addRequirements, name='addRequirements'),
	path('removeRequirements/<int:complianceProgramId>/<RequirementsIds>/', ComplianceProgramView.removeRequirements, name='removeRequirements'),
	path('addControls/<int:complianceProgramId>/<ControlsIds>/', ComplianceProgramView.addControls, name='addControls'),
	path('removeControls/<int:complianceProgramId>/<ControlsIds>/', ComplianceProgramView.removeControls, name='removeControls'),
	path('addAttestations/<int:complianceProgramId>/<AttestationsIds>/', ComplianceProgramView.addAttestations, name='addAttestations'),
	path('removeAttestations/<int:complianceProgramId>/<AttestationsIds>/', ComplianceProgramView.removeAttestations, name='removeAttestations'),
	path('addRegulations/<int:complianceProgramId>/<RegulationsIds>/', ComplianceProgramView.addRegulations, name='addRegulations'),
	path('removeRegulations/<int:complianceProgramId>/<RegulationsIds>/', ComplianceProgramView.removeRegulations, name='removeRegulations'),
]
