from django.urls import path
from hrOnDjango.views import PerformanceReviewView

urlpatterns = [
    path('', PerformanceReviewView.index, name='index'),
	path('create', PerformanceReviewView.get, name='create'),
	path('get/<int:performanceReviewId>/', PerformanceReviewView.get, name='get'),
	path('save', PerformanceReviewView.save, name='save'),
	path('getAll', PerformanceReviewView.getAll, name='getAll'),
	path('delete/<int:performanceReviewId>/', PerformanceReviewView.delete, name='delete'),
	path('assignEmployee/<int:performanceReviewId>/<int:EmployeeId>/', PerformanceReviewView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:performanceReviewId>/', PerformanceReviewView.unassignEmployee, name='unassignEmployee'),
	path('assignReviewer/<int:performanceReviewId>/<int:ReviewerId>/', PerformanceReviewView.assignReviewer, name='assignReviewer'),
	path('unassignReviewer/<int:performanceReviewId>/', PerformanceReviewView.unassignReviewer, name='unassignReviewer'),
	path('assignCycle/<int:performanceReviewId>/<int:CycleId>/', PerformanceReviewView.assignCycle, name='assignCycle'),
	path('unassignCycle/<int:performanceReviewId>/', PerformanceReviewView.unassignCycle, name='unassignCycle'),
	path('addCompetencyRatings/<int:performanceReviewId>/<CompetencyRatingsIds>/', PerformanceReviewView.addCompetencyRatings, name='addCompetencyRatings'),
	path('removeCompetencyRatings/<int:performanceReviewId>/<CompetencyRatingsIds>/', PerformanceReviewView.removeCompetencyRatings, name='removeCompetencyRatings'),
	path('addGoals/<int:performanceReviewId>/<GoalsIds>/', PerformanceReviewView.addGoals, name='addGoals'),
	path('removeGoals/<int:performanceReviewId>/<GoalsIds>/', PerformanceReviewView.removeGoals, name='removeGoals'),
]
