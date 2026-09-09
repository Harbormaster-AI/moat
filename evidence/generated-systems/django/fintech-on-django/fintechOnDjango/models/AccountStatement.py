from django.db import models

#======================================================================
# 
# Encapsulates data for model AccountStatement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountStatement Declaration
#======================================================================
class AccountStatement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	statementNumber = models.CharField(max_length=200, null=True)
	periodStart = models.DateField(null=True)
	periodEnd = models.DateField(null=True)
	openingBalance = Money
	closingBalance = Money
	generatedAt = DateTime
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.statementNumber
		str = str + self.periodStart
		str = str + self.periodEnd
		str = str + self.openingBalance
		str = str + self.closingBalance
		str = str + self.generatedAt
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AccountStatement";
    
	def objectType(self):
		return "AccountStatement";
