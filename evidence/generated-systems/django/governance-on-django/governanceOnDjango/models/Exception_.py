from django.db import models
from governanceOnDjango.models.ExceptionType import ExceptionType
from governanceOnDjango.models.ExceptionStatus import ExceptionStatus

#======================================================================
# 
# Encapsulates data for model Exception_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Exception_ Declaration
#======================================================================
class Exception_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	justification = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	retentionSchedule = models.ForeignKey('RetentionSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	control = models.ForeignKey('Control', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	risk = models.ForeignKey('Risk', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	exceptionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExceptionType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExceptionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.justification
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.exceptionType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Exception_";
    
	def objectType(self):
		return "Exception_";
