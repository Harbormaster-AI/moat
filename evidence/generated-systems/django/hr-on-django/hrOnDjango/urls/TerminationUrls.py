from django.urls import path
from hrOnDjango.views import TerminationView

urlpatterns = [
    path('', TerminationView.index, name='index'),
	path('create', TerminationView.get, name='create'),
	path('get/<int:terminationId>/', TerminationView.get, name='get'),
	path('save', TerminationView.save, name='save'),
	path('getAll', TerminationView.getAll, name='getAll'),
	path('delete/<int:terminationId>/', TerminationView.delete, name='delete'),
	path('assignEmployee/<int:terminationId>/<int:EmployeeId>/', TerminationView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:terminationId>/', TerminationView.unassignEmployee, name='unassignEmployee'),
	path('assignAssignment/<int:terminationId>/<int:AssignmentId>/', TerminationView.assignAssignment, name='assignAssignment'),
	path('unassignAssignment/<int:terminationId>/', TerminationView.unassignAssignment, name='unassignAssignment'),
]
