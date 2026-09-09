from django.urls import path
from fintechOnDjango.views import DirectDebitMandateView

urlpatterns = [
    path('', DirectDebitMandateView.index, name='index'),
	path('create', DirectDebitMandateView.get, name='create'),
	path('get/<int:directDebitMandateId>/', DirectDebitMandateView.get, name='get'),
	path('save', DirectDebitMandateView.save, name='save'),
	path('getAll', DirectDebitMandateView.getAll, name='getAll'),
	path('delete/<int:directDebitMandateId>/', DirectDebitMandateView.delete, name='delete'),
	path('assignAccount/<int:directDebitMandateId>/<int:AccountId>/', DirectDebitMandateView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:directDebitMandateId>/', DirectDebitMandateView.unassignAccount, name='unassignAccount'),
	path('assignCreditor/<int:directDebitMandateId>/<int:CreditorId>/', DirectDebitMandateView.assignCreditor, name='assignCreditor'),
	path('unassignCreditor/<int:directDebitMandateId>/', DirectDebitMandateView.unassignCreditor, name='unassignCreditor'),
]
