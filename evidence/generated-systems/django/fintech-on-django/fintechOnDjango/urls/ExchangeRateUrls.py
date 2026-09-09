from django.urls import path
from fintechOnDjango.views import ExchangeRateView

urlpatterns = [
    path('', ExchangeRateView.index, name='index'),
	path('create', ExchangeRateView.get, name='create'),
	path('get/<int:exchangeRateId>/', ExchangeRateView.get, name='get'),
	path('save', ExchangeRateView.save, name='save'),
	path('getAll', ExchangeRateView.getAll, name='getAll'),
	path('delete/<int:exchangeRateId>/', ExchangeRateView.delete, name='delete'),
	path('addUsedByQuotes/<int:exchangeRateId>/<UsedByQuotesIds>/', ExchangeRateView.addUsedByQuotes, name='addUsedByQuotes'),
	path('removeUsedByQuotes/<int:exchangeRateId>/<UsedByQuotesIds>/', ExchangeRateView.removeUsedByQuotes, name='removeUsedByQuotes'),
]
