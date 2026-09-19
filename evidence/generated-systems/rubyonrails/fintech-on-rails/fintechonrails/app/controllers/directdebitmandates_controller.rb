class DirectDebitMandatesController < ApplicationController
  def index
    @directDebitMandates = DirectDebitMandate.all
  end
 
  def show
    @directDebitMandate = DirectDebitMandate.find(params[:id])
  end
 
  def new
    @directDebitMandate = DirectDebitMandate.new
  end
 
  def edit
    @directDebitMandate = DirectDebitMandate.find(params[:id])
  end
 
  def create
    @directDebitMandate = DirectDebitMandate.new(directDebitMandate_params)
 
    if @directDebitMandate.save
      redirect_to directDebitMandates_path
    else
      render 'new'
    end
  end
 
  def update
    @directDebitMandate = DirectDebitMandate.find(params[:id])
 
    if @directDebitMandate.update(directDebitMandate_params)
      redirect_to directDebitMandates_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @directDebitMandate = DirectDebitMandate.find(params[:id])
    @directDebitMandate.destroy
    redirect_to directDebitMandates_path
  end

 
  private
    def directDebitMandate_params
      params.require(:directDebitMandate).permit(:mandateId, :signedAt, :Scheme, :Status)
    end
end