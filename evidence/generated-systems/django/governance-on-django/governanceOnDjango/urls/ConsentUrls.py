from django.urls import path
from governanceOnDjango.views import ConsentView

urlpatterns = [
    path('', ConsentView.index, name='index'),
	path('create', ConsentView.get, name='create'),
	path('get/<int:consentId>/', ConsentView.get, name='get'),
	path('save', ConsentView.save, name='save'),
	path('getAll', ConsentView.getAll, name='getAll'),
	path('delete/<int:consentId>/', ConsentView.delete, name='delete'),
	path('assignPrivacyNotice/<int:consentId>/<int:PrivacyNoticeId>/', ConsentView.assignPrivacyNotice, name='assignPrivacyNotice'),
	path('unassignPrivacyNotice/<int:consentId>/', ConsentView.unassignPrivacyNotice, name='unassignPrivacyNotice'),
	path('addProcessingActivities/<int:consentId>/<ProcessingActivitiesIds>/', ConsentView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:consentId>/<ProcessingActivitiesIds>/', ConsentView.removeProcessingActivities, name='removeProcessingActivities'),
]
