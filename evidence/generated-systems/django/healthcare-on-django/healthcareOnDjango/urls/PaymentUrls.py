from django.urls import path
from healthcareOnDjango.views import PaymentView

urlpatterns = [
    path('', PaymentView.index, name='index'),
	path('create', PaymentView.get, name='create'),
	path('get/<int:paymentId>/', PaymentView.get, name='get'),
	path('save', PaymentView.save, name='save'),
	path('getAll', PaymentView.getAll, name='getAll'),
	path('delete/<int:paymentId>/', PaymentView.delete, name='delete'),
	path('assignInvoice/<int:paymentId>/<int:InvoiceId>/', PaymentView.assignInvoice, name='assignInvoice'),
	path('unassignInvoice/<int:paymentId>/', PaymentView.unassignInvoice, name='unassignInvoice'),
	path('assignPayer/<int:paymentId>/<int:PayerId>/', PaymentView.assignPayer, name='assignPayer'),
	path('unassignPayer/<int:paymentId>/', PaymentView.unassignPayer, name='unassignPayer'),
]
