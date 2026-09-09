from django.urls import path
from governanceOnDjango.views import RoleView

urlpatterns = [
    path('', RoleView.index, name='index'),
	path('create', RoleView.get, name='create'),
	path('get/<int:roleId>/', RoleView.get, name='get'),
	path('save', RoleView.save, name='save'),
	path('getAll', RoleView.getAll, name='getAll'),
	path('delete/<int:roleId>/', RoleView.delete, name='delete'),
	path('addAssignments/<int:roleId>/<AssignmentsIds>/', RoleView.addAssignments, name='addAssignments'),
	path('removeAssignments/<int:roleId>/<AssignmentsIds>/', RoleView.removeAssignments, name='removeAssignments'),
]
