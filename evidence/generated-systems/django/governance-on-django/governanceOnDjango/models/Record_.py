from django.db import models
from governanceOnDjango.models.RecordType import RecordType
from governanceOnDjango.models.DataClassificationLevel import DataClassificationLevel
from governanceOnDjango.models.RecordStatus import RecordStatus

#======================================================================
# 
# Encapsulates data for model Record_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Record_ Declaration
#======================================================================
class Record_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	creationDate = models.DateField(null=True)
	repository = models.ForeignKey('RecordsRepository', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	retentionSchedule = models.ForeignKey('RetentionSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	dataCategories = models.ManyToManyField('DataCategory',  blank=True, related_name='+')
	legalHolds = models.ManyToManyField('LegalHold',  blank=True, related_name='+')
	dataSubjectRequests = models.ManyToManyField('DataSubjectRequest',  blank=True, related_name='+')
	recordType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RecordType])
	classification = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataClassificationLevel])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RecordStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.creationDate
		str = str + self.recordType
		str = str + self.classification
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Record_";
    
	def objectType(self):
		return "Record_";
