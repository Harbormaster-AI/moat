from django.urls import path
from advertisingOnDjango.views import CampaignView

urlpatterns = [
    path('', CampaignView.index, name='index'),
	path('create', CampaignView.get, name='create'),
	path('get/<int:campaignId>/', CampaignView.get, name='get'),
	path('save', CampaignView.save, name='save'),
	path('getAll', CampaignView.getAll, name='getAll'),
	path('delete/<int:campaignId>/', CampaignView.delete, name='delete'),
	path('assignAdAccount/<int:campaignId>/<int:AdAccountId>/', CampaignView.assignAdAccount, name='assignAdAccount'),
	path('unassignAdAccount/<int:campaignId>/', CampaignView.unassignAdAccount, name='unassignAdAccount'),
	path('assignInsertionOrder/<int:campaignId>/<int:InsertionOrderId>/', CampaignView.assignInsertionOrder, name='assignInsertionOrder'),
	path('unassignInsertionOrder/<int:campaignId>/', CampaignView.unassignInsertionOrder, name='unassignInsertionOrder'),
	path('addLineItems/<int:campaignId>/<LineItemsIds>/', CampaignView.addLineItems, name='addLineItems'),
	path('removeLineItems/<int:campaignId>/<LineItemsIds>/', CampaignView.removeLineItems, name='removeLineItems'),
	path('addKpis/<int:campaignId>/<KpisIds>/', CampaignView.addKpis, name='addKpis'),
	path('removeKpis/<int:campaignId>/<KpisIds>/', CampaignView.removeKpis, name='removeKpis'),
	path('addTrackingPixels/<int:campaignId>/<TrackingPixelsIds>/', CampaignView.addTrackingPixels, name='addTrackingPixels'),
	path('removeTrackingPixels/<int:campaignId>/<TrackingPixelsIds>/', CampaignView.removeTrackingPixels, name='removeTrackingPixels'),
	path('addAudiences/<int:campaignId>/<AudiencesIds>/', CampaignView.addAudiences, name='addAudiences'),
	path('removeAudiences/<int:campaignId>/<AudiencesIds>/', CampaignView.removeAudiences, name='removeAudiences'),
	path('addReports/<int:campaignId>/<ReportsIds>/', CampaignView.addReports, name='addReports'),
	path('removeReports/<int:campaignId>/<ReportsIds>/', CampaignView.removeReports, name='removeReports'),
]
