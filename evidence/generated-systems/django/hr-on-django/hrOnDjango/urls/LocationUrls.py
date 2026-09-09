from django.urls import path
from hrOnDjango.views import LocationView

urlpatterns = [
    path('', LocationView.index, name='index'),
	path('create', LocationView.get, name='create'),
	path('get/<int:locationId>/', LocationView.get, name='get'),
	path('save', LocationView.save, name='save'),
	path('getAll', LocationView.getAll, name='getAll'),
	path('delete/<int:locationId>/', LocationView.delete, name='delete'),
	path('assignOrganization/<int:locationId>/<int:OrganizationId>/', LocationView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:locationId>/', LocationView.unassignOrganization, name='unassignOrganization'),
	path('addDepartments/<int:locationId>/<DepartmentsIds>/', LocationView.addDepartments, name='addDepartments'),
	path('removeDepartments/<int:locationId>/<DepartmentsIds>/', LocationView.removeDepartments, name='removeDepartments'),
	path('addPositions/<int:locationId>/<PositionsIds>/', LocationView.addPositions, name='addPositions'),
	path('removePositions/<int:locationId>/<PositionsIds>/', LocationView.removePositions, name='removePositions'),
	path('addEmployees/<int:locationId>/<EmployeesIds>/', LocationView.addEmployees, name='addEmployees'),
	path('removeEmployees/<int:locationId>/<EmployeesIds>/', LocationView.removeEmployees, name='removeEmployees'),
]
