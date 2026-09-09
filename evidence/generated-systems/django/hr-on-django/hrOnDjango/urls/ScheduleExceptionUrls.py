from django.urls import path
from hrOnDjango.views import ScheduleExceptionView

urlpatterns = [
    path('', ScheduleExceptionView.index, name='index'),
	path('create', ScheduleExceptionView.get, name='create'),
	path('get/<int:scheduleExceptionId>/', ScheduleExceptionView.get, name='get'),
	path('save', ScheduleExceptionView.save, name='save'),
	path('getAll', ScheduleExceptionView.getAll, name='getAll'),
	path('delete/<int:scheduleExceptionId>/', ScheduleExceptionView.delete, name='delete'),
	path('assignWorkSchedule/<int:scheduleExceptionId>/<int:WorkScheduleId>/', ScheduleExceptionView.assignWorkSchedule, name='assignWorkSchedule'),
	path('unassignWorkSchedule/<int:scheduleExceptionId>/', ScheduleExceptionView.unassignWorkSchedule, name='unassignWorkSchedule'),
	path('assignEmployee/<int:scheduleExceptionId>/<int:EmployeeId>/', ScheduleExceptionView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:scheduleExceptionId>/', ScheduleExceptionView.unassignEmployee, name='unassignEmployee'),
]
