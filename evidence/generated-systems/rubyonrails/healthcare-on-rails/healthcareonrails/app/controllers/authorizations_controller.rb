class AuthorizationsController < ApplicationController
  def index
    @authorizations = Authorization.all
  end
 
  def show
    @authorization = Authorization.find(params[:id])
  end
 
  def new
    @authorization = Authorization.new
  end
 
  def edit
    @authorization = Authorization.find(params[:id])
  end
 
  def create
    @authorization = Authorization.new(authorization_params)
 
    if @authorization.save
      redirect_to authorizations_path
    else
      render 'new'
    end
  end
 
  def update
    @authorization = Authorization.find(params[:id])
 
    if @authorization.update(authorization_params)
      redirect_to authorizations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @authorization = Authorization.find(params[:id])
    @authorization.destroy
    redirect_to authorizations_path
  end

 
  private
    def authorization_params
      params.require(:authorization).permit(:authNumber, :requestedService, :Status)
    end
end