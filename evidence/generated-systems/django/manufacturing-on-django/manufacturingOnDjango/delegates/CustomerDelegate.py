from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Customer import Customer
from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def addEnterprises( self, customerId, enterprisesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to add elements " + str(enterprisesIds) + " for Enterprises on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = enterprisesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Enterprise		
				enterprise = EnterpriseDelegate().get(id).first();	
				# add the Enterprise
				customer.enterprises.add(enterprise)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEnterprises( self, customerId, enterprisesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to remove elements " + str(enterprisesIds) + " for Enterprises on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = enterprisesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Enterprise		
				enterprise = EnterpriseDelegate().get(id).first();	
				# add the Enterprise
				customer.enterprises.remove(enterprise)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSalesOrders( self, customerId, salesOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

		errMsg = "Failed to add elements " + str(salesOrdersIds) + " for SalesOrders on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = salesOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SalesOrder		
				salesOrder = SalesOrderDelegate().get(id).first();	
				# add the SalesOrder
				customer.salesOrders.add(salesOrder)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSalesOrders( self, customerId, salesOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

		errMsg = "Failed to remove elements " + str(salesOrdersIds) + " for SalesOrders on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = salesOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SalesOrder		
				salesOrder = SalesOrderDelegate().get(id).first();	
				# add the SalesOrder
				customer.salesOrders.remove(salesOrder)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
