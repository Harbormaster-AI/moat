from django.urls import path
from aerospaceOnDjango.views import PlantView

urlpatterns = [
    path('', PlantView.index, name='index'),
	path('create', PlantView.get, name='create'),
	path('get/<int:plantId>/', PlantView.get, name='get'),
	path('save', PlantView.save, name='save'),
	path('getAll', PlantView.getAll, name='getAll'),
	path('delete/<int:plantId>/', PlantView.delete, name='delete'),
	path('assignManufacturer/<int:plantId>/<int:ManufacturerId>/', PlantView.assignManufacturer, name='assignManufacturer'),
	path('unassignManufacturer/<int:plantId>/', PlantView.unassignManufacturer, name='unassignManufacturer'),
	path('addProductionLines/<int:plantId>/<ProductionLinesIds>/', PlantView.addProductionLines, name='addProductionLines'),
	path('removeProductionLines/<int:plantId>/<ProductionLinesIds>/', PlantView.removeProductionLines, name='removeProductionLines'),
	path('addWarehouses/<int:plantId>/<WarehousesIds>/', PlantView.addWarehouses, name='addWarehouses'),
	path('removeWarehouses/<int:plantId>/<WarehousesIds>/', PlantView.removeWarehouses, name='removeWarehouses'),
]
