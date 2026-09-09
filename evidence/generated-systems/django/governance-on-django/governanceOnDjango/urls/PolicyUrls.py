from django.urls import path
from governanceOnDjango.views import PolicyView

urlpatterns = [
    path('', PolicyView.index, name='index'),
	path('create', PolicyView.get, name='create'),
	path('get/<int:policyId>/', PolicyView.get, name='get'),
	path('save', PolicyView.save, name='save'),
	path('getAll', PolicyView.getAll, name='getAll'),
	path('delete/<int:policyId>/', PolicyView.delete, name='delete'),
	path('assignOrganization/<int:policyId>/<int:OrganizationId>/', PolicyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:policyId>/', PolicyView.unassignOrganization, name='unassignOrganization'),
	path('addOwners/<int:policyId>/<OwnersIds>/', PolicyView.addOwners, name='addOwners'),
	path('removeOwners/<int:policyId>/<OwnersIds>/', PolicyView.removeOwners, name='removeOwners'),
	path('addRelatedRequirements/<int:policyId>/<RelatedRequirementsIds>/', PolicyView.addRelatedRequirements, name='addRelatedRequirements'),
	path('removeRelatedRequirements/<int:policyId>/<RelatedRequirementsIds>/', PolicyView.removeRelatedRequirements, name='removeRelatedRequirements'),
	path('addControls/<int:policyId>/<ControlsIds>/', PolicyView.addControls, name='addControls'),
	path('removeControls/<int:policyId>/<ControlsIds>/', PolicyView.removeControls, name='removeControls'),
	path('addProcedures/<int:policyId>/<ProceduresIds>/', PolicyView.addProcedures, name='addProcedures'),
	path('removeProcedures/<int:policyId>/<ProceduresIds>/', PolicyView.removeProcedures, name='removeProcedures'),
	path('addExceptions/<int:policyId>/<ExceptionsIds>/', PolicyView.addExceptions, name='addExceptions'),
	path('removeExceptions/<int:policyId>/<ExceptionsIds>/', PolicyView.removeExceptions, name='removeExceptions'),
	path('addAttestations/<int:policyId>/<AttestationsIds>/', PolicyView.addAttestations, name='addAttestations'),
	path('removeAttestations/<int:policyId>/<AttestationsIds>/', PolicyView.removeAttestations, name='removeAttestations'),
]
