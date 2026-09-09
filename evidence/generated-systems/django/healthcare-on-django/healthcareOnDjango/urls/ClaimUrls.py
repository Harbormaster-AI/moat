from django.urls import path
from healthcareOnDjango.views import ClaimView

urlpatterns = [
    path('', ClaimView.index, name='index'),
	path('create', ClaimView.get, name='create'),
	path('get/<int:claimId>/', ClaimView.get, name='get'),
	path('save', ClaimView.save, name='save'),
	path('getAll', ClaimView.getAll, name='getAll'),
	path('delete/<int:claimId>/', ClaimView.delete, name='delete'),
	path('assignPatient/<int:claimId>/<int:PatientId>/', ClaimView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:claimId>/', ClaimView.unassignPatient, name='unassignPatient'),
	path('assignCoverage/<int:claimId>/<int:CoverageId>/', ClaimView.assignCoverage, name='assignCoverage'),
	path('unassignCoverage/<int:claimId>/', ClaimView.unassignCoverage, name='unassignCoverage'),
	path('assignEncounter/<int:claimId>/<int:EncounterId>/', ClaimView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:claimId>/', ClaimView.unassignEncounter, name='unassignEncounter'),
	path('assignPayer/<int:claimId>/<int:PayerId>/', ClaimView.assignPayer, name='assignPayer'),
	path('unassignPayer/<int:claimId>/', ClaimView.unassignPayer, name='unassignPayer'),
	path('addInvoices/<int:claimId>/<InvoicesIds>/', ClaimView.addInvoices, name='addInvoices'),
	path('removeInvoices/<int:claimId>/<InvoicesIds>/', ClaimView.removeInvoices, name='removeInvoices'),
]
