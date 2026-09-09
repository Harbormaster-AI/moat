from django.urls import path
from advertisingOnDjango.views import BrandSafetyPolicyView

urlpatterns = [
    path('', BrandSafetyPolicyView.index, name='index'),
	path('create', BrandSafetyPolicyView.get, name='create'),
	path('get/<int:brandSafetyPolicyId>/', BrandSafetyPolicyView.get, name='get'),
	path('save', BrandSafetyPolicyView.save, name='save'),
	path('getAll', BrandSafetyPolicyView.getAll, name='getAll'),
	path('delete/<int:brandSafetyPolicyId>/', BrandSafetyPolicyView.delete, name='delete'),
	path('addTargetingProfiles/<int:brandSafetyPolicyId>/<TargetingProfilesIds>/', BrandSafetyPolicyView.addTargetingProfiles, name='addTargetingProfiles'),
	path('removeTargetingProfiles/<int:brandSafetyPolicyId>/<TargetingProfilesIds>/', BrandSafetyPolicyView.removeTargetingProfiles, name='removeTargetingProfiles'),
]
