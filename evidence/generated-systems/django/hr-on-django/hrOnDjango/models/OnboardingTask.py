from django.db import models
from hrOnDjango.models.OnboardingTaskStatus import OnboardingTaskStatus

#======================================================================
# 
# Encapsulates data for model OnboardingTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OnboardingTask Declaration
#======================================================================
class OnboardingTask (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	taskNumber = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	dueDate = models.DateField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assignedTo = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dependencies = models.ManyToManyField('OnboardingTask',  blank=True, related_name='+')
	relatedOffer = models.ForeignKey('Offer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OnboardingTaskStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.taskNumber
		str = str + self.name
		str = str + self.dueDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "OnboardingTask";
    
	def objectType(self):
		return "OnboardingTask";
