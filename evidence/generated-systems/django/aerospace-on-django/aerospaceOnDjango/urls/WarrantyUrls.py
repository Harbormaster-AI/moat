from django.urls import path
from aerospaceOnDjango.views import WarrantyView

urlpatterns = [
    path('', WarrantyView.index, name='index'),
	path('create', WarrantyView.get, name='create'),
	path('get/<int:warrantyId>/', WarrantyView.get, name='get'),
	path('save', WarrantyView.save, name='save'),
	path('getAll', WarrantyView.getAll, name='getAll'),
	path('delete/<int:warrantyId>/', WarrantyView.delete, name='delete'),
	path('assignAircraft/<int:warrantyId>/<int:AircraftId>/', WarrantyView.assignAircraft, name='assignAircraft'),
	path('unassignAircraft/<int:warrantyId>/', WarrantyView.unassignAircraft, name='unassignAircraft'),
]
