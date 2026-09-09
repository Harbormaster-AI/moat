from django.urls import path
from hrOnDjango.views import WorkScheduleView

urlpatterns = [
    path('', WorkScheduleView.index, name='index'),
	path('create', WorkScheduleView.get, name='create'),
	path('get/<int:workScheduleId>/', WorkScheduleView.get, name='get'),
	path('save', WorkScheduleView.save, name='save'),
	path('getAll', WorkScheduleView.getAll, name='getAll'),
	path('delete/<int:workScheduleId>/', WorkScheduleView.delete, name='delete'),
	path('addContracts/<int:workScheduleId>/<ContractsIds>/', WorkScheduleView.addContracts, name='addContracts'),
	path('removeContracts/<int:workScheduleId>/<ContractsIds>/', WorkScheduleView.removeContracts, name='removeContracts'),
	path('addShifts/<int:workScheduleId>/<ShiftsIds>/', WorkScheduleView.addShifts, name='addShifts'),
	path('removeShifts/<int:workScheduleId>/<ShiftsIds>/', WorkScheduleView.removeShifts, name='removeShifts'),
	path('addExceptions/<int:workScheduleId>/<ExceptionsIds>/', WorkScheduleView.addExceptions, name='addExceptions'),
	path('removeExceptions/<int:workScheduleId>/<ExceptionsIds>/', WorkScheduleView.removeExceptions, name='removeExceptions'),
]
