from django.urls import path
from advertisingOnDjango.views import AudienceSegmentView

urlpatterns = [
    path('', AudienceSegmentView.index, name='index'),
	path('create', AudienceSegmentView.get, name='create'),
	path('get/<int:audienceSegmentId>/', AudienceSegmentView.get, name='get'),
	path('save', AudienceSegmentView.save, name='save'),
	path('getAll', AudienceSegmentView.getAll, name='getAll'),
	path('delete/<int:audienceSegmentId>/', AudienceSegmentView.delete, name='delete'),
	path('assignProvider/<int:audienceSegmentId>/<int:ProviderId>/', AudienceSegmentView.assignProvider, name='assignProvider'),
	path('unassignProvider/<int:audienceSegmentId>/', AudienceSegmentView.unassignProvider, name='unassignProvider'),
	path('addCampaigns/<int:audienceSegmentId>/<CampaignsIds>/', AudienceSegmentView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:audienceSegmentId>/<CampaignsIds>/', AudienceSegmentView.removeCampaigns, name='removeCampaigns'),
]
