from django.urls import path
from hrOnDjango.views import CostCenterView

urlpatterns = [
    path('', CostCenterView.index, name='index'),
	path('create', CostCenterView.get, name='create'),
	path('get/<int:costCenterId>/', CostCenterView.get, name='get'),
	path('save', CostCenterView.save, name='save'),
	path('getAll', CostCenterView.getAll, name='getAll'),
	path('delete/<int:costCenterId>/', CostCenterView.delete, name='delete'),
	path('assignOrganization/<int:costCenterId>/<int:OrganizationId>/', CostCenterView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:costCenterId>/', CostCenterView.unassignOrganization, name='unassignOrganization'),
	path('addDepartments/<int:costCenterId>/<DepartmentsIds>/', CostCenterView.addDepartments, name='addDepartments'),
	path('removeDepartments/<int:costCenterId>/<DepartmentsIds>/', CostCenterView.removeDepartments, name='removeDepartments'),
	path('addPositions/<int:costCenterId>/<PositionsIds>/', CostCenterView.addPositions, name='addPositions'),
	path('removePositions/<int:costCenterId>/<PositionsIds>/', CostCenterView.removePositions, name='removePositions'),
	path('addEmployees/<int:costCenterId>/<EmployeesIds>/', CostCenterView.addEmployees, name='addEmployees'),
	path('removeEmployees/<int:costCenterId>/<EmployeesIds>/', CostCenterView.removeEmployees, name='removeEmployees'),
]
