from django.urls import path
from aerospaceOnDjango.views import ProductionCertificateView

urlpatterns = [
    path('', ProductionCertificateView.index, name='index'),
	path('create', ProductionCertificateView.get, name='create'),
	path('get/<int:productionCertificateId>/', ProductionCertificateView.get, name='get'),
	path('save', ProductionCertificateView.save, name='save'),
	path('getAll', ProductionCertificateView.getAll, name='getAll'),
	path('delete/<int:productionCertificateId>/', ProductionCertificateView.delete, name='delete'),
	path('assignManufacturer/<int:productionCertificateId>/<int:ManufacturerId>/', ProductionCertificateView.assignManufacturer, name='assignManufacturer'),
	path('unassignManufacturer/<int:productionCertificateId>/', ProductionCertificateView.unassignManufacturer, name='unassignManufacturer'),
]
