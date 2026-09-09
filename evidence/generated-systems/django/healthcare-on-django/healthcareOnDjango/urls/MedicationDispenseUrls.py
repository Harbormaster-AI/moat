from django.urls import path
from healthcareOnDjango.views import MedicationDispenseView

urlpatterns = [
    path('', MedicationDispenseView.index, name='index'),
	path('create', MedicationDispenseView.get, name='create'),
	path('get/<int:medicationDispenseId>/', MedicationDispenseView.get, name='get'),
	path('save', MedicationDispenseView.save, name='save'),
	path('getAll', MedicationDispenseView.getAll, name='getAll'),
	path('delete/<int:medicationDispenseId>/', MedicationDispenseView.delete, name='delete'),
	path('assignMedicationOrder/<int:medicationDispenseId>/<int:MedicationOrderId>/', MedicationDispenseView.assignMedicationOrder, name='assignMedicationOrder'),
	path('unassignMedicationOrder/<int:medicationDispenseId>/', MedicationDispenseView.unassignMedicationOrder, name='unassignMedicationOrder'),
	path('assignPharmacy/<int:medicationDispenseId>/<int:PharmacyId>/', MedicationDispenseView.assignPharmacy, name='assignPharmacy'),
	path('unassignPharmacy/<int:medicationDispenseId>/', MedicationDispenseView.unassignPharmacy, name='unassignPharmacy'),
	path('assignPatient/<int:medicationDispenseId>/<int:PatientId>/', MedicationDispenseView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:medicationDispenseId>/', MedicationDispenseView.unassignPatient, name='unassignPatient'),
]
