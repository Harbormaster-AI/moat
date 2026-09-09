from django.urls import path
from aerospaceOnDjango.views import RegistrationView

urlpatterns = [
    path('', RegistrationView.index, name='index'),
	path('create', RegistrationView.get, name='create'),
	path('get/<int:registrationId>/', RegistrationView.get, name='get'),
	path('save', RegistrationView.save, name='save'),
	path('getAll', RegistrationView.getAll, name='getAll'),
	path('delete/<int:registrationId>/', RegistrationView.delete, name='delete'),
	path('assignAircraft/<int:registrationId>/<int:AircraftId>/', RegistrationView.assignAircraft, name='assignAircraft'),
	path('unassignAircraft/<int:registrationId>/', RegistrationView.unassignAircraft, name='unassignAircraft'),
]
