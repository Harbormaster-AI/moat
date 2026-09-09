from django.db import models
from hrOnDjango.models.CycleStatus import CycleStatus

#======================================================================
# 
# Encapsulates data for model PerformanceCycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceCycle Declaration
#======================================================================
class PerformanceCycle (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reviews = models.ManyToManyField('PerformanceReview',  blank=True, related_name='+')
	goals = models.ManyToManyField('Goal',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CycleStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PerformanceCycle";
    
	def objectType(self):
		return "PerformanceCycle";
