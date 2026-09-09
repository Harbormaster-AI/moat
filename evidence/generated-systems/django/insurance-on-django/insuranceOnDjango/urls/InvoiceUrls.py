from django.urls import path
from insuranceOnDjango.views import InvoiceView

urlpatterns = [
    path('', InvoiceView.index, name='index'),
	path('create', InvoiceView.get, name='create'),
	path('get/<int:invoiceId>/', InvoiceView.get, name='get'),
	path('save', InvoiceView.save, name='save'),
	path('getAll', InvoiceView.getAll, name='getAll'),
	path('delete/<int:invoiceId>/', InvoiceView.delete, name='delete'),
	path('assignBillingAccount/<int:invoiceId>/<int:BillingAccountId>/', InvoiceView.assignBillingAccount, name='assignBillingAccount'),
	path('unassignBillingAccount/<int:invoiceId>/', InvoiceView.unassignBillingAccount, name='unassignBillingAccount'),
	path('assignPolicy/<int:invoiceId>/<int:PolicyId>/', InvoiceView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:invoiceId>/', InvoiceView.unassignPolicy, name='unassignPolicy'),
	path('addPayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.addPayments, name='addPayments'),
	path('removePayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.removePayments, name='removePayments'),
]
