from django.urls import path
from advertisingOnDjango.views import ReportView

urlpatterns = [
    path('', ReportView.index, name='index'),
	path('create', ReportView.get, name='create'),
	path('get/<int:reportId>/', ReportView.get, name='get'),
	path('save', ReportView.save, name='save'),
	path('getAll', ReportView.getAll, name='getAll'),
	path('delete/<int:reportId>/', ReportView.delete, name='delete'),
	path('assignAdAccount/<int:reportId>/<int:AdAccountId>/', ReportView.assignAdAccount, name='assignAdAccount'),
	path('unassignAdAccount/<int:reportId>/', ReportView.unassignAdAccount, name='unassignAdAccount'),
	path('assignCampaign/<int:reportId>/<int:CampaignId>/', ReportView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:reportId>/', ReportView.unassignCampaign, name='unassignCampaign'),
	path('assignLineItem/<int:reportId>/<int:LineItemId>/', ReportView.assignLineItem, name='assignLineItem'),
	path('unassignLineItem/<int:reportId>/', ReportView.unassignLineItem, name='unassignLineItem'),
]
