from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Application import Application
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.models.Distributor import Distributor
from insuranceOnDjango.models.Quote import Quote
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Application
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApplicationDelegate Declaration
#======================================================================
class ApplicationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, applicationId ):
		try:	
			application = Application.objects.filter(id=applicationId)
			return application.first();
		except Application.DoesNotExist:
			raise ProcessingError("Application with id " + str(applicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, application):
		for model in serializers.deserialize("json", application):
			model.save()
			return model;

	def create(self, application):
		application.save()
		return application;

	def saveFromJson(self, application):
		for model in serializers.deserialize("json", application):
			model.save()
			return application;
	
	def save(self, application):
		application.save()
		return application;
	
	def delete(self, applicationId ):
		errMsg = "Failed to delete Application from db using id " + str(applicationId)
		
		try:
			application = Application.objects.get(id=applicationId)
			application.delete()
			return True
		except Application.DoesNotExist:
			raise ProcessingError("Application with id " + str(applicationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Application.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Application from db")
		except Exception:
			return None;
		
	def assignCustomer( self, applicationId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			application.customer = customer
			
			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, applicationId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# assign to None for unassignment
			application.customer = None			

			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, applicationId, productId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# get the InsuranceProduct from db
			insuranceProduct = InsuranceProductDelegate().get(productId).first();
			
			# assign the Product		
			application.product = insuranceProduct
			
			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, applicationId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# assign to None for unassignment
			application.insuranceProduct = None			

			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDistributor( self, applicationId, distributorId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

		errMsg = "Failed to assign element " + str(distributorId) + " for Distributor on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# get the Distributor from db
			distributor = DistributorDelegate().get(distributorId).first();
			
			# assign the Distributor		
			application.distributor = distributor
			
			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDistributor( self, applicationId ):
		errMsg = "Failed to unassign element " + str(distributorId) + " for Distributor on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# assign to None for unassignment
			application.distributor = None			

			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSelectedQuote( self, applicationId, selectedQuoteId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to assign element " + str(selectedQuoteId) + " for SelectedQuote on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# get the Quote from db
			quote = QuoteDelegate().get(selectedQuoteId).first();
			
			# assign the SelectedQuote		
			application.selectedQuote = quote
			
			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(selectedQuoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSelectedQuote( self, applicationId ):
		errMsg = "Failed to unassign element " + str(selectedQuoteId) + " for SelectedQuote on Application"

		try:
			# get the Application from db
			application = self.get( applicationId ).first()	
			
			# assign to None for unassignment
			application.quote = None			

			#save it
			application.save()

			# reload and return the appropriate version					
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
		
	def addQuotes( self, applicationId, quotesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on Application"

		try:
			# get the Application
			application = self.get( applicationId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				application.quotes.add(quote)
				
			# save it		
			application.save()
			
			# reload and return the appropriate version
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, applicationId, quotesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on Application"

		try:
			# get the Application
			application = self.get( applicationId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				application.quotes.remove(quote)
				
			# save it		
			application.save()
			
			# reload and return the appropriate version
			return self.get( applicationId );
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
