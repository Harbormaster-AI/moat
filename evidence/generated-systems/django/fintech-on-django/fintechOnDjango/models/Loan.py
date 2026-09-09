from django.db import models
from fintechOnDjango.models.InterestRateType import InterestRateType
from fintechOnDjango.models.LoanStatus import LoanStatus

#======================================================================
# 
# Encapsulates data for model Loan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Loan Declaration
#======================================================================
class Loan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	loanNumber = models.CharField(max_length=200, null=True)
	principal = Money
	interestRate = models.CharField(max_length=64, null=True)
	originationDate = models.DateField(null=True)
	maturityDate = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	schedule = models.ManyToManyField('RepaymentSchedule',  blank=True, related_name='+')
	collateral = models.ManyToManyField('Collateral',  blank=True, related_name='+')
	transactions = models.ManyToManyField('LoanTransaction',  blank=True, related_name='+')
	rateType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InterestRateType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LoanStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.loanNumber
		str = str + self.principal
		str = str + self.interestRate
		str = str + self.originationDate
		str = str + self.maturityDate
		str = str + self.rateType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Loan";
    
	def objectType(self):
		return "Loan";
