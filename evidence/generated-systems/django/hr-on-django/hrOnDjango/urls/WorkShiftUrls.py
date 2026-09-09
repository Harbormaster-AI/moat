from django.urls import path
from hrOnDjango.views import WorkShiftView

urlpatterns = [
    path('', WorkShiftView.index, name='index'),
	path('create', WorkShiftView.get, name='create'),
	path('get/<int:workShiftId>/', WorkShiftView.get, name='get'),
	path('save', WorkShiftView.save, name='save'),
	path('getAll', WorkShiftView.getAll, name='getAll'),
	path('delete/<int:workShiftId>/', WorkShiftView.delete, name='delete'),
	path('assignWorkSchedule/<int:workShiftId>/<int:WorkScheduleId>/', WorkShiftView.assignWorkSchedule, name='assignWorkSchedule'),
	path('unassignWorkSchedule/<int:workShiftId>/', WorkShiftView.unassignWorkSchedule, name='unassignWorkSchedule'),
]
