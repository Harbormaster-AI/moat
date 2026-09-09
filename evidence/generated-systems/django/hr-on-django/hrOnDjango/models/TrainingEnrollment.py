from django.db import models
from hrOnDjango.models.TrainingStatus import TrainingStatus

#======================================================================
# 
# Encapsulates data for model TrainingEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingEnrollment Declaration
#======================================================================
class TrainingEnrollment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	enrollmentNumber = models.CharField(max_length=200, null=True)
	completionDate = models.DateField(null=True)
	score = models.CharField(max_length=64, null=True)
	course = models.ForeignKey('TrainingCourse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	instructor = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TrainingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.enrollmentNumber
		str = str + self.completionDate
		str = str + self.score
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TrainingEnrollment";
    
	def objectType(self):
		return "TrainingEnrollment";
