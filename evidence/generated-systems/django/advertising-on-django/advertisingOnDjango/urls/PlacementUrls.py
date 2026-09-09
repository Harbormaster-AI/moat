from django.urls import path
from advertisingOnDjango.views import PlacementView

urlpatterns = [
    path('', PlacementView.index, name='index'),
	path('create', PlacementView.get, name='create'),
	path('get/<int:placementId>/', PlacementView.get, name='get'),
	path('save', PlacementView.save, name='save'),
	path('getAll', PlacementView.getAll, name='getAll'),
	path('delete/<int:placementId>/', PlacementView.delete, name='delete'),
	path('assignLineItem/<int:placementId>/<int:LineItemId>/', PlacementView.assignLineItem, name='assignLineItem'),
	path('unassignLineItem/<int:placementId>/', PlacementView.unassignLineItem, name='unassignLineItem'),
	path('assignAdSlot/<int:placementId>/<int:AdSlotId>/', PlacementView.assignAdSlot, name='assignAdSlot'),
	path('unassignAdSlot/<int:placementId>/', PlacementView.unassignAdSlot, name='unassignAdSlot'),
	path('assignDeal/<int:placementId>/<int:DealId>/', PlacementView.assignDeal, name='assignDeal'),
	path('unassignDeal/<int:placementId>/', PlacementView.unassignDeal, name='unassignDeal'),
]
