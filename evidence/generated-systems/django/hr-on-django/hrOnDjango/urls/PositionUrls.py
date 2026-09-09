from django.urls import path
from hrOnDjango.views import PositionView

urlpatterns = [
    path('', PositionView.index, name='index'),
	path('create', PositionView.get, name='create'),
	path('get/<int:positionId>/', PositionView.get, name='get'),
	path('save', PositionView.save, name='save'),
	path('getAll', PositionView.getAll, name='getAll'),
	path('delete/<int:positionId>/', PositionView.delete, name='delete'),
	path('assignDepartment/<int:positionId>/<int:DepartmentId>/', PositionView.assignDepartment, name='assignDepartment'),
	path('unassignDepartment/<int:positionId>/', PositionView.unassignDepartment, name='unassignDepartment'),
	path('assignJobProfile/<int:positionId>/<int:JobProfileId>/', PositionView.assignJobProfile, name='assignJobProfile'),
	path('unassignJobProfile/<int:positionId>/', PositionView.unassignJobProfile, name='unassignJobProfile'),
	path('assignCostCenter/<int:positionId>/<int:CostCenterId>/', PositionView.assignCostCenter, name='assignCostCenter'),
	path('unassignCostCenter/<int:positionId>/', PositionView.unassignCostCenter, name='unassignCostCenter'),
	path('assignLocation/<int:positionId>/<int:LocationId>/', PositionView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:positionId>/', PositionView.unassignLocation, name='unassignLocation'),
	path('assignManagerPosition/<int:positionId>/<int:ManagerPositionId>/', PositionView.assignManagerPosition, name='assignManagerPosition'),
	path('unassignManagerPosition/<int:positionId>/', PositionView.unassignManagerPosition, name='unassignManagerPosition'),
	path('addDirectReports/<int:positionId>/<DirectReportsIds>/', PositionView.addDirectReports, name='addDirectReports'),
	path('removeDirectReports/<int:positionId>/<DirectReportsIds>/', PositionView.removeDirectReports, name='removeDirectReports'),
	path('addAssignments/<int:positionId>/<AssignmentsIds>/', PositionView.addAssignments, name='addAssignments'),
	path('removeAssignments/<int:positionId>/<AssignmentsIds>/', PositionView.removeAssignments, name='removeAssignments'),
]
