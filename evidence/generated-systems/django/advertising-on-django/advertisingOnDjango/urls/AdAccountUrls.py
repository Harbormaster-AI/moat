from django.urls import path
from advertisingOnDjango.views import AdAccountView

urlpatterns = [
    path('', AdAccountView.index, name='index'),
	path('create', AdAccountView.get, name='create'),
	path('get/<int:adAccountId>/', AdAccountView.get, name='get'),
	path('save', AdAccountView.save, name='save'),
	path('getAll', AdAccountView.getAll, name='getAll'),
	path('delete/<int:adAccountId>/', AdAccountView.delete, name='delete'),
	path('assignAdvertiser/<int:adAccountId>/<int:AdvertiserId>/', AdAccountView.assignAdvertiser, name='assignAdvertiser'),
	path('unassignAdvertiser/<int:adAccountId>/', AdAccountView.unassignAdvertiser, name='unassignAdvertiser'),
	path('assignBillingProfile/<int:adAccountId>/<int:BillingProfileId>/', AdAccountView.assignBillingProfile, name='assignBillingProfile'),
	path('unassignBillingProfile/<int:adAccountId>/', AdAccountView.unassignBillingProfile, name='unassignBillingProfile'),
	path('assignDsp/<int:adAccountId>/<int:DspId>/', AdAccountView.assignDsp, name='assignDsp'),
	path('unassignDsp/<int:adAccountId>/', AdAccountView.unassignDsp, name='unassignDsp'),
	path('addUsers/<int:adAccountId>/<UsersIds>/', AdAccountView.addUsers, name='addUsers'),
	path('removeUsers/<int:adAccountId>/<UsersIds>/', AdAccountView.removeUsers, name='removeUsers'),
	path('addCampaigns/<int:adAccountId>/<CampaignsIds>/', AdAccountView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:adAccountId>/<CampaignsIds>/', AdAccountView.removeCampaigns, name='removeCampaigns'),
	path('addPerformanceMetrics/<int:adAccountId>/<PerformanceMetricsIds>/', AdAccountView.addPerformanceMetrics, name='addPerformanceMetrics'),
	path('removePerformanceMetrics/<int:adAccountId>/<PerformanceMetricsIds>/', AdAccountView.removePerformanceMetrics, name='removePerformanceMetrics'),
]
