from django.db import models

#======================================================================
# 
# Encapsulates data for model Regulation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Regulation Declaration
#======================================================================
class Regulation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	citation = models.CharField(max_length=200, null=True)
	jurisdiction = models.CharField(max_length=200, null=True)
	publicationUrl = URL
	obligations = models.ManyToManyField('Obligation',  blank=True, related_name='+')
	compliancePrograms = models.ManyToManyField('ComplianceProgram',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.citation
		str = str + self.jurisdiction
		str = str + self.publicationUrl
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Regulation";
    
	def objectType(self):
		return "Regulation";
