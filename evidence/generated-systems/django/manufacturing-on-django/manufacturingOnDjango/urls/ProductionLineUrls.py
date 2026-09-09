from django.urls import path
from manufacturingOnDjango.views import ProductionLineView

urlpatterns = [
    path('', ProductionLineView.index, name='index'),
	path('create', ProductionLineView.get, name='create'),
	path('get/<int:productionLineId>/', ProductionLineView.get, name='get'),
	path('save', ProductionLineView.save, name='save'),
	path('getAll', ProductionLineView.getAll, name='getAll'),
	path('delete/<int:productionLineId>/', ProductionLineView.delete, name='delete'),
	path('assignPlant/<int:productionLineId>/<int:PlantId>/', ProductionLineView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:productionLineId>/', ProductionLineView.unassignPlant, name='unassignPlant'),
	path('addWorkCenters/<int:productionLineId>/<WorkCentersIds>/', ProductionLineView.addWorkCenters, name='addWorkCenters'),
	path('removeWorkCenters/<int:productionLineId>/<WorkCentersIds>/', ProductionLineView.removeWorkCenters, name='removeWorkCenters'),
]
