from django.db import models
from hrOnDjango.models.InterviewStage import InterviewStage
from hrOnDjango.models.InterviewResult import InterviewResult

#======================================================================
# 
# Encapsulates data for model Interview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Interview Declaration
#======================================================================
class Interview (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	interviewDate = models.DateField(null=True)
	feedback = models.CharField(max_length=200, null=True)
	requisition = models.ForeignKey('JobRequisition', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	candidate = models.ForeignKey('Candidate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	interviewers = models.ManyToManyField('Employee',  blank=True, related_name='+')
	stage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InterviewStage])
	result = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InterviewResult])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.interviewDate
		str = str + self.feedback
		str = str + self.stage
		str = str + self.result
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Interview";
    
	def objectType(self):
		return "Interview";
