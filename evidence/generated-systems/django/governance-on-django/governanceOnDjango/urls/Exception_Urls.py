from django.urls import path
from governanceOnDjango.views import Exception_View

urlpatterns = [
    path('', Exception_View.index, name='index'),
	path('create', Exception_View.get, name='create'),
	path('get/<int:exception_Id>/', Exception_View.get, name='get'),
	path('save', Exception_View.save, name='save'),
	path('getAll', Exception_View.getAll, name='getAll'),
	path('delete/<int:exception_Id>/', Exception_View.delete, name='delete'),
	path('assignRetentionSchedule/<int:exception_Id>/<int:RetentionScheduleId>/', Exception_View.assignRetentionSchedule, name='assignRetentionSchedule'),
	path('unassignRetentionSchedule/<int:exception_Id>/', Exception_View.unassignRetentionSchedule, name='unassignRetentionSchedule'),
	path('assignPolicy/<int:exception_Id>/<int:PolicyId>/', Exception_View.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:exception_Id>/', Exception_View.unassignPolicy, name='unassignPolicy'),
	path('assignControl/<int:exception_Id>/<int:ControlId>/', Exception_View.assignControl, name='assignControl'),
	path('unassignControl/<int:exception_Id>/', Exception_View.unassignControl, name='unassignControl'),
	path('assignRisk/<int:exception_Id>/<int:RiskId>/', Exception_View.assignRisk, name='assignRisk'),
	path('unassignRisk/<int:exception_Id>/', Exception_View.unassignRisk, name='unassignRisk'),
]
