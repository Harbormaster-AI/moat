class CertificationsController < ApplicationController
  def index
    @certifications = Certification.all
  end
 
  def show
    @certification = Certification.find(params[:id])
  end
 
  def new
    @certification = Certification.new
  end
 
  def edit
    @certification = Certification.find(params[:id])
  end
 
  def create
    @certification = Certification.new(certification_params)
 
    if @certification.save
      redirect_to certifications_path
    else
      render 'new'
    end
  end
 
  def update
    @certification = Certification.find(params[:id])
 
    if @certification.update(certification_params)
      redirect_to certifications_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @certification = Certification.find(params[:id])
    @certification.destroy
    redirect_to certifications_path
  end

 
  private
    def certification_params
      params.require(:certification).permit(:name, :issuer, :validFrom, :validTo, :credentialId)
    end
end