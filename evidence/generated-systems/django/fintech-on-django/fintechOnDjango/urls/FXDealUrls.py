from django.urls import path
from fintechOnDjango.views import FXDealView

urlpatterns = [
    path('', FXDealView.index, name='index'),
	path('create', FXDealView.get, name='create'),
	path('get/<int:fXDealId>/', FXDealView.get, name='get'),
	path('save', FXDealView.save, name='save'),
	path('getAll', FXDealView.getAll, name='getAll'),
	path('delete/<int:fXDealId>/', FXDealView.delete, name='delete'),
	path('assignQuote/<int:fXDealId>/<int:QuoteId>/', FXDealView.assignQuote, name='assignQuote'),
	path('unassignQuote/<int:fXDealId>/', FXDealView.unassignQuote, name='unassignQuote'),
	path('addPaymentOrders/<int:fXDealId>/<PaymentOrdersIds>/', FXDealView.addPaymentOrders, name='addPaymentOrders'),
	path('removePaymentOrders/<int:fXDealId>/<PaymentOrdersIds>/', FXDealView.removePaymentOrders, name='removePaymentOrders'),
]
