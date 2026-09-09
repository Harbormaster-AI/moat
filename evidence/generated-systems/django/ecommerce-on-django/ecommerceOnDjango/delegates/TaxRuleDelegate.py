from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.TaxRule import TaxRule
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TaxRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxRuleDelegate Declaration
#======================================================================
class TaxRuleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, taxRuleId ):
		try:	
			taxRule = TaxRule.objects.filter(id=taxRuleId)
			return taxRule.first();
		except TaxRule.DoesNotExist:
			raise ProcessingError("TaxRule with id " + str(taxRuleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, taxRule):
		for model in serializers.deserialize("json", taxRule):
			model.save()
			return model;

	def create(self, taxRule):
		taxRule.save()
		return taxRule;

	def saveFromJson(self, taxRule):
		for model in serializers.deserialize("json", taxRule):
			model.save()
			return taxRule;
	
	def save(self, taxRule):
		taxRule.save()
		return taxRule;
	
	def delete(self, taxRuleId ):
		errMsg = "Failed to delete TaxRule from db using id " + str(taxRuleId)
		
		try:
			taxRule = TaxRule.objects.get(id=taxRuleId)
			taxRule.delete()
			return True
		except TaxRule.DoesNotExist:
			raise ProcessingError("TaxRule with id " + str(taxRuleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TaxRule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TaxRule from db")
		except Exception:
			return None;
		
	def assignMerchant( self, taxRuleId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on TaxRule"

		try:
			# get the TaxRule from db
			taxRule = self.get( taxRuleId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			taxRule.merchant = merchant
			
			#save it
			taxRule.save()

			# reload and return the appropriate version					
			return self.get( taxRuleId );
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule with id " + str(taxRuleId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, taxRuleId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on TaxRule"

		try:
			# get the TaxRule from db
			taxRule = self.get( taxRuleId ).first()	
			
			# assign to None for unassignment
			taxRule.merchant = None			

			#save it
			taxRule.save()

			# reload and return the appropriate version					
			return self.get( taxRuleId );
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule with id " + str(taxRuleId) + " does not exist.")
		except Exception:
			return None;
		
	def addChannels( self, taxRuleId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to add elements " + str(channelsIds) + " for Channels on TaxRule"

		try:
			# get the TaxRule
			taxRule = self.get( taxRuleId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				taxRule.channels.add(channel)
				
			# save it		
			taxRule.save()
			
			# reload and return the appropriate version
			return self.get( taxRuleId );
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule with id " + str(taxRuleId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChannels( self, taxRuleId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to remove elements " + str(channelsIds) + " for Channels on TaxRule"

		try:
			# get the TaxRule
			taxRule = self.get( taxRuleId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				taxRule.channels.remove(channel)
				
			# save it		
			taxRule.save()
			
			# reload and return the appropriate version
			return self.get( taxRuleId );
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule with id " + str(taxRuleId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
