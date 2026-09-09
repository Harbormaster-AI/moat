from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Subscription import Subscription
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Subscription
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriptionDelegate Declaration
#======================================================================
class SubscriptionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, subscriptionId ):
		try:	
			subscription = Subscription.objects.filter(id=subscriptionId)
			return subscription.first();
		except Subscription.DoesNotExist:
			raise ProcessingError("Subscription with id " + str(subscriptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, subscription):
		for model in serializers.deserialize("json", subscription):
			model.save()
			return model;

	def create(self, subscription):
		subscription.save()
		return subscription;

	def saveFromJson(self, subscription):
		for model in serializers.deserialize("json", subscription):
			model.save()
			return subscription;
	
	def save(self, subscription):
		subscription.save()
		return subscription;
	
	def delete(self, subscriptionId ):
		errMsg = "Failed to delete Subscription from db using id " + str(subscriptionId)
		
		try:
			subscription = Subscription.objects.get(id=subscriptionId)
			subscription.delete()
			return True
		except Subscription.DoesNotExist:
			raise ProcessingError("Subscription with id " + str(subscriptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Subscription.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Subscription from db")
		except Exception:
			return None;
		
	def assignCustomer( self, subscriptionId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			subscription.customer = customer
			
			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, subscriptionId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# assign to None for unassignment
			subscription.customer = None			

			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, subscriptionId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			subscription.variant = productVariant
			
			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, subscriptionId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# assign to None for unassignment
			subscription.productVariant = None			

			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPaymentProvider( self, subscriptionId, paymentProviderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to assign element " + str(paymentProviderId) + " for PaymentProvider on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# get the PaymentProvider from db
			paymentProvider = PaymentProviderDelegate().get(paymentProviderId).first();
			
			# assign the PaymentProvider		
			subscription.paymentProvider = paymentProvider
			
			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPaymentProvider( self, subscriptionId ):
		errMsg = "Failed to unassign element " + str(paymentProviderId) + " for PaymentProvider on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# assign to None for unassignment
			subscription.paymentProvider = None			

			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignChannel( self, subscriptionId, channelId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to assign element " + str(channelId) + " for Channel on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# get the Channel from db
			channel = ChannelDelegate().get(channelId).first();
			
			# assign the Channel		
			subscription.channel = channel
			
			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChannel( self, subscriptionId ):
		errMsg = "Failed to unassign element " + str(channelId) + " for Channel on Subscription"

		try:
			# get the Subscription from db
			subscription = self.get( subscriptionId ).first()	
			
			# assign to None for unassignment
			subscription.channel = None			

			#save it
			subscription.save()

			# reload and return the appropriate version					
			return self.get( subscriptionId );
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription with id " + str(subscriptionId) + " does not exist.")
		except Exception:
			return None;
		
