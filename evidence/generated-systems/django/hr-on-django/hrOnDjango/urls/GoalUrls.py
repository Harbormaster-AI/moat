from django.urls import path
from hrOnDjango.views import GoalView

urlpatterns = [
    path('', GoalView.index, name='index'),
	path('create', GoalView.get, name='create'),
	path('get/<int:goalId>/', GoalView.get, name='get'),
	path('save', GoalView.save, name='save'),
	path('getAll', GoalView.getAll, name='getAll'),
	path('delete/<int:goalId>/', GoalView.delete, name='delete'),
	path('assignEmployee/<int:goalId>/<int:EmployeeId>/', GoalView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:goalId>/', GoalView.unassignEmployee, name='unassignEmployee'),
	path('assignCycle/<int:goalId>/<int:CycleId>/', GoalView.assignCycle, name='assignCycle'),
	path('unassignCycle/<int:goalId>/', GoalView.unassignCycle, name='unassignCycle'),
	path('assignParentGoal/<int:goalId>/<int:ParentGoalId>/', GoalView.assignParentGoal, name='assignParentGoal'),
	path('unassignParentGoal/<int:goalId>/', GoalView.unassignParentGoal, name='unassignParentGoal'),
	path('addChildGoals/<int:goalId>/<ChildGoalsIds>/', GoalView.addChildGoals, name='addChildGoals'),
	path('removeChildGoals/<int:goalId>/<ChildGoalsIds>/', GoalView.removeChildGoals, name='removeChildGoals'),
]
