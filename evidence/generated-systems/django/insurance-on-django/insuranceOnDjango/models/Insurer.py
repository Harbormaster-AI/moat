from django.db import models

#======================================================================
# 
# Encapsulates data for model Insurer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Insurer Declaration
#======================================================================
class Insurer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	domicileCountry = models.CharField(max_length=200, null=True)
	naicNumber = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	products = models.ManyToManyField('InsuranceProduct',  blank=True, related_name='+')
	distributionPartners = models.ManyToManyField('Distributor',  blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	reinsuranceAgreements = models.ManyToManyField('ReinsuranceAgreement',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.domicileCountry
		str = str + self.naicNumber
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Insurer";
    
	def objectType(self):
		return "Insurer";
