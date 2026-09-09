from django.db import models
from hrOnDjango.models.DeliveryMethod import DeliveryMethod

#======================================================================
# 
# Encapsulates data for model TrainingCourse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingCourse Declaration
#======================================================================
class TrainingCourse (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	title = models.CharField(max_length=200, null=True)
	durationHours = models.CharField(max_length=64, null=True)
	prerequisites = models.ManyToManyField('TrainingCourse',  blank=True, related_name='+')
	enrollments = models.ManyToManyField('TrainingEnrollment',  blank=True, related_name='+')
	jobProfiles = models.ManyToManyField('JobProfile',  blank=True, related_name='+')
	deliveryMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DeliveryMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.title
		str = str + self.durationHours
		str = str + self.deliveryMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TrainingCourse";
    
	def objectType(self):
		return "TrainingCourse";
