from django.db import models
from healthcareOnDjango.models.TaskStatus import TaskStatus
from healthcareOnDjango.models.Priority import Priority

#======================================================================
# 
# Encapsulates data for model CareTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTask Declaration
#======================================================================
class CareTask (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	description = models.CharField(max_length=200, null=True)
	dueDate = models.DateField(null=True)
	carePlan = models.ForeignKey('CarePlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assignedTo = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TaskStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Priority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.description
		str = str + self.dueDate
		str = str + self.status
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CareTask";
    
	def objectType(self):
		return "CareTask";
