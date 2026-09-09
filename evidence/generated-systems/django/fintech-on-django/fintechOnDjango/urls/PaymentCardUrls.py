from django.urls import path
from fintechOnDjango.views import PaymentCardView

urlpatterns = [
    path('', PaymentCardView.index, name='index'),
	path('create', PaymentCardView.get, name='create'),
	path('get/<int:paymentCardId>/', PaymentCardView.get, name='get'),
	path('save', PaymentCardView.save, name='save'),
	path('getAll', PaymentCardView.getAll, name='getAll'),
	path('delete/<int:paymentCardId>/', PaymentCardView.delete, name='delete'),
	path('assignCustomer/<int:paymentCardId>/<int:CustomerId>/', PaymentCardView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:paymentCardId>/', PaymentCardView.unassignCustomer, name='unassignCustomer'),
	path('assignAccount/<int:paymentCardId>/<int:AccountId>/', PaymentCardView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:paymentCardId>/', PaymentCardView.unassignAccount, name='unassignAccount'),
	path('addTokenizations/<int:paymentCardId>/<TokenizationsIds>/', PaymentCardView.addTokenizations, name='addTokenizations'),
	path('removeTokenizations/<int:paymentCardId>/<TokenizationsIds>/', PaymentCardView.removeTokenizations, name='removeTokenizations'),
	path('addDisputes/<int:paymentCardId>/<DisputesIds>/', PaymentCardView.addDisputes, name='addDisputes'),
	path('removeDisputes/<int:paymentCardId>/<DisputesIds>/', PaymentCardView.removeDisputes, name='removeDisputes'),
]
