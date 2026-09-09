from django.urls import path
from aerospaceOnDjango.views import AircraftOptionView

urlpatterns = [
    path('', AircraftOptionView.index, name='index'),
	path('create', AircraftOptionView.get, name='create'),
	path('get/<int:aircraftOptionId>/', AircraftOptionView.get, name='get'),
	path('save', AircraftOptionView.save, name='save'),
	path('getAll', AircraftOptionView.getAll, name='getAll'),
	path('delete/<int:aircraftOptionId>/', AircraftOptionView.delete, name='delete'),
	path('addVariants/<int:aircraftOptionId>/<VariantsIds>/', AircraftOptionView.addVariants, name='addVariants'),
	path('removeVariants/<int:aircraftOptionId>/<VariantsIds>/', AircraftOptionView.removeVariants, name='removeVariants'),
	path('addPackages/<int:aircraftOptionId>/<PackagesIds>/', AircraftOptionView.addPackages, name='addPackages'),
	path('removePackages/<int:aircraftOptionId>/<PackagesIds>/', AircraftOptionView.removePackages, name='removePackages'),
]
