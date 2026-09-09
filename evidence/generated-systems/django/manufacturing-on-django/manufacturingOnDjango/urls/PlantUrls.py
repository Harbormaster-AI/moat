from django.urls import path
from manufacturingOnDjango.views import PlantView

urlpatterns = [
    path('', PlantView.index, name='index'),
	path('create', PlantView.get, name='create'),
	path('get/<int:plantId>/', PlantView.get, name='get'),
	path('save', PlantView.save, name='save'),
	path('getAll', PlantView.getAll, name='getAll'),
	path('delete/<int:plantId>/', PlantView.delete, name='delete'),
	path('assignEnterprise/<int:plantId>/<int:EnterpriseId>/', PlantView.assignEnterprise, name='assignEnterprise'),
	path('unassignEnterprise/<int:plantId>/', PlantView.unassignEnterprise, name='unassignEnterprise'),
	path('addProductionLines/<int:plantId>/<ProductionLinesIds>/', PlantView.addProductionLines, name='addProductionLines'),
	path('removeProductionLines/<int:plantId>/<ProductionLinesIds>/', PlantView.removeProductionLines, name='removeProductionLines'),
	path('addWorkCenters/<int:plantId>/<WorkCentersIds>/', PlantView.addWorkCenters, name='addWorkCenters'),
	path('removeWorkCenters/<int:plantId>/<WorkCentersIds>/', PlantView.removeWorkCenters, name='removeWorkCenters'),
	path('addWarehouses/<int:plantId>/<WarehousesIds>/', PlantView.addWarehouses, name='addWarehouses'),
	path('removeWarehouses/<int:plantId>/<WarehousesIds>/', PlantView.removeWarehouses, name='removeWarehouses'),
	path('addAssets/<int:plantId>/<AssetsIds>/', PlantView.addAssets, name='addAssets'),
	path('removeAssets/<int:plantId>/<AssetsIds>/', PlantView.removeAssets, name='removeAssets'),
	path('addProductionSchedules/<int:plantId>/<ProductionSchedulesIds>/', PlantView.addProductionSchedules, name='addProductionSchedules'),
	path('removeProductionSchedules/<int:plantId>/<ProductionSchedulesIds>/', PlantView.removeProductionSchedules, name='removeProductionSchedules'),
]
