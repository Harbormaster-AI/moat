from django.urls import path
from fintechOnDjango.views import PaymentOrderView

urlpatterns = [
    path('', PaymentOrderView.index, name='index'),
	path('create', PaymentOrderView.get, name='create'),
	path('get/<int:paymentOrderId>/', PaymentOrderView.get, name='get'),
	path('save', PaymentOrderView.save, name='save'),
	path('getAll', PaymentOrderView.getAll, name='getAll'),
	path('delete/<int:paymentOrderId>/', PaymentOrderView.delete, name='delete'),
	path('assignSourceAccount/<int:paymentOrderId>/<int:SourceAccountId>/', PaymentOrderView.assignSourceAccount, name='assignSourceAccount'),
	path('unassignSourceAccount/<int:paymentOrderId>/', PaymentOrderView.unassignSourceAccount, name='unassignSourceAccount'),
	path('assignDestinationAccount/<int:paymentOrderId>/<int:DestinationAccountId>/', PaymentOrderView.assignDestinationAccount, name='assignDestinationAccount'),
	path('unassignDestinationAccount/<int:paymentOrderId>/', PaymentOrderView.unassignDestinationAccount, name='unassignDestinationAccount'),
	path('assignBeneficiary/<int:paymentOrderId>/<int:BeneficiaryId>/', PaymentOrderView.assignBeneficiary, name='assignBeneficiary'),
	path('unassignBeneficiary/<int:paymentOrderId>/', PaymentOrderView.unassignBeneficiary, name='unassignBeneficiary'),
	path('assignFxDeal/<int:paymentOrderId>/<int:FxDealId>/', PaymentOrderView.assignFxDeal, name='assignFxDeal'),
	path('unassignFxDeal/<int:paymentOrderId>/', PaymentOrderView.unassignFxDeal, name='unassignFxDeal'),
	path('addTransactions/<int:paymentOrderId>/<TransactionsIds>/', PaymentOrderView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:paymentOrderId>/<TransactionsIds>/', PaymentOrderView.removeTransactions, name='removeTransactions'),
	path('addFees/<int:paymentOrderId>/<FeesIds>/', PaymentOrderView.addFees, name='addFees'),
	path('removeFees/<int:paymentOrderId>/<FeesIds>/', PaymentOrderView.removeFees, name='removeFees'),
]
