from django.urls import path
from governanceOnDjango.views import GovernanceBodyView

urlpatterns = [
    path('', GovernanceBodyView.index, name='index'),
	path('create', GovernanceBodyView.get, name='create'),
	path('get/<int:governanceBodyId>/', GovernanceBodyView.get, name='get'),
	path('save', GovernanceBodyView.save, name='save'),
	path('getAll', GovernanceBodyView.getAll, name='getAll'),
	path('delete/<int:governanceBodyId>/', GovernanceBodyView.delete, name='delete'),
	path('assignOrganization/<int:governanceBodyId>/<int:OrganizationId>/', GovernanceBodyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:governanceBodyId>/', GovernanceBodyView.unassignOrganization, name='unassignOrganization'),
	path('addRoleAssignments/<int:governanceBodyId>/<RoleAssignmentsIds>/', GovernanceBodyView.addRoleAssignments, name='addRoleAssignments'),
	path('removeRoleAssignments/<int:governanceBodyId>/<RoleAssignmentsIds>/', GovernanceBodyView.removeRoleAssignments, name='removeRoleAssignments'),
	path('addPolicies/<int:governanceBodyId>/<PoliciesIds>/', GovernanceBodyView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:governanceBodyId>/<PoliciesIds>/', GovernanceBodyView.removePolicies, name='removePolicies'),
]
