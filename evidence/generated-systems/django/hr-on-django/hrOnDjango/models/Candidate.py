from django.db import models
from hrOnDjango.models.CandidateSource import CandidateSource

#======================================================================
# 
# Encapsulates data for model Candidate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Candidate Declaration
#======================================================================
class Candidate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = PersonName
	email = Email
	phone = PhoneNumber
	applications = models.ManyToManyField('JobApplication',  blank=True, related_name='+')
	interviews = models.ManyToManyField('Interview',  blank=True, related_name='+')
	offers = models.ManyToManyField('Offer',  blank=True, related_name='+')
	documents = models.ManyToManyField('Document',  blank=True, related_name='+')
	source = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CandidateSource])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.email
		str = str + self.phone
		str = str + self.source
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Candidate";
    
	def objectType(self):
		return "Candidate";
