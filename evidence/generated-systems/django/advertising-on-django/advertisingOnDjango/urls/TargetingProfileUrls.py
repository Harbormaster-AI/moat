from django.urls import path
from advertisingOnDjango.views import TargetingProfileView

urlpatterns = [
    path('', TargetingProfileView.index, name='index'),
	path('create', TargetingProfileView.get, name='create'),
	path('get/<int:targetingProfileId>/', TargetingProfileView.get, name='get'),
	path('save', TargetingProfileView.save, name='save'),
	path('getAll', TargetingProfileView.getAll, name='getAll'),
	path('delete/<int:targetingProfileId>/', TargetingProfileView.delete, name='delete'),
	path('assignBrandSafetyPolicy/<int:targetingProfileId>/<int:BrandSafetyPolicyId>/', TargetingProfileView.assignBrandSafetyPolicy, name='assignBrandSafetyPolicy'),
	path('unassignBrandSafetyPolicy/<int:targetingProfileId>/', TargetingProfileView.unassignBrandSafetyPolicy, name='unassignBrandSafetyPolicy'),
	path('addAudienceSegments/<int:targetingProfileId>/<AudienceSegmentsIds>/', TargetingProfileView.addAudienceSegments, name='addAudienceSegments'),
	path('removeAudienceSegments/<int:targetingProfileId>/<AudienceSegmentsIds>/', TargetingProfileView.removeAudienceSegments, name='removeAudienceSegments'),
	path('addGeoRegions/<int:targetingProfileId>/<GeoRegionsIds>/', TargetingProfileView.addGeoRegions, name='addGeoRegions'),
	path('removeGeoRegions/<int:targetingProfileId>/<GeoRegionsIds>/', TargetingProfileView.removeGeoRegions, name='removeGeoRegions'),
	path('addContentCategories/<int:targetingProfileId>/<ContentCategoriesIds>/', TargetingProfileView.addContentCategories, name='addContentCategories'),
	path('removeContentCategories/<int:targetingProfileId>/<ContentCategoriesIds>/', TargetingProfileView.removeContentCategories, name='removeContentCategories'),
	path('addDeviceCriteria/<int:targetingProfileId>/<DeviceCriteriaIds>/', TargetingProfileView.addDeviceCriteria, name='addDeviceCriteria'),
	path('removeDeviceCriteria/<int:targetingProfileId>/<DeviceCriteriaIds>/', TargetingProfileView.removeDeviceCriteria, name='removeDeviceCriteria'),
]
