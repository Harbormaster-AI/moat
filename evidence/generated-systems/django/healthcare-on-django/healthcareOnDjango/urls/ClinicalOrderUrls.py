from django.urls import path
from healthcareOnDjango.views import ClinicalOrderView

urlpatterns = [
    path('', ClinicalOrderView.index, name='index'),
	path('create', ClinicalOrderView.get, name='create'),
	path('get/<int:clinicalOrderId>/', ClinicalOrderView.get, name='get'),
	path('save', ClinicalOrderView.save, name='save'),
	path('getAll', ClinicalOrderView.getAll, name='getAll'),
	path('delete/<int:clinicalOrderId>/', ClinicalOrderView.delete, name='delete'),
	path('assignPatient/<int:clinicalOrderId>/<int:PatientId>/', ClinicalOrderView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:clinicalOrderId>/', ClinicalOrderView.unassignPatient, name='unassignPatient'),
	path('assignEncounter/<int:clinicalOrderId>/<int:EncounterId>/', ClinicalOrderView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:clinicalOrderId>/', ClinicalOrderView.unassignEncounter, name='unassignEncounter'),
	path('assignOrderingClinician/<int:clinicalOrderId>/<int:OrderingClinicianId>/', ClinicalOrderView.assignOrderingClinician, name='assignOrderingClinician'),
	path('unassignOrderingClinician/<int:clinicalOrderId>/', ClinicalOrderView.unassignOrderingClinician, name='unassignOrderingClinician'),
	path('addMedicationOrders/<int:clinicalOrderId>/<MedicationOrdersIds>/', ClinicalOrderView.addMedicationOrders, name='addMedicationOrders'),
	path('removeMedicationOrders/<int:clinicalOrderId>/<MedicationOrdersIds>/', ClinicalOrderView.removeMedicationOrders, name='removeMedicationOrders'),
	path('addLaboratoryOrders/<int:clinicalOrderId>/<LaboratoryOrdersIds>/', ClinicalOrderView.addLaboratoryOrders, name='addLaboratoryOrders'),
	path('removeLaboratoryOrders/<int:clinicalOrderId>/<LaboratoryOrdersIds>/', ClinicalOrderView.removeLaboratoryOrders, name='removeLaboratoryOrders'),
	path('addImagingOrders/<int:clinicalOrderId>/<ImagingOrdersIds>/', ClinicalOrderView.addImagingOrders, name='addImagingOrders'),
	path('removeImagingOrders/<int:clinicalOrderId>/<ImagingOrdersIds>/', ClinicalOrderView.removeImagingOrders, name='removeImagingOrders'),
	path('addProcedureOrders/<int:clinicalOrderId>/<ProcedureOrdersIds>/', ClinicalOrderView.addProcedureOrders, name='addProcedureOrders'),
	path('removeProcedureOrders/<int:clinicalOrderId>/<ProcedureOrdersIds>/', ClinicalOrderView.removeProcedureOrders, name='removeProcedureOrders'),
	path('addAuthorizations/<int:clinicalOrderId>/<AuthorizationsIds>/', ClinicalOrderView.addAuthorizations, name='addAuthorizations'),
	path('removeAuthorizations/<int:clinicalOrderId>/<AuthorizationsIds>/', ClinicalOrderView.removeAuthorizations, name='removeAuthorizations'),
]
