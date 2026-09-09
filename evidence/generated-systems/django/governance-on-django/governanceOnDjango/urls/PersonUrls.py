from django.urls import path
from governanceOnDjango.views import PersonView

urlpatterns = [
    path('', PersonView.index, name='index'),
	path('create', PersonView.get, name='create'),
	path('get/<int:personId>/', PersonView.get, name='get'),
	path('save', PersonView.save, name='save'),
	path('getAll', PersonView.getAll, name='getAll'),
	path('delete/<int:personId>/', PersonView.delete, name='delete'),
	path('addRoleAssignments/<int:personId>/<RoleAssignmentsIds>/', PersonView.addRoleAssignments, name='addRoleAssignments'),
	path('removeRoleAssignments/<int:personId>/<RoleAssignmentsIds>/', PersonView.removeRoleAssignments, name='removeRoleAssignments'),
	path('addOwnedPolicies/<int:personId>/<OwnedPoliciesIds>/', PersonView.addOwnedPolicies, name='addOwnedPolicies'),
	path('removeOwnedPolicies/<int:personId>/<OwnedPoliciesIds>/', PersonView.removeOwnedPolicies, name='removeOwnedPolicies'),
	path('addCorrectiveActions/<int:personId>/<CorrectiveActionsIds>/', PersonView.addCorrectiveActions, name='addCorrectiveActions'),
	path('removeCorrectiveActions/<int:personId>/<CorrectiveActionsIds>/', PersonView.removeCorrectiveActions, name='removeCorrectiveActions'),
]
