from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.ImagingCenter import ImagingCenter
from healthcareOnDjango.models.ImagingReport import ImagingReport
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ImagingOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingOrderDelegate Declaration
#======================================================================
class ImagingOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, imagingOrderId ):
		try:	
			imagingOrder = ImagingOrder.objects.filter(id=imagingOrderId)
			return imagingOrder.first();
		except ImagingOrder.DoesNotExist:
			raise ProcessingError("ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, imagingOrder):
		for model in serializers.deserialize("json", imagingOrder):
			model.save()
			return model;

	def create(self, imagingOrder):
		imagingOrder.save()
		return imagingOrder;

	def saveFromJson(self, imagingOrder):
		for model in serializers.deserialize("json", imagingOrder):
			model.save()
			return imagingOrder;
	
	def save(self, imagingOrder):
		imagingOrder.save()
		return imagingOrder;
	
	def delete(self, imagingOrderId ):
		errMsg = "Failed to delete ImagingOrder from db using id " + str(imagingOrderId)
		
		try:
			imagingOrder = ImagingOrder.objects.get(id=imagingOrderId)
			imagingOrder.delete()
			return True
		except ImagingOrder.DoesNotExist:
			raise ProcessingError("ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ImagingOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ImagingOrder from db")
		except Exception:
			return None;
		
	def assignOrder( self, imagingOrderId, orderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on ImagingOrder"

		try:
			# get the ImagingOrder from db
			imagingOrder = self.get( imagingOrderId ).first()	
			
			# get the ClinicalOrder from db
			clinicalOrder = ClinicalOrderDelegate().get(orderId).first();
			
			# assign the Order		
			imagingOrder.order = clinicalOrder
			
			#save it
			imagingOrder.save()

			# reload and return the appropriate version					
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, imagingOrderId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on ImagingOrder"

		try:
			# get the ImagingOrder from db
			imagingOrder = self.get( imagingOrderId ).first()	
			
			# assign to None for unassignment
			imagingOrder.clinicalOrder = None			

			#save it
			imagingOrder.save()

			# reload and return the appropriate version					
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignImagingCenter( self, imagingOrderId, imagingCenterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

		errMsg = "Failed to assign element " + str(imagingCenterId) + " for ImagingCenter on ImagingOrder"

		try:
			# get the ImagingOrder from db
			imagingOrder = self.get( imagingOrderId ).first()	
			
			# get the ImagingCenter from db
			imagingCenter = ImagingCenterDelegate().get(imagingCenterId).first();
			
			# assign the ImagingCenter		
			imagingOrder.imagingCenter = imagingCenter
			
			#save it
			imagingOrder.save()

			# reload and return the appropriate version					
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except ImagingCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingCenter with id " + str(imagingCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignImagingCenter( self, imagingOrderId ):
		errMsg = "Failed to unassign element " + str(imagingCenterId) + " for ImagingCenter on ImagingOrder"

		try:
			# get the ImagingOrder from db
			imagingOrder = self.get( imagingOrderId ).first()	
			
			# assign to None for unassignment
			imagingOrder.imagingCenter = None			

			#save it
			imagingOrder.save()

			# reload and return the appropriate version					
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addReports( self, imagingOrderId, reportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on ImagingOrder"

		try:
			# get the ImagingOrder
			imagingOrder = self.get( imagingOrderId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				imagingOrder.reports.add(imagingReport)
				
			# save it		
			imagingOrder.save()
			
			# reload and return the appropriate version
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, imagingOrderId, reportsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on ImagingOrder"

		try:
			# get the ImagingOrder
			imagingOrder = self.get( imagingOrderId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ImagingReport		
				imagingReport = ImagingReportDelegate().get(id).first();	
				# add the ImagingReport
				imagingOrder.reports.remove(imagingReport)
				
			# save it		
			imagingOrder.save()
			
			# reload and return the appropriate version
			return self.get( imagingOrderId );
		except ImagingOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingOrder with id " + str(imagingOrderId) + " does not exist.")
		except ImagingReport.DoesNotExist:
			raise ProcessingError(errMsg + " : ImagingReport does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
