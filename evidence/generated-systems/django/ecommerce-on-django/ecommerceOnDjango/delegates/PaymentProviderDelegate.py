from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.models.Subscription import Subscription
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProviderDelegate Declaration
#======================================================================
class PaymentProviderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentProviderId ):
		try:	
			paymentProvider = PaymentProvider.objects.filter(id=paymentProviderId)
			return paymentProvider.first();
		except PaymentProvider.DoesNotExist:
			raise ProcessingError("PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentProvider):
		for model in serializers.deserialize("json", paymentProvider):
			model.save()
			return model;

	def create(self, paymentProvider):
		paymentProvider.save()
		return paymentProvider;

	def saveFromJson(self, paymentProvider):
		for model in serializers.deserialize("json", paymentProvider):
			model.save()
			return paymentProvider;
	
	def save(self, paymentProvider):
		paymentProvider.save()
		return paymentProvider;
	
	def delete(self, paymentProviderId ):
		errMsg = "Failed to delete PaymentProvider from db using id " + str(paymentProviderId)
		
		try:
			paymentProvider = PaymentProvider.objects.get(id=paymentProviderId)
			paymentProvider.delete()
			return True
		except PaymentProvider.DoesNotExist:
			raise ProcessingError("PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentProvider.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentProvider from db")
		except Exception:
			return None;
		
	def assignMerchant( self, paymentProviderId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on PaymentProvider"

		try:
			# get the PaymentProvider from db
			paymentProvider = self.get( paymentProviderId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			paymentProvider.merchant = merchant
			
			#save it
			paymentProvider.save()

			# reload and return the appropriate version					
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, paymentProviderId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on PaymentProvider"

		try:
			# get the PaymentProvider from db
			paymentProvider = self.get( paymentProviderId ).first()	
			
			# assign to None for unassignment
			paymentProvider.merchant = None			

			#save it
			paymentProvider.save()

			# reload and return the appropriate version					
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Exception:
			return None;
		
	def addChannels( self, paymentProviderId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to add elements " + str(channelsIds) + " for Channels on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				paymentProvider.channels.add(channel)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChannels( self, paymentProviderId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to remove elements " + str(channelsIds) + " for Channels on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				paymentProvider.channels.remove(channel)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayments( self, paymentProviderId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				paymentProvider.payments.add(payment)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, paymentProviderId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				paymentProvider.payments.remove(payment)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSubscriptions( self, paymentProviderId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to add elements " + str(subscriptionsIds) + " for Subscriptions on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				paymentProvider.subscriptions.add(subscription)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubscriptions( self, paymentProviderId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to remove elements " + str(subscriptionsIds) + " for Subscriptions on PaymentProvider"

		try:
			# get the PaymentProvider
			paymentProvider = self.get( paymentProviderId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				paymentProvider.subscriptions.remove(subscription)
				
			# save it		
			paymentProvider.save()
			
			# reload and return the appropriate version
			return self.get( paymentProviderId );
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
