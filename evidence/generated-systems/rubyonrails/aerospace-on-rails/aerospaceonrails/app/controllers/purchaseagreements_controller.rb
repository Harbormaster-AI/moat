class PurchaseAgreementsController < ApplicationController
  def index
    @purchaseAgreements = PurchaseAgreement.all
  end
 
  def show
    @purchaseAgreement = PurchaseAgreement.find(params[:id])
  end
 
  def new
    @purchaseAgreement = PurchaseAgreement.new
  end
 
  def edit
    @purchaseAgreement = PurchaseAgreement.find(params[:id])
  end
 
  def create
    @purchaseAgreement = PurchaseAgreement.new(purchaseAgreement_params)
 
    if @purchaseAgreement.save
      redirect_to purchaseAgreements_path
    else
      render 'new'
    end
  end
 
  def update
    @purchaseAgreement = PurchaseAgreement.find(params[:id])
 
    if @purchaseAgreement.update(purchaseAgreement_params)
      redirect_to purchaseAgreements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @purchaseAgreement = PurchaseAgreement.find(params[:id])
    @purchaseAgreement.destroy
    redirect_to purchaseAgreements_path
  end

 
  private
    def purchaseAgreement_params
      params.require(:purchaseAgreement).permit(:agreementNumber, :effectiveDate)
    end
end