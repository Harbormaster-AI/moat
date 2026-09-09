from django.urls import path
from aerospaceOnDjango.views import AircraftFamilyView

urlpatterns = [
    path('', AircraftFamilyView.index, name='index'),
	path('create', AircraftFamilyView.get, name='create'),
	path('get/<int:aircraftFamilyId>/', AircraftFamilyView.get, name='get'),
	path('save', AircraftFamilyView.save, name='save'),
	path('getAll', AircraftFamilyView.getAll, name='getAll'),
	path('delete/<int:aircraftFamilyId>/', AircraftFamilyView.delete, name='delete'),
	path('assignProgram/<int:aircraftFamilyId>/<int:ProgramId>/', AircraftFamilyView.assignProgram, name='assignProgram'),
	path('unassignProgram/<int:aircraftFamilyId>/', AircraftFamilyView.unassignProgram, name='unassignProgram'),
	path('addAircraftModels/<int:aircraftFamilyId>/<AircraftModelsIds>/', AircraftFamilyView.addAircraftModels, name='addAircraftModels'),
	path('removeAircraftModels/<int:aircraftFamilyId>/<AircraftModelsIds>/', AircraftFamilyView.removeAircraftModels, name='removeAircraftModels'),
]
