from django.urls import path
from advertisingOnDjango.views import TrackingPixelView

urlpatterns = [
    path('', TrackingPixelView.index, name='index'),
	path('create', TrackingPixelView.get, name='create'),
	path('get/<int:trackingPixelId>/', TrackingPixelView.get, name='get'),
	path('save', TrackingPixelView.save, name='save'),
	path('getAll', TrackingPixelView.getAll, name='getAll'),
	path('delete/<int:trackingPixelId>/', TrackingPixelView.delete, name='delete'),
	path('assignCampaign/<int:trackingPixelId>/<int:CampaignId>/', TrackingPixelView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:trackingPixelId>/', TrackingPixelView.unassignCampaign, name='unassignCampaign'),
	path('assignAdvertiser/<int:trackingPixelId>/<int:AdvertiserId>/', TrackingPixelView.assignAdvertiser, name='assignAdvertiser'),
	path('unassignAdvertiser/<int:trackingPixelId>/', TrackingPixelView.unassignAdvertiser, name='unassignAdvertiser'),
	path('addConversionEvents/<int:trackingPixelId>/<ConversionEventsIds>/', TrackingPixelView.addConversionEvents, name='addConversionEvents'),
	path('removeConversionEvents/<int:trackingPixelId>/<ConversionEventsIds>/', TrackingPixelView.removeConversionEvents, name='removeConversionEvents'),
]
