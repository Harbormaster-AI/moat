from django.urls import path
from governanceOnDjango.views import PrivacyNoticeView

urlpatterns = [
    path('', PrivacyNoticeView.index, name='index'),
	path('create', PrivacyNoticeView.get, name='create'),
	path('get/<int:privacyNoticeId>/', PrivacyNoticeView.get, name='get'),
	path('save', PrivacyNoticeView.save, name='save'),
	path('getAll', PrivacyNoticeView.getAll, name='getAll'),
	path('delete/<int:privacyNoticeId>/', PrivacyNoticeView.delete, name='delete'),
	path('assignOrganization/<int:privacyNoticeId>/<int:OrganizationId>/', PrivacyNoticeView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:privacyNoticeId>/', PrivacyNoticeView.unassignOrganization, name='unassignOrganization'),
	path('addProcessingActivities/<int:privacyNoticeId>/<ProcessingActivitiesIds>/', PrivacyNoticeView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:privacyNoticeId>/<ProcessingActivitiesIds>/', PrivacyNoticeView.removeProcessingActivities, name='removeProcessingActivities'),
	path('addConsents/<int:privacyNoticeId>/<ConsentsIds>/', PrivacyNoticeView.addConsents, name='addConsents'),
	path('removeConsents/<int:privacyNoticeId>/<ConsentsIds>/', PrivacyNoticeView.removeConsents, name='removeConsents'),
]
