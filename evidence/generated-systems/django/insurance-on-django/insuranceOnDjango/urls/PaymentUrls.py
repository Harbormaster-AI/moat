from django.urls import path
from insuranceOnDjango.views import PaymentView

urlpatterns = [
    path('', PaymentView.index, name='index'),
	path('create', PaymentView.get, name='create'),
	path('get/<int:paymentId>/', PaymentView.get, name='get'),
	path('save', PaymentView.save, name='save'),
	path('getAll', PaymentView.getAll, name='getAll'),
	path('delete/<int:paymentId>/', PaymentView.delete, name='delete'),
	path('assignInvoice/<int:paymentId>/<int:InvoiceId>/', PaymentView.assignInvoice, name='assignInvoice'),
	path('unassignInvoice/<int:paymentId>/', PaymentView.unassignInvoice, name='unassignInvoice'),
	path('assignBillingAccount/<int:paymentId>/<int:BillingAccountId>/', PaymentView.assignBillingAccount, name='assignBillingAccount'),
	path('unassignBillingAccount/<int:paymentId>/', PaymentView.unassignBillingAccount, name='unassignBillingAccount'),
	path('assignPolicy/<int:paymentId>/<int:PolicyId>/', PaymentView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:paymentId>/', PaymentView.unassignPolicy, name='unassignPolicy'),
]
