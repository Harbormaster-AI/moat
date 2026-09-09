from django.urls import path
from advertisingOnDjango.views import RateView

urlpatterns = [
    path('', RateView.index, name='index'),
	path('create', RateView.get, name='create'),
	path('get/<int:rateId>/', RateView.get, name='get'),
	path('save', RateView.save, name='save'),
	path('getAll', RateView.getAll, name='getAll'),
	path('delete/<int:rateId>/', RateView.delete, name='delete'),
	path('assignRateCard/<int:rateId>/<int:RateCardId>/', RateView.assignRateCard, name='assignRateCard'),
	path('unassignRateCard/<int:rateId>/', RateView.unassignRateCard, name='unassignRateCard'),
	path('assignAdSlot/<int:rateId>/<int:AdSlotId>/', RateView.assignAdSlot, name='assignAdSlot'),
	path('unassignAdSlot/<int:rateId>/', RateView.unassignAdSlot, name='unassignAdSlot'),
]
