from django.urls import path
from hrOnDjango.views import CandidateView

urlpatterns = [
    path('', CandidateView.index, name='index'),
	path('create', CandidateView.get, name='create'),
	path('get/<int:candidateId>/', CandidateView.get, name='get'),
	path('save', CandidateView.save, name='save'),
	path('getAll', CandidateView.getAll, name='getAll'),
	path('delete/<int:candidateId>/', CandidateView.delete, name='delete'),
	path('addApplications/<int:candidateId>/<ApplicationsIds>/', CandidateView.addApplications, name='addApplications'),
	path('removeApplications/<int:candidateId>/<ApplicationsIds>/', CandidateView.removeApplications, name='removeApplications'),
	path('addInterviews/<int:candidateId>/<InterviewsIds>/', CandidateView.addInterviews, name='addInterviews'),
	path('removeInterviews/<int:candidateId>/<InterviewsIds>/', CandidateView.removeInterviews, name='removeInterviews'),
	path('addOffers/<int:candidateId>/<OffersIds>/', CandidateView.addOffers, name='addOffers'),
	path('removeOffers/<int:candidateId>/<OffersIds>/', CandidateView.removeOffers, name='removeOffers'),
	path('addDocuments/<int:candidateId>/<DocumentsIds>/', CandidateView.addDocuments, name='addDocuments'),
	path('removeDocuments/<int:candidateId>/<DocumentsIds>/', CandidateView.removeDocuments, name='removeDocuments'),
]
