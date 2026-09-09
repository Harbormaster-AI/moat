from django.urls import path
from advertisingOnDjango.views import DataProviderView

urlpatterns = [
    path('', DataProviderView.index, name='index'),
	path('create', DataProviderView.get, name='create'),
	path('get/<int:dataProviderId>/', DataProviderView.get, name='get'),
	path('save', DataProviderView.save, name='save'),
	path('getAll', DataProviderView.getAll, name='getAll'),
	path('delete/<int:dataProviderId>/', DataProviderView.delete, name='delete'),
	path('addAudienceSegments/<int:dataProviderId>/<AudienceSegmentsIds>/', DataProviderView.addAudienceSegments, name='addAudienceSegments'),
	path('removeAudienceSegments/<int:dataProviderId>/<AudienceSegmentsIds>/', DataProviderView.removeAudienceSegments, name='removeAudienceSegments'),
]
