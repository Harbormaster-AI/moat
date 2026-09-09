from django.urls import path
from fintechOnDjango.views import InvoiceView

urlpatterns = [
    path('', InvoiceView.index, name='index'),
	path('create', InvoiceView.get, name='create'),
	path('get/<int:invoiceId>/', InvoiceView.get, name='get'),
	path('save', InvoiceView.save, name='save'),
	path('getAll', InvoiceView.getAll, name='getAll'),
	path('delete/<int:invoiceId>/', InvoiceView.delete, name='delete'),
	path('assignMerchant/<int:invoiceId>/<int:MerchantId>/', InvoiceView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:invoiceId>/', InvoiceView.unassignMerchant, name='unassignMerchant'),
	path('addPayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.addPayments, name='addPayments'),
	path('removePayments/<int:invoiceId>/<PaymentsIds>/', InvoiceView.removePayments, name='removePayments'),
]
