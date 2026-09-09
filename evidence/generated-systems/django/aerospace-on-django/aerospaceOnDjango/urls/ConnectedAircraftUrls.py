from django.urls import path
from aerospaceOnDjango.views import ConnectedAircraftView

urlpatterns = [
    path('', ConnectedAircraftView.index, name='index'),
	path('create', ConnectedAircraftView.get, name='create'),
	path('get/<int:connectedAircraftId>/', ConnectedAircraftView.get, name='get'),
	path('save', ConnectedAircraftView.save, name='save'),
	path('getAll', ConnectedAircraftView.getAll, name='getAll'),
	path('delete/<int:connectedAircraftId>/', ConnectedAircraftView.delete, name='delete'),
	path('assignAircraft/<int:connectedAircraftId>/<int:AircraftId>/', ConnectedAircraftView.assignAircraft, name='assignAircraft'),
	path('unassignAircraft/<int:connectedAircraftId>/', ConnectedAircraftView.unassignAircraft, name='unassignAircraft'),
	path('addFlightHealthEvents/<int:connectedAircraftId>/<FlightHealthEventsIds>/', ConnectedAircraftView.addFlightHealthEvents, name='addFlightHealthEvents'),
	path('removeFlightHealthEvents/<int:connectedAircraftId>/<FlightHealthEventsIds>/', ConnectedAircraftView.removeFlightHealthEvents, name='removeFlightHealthEvents'),
	path('addSoftwareLoads/<int:connectedAircraftId>/<SoftwareLoadsIds>/', ConnectedAircraftView.addSoftwareLoads, name='addSoftwareLoads'),
	path('removeSoftwareLoads/<int:connectedAircraftId>/<SoftwareLoadsIds>/', ConnectedAircraftView.removeSoftwareLoads, name='removeSoftwareLoads'),
]
