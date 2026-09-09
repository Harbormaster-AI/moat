from django.db import models
from fintechOnDjango.models.WalletProvider import WalletProvider
from fintechOnDjango.models.TokenizationStatus import TokenizationStatus

#======================================================================
# 
# Encapsulates data for model CardTokenization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CardTokenization Declaration
#======================================================================
class CardTokenization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	tokenReference = models.CharField(max_length=200, null=True)
	createdAt = DateTime
	card = models.ForeignKey('PaymentCard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	walletProvider = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WalletProvider])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TokenizationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.tokenReference
		str = str + self.createdAt
		str = str + self.walletProvider
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CardTokenization";
    
	def objectType(self):
		return "CardTokenization";
