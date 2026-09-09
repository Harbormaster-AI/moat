from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.Brand import Brand
from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.models.TaxRule import TaxRule
from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MerchantDelegate Declaration
#======================================================================
class MerchantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, merchantId ):
		try:	
			merchant = Merchant.objects.filter(id=merchantId)
			return merchant.first();
		except Merchant.DoesNotExist:
			raise ProcessingError("Merchant with id " + str(merchantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, merchant):
		for model in serializers.deserialize("json", merchant):
			model.save()
			return model;

	def create(self, merchant):
		merchant.save()
		return merchant;

	def saveFromJson(self, merchant):
		for model in serializers.deserialize("json", merchant):
			model.save()
			return merchant;
	
	def save(self, merchant):
		merchant.save()
		return merchant;
	
	def delete(self, merchantId ):
		errMsg = "Failed to delete Merchant from db using id " + str(merchantId)
		
		try:
			merchant = Merchant.objects.get(id=merchantId)
			merchant.delete()
			return True
		except Merchant.DoesNotExist:
			raise ProcessingError("Merchant with id " + str(merchantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Merchant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Merchant from db")
		except Exception:
			return None;
		
	def addChannels( self, merchantId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to add elements " + str(channelsIds) + " for Channels on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				merchant.channels.add(channel)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChannels( self, merchantId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to remove elements " + str(channelsIds) + " for Channels on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				merchant.channels.remove(channel)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBrands( self, merchantId, brandsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.BrandDelegate import BrandDelegate

		errMsg = "Failed to add elements " + str(brandsIds) + " for Brands on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = brandsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Brand		
				brand = BrandDelegate().get(id).first();	
				# add the Brand
				merchant.brands.add(brand)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBrands( self, merchantId, brandsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.BrandDelegate import BrandDelegate

		errMsg = "Failed to remove elements " + str(brandsIds) + " for Brands on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = brandsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Brand		
				brand = BrandDelegate().get(id).first();	
				# add the Brand
				merchant.brands.remove(brand)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFulfillmentCenters( self, merchantId, fulfillmentCentersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to add elements " + str(fulfillmentCentersIds) + " for FulfillmentCenters on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = fulfillmentCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FulfillmentCenter		
				fulfillmentCenter = FulfillmentCenterDelegate().get(id).first();	
				# add the FulfillmentCenter
				merchant.fulfillmentCenters.add(fulfillmentCenter)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFulfillmentCenters( self, merchantId, fulfillmentCentersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to remove elements " + str(fulfillmentCentersIds) + " for FulfillmentCenters on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = fulfillmentCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FulfillmentCenter		
				fulfillmentCenter = FulfillmentCenterDelegate().get(id).first();	
				# add the FulfillmentCenter
				merchant.fulfillmentCenters.remove(fulfillmentCenter)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTaxRules( self, merchantId, taxRulesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.TaxRuleDelegate import TaxRuleDelegate

		errMsg = "Failed to add elements " + str(taxRulesIds) + " for TaxRules on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = taxRulesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TaxRule		
				taxRule = TaxRuleDelegate().get(id).first();	
				# add the TaxRule
				merchant.taxRules.add(taxRule)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTaxRules( self, merchantId, taxRulesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.TaxRuleDelegate import TaxRuleDelegate

		errMsg = "Failed to remove elements " + str(taxRulesIds) + " for TaxRules on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = taxRulesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TaxRule		
				taxRule = TaxRuleDelegate().get(id).first();	
				# add the TaxRule
				merchant.taxRules.remove(taxRule)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except TaxRule.DoesNotExist:
			raise ProcessingError(errMsg + " : TaxRule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPaymentProviders( self, merchantId, paymentProvidersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to add elements " + str(paymentProvidersIds) + " for PaymentProviders on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = paymentProvidersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentProvider		
				paymentProvider = PaymentProviderDelegate().get(id).first();	
				# add the PaymentProvider
				merchant.paymentProviders.add(paymentProvider)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentProviders( self, merchantId, paymentProvidersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to remove elements " + str(paymentProvidersIds) + " for PaymentProviders on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = paymentProvidersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentProvider		
				paymentProvider = PaymentProviderDelegate().get(id).first();	
				# add the PaymentProvider
				merchant.paymentProviders.remove(paymentProvider)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSellers( self, merchantId, sellersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

		errMsg = "Failed to add elements " + str(sellersIds) + " for Sellers on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = sellersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Seller		
				seller = SellerDelegate().get(id).first();	
				# add the Seller
				merchant.sellers.add(seller)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSellers( self, merchantId, sellersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

		errMsg = "Failed to remove elements " + str(sellersIds) + " for Sellers on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = sellersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Seller		
				seller = SellerDelegate().get(id).first();	
				# add the Seller
				merchant.sellers.remove(seller)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPromotions( self, merchantId, promotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(promotionsIds) + " for Promotions on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = promotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				merchant.promotions.add(promotion)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePromotions( self, merchantId, promotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(promotionsIds) + " for Promotions on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = promotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				merchant.promotions.remove(promotion)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
