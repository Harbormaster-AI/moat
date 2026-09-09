from django.urls import path
from aerospaceOnDjango.views import AircraftVariantView

urlpatterns = [
    path('', AircraftVariantView.index, name='index'),
	path('create', AircraftVariantView.get, name='create'),
	path('get/<int:aircraftVariantId>/', AircraftVariantView.get, name='get'),
	path('save', AircraftVariantView.save, name='save'),
	path('getAll', AircraftVariantView.getAll, name='getAll'),
	path('delete/<int:aircraftVariantId>/', AircraftVariantView.delete, name='delete'),
	path('assignModel/<int:aircraftVariantId>/<int:ModelId>/', AircraftVariantView.assignModel, name='assignModel'),
	path('unassignModel/<int:aircraftVariantId>/', AircraftVariantView.unassignModel, name='unassignModel'),
	path('assignEngineType/<int:aircraftVariantId>/<int:EngineTypeId>/', AircraftVariantView.assignEngineType, name='assignEngineType'),
	path('unassignEngineType/<int:aircraftVariantId>/', AircraftVariantView.unassignEngineType, name='unassignEngineType'),
	path('assignAvionicsSuite/<int:aircraftVariantId>/<int:AvionicsSuiteId>/', AircraftVariantView.assignAvionicsSuite, name='assignAvionicsSuite'),
	path('unassignAvionicsSuite/<int:aircraftVariantId>/', AircraftVariantView.unassignAvionicsSuite, name='unassignAvionicsSuite'),
	path('assignApu/<int:aircraftVariantId>/<int:ApuId>/', AircraftVariantView.assignApu, name='assignApu'),
	path('unassignApu/<int:aircraftVariantId>/', AircraftVariantView.unassignApu, name='unassignApu'),
	path('assignLandingGear/<int:aircraftVariantId>/<int:LandingGearId>/', AircraftVariantView.assignLandingGear, name='assignLandingGear'),
	path('unassignLandingGear/<int:aircraftVariantId>/', AircraftVariantView.unassignLandingGear, name='unassignLandingGear'),
	path('addCabinLayouts/<int:aircraftVariantId>/<CabinLayoutsIds>/', AircraftVariantView.addCabinLayouts, name='addCabinLayouts'),
	path('removeCabinLayouts/<int:aircraftVariantId>/<CabinLayoutsIds>/', AircraftVariantView.removeCabinLayouts, name='removeCabinLayouts'),
	path('addOptions/<int:aircraftVariantId>/<OptionsIds>/', AircraftVariantView.addOptions, name='addOptions'),
	path('removeOptions/<int:aircraftVariantId>/<OptionsIds>/', AircraftVariantView.removeOptions, name='removeOptions'),
	path('addPackages/<int:aircraftVariantId>/<PackagesIds>/', AircraftVariantView.addPackages, name='addPackages'),
	path('removePackages/<int:aircraftVariantId>/<PackagesIds>/', AircraftVariantView.removePackages, name='removePackages'),
]
