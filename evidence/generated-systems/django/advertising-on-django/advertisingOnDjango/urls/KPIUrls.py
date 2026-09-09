from django.urls import path
from advertisingOnDjango.views import KPIView

urlpatterns = [
    path('', KPIView.index, name='index'),
	path('create', KPIView.get, name='create'),
	path('get/<int:kPIId>/', KPIView.get, name='get'),
	path('save', KPIView.save, name='save'),
	path('getAll', KPIView.getAll, name='getAll'),
	path('delete/<int:kPIId>/', KPIView.delete, name='delete'),
	path('assignCampaign/<int:kPIId>/<int:CampaignId>/', KPIView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:kPIId>/', KPIView.unassignCampaign, name='unassignCampaign'),
]
