from django.db import models
from hrOnDjango.models.ApplicationStatus import ApplicationStatus

#======================================================================
# 
# Encapsulates data for model JobApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobApplication Declaration
#======================================================================
class JobApplication (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	applicationNumber = models.CharField(max_length=200, null=True)
	appliedDate = models.DateField(null=True)
	resumeUrl = models.CharField(max_length=200, null=True)
	candidate = models.ForeignKey('Candidate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	requisition = models.ForeignKey('JobRequisition', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	screenings = models.ManyToManyField('Screening',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ApplicationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.applicationNumber
		str = str + self.appliedDate
		str = str + self.resumeUrl
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "JobApplication";
    
	def objectType(self):
		return "JobApplication";
