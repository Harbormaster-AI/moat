from django.urls import path
from aerospaceOnDjango.views import SoftwareLoadView

urlpatterns = [
    path('', SoftwareLoadView.index, name='index'),
	path('create', SoftwareLoadView.get, name='create'),
	path('get/<int:softwareLoadId>/', SoftwareLoadView.get, name='get'),
	path('save', SoftwareLoadView.save, name='save'),
	path('getAll', SoftwareLoadView.getAll, name='getAll'),
	path('delete/<int:softwareLoadId>/', SoftwareLoadView.delete, name='delete'),
	path('assignConnectedAircraft/<int:softwareLoadId>/<int:ConnectedAircraftId>/', SoftwareLoadView.assignConnectedAircraft, name='assignConnectedAircraft'),
	path('unassignConnectedAircraft/<int:softwareLoadId>/', SoftwareLoadView.unassignConnectedAircraft, name='unassignConnectedAircraft'),
	path('assignAvionicsSuite/<int:softwareLoadId>/<int:AvionicsSuiteId>/', SoftwareLoadView.assignAvionicsSuite, name='assignAvionicsSuite'),
	path('unassignAvionicsSuite/<int:softwareLoadId>/', SoftwareLoadView.unassignAvionicsSuite, name='unassignAvionicsSuite'),
]
