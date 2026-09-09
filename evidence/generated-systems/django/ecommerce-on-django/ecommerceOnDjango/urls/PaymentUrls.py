from django.urls import path
from ecommerceOnDjango.views import PaymentView

urlpatterns = [
    path('', PaymentView.index, name='index'),
	path('create', PaymentView.get, name='create'),
	path('get/<int:paymentId>/', PaymentView.get, name='get'),
	path('save', PaymentView.save, name='save'),
	path('getAll', PaymentView.getAll, name='getAll'),
	path('delete/<int:paymentId>/', PaymentView.delete, name='delete'),
	path('assignOrder/<int:paymentId>/<int:OrderId>/', PaymentView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:paymentId>/', PaymentView.unassignOrder, name='unassignOrder'),
	path('assignCustomer/<int:paymentId>/<int:CustomerId>/', PaymentView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:paymentId>/', PaymentView.unassignCustomer, name='unassignCustomer'),
	path('assignPaymentProvider/<int:paymentId>/<int:PaymentProviderId>/', PaymentView.assignPaymentProvider, name='assignPaymentProvider'),
	path('unassignPaymentProvider/<int:paymentId>/', PaymentView.unassignPaymentProvider, name='unassignPaymentProvider'),
	path('addRefunds/<int:paymentId>/<RefundsIds>/', PaymentView.addRefunds, name='addRefunds'),
	path('removeRefunds/<int:paymentId>/<RefundsIds>/', PaymentView.removeRefunds, name='removeRefunds'),
]
