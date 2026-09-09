from django.urls import path
from advertisingOnDjango.views import RateCardView

urlpatterns = [
    path('', RateCardView.index, name='index'),
	path('create', RateCardView.get, name='create'),
	path('get/<int:rateCardId>/', RateCardView.get, name='get'),
	path('save', RateCardView.save, name='save'),
	path('getAll', RateCardView.getAll, name='getAll'),
	path('delete/<int:rateCardId>/', RateCardView.delete, name='delete'),
	path('assignPublisher/<int:rateCardId>/<int:PublisherId>/', RateCardView.assignPublisher, name='assignPublisher'),
	path('unassignPublisher/<int:rateCardId>/', RateCardView.unassignPublisher, name='unassignPublisher'),
	path('addRates/<int:rateCardId>/<RatesIds>/', RateCardView.addRates, name='addRates'),
	path('removeRates/<int:rateCardId>/<RatesIds>/', RateCardView.removeRates, name='removeRates'),
]
