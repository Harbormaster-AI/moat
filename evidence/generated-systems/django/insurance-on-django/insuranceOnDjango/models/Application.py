from django.db import models
from insuranceOnDjango.models.ApplicationStatus import ApplicationStatus

#======================================================================
# 
# Encapsulates data for model Application
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Application Declaration
#======================================================================
class Application (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	applicationNumber = models.CharField(max_length=200, null=True)
	submissionDate = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.ForeignKey('InsuranceProduct', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	distributor = models.ForeignKey('Distributor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	selectedQuote = models.OneToOneField('Quote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ApplicationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.applicationNumber
		str = str + self.submissionDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Application";
    
	def objectType(self):
		return "Application";
