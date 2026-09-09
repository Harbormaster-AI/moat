from django.db import models
from governanceOnDjango.models.LegalHoldStatus import LegalHoldStatus

#======================================================================
# 
# Encapsulates data for model LegalHold
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LegalHold Declaration
#======================================================================
class LegalHold (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	reason = models.CharField(max_length=200, null=True)
	issuedDate = models.DateField(null=True)
	releaseDate = models.DateField(null=True)
	repositories = models.ManyToManyField('RecordsRepository',  blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	matter = models.ForeignKey('Matter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	holdStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LegalHoldStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.reason
		str = str + self.issuedDate
		str = str + self.releaseDate
		str = str + self.holdStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LegalHold";
    
	def objectType(self):
		return "LegalHold";
