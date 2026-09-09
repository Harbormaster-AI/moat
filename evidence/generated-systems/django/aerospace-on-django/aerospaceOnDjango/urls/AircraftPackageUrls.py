from django.urls import path
from aerospaceOnDjango.views import AircraftPackageView

urlpatterns = [
    path('', AircraftPackageView.index, name='index'),
	path('create', AircraftPackageView.get, name='create'),
	path('get/<int:aircraftPackageId>/', AircraftPackageView.get, name='get'),
	path('save', AircraftPackageView.save, name='save'),
	path('getAll', AircraftPackageView.getAll, name='getAll'),
	path('delete/<int:aircraftPackageId>/', AircraftPackageView.delete, name='delete'),
	path('addOptions/<int:aircraftPackageId>/<OptionsIds>/', AircraftPackageView.addOptions, name='addOptions'),
	path('removeOptions/<int:aircraftPackageId>/<OptionsIds>/', AircraftPackageView.removeOptions, name='removeOptions'),
	path('addVariants/<int:aircraftPackageId>/<VariantsIds>/', AircraftPackageView.addVariants, name='addVariants'),
	path('removeVariants/<int:aircraftPackageId>/<VariantsIds>/', AircraftPackageView.removeVariants, name='removeVariants'),
]
