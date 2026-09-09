from django.urls import path
from fintechOnDjango.views import AccountView

urlpatterns = [
    path('', AccountView.index, name='index'),
	path('create', AccountView.get, name='create'),
	path('get/<int:accountId>/', AccountView.get, name='get'),
	path('save', AccountView.save, name='save'),
	path('getAll', AccountView.getAll, name='getAll'),
	path('delete/<int:accountId>/', AccountView.delete, name='delete'),
	path('assignCustomer/<int:accountId>/<int:CustomerId>/', AccountView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:accountId>/', AccountView.unassignCustomer, name='unassignCustomer'),
	path('assignInstitution/<int:accountId>/<int:InstitutionId>/', AccountView.assignInstitution, name='assignInstitution'),
	path('unassignInstitution/<int:accountId>/', AccountView.unassignInstitution, name='unassignInstitution'),
	path('addTransactions/<int:accountId>/<TransactionsIds>/', AccountView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:accountId>/<TransactionsIds>/', AccountView.removeTransactions, name='removeTransactions'),
	path('addCards/<int:accountId>/<CardsIds>/', AccountView.addCards, name='addCards'),
	path('removeCards/<int:accountId>/<CardsIds>/', AccountView.removeCards, name='removeCards'),
	path('addStatements/<int:accountId>/<StatementsIds>/', AccountView.addStatements, name='addStatements'),
	path('removeStatements/<int:accountId>/<StatementsIds>/', AccountView.removeStatements, name='removeStatements'),
	path('addMandates/<int:accountId>/<MandatesIds>/', AccountView.addMandates, name='addMandates'),
	path('removeMandates/<int:accountId>/<MandatesIds>/', AccountView.removeMandates, name='removeMandates'),
]
