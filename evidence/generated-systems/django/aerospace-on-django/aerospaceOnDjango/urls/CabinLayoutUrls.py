from django.urls import path
from aerospaceOnDjango.views import CabinLayoutView

urlpatterns = [
    path('', CabinLayoutView.index, name='index'),
	path('create', CabinLayoutView.get, name='create'),
	path('get/<int:cabinLayoutId>/', CabinLayoutView.get, name='get'),
	path('save', CabinLayoutView.save, name='save'),
	path('getAll', CabinLayoutView.getAll, name='getAll'),
	path('delete/<int:cabinLayoutId>/', CabinLayoutView.delete, name='delete'),
	path('assignVariant/<int:cabinLayoutId>/<int:VariantId>/', CabinLayoutView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:cabinLayoutId>/', CabinLayoutView.unassignVariant, name='unassignVariant'),
	path('addAircraft/<int:cabinLayoutId>/<AircraftIds>/', CabinLayoutView.addAircraft, name='addAircraft'),
	path('removeAircraft/<int:cabinLayoutId>/<AircraftIds>/', CabinLayoutView.removeAircraft, name='removeAircraft'),
	path('addOptions/<int:cabinLayoutId>/<OptionsIds>/', CabinLayoutView.addOptions, name='addOptions'),
	path('removeOptions/<int:cabinLayoutId>/<OptionsIds>/', CabinLayoutView.removeOptions, name='removeOptions'),
]
