from django.db import models
from fintechOnDjango.models.LoanProductType import LoanProductType
from fintechOnDjango.models.LoanPurpose import LoanPurpose
from fintechOnDjango.models.ApplicationStatus import ApplicationStatus

#======================================================================
# 
# Encapsulates data for model LoanApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanApplication Declaration
#======================================================================
class LoanApplication (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	applicationNumber = models.CharField(max_length=200, null=True)
	amountRequested = Money
	termMonths = models.IntegerField(null=True)
	submittedAt = DateTime
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	riskAssessment = models.ForeignKey('RiskAssessment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	loan = models.ForeignKey('Loan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LoanProductType])
	purpose = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LoanPurpose])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ApplicationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.applicationNumber
		str = str + self.amountRequested
		str = str + self.termMonths
		str = str + self.submittedAt
		str = str + self.product
		str = str + self.purpose
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LoanApplication";
    
	def objectType(self):
		return "LoanApplication";
