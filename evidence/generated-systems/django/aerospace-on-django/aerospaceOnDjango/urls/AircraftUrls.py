from django.urls import path
from aerospaceOnDjango.views import AircraftView

urlpatterns = [
    path('', AircraftView.index, name='index'),
	path('create', AircraftView.get, name='create'),
	path('get/<int:aircraftId>/', AircraftView.get, name='get'),
	path('save', AircraftView.save, name='save'),
	path('getAll', AircraftView.getAll, name='getAll'),
	path('delete/<int:aircraftId>/', AircraftView.delete, name='delete'),
	path('assignVariant/<int:aircraftId>/<int:VariantId>/', AircraftView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:aircraftId>/', AircraftView.unassignVariant, name='unassignVariant'),
	path('assignOperator/<int:aircraftId>/<int:OperatorId>/', AircraftView.assignOperator, name='assignOperator'),
	path('unassignOperator/<int:aircraftId>/', AircraftView.unassignOperator, name='unassignOperator'),
	path('assignRegistration/<int:aircraftId>/<int:RegistrationId>/', AircraftView.assignRegistration, name='assignRegistration'),
	path('unassignRegistration/<int:aircraftId>/', AircraftView.unassignRegistration, name='unassignRegistration'),
	path('assignWarranty/<int:aircraftId>/<int:WarrantyId>/', AircraftView.assignWarranty, name='assignWarranty'),
	path('unassignWarranty/<int:aircraftId>/', AircraftView.unassignWarranty, name='unassignWarranty'),
	path('assignConnectedAircraft/<int:aircraftId>/<int:ConnectedAircraftId>/', AircraftView.assignConnectedAircraft, name='assignConnectedAircraft'),
	path('unassignConnectedAircraft/<int:aircraftId>/', AircraftView.unassignConnectedAircraft, name='unassignConnectedAircraft'),
	path('assignCabinLayout/<int:aircraftId>/<int:CabinLayoutId>/', AircraftView.assignCabinLayout, name='assignCabinLayout'),
	path('unassignCabinLayout/<int:aircraftId>/', AircraftView.unassignCabinLayout, name='unassignCabinLayout'),
	path('addMaintenanceRecords/<int:aircraftId>/<MaintenanceRecordsIds>/', AircraftView.addMaintenanceRecords, name='addMaintenanceRecords'),
	path('removeMaintenanceRecords/<int:aircraftId>/<MaintenanceRecordsIds>/', AircraftView.removeMaintenanceRecords, name='removeMaintenanceRecords'),
]
