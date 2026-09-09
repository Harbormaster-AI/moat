from django.urls import path
from advertisingOnDjango.views import PerformanceMetricView

urlpatterns = [
    path('', PerformanceMetricView.index, name='index'),
	path('create', PerformanceMetricView.get, name='create'),
	path('get/<int:performanceMetricId>/', PerformanceMetricView.get, name='get'),
	path('save', PerformanceMetricView.save, name='save'),
	path('getAll', PerformanceMetricView.getAll, name='getAll'),
	path('delete/<int:performanceMetricId>/', PerformanceMetricView.delete, name='delete'),
	path('assignAdAccount/<int:performanceMetricId>/<int:AdAccountId>/', PerformanceMetricView.assignAdAccount, name='assignAdAccount'),
	path('unassignAdAccount/<int:performanceMetricId>/', PerformanceMetricView.unassignAdAccount, name='unassignAdAccount'),
	path('assignCampaign/<int:performanceMetricId>/<int:CampaignId>/', PerformanceMetricView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:performanceMetricId>/', PerformanceMetricView.unassignCampaign, name='unassignCampaign'),
	path('assignLineItem/<int:performanceMetricId>/<int:LineItemId>/', PerformanceMetricView.assignLineItem, name='assignLineItem'),
	path('unassignLineItem/<int:performanceMetricId>/', PerformanceMetricView.unassignLineItem, name='unassignLineItem'),
	path('assignPlacement/<int:performanceMetricId>/<int:PlacementId>/', PerformanceMetricView.assignPlacement, name='assignPlacement'),
	path('unassignPlacement/<int:performanceMetricId>/', PerformanceMetricView.unassignPlacement, name='unassignPlacement'),
	path('assignCreativeAsset/<int:performanceMetricId>/<int:CreativeAssetId>/', PerformanceMetricView.assignCreativeAsset, name='assignCreativeAsset'),
	path('unassignCreativeAsset/<int:performanceMetricId>/', PerformanceMetricView.unassignCreativeAsset, name='unassignCreativeAsset'),
]
