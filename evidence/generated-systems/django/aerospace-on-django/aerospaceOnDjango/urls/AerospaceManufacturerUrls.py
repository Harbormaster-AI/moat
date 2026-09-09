from django.urls import path
from aerospaceOnDjango.views import AerospaceManufacturerView

urlpatterns = [
    path('', AerospaceManufacturerView.index, name='index'),
	path('create', AerospaceManufacturerView.get, name='create'),
	path('get/<int:aerospaceManufacturerId>/', AerospaceManufacturerView.get, name='get'),
	path('save', AerospaceManufacturerView.save, name='save'),
	path('getAll', AerospaceManufacturerView.getAll, name='getAll'),
	path('delete/<int:aerospaceManufacturerId>/', AerospaceManufacturerView.delete, name='delete'),
	path('addPrograms/<int:aerospaceManufacturerId>/<ProgramsIds>/', AerospaceManufacturerView.addPrograms, name='addPrograms'),
	path('removePrograms/<int:aerospaceManufacturerId>/<ProgramsIds>/', AerospaceManufacturerView.removePrograms, name='removePrograms'),
	path('addPlants/<int:aerospaceManufacturerId>/<PlantsIds>/', AerospaceManufacturerView.addPlants, name='addPlants'),
	path('removePlants/<int:aerospaceManufacturerId>/<PlantsIds>/', AerospaceManufacturerView.removePlants, name='removePlants'),
	path('addSuppliers/<int:aerospaceManufacturerId>/<SuppliersIds>/', AerospaceManufacturerView.addSuppliers, name='addSuppliers'),
	path('removeSuppliers/<int:aerospaceManufacturerId>/<SuppliersIds>/', AerospaceManufacturerView.removeSuppliers, name='removeSuppliers'),
	path('addProductionCertificates/<int:aerospaceManufacturerId>/<ProductionCertificatesIds>/', AerospaceManufacturerView.addProductionCertificates, name='addProductionCertificates'),
	path('removeProductionCertificates/<int:aerospaceManufacturerId>/<ProductionCertificatesIds>/', AerospaceManufacturerView.removeProductionCertificates, name='removeProductionCertificates'),
]
