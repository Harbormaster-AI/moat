from django.urls import path
from aerospaceOnDjango.views import FlightHealthEventView

urlpatterns = [
    path('', FlightHealthEventView.index, name='index'),
	path('create', FlightHealthEventView.get, name='create'),
	path('get/<int:flightHealthEventId>/', FlightHealthEventView.get, name='get'),
	path('save', FlightHealthEventView.save, name='save'),
	path('getAll', FlightHealthEventView.getAll, name='getAll'),
	path('delete/<int:flightHealthEventId>/', FlightHealthEventView.delete, name='delete'),
	path('assignConnectedAircraft/<int:flightHealthEventId>/<int:ConnectedAircraftId>/', FlightHealthEventView.assignConnectedAircraft, name='assignConnectedAircraft'),
	path('unassignConnectedAircraft/<int:flightHealthEventId>/', FlightHealthEventView.unassignConnectedAircraft, name='unassignConnectedAircraft'),
]
