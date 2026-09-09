from django.db import models
from hrOnDjango.models.DependentRelationship import DependentRelationship

#======================================================================
# 
# Encapsulates data for model Dependent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Dependent Declaration
#======================================================================
class Dependent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	birthDate = models.DateField(null=True)
	benefitEnrollment = models.ForeignKey('BenefitEnrollment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	relationship = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DependentRelationship])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.birthDate
		str = str + self.relationship
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Dependent";
    
	def objectType(self):
		return "Dependent";
