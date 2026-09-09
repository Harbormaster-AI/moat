from django.db import models
from governanceOnDjango.models.RepositoryType import RepositoryType

#======================================================================
# 
# Encapsulates data for model RecordsRepository
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordsRepository Declaration
#======================================================================
class RecordsRepository (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	location = models.CharField(max_length=200, null=True)
	ownerDepartment = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	systems = models.ManyToManyField('System_',  blank=True, related_name='+')
	retentionSchedules = models.ManyToManyField('RetentionSchedule',  blank=True, related_name='+')
	legalHolds = models.ManyToManyField('LegalHold',  blank=True, related_name='+')
	repositoryType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RepositoryType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.location
		str = str + self.ownerDepartment
		str = str + self.repositoryType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RecordsRepository";
    
	def objectType(self):
		return "RecordsRepository";
