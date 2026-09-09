from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.CustomerAddress import CustomerAddress
from ecommerceOnDjango.models.Cart import Cart
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.models.Review import Review
from ecommerceOnDjango.models.Wishlist import Wishlist
from ecommerceOnDjango.models.Subscription import Subscription
from ecommerceOnDjango.models.CouponRedemption import CouponRedemption
from ecommerceOnDjango.models.GiftCard import GiftCard
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerDelegate Declaration
#======================================================================
class CustomerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, customerId ):
		try:	
			customer = Customer.objects.filter(id=customerId)
			return customer.first();
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return model;

	def create(self, customer):
		customer.save()
		return customer;

	def saveFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return customer;
	
	def save(self, customer):
		customer.save()
		return customer;
	
	def delete(self, customerId ):
		errMsg = "Failed to delete Customer from db using id " + str(customerId)
		
		try:
			customer = Customer.objects.get(id=customerId)
			customer.delete()
			return True
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Customer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Customer from db")
		except Exception:
			return None;
		
	def addAddresses( self, customerId, addressesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerAddressDelegate import CustomerAddressDelegate

		errMsg = "Failed to add elements " + str(addressesIds) + " for Addresses on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = addressesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CustomerAddress		
				customerAddress = CustomerAddressDelegate().get(id).first();	
				# add the CustomerAddress
				customer.addresses.add(customerAddress)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except CustomerAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : CustomerAddress does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAddresses( self, customerId, addressesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerAddressDelegate import CustomerAddressDelegate

		errMsg = "Failed to remove elements " + str(addressesIds) + " for Addresses on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = addressesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CustomerAddress		
				customerAddress = CustomerAddressDelegate().get(id).first();	
				# add the CustomerAddress
				customer.addresses.remove(customerAddress)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except CustomerAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : CustomerAddress does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCarts( self, customerId, cartsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartDelegate import CartDelegate

		errMsg = "Failed to add elements " + str(cartsIds) + " for Carts on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = cartsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Cart		
				cart = CartDelegate().get(id).first();	
				# add the Cart
				customer.carts.add(cart)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCarts( self, customerId, cartsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartDelegate import CartDelegate

		errMsg = "Failed to remove elements " + str(cartsIds) + " for Carts on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = cartsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Cart		
				cart = CartDelegate().get(id).first();	
				# add the Cart
				customer.carts.remove(cart)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, customerId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				customer.orders.add(order)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, customerId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				customer.orders.remove(order)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayments( self, customerId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				customer.payments.add(payment)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, customerId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				customer.payments.remove(payment)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReviews( self, customerId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

		errMsg = "Failed to add elements " + str(reviewsIds) + " for Reviews on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Review		
				review = ReviewDelegate().get(id).first();	
				# add the Review
				customer.reviews.add(review)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReviews( self, customerId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

		errMsg = "Failed to remove elements " + str(reviewsIds) + " for Reviews on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Review		
				review = ReviewDelegate().get(id).first();	
				# add the Review
				customer.reviews.remove(review)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWishlists( self, customerId, wishlistsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistDelegate import WishlistDelegate

		errMsg = "Failed to add elements " + str(wishlistsIds) + " for Wishlists on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = wishlistsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Wishlist		
				wishlist = WishlistDelegate().get(id).first();	
				# add the Wishlist
				customer.wishlists.add(wishlist)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWishlists( self, customerId, wishlistsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistDelegate import WishlistDelegate

		errMsg = "Failed to remove elements " + str(wishlistsIds) + " for Wishlists on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = wishlistsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Wishlist		
				wishlist = WishlistDelegate().get(id).first();	
				# add the Wishlist
				customer.wishlists.remove(wishlist)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSubscriptions( self, customerId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to add elements " + str(subscriptionsIds) + " for Subscriptions on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				customer.subscriptions.add(subscription)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubscriptions( self, customerId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to remove elements " + str(subscriptionsIds) + " for Subscriptions on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				customer.subscriptions.remove(subscription)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCouponRedemptions( self, customerId, couponRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to add elements " + str(couponRedemptionsIds) + " for CouponRedemptions on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = couponRedemptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				customer.couponRedemptions.add(couponRedemption)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCouponRedemptions( self, customerId, couponRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to remove elements " + str(couponRedemptionsIds) + " for CouponRedemptions on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = couponRedemptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				customer.couponRedemptions.remove(couponRedemption)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGiftCards( self, customerId, giftCardsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardDelegate import GiftCardDelegate

		errMsg = "Failed to add elements " + str(giftCardsIds) + " for GiftCards on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = giftCardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GiftCard		
				giftCard = GiftCardDelegate().get(id).first();	
				# add the GiftCard
				customer.giftCards.add(giftCard)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGiftCards( self, customerId, giftCardsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardDelegate import GiftCardDelegate

		errMsg = "Failed to remove elements " + str(giftCardsIds) + " for GiftCards on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = giftCardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GiftCard		
				giftCard = GiftCardDelegate().get(id).first();	
				# add the GiftCard
				customer.giftCards.remove(giftCard)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
