from django.db import models
from hrOnDjango.models.BackgroundCheckStatus import BackgroundCheckStatus

#======================================================================
# 
# Encapsulates data for model BackgroundCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BackgroundCheck Declaration
#======================================================================
class BackgroundCheck (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	checkNumber = models.CharField(max_length=200, null=True)
	provider = models.CharField(max_length=200, null=True)
	completedDate = models.DateField(null=True)
	candidate = models.ForeignKey('Candidate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	requisition = models.ForeignKey('JobRequisition', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	report = models.OneToOneField('Document', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BackgroundCheckStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.checkNumber
		str = str + self.provider
		str = str + self.completedDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BackgroundCheck";
    
	def objectType(self):
		return "BackgroundCheck";
