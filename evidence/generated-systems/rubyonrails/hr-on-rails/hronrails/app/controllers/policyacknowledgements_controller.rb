class PolicyAcknowledgementsController < ApplicationController
  def index
    @policyAcknowledgements = PolicyAcknowledgement.all
  end
 
  def show
    @policyAcknowledgement = PolicyAcknowledgement.find(params[:id])
  end
 
  def new
    @policyAcknowledgement = PolicyAcknowledgement.new
  end
 
  def edit
    @policyAcknowledgement = PolicyAcknowledgement.find(params[:id])
  end
 
  def create
    @policyAcknowledgement = PolicyAcknowledgement.new(policyAcknowledgement_params)
 
    if @policyAcknowledgement.save
      redirect_to policyAcknowledgements_path
    else
      render 'new'
    end
  end
 
  def update
    @policyAcknowledgement = PolicyAcknowledgement.find(params[:id])
 
    if @policyAcknowledgement.update(policyAcknowledgement_params)
      redirect_to policyAcknowledgements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @policyAcknowledgement = PolicyAcknowledgement.find(params[:id])
    @policyAcknowledgement.destroy
    redirect_to policyAcknowledgements_path
  end

 
  private
    def policyAcknowledgement_params
      params.require(:policyAcknowledgement).permit(:acknowledgementDate, :Status)
    end
end