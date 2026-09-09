from django.db import models
from governanceOnDjango.models.RetentionTrigger import RetentionTrigger
from governanceOnDjango.models.DispositionAction import DispositionAction
from governanceOnDjango.models.RetentionStatus import RetentionStatus

#======================================================================
# 
# Encapsulates data for model RetentionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionSchedule Declaration
#======================================================================
class RetentionSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	retentionPeriodMonths = models.IntegerField(null=True)
	repositories = models.ManyToManyField('RecordsRepository',  blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	exceptions = models.ManyToManyField('Exception_',  blank=True, related_name='+')
	dispositionReviews = models.ManyToManyField('DispositionReview',  blank=True, related_name='+')
	retentionTrigger = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RetentionTrigger])
	dispositionAction = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DispositionAction])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RetentionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.retentionPeriodMonths
		str = str + self.retentionTrigger
		str = str + self.dispositionAction
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RetentionSchedule";
    
	def objectType(self):
		return "RetentionSchedule";
