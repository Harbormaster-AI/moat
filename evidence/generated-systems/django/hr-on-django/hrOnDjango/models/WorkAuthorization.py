from django.db import models
from hrOnDjango.models.WorkAuthorizationStatus import WorkAuthorizationStatus

#======================================================================
# 
# Encapsulates data for model WorkAuthorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkAuthorization Declaration
#======================================================================
class WorkAuthorization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	country = models.CharField(max_length=200, null=True)
	expirationDate = models.DateField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	documents = models.ManyToManyField('Document',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WorkAuthorizationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.country
		str = str + self.expirationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkAuthorization";
    
	def objectType(self):
		return "WorkAuthorization";
