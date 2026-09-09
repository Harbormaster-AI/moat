from django.urls import path
from hrOnDjango.views import BankAccountView

urlpatterns = [
    path('', BankAccountView.index, name='index'),
	path('create', BankAccountView.get, name='create'),
	path('get/<int:bankAccountId>/', BankAccountView.get, name='get'),
	path('save', BankAccountView.save, name='save'),
	path('getAll', BankAccountView.getAll, name='getAll'),
	path('delete/<int:bankAccountId>/', BankAccountView.delete, name='delete'),
]
