from django.urls import path
from aerospaceOnDjango.views import WorkCenterView

urlpatterns = [
    path('', WorkCenterView.index, name='index'),
	path('create', WorkCenterView.get, name='create'),
	path('get/<int:workCenterId>/', WorkCenterView.get, name='get'),
	path('save', WorkCenterView.save, name='save'),
	path('getAll', WorkCenterView.getAll, name='getAll'),
	path('delete/<int:workCenterId>/', WorkCenterView.delete, name='delete'),
	path('assignProductionLine/<int:workCenterId>/<int:ProductionLineId>/', WorkCenterView.assignProductionLine, name='assignProductionLine'),
	path('unassignProductionLine/<int:workCenterId>/', WorkCenterView.unassignProductionLine, name='unassignProductionLine'),
]
