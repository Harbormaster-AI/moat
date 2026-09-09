from django.urls import path
from hrOnDjango.views import PaymentMethodView

urlpatterns = [
    path('', PaymentMethodView.index, name='index'),
	path('create', PaymentMethodView.get, name='create'),
	path('get/<int:paymentMethodId>/', PaymentMethodView.get, name='get'),
	path('save', PaymentMethodView.save, name='save'),
	path('getAll', PaymentMethodView.getAll, name='getAll'),
	path('delete/<int:paymentMethodId>/', PaymentMethodView.delete, name='delete'),
	path('assignEmployee/<int:paymentMethodId>/<int:EmployeeId>/', PaymentMethodView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:paymentMethodId>/', PaymentMethodView.unassignEmployee, name='unassignEmployee'),
	path('assignBankAccount/<int:paymentMethodId>/<int:BankAccountId>/', PaymentMethodView.assignBankAccount, name='assignBankAccount'),
	path('unassignBankAccount/<int:paymentMethodId>/', PaymentMethodView.unassignBankAccount, name='unassignBankAccount'),
]
