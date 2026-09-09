from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Review import Review
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Review
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReviewDelegate Declaration
#======================================================================
class ReviewDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, reviewId ):
		try:	
			review = Review.objects.filter(id=reviewId)
			return review.first();
		except Review.DoesNotExist:
			raise ProcessingError("Review with id " + str(reviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, review):
		for model in serializers.deserialize("json", review):
			model.save()
			return model;

	def create(self, review):
		review.save()
		return review;

	def saveFromJson(self, review):
		for model in serializers.deserialize("json", review):
			model.save()
			return review;
	
	def save(self, review):
		review.save()
		return review;
	
	def delete(self, reviewId ):
		errMsg = "Failed to delete Review from db using id " + str(reviewId)
		
		try:
			review = Review.objects.get(id=reviewId)
			review.delete()
			return True
		except Review.DoesNotExist:
			raise ProcessingError("Review with id " + str(reviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Review.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Review from db")
		except Exception:
			return None;
		
	def assignProduct( self, reviewId, productId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			review.product = product
			
			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, reviewId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# assign to None for unassignment
			review.product = None			

			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, reviewId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			review.customer = customer
			
			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, reviewId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# assign to None for unassignment
			review.customer = None			

			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, reviewId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			review.order = order
			
			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, reviewId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Review"

		try:
			# get the Review from db
			review = self.get( reviewId ).first()	
			
			# assign to None for unassignment
			review.order = None			

			#save it
			review.save()

			# reload and return the appropriate version					
			return self.get( reviewId );
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review with id " + str(reviewId) + " does not exist.")
		except Exception:
			return None;
		
