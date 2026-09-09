from django.urls import path
from healthcareOnDjango.views import SoftwareUpdateView

urlpatterns = [
    path('', SoftwareUpdateView.index, name='index'),
	path('create', SoftwareUpdateView.get, name='create'),
	path('get/<int:softwareUpdateId>/', SoftwareUpdateView.get, name='get'),
	path('save', SoftwareUpdateView.save, name='save'),
	path('getAll', SoftwareUpdateView.getAll, name='getAll'),
	path('delete/<int:softwareUpdateId>/', SoftwareUpdateView.delete, name='delete'),
	path('assignDevice/<int:softwareUpdateId>/<int:DeviceId>/', SoftwareUpdateView.assignDevice, name='assignDevice'),
	path('unassignDevice/<int:softwareUpdateId>/', SoftwareUpdateView.unassignDevice, name='unassignDevice'),
]
