from django.urls import path
from healthcareOnDjango.views import DepartmentView

urlpatterns = [
    path('', DepartmentView.index, name='index'),
	path('create', DepartmentView.get, name='create'),
	path('get/<int:departmentId>/', DepartmentView.get, name='get'),
	path('save', DepartmentView.save, name='save'),
	path('getAll', DepartmentView.getAll, name='getAll'),
	path('delete/<int:departmentId>/', DepartmentView.delete, name='delete'),
	path('assignFacility/<int:departmentId>/<int:FacilityId>/', DepartmentView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:departmentId>/', DepartmentView.unassignFacility, name='unassignFacility'),
	path('addCareTeams/<int:departmentId>/<CareTeamsIds>/', DepartmentView.addCareTeams, name='addCareTeams'),
	path('removeCareTeams/<int:departmentId>/<CareTeamsIds>/', DepartmentView.removeCareTeams, name='removeCareTeams'),
]
