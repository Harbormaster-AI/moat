class ReinsuranceAgreementsController < ApplicationController
  def index
    @reinsuranceAgreements = ReinsuranceAgreement.all
  end
 
  def show
    @reinsuranceAgreement = ReinsuranceAgreement.find(params[:id])
  end
 
  def new
    @reinsuranceAgreement = ReinsuranceAgreement.new
  end
 
  def edit
    @reinsuranceAgreement = ReinsuranceAgreement.find(params[:id])
  end
 
  def create
    @reinsuranceAgreement = ReinsuranceAgreement.new(reinsuranceAgreement_params)
 
    if @reinsuranceAgreement.save
      redirect_to reinsuranceAgreements_path
    else
      render 'new'
    end
  end
 
  def update
    @reinsuranceAgreement = ReinsuranceAgreement.find(params[:id])
 
    if @reinsuranceAgreement.update(reinsuranceAgreement_params)
      redirect_to reinsuranceAgreements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @reinsuranceAgreement = ReinsuranceAgreement.find(params[:id])
    @reinsuranceAgreement.destroy
    redirect_to reinsuranceAgreements_path
  end

 
  private
    def reinsuranceAgreement_params
      params.require(:reinsuranceAgreement).permit(:agreementNumber, :effectivePeriod, :retention, :limit, :cessionPercentage, :ReinsuranceType, :TreatyType)
    end
end