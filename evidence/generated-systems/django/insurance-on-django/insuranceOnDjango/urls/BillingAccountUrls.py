from django.urls import path
from insuranceOnDjango.views import BillingAccountView

urlpatterns = [
    path('', BillingAccountView.index, name='index'),
	path('create', BillingAccountView.get, name='create'),
	path('get/<int:billingAccountId>/', BillingAccountView.get, name='get'),
	path('save', BillingAccountView.save, name='save'),
	path('getAll', BillingAccountView.getAll, name='getAll'),
	path('delete/<int:billingAccountId>/', BillingAccountView.delete, name='delete'),
	path('assignCustomer/<int:billingAccountId>/<int:CustomerId>/', BillingAccountView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:billingAccountId>/', BillingAccountView.unassignCustomer, name='unassignCustomer'),
	path('addPolicies/<int:billingAccountId>/<PoliciesIds>/', BillingAccountView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:billingAccountId>/<PoliciesIds>/', BillingAccountView.removePolicies, name='removePolicies'),
	path('addInvoices/<int:billingAccountId>/<InvoicesIds>/', BillingAccountView.addInvoices, name='addInvoices'),
	path('removeInvoices/<int:billingAccountId>/<InvoicesIds>/', BillingAccountView.removeInvoices, name='removeInvoices'),
	path('addPayments/<int:billingAccountId>/<PaymentsIds>/', BillingAccountView.addPayments, name='addPayments'),
	path('removePayments/<int:billingAccountId>/<PaymentsIds>/', BillingAccountView.removePayments, name='removePayments'),
]
