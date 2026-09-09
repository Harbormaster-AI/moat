from django.urls import path
from healthcareOnDjango.views import MedicationOrderView

urlpatterns = [
    path('', MedicationOrderView.index, name='index'),
	path('create', MedicationOrderView.get, name='create'),
	path('get/<int:medicationOrderId>/', MedicationOrderView.get, name='get'),
	path('save', MedicationOrderView.save, name='save'),
	path('getAll', MedicationOrderView.getAll, name='getAll'),
	path('delete/<int:medicationOrderId>/', MedicationOrderView.delete, name='delete'),
	path('assignOrder/<int:medicationOrderId>/<int:OrderId>/', MedicationOrderView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:medicationOrderId>/', MedicationOrderView.unassignOrder, name='unassignOrder'),
	path('assignPharmacy/<int:medicationOrderId>/<int:PharmacyId>/', MedicationOrderView.assignPharmacy, name='assignPharmacy'),
	path('unassignPharmacy/<int:medicationOrderId>/', MedicationOrderView.unassignPharmacy, name='unassignPharmacy'),
	path('addDispenses/<int:medicationOrderId>/<DispensesIds>/', MedicationOrderView.addDispenses, name='addDispenses'),
	path('removeDispenses/<int:medicationOrderId>/<DispensesIds>/', MedicationOrderView.removeDispenses, name='removeDispenses'),
]
