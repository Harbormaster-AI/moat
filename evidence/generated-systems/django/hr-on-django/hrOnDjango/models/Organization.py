from django.db import models

#======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Organization Declaration
#======================================================================
class Organization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	registrationCountry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	departments = models.ManyToManyField('Department',  blank=True, related_name='+')
	locations = models.ManyToManyField('Location',  blank=True, related_name='+')
	jobFamilies = models.ManyToManyField('JobFamily',  blank=True, related_name='+')
	benefitPlans = models.ManyToManyField('BenefitPlan',  blank=True, related_name='+')
	costCenters = models.ManyToManyField('CostCenter',  blank=True, related_name='+')
	payrollCalendars = models.ManyToManyField('PayrollCalendar',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.registrationCountry
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Organization";
    
	def objectType(self):
		return "Organization";
