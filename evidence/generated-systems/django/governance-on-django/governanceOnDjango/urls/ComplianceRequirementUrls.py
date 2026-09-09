from django.urls import path
from governanceOnDjango.views import ComplianceRequirementView

urlpatterns = [
    path('', ComplianceRequirementView.index, name='index'),
	path('create', ComplianceRequirementView.get, name='create'),
	path('get/<int:complianceRequirementId>/', ComplianceRequirementView.get, name='get'),
	path('save', ComplianceRequirementView.save, name='save'),
	path('getAll', ComplianceRequirementView.getAll, name='getAll'),
	path('delete/<int:complianceRequirementId>/', ComplianceRequirementView.delete, name='delete'),
	path('assignComplianceProgram/<int:complianceRequirementId>/<int:ComplianceProgramId>/', ComplianceRequirementView.assignComplianceProgram, name='assignComplianceProgram'),
	path('unassignComplianceProgram/<int:complianceRequirementId>/', ComplianceRequirementView.unassignComplianceProgram, name='unassignComplianceProgram'),
	path('addPolicies/<int:complianceRequirementId>/<PoliciesIds>/', ComplianceRequirementView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:complianceRequirementId>/<PoliciesIds>/', ComplianceRequirementView.removePolicies, name='removePolicies'),
	path('addControls/<int:complianceRequirementId>/<ControlsIds>/', ComplianceRequirementView.addControls, name='addControls'),
	path('removeControls/<int:complianceRequirementId>/<ControlsIds>/', ComplianceRequirementView.removeControls, name='removeControls'),
	path('addObligations/<int:complianceRequirementId>/<ObligationsIds>/', ComplianceRequirementView.addObligations, name='addObligations'),
	path('removeObligations/<int:complianceRequirementId>/<ObligationsIds>/', ComplianceRequirementView.removeObligations, name='removeObligations'),
]
