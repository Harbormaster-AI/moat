from django.urls import path
from aerospaceOnDjango.views import AircraftModelView

urlpatterns = [
    path('', AircraftModelView.index, name='index'),
	path('create', AircraftModelView.get, name='create'),
	path('get/<int:aircraftModelId>/', AircraftModelView.get, name='get'),
	path('save', AircraftModelView.save, name='save'),
	path('getAll', AircraftModelView.getAll, name='getAll'),
	path('delete/<int:aircraftModelId>/', AircraftModelView.delete, name='delete'),
	path('assignFamily/<int:aircraftModelId>/<int:FamilyId>/', AircraftModelView.assignFamily, name='assignFamily'),
	path('unassignFamily/<int:aircraftModelId>/', AircraftModelView.unassignFamily, name='unassignFamily'),
	path('addVariants/<int:aircraftModelId>/<VariantsIds>/', AircraftModelView.addVariants, name='addVariants'),
	path('removeVariants/<int:aircraftModelId>/<VariantsIds>/', AircraftModelView.removeVariants, name='removeVariants'),
	path('addEngineTypes/<int:aircraftModelId>/<EngineTypesIds>/', AircraftModelView.addEngineTypes, name='addEngineTypes'),
	path('removeEngineTypes/<int:aircraftModelId>/<EngineTypesIds>/', AircraftModelView.removeEngineTypes, name='removeEngineTypes'),
]
