from django.urls import path
from hrOnDjango.views import EmploymentAssignmentView

urlpatterns = [
    path('', EmploymentAssignmentView.index, name='index'),
	path('create', EmploymentAssignmentView.get, name='create'),
	path('get/<int:employmentAssignmentId>/', EmploymentAssignmentView.get, name='get'),
	path('save', EmploymentAssignmentView.save, name='save'),
	path('getAll', EmploymentAssignmentView.getAll, name='getAll'),
	path('delete/<int:employmentAssignmentId>/', EmploymentAssignmentView.delete, name='delete'),
	path('assignEmployee/<int:employmentAssignmentId>/<int:EmployeeId>/', EmploymentAssignmentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:employmentAssignmentId>/', EmploymentAssignmentView.unassignEmployee, name='unassignEmployee'),
	path('assignPosition/<int:employmentAssignmentId>/<int:PositionId>/', EmploymentAssignmentView.assignPosition, name='assignPosition'),
	path('unassignPosition/<int:employmentAssignmentId>/', EmploymentAssignmentView.unassignPosition, name='unassignPosition'),
	path('assignSupervisor/<int:employmentAssignmentId>/<int:SupervisorId>/', EmploymentAssignmentView.assignSupervisor, name='assignSupervisor'),
	path('unassignSupervisor/<int:employmentAssignmentId>/', EmploymentAssignmentView.unassignSupervisor, name='unassignSupervisor'),
]
