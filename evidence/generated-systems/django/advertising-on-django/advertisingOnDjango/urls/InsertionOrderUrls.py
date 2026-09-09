from django.urls import path
from advertisingOnDjango.views import InsertionOrderView

urlpatterns = [
    path('', InsertionOrderView.index, name='index'),
	path('create', InsertionOrderView.get, name='create'),
	path('get/<int:insertionOrderId>/', InsertionOrderView.get, name='get'),
	path('save', InsertionOrderView.save, name='save'),
	path('getAll', InsertionOrderView.getAll, name='getAll'),
	path('delete/<int:insertionOrderId>/', InsertionOrderView.delete, name='delete'),
	path('assignAdvertiser/<int:insertionOrderId>/<int:AdvertiserId>/', InsertionOrderView.assignAdvertiser, name='assignAdvertiser'),
	path('unassignAdvertiser/<int:insertionOrderId>/', InsertionOrderView.unassignAdvertiser, name='unassignAdvertiser'),
	path('assignAgency/<int:insertionOrderId>/<int:AgencyId>/', InsertionOrderView.assignAgency, name='assignAgency'),
	path('unassignAgency/<int:insertionOrderId>/', InsertionOrderView.unassignAgency, name='unassignAgency'),
	path('assignPublisher/<int:insertionOrderId>/<int:PublisherId>/', InsertionOrderView.assignPublisher, name='assignPublisher'),
	path('unassignPublisher/<int:insertionOrderId>/', InsertionOrderView.unassignPublisher, name='unassignPublisher'),
	path('addCampaigns/<int:insertionOrderId>/<CampaignsIds>/', InsertionOrderView.addCampaigns, name='addCampaigns'),
	path('removeCampaigns/<int:insertionOrderId>/<CampaignsIds>/', InsertionOrderView.removeCampaigns, name='removeCampaigns'),
]
