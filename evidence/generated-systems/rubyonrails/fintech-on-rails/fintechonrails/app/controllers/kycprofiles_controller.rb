class KYCProfilesController < ApplicationController
  def index
    @kYCProfiles = KYCProfile.all
  end
 
  def show
    @kYCProfile = KYCProfile.find(params[:id])
  end
 
  def new
    @kYCProfile = KYCProfile.new
  end
 
  def edit
    @kYCProfile = KYCProfile.find(params[:id])
  end
 
  def create
    @kYCProfile = KYCProfile.new(kYCProfile_params)
 
    if @kYCProfile.save
      redirect_to kYCProfiles_path
    else
      render 'new'
    end
  end
 
  def update
    @kYCProfile = KYCProfile.find(params[:id])
 
    if @kYCProfile.update(kYCProfile_params)
      redirect_to kYCProfiles_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @kYCProfile = KYCProfile.find(params[:id])
    @kYCProfile.destroy
    redirect_to kYCProfiles_path
  end

 
  private
    def kYCProfile_params
      params.require(:kYCProfile).permit(:profileId, :createdAt, :Status, :VerificationLevel)
    end
end