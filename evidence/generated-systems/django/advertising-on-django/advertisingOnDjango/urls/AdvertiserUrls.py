from django.urls import path
from advertisingOnDjango.views import AdvertiserView

urlpatterns = [
    path('', AdvertiserView.index, name='index'),
	path('create', AdvertiserView.get, name='create'),
	path('get/<int:advertiserId>/', AdvertiserView.get, name='get'),
	path('save', AdvertiserView.save, name='save'),
	path('getAll', AdvertiserView.getAll, name='getAll'),
	path('delete/<int:advertiserId>/', AdvertiserView.delete, name='delete'),
	path('assignAgency/<int:advertiserId>/<int:AgencyId>/', AdvertiserView.assignAgency, name='assignAgency'),
	path('unassignAgency/<int:advertiserId>/', AdvertiserView.unassignAgency, name='unassignAgency'),
	path('addAdAccounts/<int:advertiserId>/<AdAccountsIds>/', AdvertiserView.addAdAccounts, name='addAdAccounts'),
	path('removeAdAccounts/<int:advertiserId>/<AdAccountsIds>/', AdvertiserView.removeAdAccounts, name='removeAdAccounts'),
	path('addBillingProfiles/<int:advertiserId>/<BillingProfilesIds>/', AdvertiserView.addBillingProfiles, name='addBillingProfiles'),
	path('removeBillingProfiles/<int:advertiserId>/<BillingProfilesIds>/', AdvertiserView.removeBillingProfiles, name='removeBillingProfiles'),
	path('addCampaigns/<int:advertiserId>/<CampaignsIds>/', AdvertiserView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:advertiserId>/<CampaignsIds>/', AdvertiserView.removeCampaigns, name='removeCampaigns'),
	path('addTrackingPixels/<int:advertiserId>/<TrackingPixelsIds>/', AdvertiserView.addTrackingPixels, name='addTrackingPixels'),
	path('removeTrackingPixels/<int:advertiserId>/<TrackingPixelsIds>/', AdvertiserView.removeTrackingPixels, name='removeTrackingPixels'),
]
