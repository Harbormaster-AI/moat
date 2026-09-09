from django.urls import path
from healthcareOnDjango.views import HealthSystemView

urlpatterns = [
    path('', HealthSystemView.index, name='index'),
	path('create', HealthSystemView.get, name='create'),
	path('get/<int:healthSystemId>/', HealthSystemView.get, name='get'),
	path('save', HealthSystemView.save, name='save'),
	path('getAll', HealthSystemView.getAll, name='getAll'),
	path('delete/<int:healthSystemId>/', HealthSystemView.delete, name='delete'),
	path('addFacilities/<int:healthSystemId>/<FacilitiesIds>/', HealthSystemView.addFacilities, name='addFacilities'),
	path('removeFacilities/<int:healthSystemId>/<FacilitiesIds>/', HealthSystemView.removeFacilities, name='removeFacilities'),
	path('addSuppliers/<int:healthSystemId>/<SuppliersIds>/', HealthSystemView.addSuppliers, name='addSuppliers'),
	path('removeSuppliers/<int:healthSystemId>/<SuppliersIds>/', HealthSystemView.removeSuppliers, name='removeSuppliers'),
]
