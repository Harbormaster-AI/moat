from django.urls import path
from advertisingOnDjango.views import ConversionEventView

urlpatterns = [
    path('', ConversionEventView.index, name='index'),
	path('create', ConversionEventView.get, name='create'),
	path('get/<int:conversionEventId>/', ConversionEventView.get, name='get'),
	path('save', ConversionEventView.save, name='save'),
	path('getAll', ConversionEventView.getAll, name='getAll'),
	path('delete/<int:conversionEventId>/', ConversionEventView.delete, name='delete'),
	path('assignCampaign/<int:conversionEventId>/<int:CampaignId>/', ConversionEventView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:conversionEventId>/', ConversionEventView.unassignCampaign, name='unassignCampaign'),
	path('assignLineItem/<int:conversionEventId>/<int:LineItemId>/', ConversionEventView.assignLineItem, name='assignLineItem'),
	path('unassignLineItem/<int:conversionEventId>/', ConversionEventView.unassignLineItem, name='unassignLineItem'),
	path('assignTrackingPixel/<int:conversionEventId>/<int:TrackingPixelId>/', ConversionEventView.assignTrackingPixel, name='assignTrackingPixel'),
	path('unassignTrackingPixel/<int:conversionEventId>/', ConversionEventView.unassignTrackingPixel, name='unassignTrackingPixel'),
]
