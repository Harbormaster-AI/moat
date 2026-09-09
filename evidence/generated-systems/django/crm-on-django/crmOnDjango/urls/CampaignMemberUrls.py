from django.urls import path
from crmOnDjango.views import CampaignMemberView

urlpatterns = [
    path('', CampaignMemberView.index, name='index'),
	path('create', CampaignMemberView.get, name='create'),
	path('get/<int:campaignMemberId>/', CampaignMemberView.get, name='get'),
	path('save', CampaignMemberView.save, name='save'),
	path('getAll', CampaignMemberView.getAll, name='getAll'),
	path('delete/<int:campaignMemberId>/', CampaignMemberView.delete, name='delete'),
	path('assignCampaign/<int:campaignMemberId>/<int:CampaignId>/', CampaignMemberView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:campaignMemberId>/', CampaignMemberView.unassignCampaign, name='unassignCampaign'),
	path('assignLead/<int:campaignMemberId>/<int:LeadId>/', CampaignMemberView.assignLead, name='assignLead'),
	path('unassignLead/<int:campaignMemberId>/', CampaignMemberView.unassignLead, name='unassignLead'),
	path('assignContact/<int:campaignMemberId>/<int:ContactId>/', CampaignMemberView.assignContact, name='assignContact'),
	path('unassignContact/<int:campaignMemberId>/', CampaignMemberView.unassignContact, name='unassignContact'),
]
