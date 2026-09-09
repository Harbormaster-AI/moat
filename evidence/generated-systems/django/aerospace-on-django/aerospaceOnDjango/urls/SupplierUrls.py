from django.urls import path
from aerospaceOnDjango.views import SupplierView

urlpatterns = [
    path('', SupplierView.index, name='index'),
	path('create', SupplierView.get, name='create'),
	path('get/<int:supplierId>/', SupplierView.get, name='get'),
	path('save', SupplierView.save, name='save'),
	path('getAll', SupplierView.getAll, name='getAll'),
	path('delete/<int:supplierId>/', SupplierView.delete, name='delete'),
	path('addManufacturers/<int:supplierId>/<ManufacturersIds>/', SupplierView.addManufacturers, name='addManufacturers'),
	path('removeManufacturers/<int:supplierId>/<ManufacturersIds>/', SupplierView.removeManufacturers, name='removeManufacturers'),
	path('addComponents/<int:supplierId>/<ComponentsIds>/', SupplierView.addComponents, name='addComponents'),
	path('removeComponents/<int:supplierId>/<ComponentsIds>/', SupplierView.removeComponents, name='removeComponents'),
	path('addEngineTypes/<int:supplierId>/<EngineTypesIds>/', SupplierView.addEngineTypes, name='addEngineTypes'),
	path('removeEngineTypes/<int:supplierId>/<EngineTypesIds>/', SupplierView.removeEngineTypes, name='removeEngineTypes'),
	path('addAvionicsSuites/<int:supplierId>/<AvionicsSuitesIds>/', SupplierView.addAvionicsSuites, name='addAvionicsSuites'),
	path('removeAvionicsSuites/<int:supplierId>/<AvionicsSuitesIds>/', SupplierView.removeAvionicsSuites, name='removeAvionicsSuites'),
	path('addApus/<int:supplierId>/<ApusIds>/', SupplierView.addApus, name='addApus'),
	path('removeApus/<int:supplierId>/<ApusIds>/', SupplierView.removeApus, name='removeApus'),
	path('addLandingGears/<int:supplierId>/<LandingGearsIds>/', SupplierView.addLandingGears, name='addLandingGears'),
	path('removeLandingGears/<int:supplierId>/<LandingGearsIds>/', SupplierView.removeLandingGears, name='removeLandingGears'),
]
