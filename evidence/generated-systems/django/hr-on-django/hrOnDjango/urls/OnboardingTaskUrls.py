from django.urls import path
from hrOnDjango.views import OnboardingTaskView

urlpatterns = [
    path('', OnboardingTaskView.index, name='index'),
	path('create', OnboardingTaskView.get, name='create'),
	path('get/<int:onboardingTaskId>/', OnboardingTaskView.get, name='get'),
	path('save', OnboardingTaskView.save, name='save'),
	path('getAll', OnboardingTaskView.getAll, name='getAll'),
	path('delete/<int:onboardingTaskId>/', OnboardingTaskView.delete, name='delete'),
	path('assignEmployee/<int:onboardingTaskId>/<int:EmployeeId>/', OnboardingTaskView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:onboardingTaskId>/', OnboardingTaskView.unassignEmployee, name='unassignEmployee'),
	path('assignAssignedTo/<int:onboardingTaskId>/<int:AssignedToId>/', OnboardingTaskView.assignAssignedTo, name='assignAssignedTo'),
	path('unassignAssignedTo/<int:onboardingTaskId>/', OnboardingTaskView.unassignAssignedTo, name='unassignAssignedTo'),
	path('assignRelatedOffer/<int:onboardingTaskId>/<int:RelatedOfferId>/', OnboardingTaskView.assignRelatedOffer, name='assignRelatedOffer'),
	path('unassignRelatedOffer/<int:onboardingTaskId>/', OnboardingTaskView.unassignRelatedOffer, name='unassignRelatedOffer'),
	path('addDependencies/<int:onboardingTaskId>/<DependenciesIds>/', OnboardingTaskView.addDependencies, name='addDependencies'),
	path('removeDependencies/<int:onboardingTaskId>/<DependenciesIds>/', OnboardingTaskView.removeDependencies, name='removeDependencies'),
]
