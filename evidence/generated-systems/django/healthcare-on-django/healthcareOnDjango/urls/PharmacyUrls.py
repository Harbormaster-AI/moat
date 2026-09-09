from django.urls import path
from healthcareOnDjango.views import PharmacyView

urlpatterns = [
    path('', PharmacyView.index, name='index'),
	path('create', PharmacyView.get, name='create'),
	path('get/<int:pharmacyId>/', PharmacyView.get, name='get'),
	path('save', PharmacyView.save, name='save'),
	path('getAll', PharmacyView.getAll, name='getAll'),
	path('delete/<int:pharmacyId>/', PharmacyView.delete, name='delete'),
	path('assignFacility/<int:pharmacyId>/<int:FacilityId>/', PharmacyView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:pharmacyId>/', PharmacyView.unassignFacility, name='unassignFacility'),
	path('addMedicationDispenses/<int:pharmacyId>/<MedicationDispensesIds>/', PharmacyView.addMedicationDispenses, name='addMedicationDispenses'),
	path('removeMedicationDispenses/<int:pharmacyId>/<MedicationDispensesIds>/', PharmacyView.removeMedicationDispenses, name='removeMedicationDispenses'),
	path('addMedicationOrders/<int:pharmacyId>/<MedicationOrdersIds>/', PharmacyView.addMedicationOrders, name='addMedicationOrders'),
	path('removeMedicationOrders/<int:pharmacyId>/<MedicationOrdersIds>/', PharmacyView.removeMedicationOrders, name='removeMedicationOrders'),
]
