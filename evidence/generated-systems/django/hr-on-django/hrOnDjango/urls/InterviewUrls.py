from django.urls import path
from hrOnDjango.views import InterviewView

urlpatterns = [
    path('', InterviewView.index, name='index'),
	path('create', InterviewView.get, name='create'),
	path('get/<int:interviewId>/', InterviewView.get, name='get'),
	path('save', InterviewView.save, name='save'),
	path('getAll', InterviewView.getAll, name='getAll'),
	path('delete/<int:interviewId>/', InterviewView.delete, name='delete'),
	path('assignRequisition/<int:interviewId>/<int:RequisitionId>/', InterviewView.assignRequisition, name='assignRequisition'),
	path('unassignRequisition/<int:interviewId>/', InterviewView.unassignRequisition, name='unassignRequisition'),
	path('assignCandidate/<int:interviewId>/<int:CandidateId>/', InterviewView.assignCandidate, name='assignCandidate'),
	path('unassignCandidate/<int:interviewId>/', InterviewView.unassignCandidate, name='unassignCandidate'),
	path('addInterviewers/<int:interviewId>/<InterviewersIds>/', InterviewView.addInterviewers, name='addInterviewers'),
	path('removeInterviewers/<int:interviewId>/<InterviewersIds>/', InterviewView.removeInterviewers, name='removeInterviewers'),
]
