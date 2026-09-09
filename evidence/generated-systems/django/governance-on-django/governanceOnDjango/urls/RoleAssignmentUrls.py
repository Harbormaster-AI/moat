from django.urls import path
from governanceOnDjango.views import RoleAssignmentView

urlpatterns = [
    path('', RoleAssignmentView.index, name='index'),
	path('create', RoleAssignmentView.get, name='create'),
	path('get/<int:roleAssignmentId>/', RoleAssignmentView.get, name='get'),
	path('save', RoleAssignmentView.save, name='save'),
	path('getAll', RoleAssignmentView.getAll, name='getAll'),
	path('delete/<int:roleAssignmentId>/', RoleAssignmentView.delete, name='delete'),
	path('assignPerson/<int:roleAssignmentId>/<int:PersonId>/', RoleAssignmentView.assignPerson, name='assignPerson'),
	path('unassignPerson/<int:roleAssignmentId>/', RoleAssignmentView.unassignPerson, name='unassignPerson'),
	path('assignRole/<int:roleAssignmentId>/<int:RoleId>/', RoleAssignmentView.assignRole, name='assignRole'),
	path('unassignRole/<int:roleAssignmentId>/', RoleAssignmentView.unassignRole, name='unassignRole'),
	path('assignGovernanceBody/<int:roleAssignmentId>/<int:GovernanceBodyId>/', RoleAssignmentView.assignGovernanceBody, name='assignGovernanceBody'),
	path('unassignGovernanceBody/<int:roleAssignmentId>/', RoleAssignmentView.unassignGovernanceBody, name='unassignGovernanceBody'),
	path('assignOrganization/<int:roleAssignmentId>/<int:OrganizationId>/', RoleAssignmentView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:roleAssignmentId>/', RoleAssignmentView.unassignOrganization, name='unassignOrganization'),
]
