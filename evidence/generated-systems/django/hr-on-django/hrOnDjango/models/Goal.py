from django.db import models
from hrOnDjango.models.GoalStatus import GoalStatus

#======================================================================
# 
# Encapsulates data for model Goal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Goal Declaration
#======================================================================
class Goal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	targetDate = models.DateField(null=True)
	weight = Percentage
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	cycle = models.ForeignKey('PerformanceCycle', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	parentGoal = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	childGoals = models.ManyToManyField('Goal',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in GoalStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.description
		str = str + self.targetDate
		str = str + self.weight
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Goal";
    
	def objectType(self):
		return "Goal";
