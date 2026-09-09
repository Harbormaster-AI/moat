from django.db import models

#======================================================================
# 
# Encapsulates data for model Underwriter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Underwriter Declaration
#======================================================================
class Underwriter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	employeeId = models.CharField(max_length=200, null=True)
	authorityLimit = Money
	decisions = models.ManyToManyField('UnderwritingDecision',  blank=True, related_name='+')
	insurer = models.ForeignKey('Insurer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.employeeId
		str = str + self.authorityLimit
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Underwriter";
    
	def objectType(self):
		return "Underwriter";
