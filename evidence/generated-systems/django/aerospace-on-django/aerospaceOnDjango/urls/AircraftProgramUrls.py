from django.urls import path
from aerospaceOnDjango.views import AircraftProgramView

urlpatterns = [
    path('', AircraftProgramView.index, name='index'),
	path('create', AircraftProgramView.get, name='create'),
	path('get/<int:aircraftProgramId>/', AircraftProgramView.get, name='get'),
	path('save', AircraftProgramView.save, name='save'),
	path('getAll', AircraftProgramView.getAll, name='getAll'),
	path('delete/<int:aircraftProgramId>/', AircraftProgramView.delete, name='delete'),
	path('assignManufacturer/<int:aircraftProgramId>/<int:ManufacturerId>/', AircraftProgramView.assignManufacturer, name='assignManufacturer'),
	path('unassignManufacturer/<int:aircraftProgramId>/', AircraftProgramView.unassignManufacturer, name='unassignManufacturer'),
	path('assignTypeCertificate/<int:aircraftProgramId>/<int:TypeCertificateId>/', AircraftProgramView.assignTypeCertificate, name='assignTypeCertificate'),
	path('unassignTypeCertificate/<int:aircraftProgramId>/', AircraftProgramView.unassignTypeCertificate, name='unassignTypeCertificate'),
	path('addAircraftFamilies/<int:aircraftProgramId>/<AircraftFamiliesIds>/', AircraftProgramView.addAircraftFamilies, name='addAircraftFamilies'),
	path('removeAircraftFamilies/<int:aircraftProgramId>/<AircraftFamiliesIds>/', AircraftProgramView.removeAircraftFamilies, name='removeAircraftFamilies'),
	path('addKeySuppliers/<int:aircraftProgramId>/<KeySuppliersIds>/', AircraftProgramView.addKeySuppliers, name='addKeySuppliers'),
	path('removeKeySuppliers/<int:aircraftProgramId>/<KeySuppliersIds>/', AircraftProgramView.removeKeySuppliers, name='removeKeySuppliers'),
]
