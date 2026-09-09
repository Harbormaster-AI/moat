from django.urls import path
from healthcareOnDjango.views import InvoiceView

urlpatterns = [
    path('', InvoiceView.index, name='index'),
	path('create', InvoiceView.get, name='create'),
	path('get/<int:invoiceId>/', InvoiceView.get, name='get'),
	path('save', InvoiceView.save, name='save'),
	path('getAll', InvoiceView.getAll, name='getAll'),
	path('delete/<int:invoiceId>/', InvoiceView.delete, name='delete'),
	path('assignPatient/<int:invoiceId>/<int:PatientId>/', InvoiceView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:invoiceId>/', InvoiceView.unassignPatient, name='unassignPatient'),
	path('assignClaim/<int:invoiceId>/<int:ClaimId>/', InvoiceView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:invoiceId>/', InvoiceView.unassignClaim, name='unassignClaim'),
	path('addPayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.addPayments, name='addPayments'),
	path('removePayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.removePayments, name='removePayments'),
]
