from django.urls import path
from hrOnDjango.views import DepartmentView

urlpatterns = [
    path('', DepartmentView.index, name='index'),
	path('create', DepartmentView.get, name='create'),
	path('get/<int:departmentId>/', DepartmentView.get, name='get'),
	path('save', DepartmentView.save, name='save'),
	path('getAll', DepartmentView.getAll, name='getAll'),
	path('delete/<int:departmentId>/', DepartmentView.delete, name='delete'),
	path('assignOrganization/<int:departmentId>/<int:OrganizationId>/', DepartmentView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:departmentId>/', DepartmentView.unassignOrganization, name='unassignOrganization'),
	path('assignManager/<int:departmentId>/<int:ManagerId>/', DepartmentView.assignManager, name='assignManager'),
	path('unassignManager/<int:departmentId>/', DepartmentView.unassignManager, name='unassignManager'),
	path('assignCostCenter/<int:departmentId>/<int:CostCenterId>/', DepartmentView.assignCostCenter, name='assignCostCenter'),
	path('unassignCostCenter/<int:departmentId>/', DepartmentView.unassignCostCenter, name='unassignCostCenter'),
	path('addPositions/<int:departmentId>/<PositionsIds>/', DepartmentView.addPositions, name='addPositions'),
	path('removePositions/<int:departmentId>/<PositionsIds>/', DepartmentView.removePositions, name='removePositions'),
	path('addEmployees/<int:departmentId>/<EmployeesIds>/', DepartmentView.addEmployees, name='addEmployees'),
	path('removeEmployees/<int:departmentId>/<EmployeesIds>/', DepartmentView.removeEmployees, name='removeEmployees'),
]
