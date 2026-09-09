from django.db import models
from hrOnDjango.models.JobLevel import JobLevel
from hrOnDjango.models.ExemptStatus import ExemptStatus

#======================================================================
# 
# Encapsulates data for model JobProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobProfile Declaration
#======================================================================
class JobProfile (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	jobCode = models.CharField(max_length=200, null=True)
	jobFamily = models.ForeignKey('JobFamily', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	competencies = models.ManyToManyField('Competency',  blank=True, related_name='+')
	trainingRecommendations = models.ManyToManyField('TrainingCourse',  blank=True, related_name='+')
	positions = models.ManyToManyField('Position',  blank=True, related_name='+')
	jobLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in JobLevel])
	exemptStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExemptStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.jobCode
		str = str + self.jobLevel
		str = str + self.exemptStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "JobProfile";
    
	def objectType(self):
		return "JobProfile";
