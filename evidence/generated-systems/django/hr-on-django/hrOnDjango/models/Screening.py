from django.db import models
from hrOnDjango.models.BackgroundCheckStatus import BackgroundCheckStatus

#======================================================================
# 
# Encapsulates data for model Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Screening Declaration
#======================================================================
class Screening (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	completedDate = models.DateField(null=True)
	application = models.ForeignKey('JobApplication', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BackgroundCheckStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.completedDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Screening";
    
	def objectType(self):
		return "Screening";
