from django.urls import path
from advertisingOnDjango.views import LineItemView

urlpatterns = [
    path('', LineItemView.index, name='index'),
	path('create', LineItemView.get, name='create'),
	path('get/<int:lineItemId>/', LineItemView.get, name='get'),
	path('save', LineItemView.save, name='save'),
	path('getAll', LineItemView.getAll, name='getAll'),
	path('delete/<int:lineItemId>/', LineItemView.delete, name='delete'),
	path('assignCampaign/<int:lineItemId>/<int:CampaignId>/', LineItemView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:lineItemId>/', LineItemView.unassignCampaign, name='unassignCampaign'),
	path('assignTargetingProfile/<int:lineItemId>/<int:TargetingProfileId>/', LineItemView.assignTargetingProfile, name='assignTargetingProfile'),
	path('unassignTargetingProfile/<int:lineItemId>/', LineItemView.unassignTargetingProfile, name='unassignTargetingProfile'),
	path('assignDeal/<int:lineItemId>/<int:DealId>/', LineItemView.assignDeal, name='assignDeal'),
	path('unassignDeal/<int:lineItemId>/', LineItemView.unassignDeal, name='unassignDeal'),
	path('addPlacements/<int:lineItemId>/<PlacementsIds>/', LineItemView.addPlacements, name='addPlacements'),
	path('removePlacements/<int:lineItemId>/<PlacementsIds>/', LineItemView.removePlacements, name='removePlacements'),
	path('addCreatives/<int:lineItemId>/<CreativesIds>/', LineItemView.addCreatives, name='addCreatives'),
	path('removeCreatives/<int:lineItemId>/<CreativesIds>/', LineItemView.removeCreatives, name='removeCreatives'),
	path('addPerformanceMetrics/<int:lineItemId>/<PerformanceMetricsIds>/', LineItemView.addPerformanceMetrics, name='addPerformanceMetrics'),
	path('removePerformanceMetrics/<int:lineItemId>/<PerformanceMetricsIds>/', LineItemView.removePerformanceMetrics, name='removePerformanceMetrics'),
]
